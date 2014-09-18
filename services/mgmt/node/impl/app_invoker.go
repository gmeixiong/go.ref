package impl

// The app invoker is responsible for managing the state of applications on the
// node manager.  The node manager manages the applications it installs and runs
// using the following directory structure:
//
// TODO(caprita): Not all is yet implemented.
//
// <config.Root>/
//   app-<hash 1>/                  - the application dir is named using a hash of the application title
//     installation-<id 1>/         - installations are labelled with ids
//       <status>                   - one of the values for installationState enum
//       origin                     - object name for application envelope
//       <version 1 timestamp>/     - timestamp of when the version was downloaded
//         bin                      - application binary
//         previous                 - symbolic link to previous version directory
//         envelope                 - application envelope (JSON-encoded)
//       <version 2 timestamp>
//       ...
//       current                    - symbolic link to the current version
//       instances/
//         instance-<id a>/         - instances are labelled with ids
//           root/                  - workspace that the instance is run from
//           logs/                  - stderr/stdout and log files generated by instance
//           info                   - app manager name and process id for the instance (if running)
//           version                - symbolic link to installation version for the instance
//           <status>               - one of the values for instanceState enum
//         instance-<id b>
//         ...
//     installation-<id 2>
//     ...
//   app-<hash 2>
//   ...
//
// When node manager starts up, it goes through all instances and resumes the
// ones that are not suspended.  If the application was still running, it
// suspends it first.  If an application fails to resume, it stays suspended.
//
// When node manager shuts down, it suspends all running instances.
//
// Start starts an instance.  Suspend kills the process but leaves the workspace
// untouched. Resume restarts the process. Stop kills the process and prevents
// future resumes (it also eventually gc's the workspace).
//
// If the process dies on its own, it stays dead and is assumed suspended.
// TODO(caprita): Later, we'll add auto-restart option.
//
// Concurrency model: installations can be created independently of one another;
// installations can be removed at any time (any running instances will be
// stopped). The first call to Uninstall will rename the installation dir as a
// first step; subsequent Uninstall's will fail. Instances can be created
// independently of one another, as long as the installation exists (if it gets
// Uninstall'ed during an instance Start, the Start may fail).
//
// The status file present in each instance is used to flag the state of the
// instance and prevent concurrent operations against the instance:
//
// - when an instance is created with Start, it is placed in state 'suspended'.
// To run the instance, Start transitions 'suspended' to 'starting' and then
// 'started' (upon success) or the instance is deleted (upon failure).
//
// - Suspend attempts to transition from 'started' to 'suspending' (if the
// instance was not in 'started' state, Suspend fails). From 'suspending', the
// instance transitions to 'suspended' upon success or back to 'started' upon
// failure.
//
// - Resume attempts to transition from 'suspended' to 'starting' (if the
// instance was not in 'suspended' state, Resume fails). From 'starting', the
// instance transitions to 'started' upon success or back to 'suspended' upon
// failure.
//
// - Stop attempts to transition from 'started' to 'stopping' and then to
// 'stopped' (upon success) or back to 'started' (upon failure); or from
// 'suspended' to 'stopped'.  If the initial state is neither 'started' or
// 'suspended', Stop fails.
//
// TODO(caprita): There is room for synergy between how node manager organizes
// its own workspace and that for the applications it runs.  In particular,
// previous, origin, and envelope could be part of a single config.  We'll
// refine that later.

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/crc64"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"veyron.io/veyron/veyron/lib/config"
	vexec "veyron.io/veyron/veyron/services/mgmt/lib/exec"
	iconfig "veyron.io/veyron/veyron/services/mgmt/node/config"

	"veyron.io/veyron/veyron2/context"
	"veyron.io/veyron/veyron2/ipc"
	"veyron.io/veyron/veyron2/mgmt"
	"veyron.io/veyron/veyron2/naming"
	"veyron.io/veyron/veyron2/rt"
	"veyron.io/veyron/veyron2/services/mgmt/appcycle"
	"veyron.io/veyron/veyron2/services/mgmt/application"
	"veyron.io/veyron/veyron2/vlog"
)

// instanceInfo holds state about a running instance.
type instanceInfo struct {
	AppCycleMgrName string
	Pid             int
}

func saveInstanceInfo(dir string, info *instanceInfo) error {
	jsonInfo, err := json.Marshal(info)
	if err != nil {
		vlog.Errorf("Marshal(%v) failed: %v", info, err)
		return errOperationFailed
	}
	infoPath := filepath.Join(dir, "info")
	if err := ioutil.WriteFile(infoPath, jsonInfo, 0600); err != nil {
		vlog.Errorf("WriteFile(%v) failed: %v", infoPath, err)
		return errOperationFailed
	}
	return nil
}

func loadInstanceInfo(dir string) (*instanceInfo, error) {
	infoPath := filepath.Join(dir, "info")
	info := new(instanceInfo)
	if infoBytes, err := ioutil.ReadFile(infoPath); err != nil {
		vlog.Errorf("ReadFile(%v) failed: %v", infoPath, err)
		return nil, errOperationFailed
	} else if err := json.Unmarshal(infoBytes, info); err != nil {
		vlog.Errorf("Unmarshal(%v) failed: %v", infoBytes, err)
		return nil, errOperationFailed
	}
	return info, nil
}

// appInvoker holds the state of an application-related method invocation.
type appInvoker struct {
	callback *callbackState
	config   *iconfig.State
	// suffix contains the name components of the current invocation name
	// suffix.  It is used to identify an application, installation, or
	// instance.
	suffix []string
}

func saveEnvelope(dir string, envelope *application.Envelope) error {
	jsonEnvelope, err := json.Marshal(envelope)
	if err != nil {
		vlog.Errorf("Marshal(%v) failed: %v", envelope, err)
		return errOperationFailed
	}
	path := filepath.Join(dir, "envelope")
	if err := ioutil.WriteFile(path, jsonEnvelope, 0600); err != nil {
		vlog.Errorf("WriteFile(%v) failed: %v", path, err)
		return errOperationFailed
	}
	return nil
}

func loadEnvelope(dir string) (*application.Envelope, error) {
	path := filepath.Join(dir, "envelope")
	envelope := new(application.Envelope)
	if envelopeBytes, err := ioutil.ReadFile(path); err != nil {
		vlog.Errorf("ReadFile(%v) failed: %v", path, err)
		return nil, errOperationFailed
	} else if err := json.Unmarshal(envelopeBytes, envelope); err != nil {
		vlog.Errorf("Unmarshal(%v) failed: %v", envelopeBytes, err)
		return nil, errOperationFailed
	}
	return envelope, nil
}

func saveOrigin(dir, originVON string) error {
	path := filepath.Join(dir, "origin")
	if err := ioutil.WriteFile(path, []byte(originVON), 0600); err != nil {
		vlog.Errorf("WriteFile(%v) failed: %v", path, err)
		return errOperationFailed
	}
	return nil
}

func loadOrigin(dir string) (string, error) {
	path := filepath.Join(dir, "origin")
	if originBytes, err := ioutil.ReadFile(path); err != nil {
		vlog.Errorf("ReadFile(%v) failed: %v", path, err)
		return "", errOperationFailed
	} else {
		return string(originBytes), nil
	}
}

// generateID returns a new unique id string.  The uniqueness is based on the
// current timestamp.  Not cryptographically secure.
func generateID() string {
	timestamp := fmt.Sprintf("%v", time.Now().Format(time.RFC3339Nano))
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	h.Write([]byte(timestamp))
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(h.Sum64()))
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

// TODO(caprita): Nothing prevents different applications from sharing the same
// title, and thereby being installed in the same app dir.  Do we want to
// prevent that for the same user or across users?

// applicationDirName generates a cryptographic hash of the application title,
// to be used as a directory name for installations of the application with the
// given title.
func applicationDirName(title string) string {
	h := md5.New()
	h.Write([]byte(title))
	hash := strings.TrimRight(base64.URLEncoding.EncodeToString(h.Sum(nil)), "=")
	return "app-" + hash
}

func installationDirName(installationID string) string {
	return "installation-" + installationID
}

func instanceDirName(instanceID string) string {
	return "instance-" + instanceID
}

func mkdir(dir string) error {
	perm := os.FileMode(0700)
	if err := os.MkdirAll(dir, perm); err != nil {
		vlog.Errorf("MkdirAll(%v, %v) failed: %v", dir, perm, err)
		return err
	}
	return nil
}

func fetchAppEnvelope(ctx context.T, origin string) (*application.Envelope, error) {
	envelope, err := fetchEnvelope(ctx, origin)
	if err != nil {
		return nil, err
	}
	if envelope.Title == application.NodeManagerTitle {
		// Disallow node manager apps from being installed like a
		// regular app.
		return nil, errInvalidOperation
	}
	return envelope, nil
}

// newVersion sets up the directory for a new application version.
func newVersion(installationDir string, envelope *application.Envelope, oldVersionDir string) (string, error) {
	versionDir := filepath.Join(installationDir, generateVersionDirName())
	if err := mkdir(versionDir); err != nil {
		return "", errOperationFailed
	}
	// TODO(caprita): Share binaries if already existing locally.
	if err := generateBinary(versionDir, "bin", envelope, true); err != nil {
		return versionDir, err
	}
	if err := saveEnvelope(versionDir, envelope); err != nil {
		return versionDir, err
	}
	if oldVersionDir != "" {
		previousLink := filepath.Join(versionDir, "previous")
		if err := os.Symlink(oldVersionDir, previousLink); err != nil {
			vlog.Errorf("Symlink(%v, %v) failed: %v", oldVersionDir, previousLink, err)
			return versionDir, errOperationFailed
		}
	}
	// updateLink should be the last thing we do, after we've ensured the
	// new version is viable (currently, that just means it installs
	// properly).
	return versionDir, updateLink(versionDir, filepath.Join(installationDir, "current"))
}

func (i *appInvoker) Install(_ ipc.ServerContext, applicationVON string) (string, error) {
	if len(i.suffix) > 0 {
		return "", errInvalidSuffix
	}
	ctx, cancel := rt.R().NewContext().WithTimeout(time.Minute)
	defer cancel()
	envelope, err := fetchAppEnvelope(ctx, applicationVON)
	if err != nil {
		return "", err
	}
	installationID := generateID()
	installationDir := filepath.Join(i.config.Root, applicationDirName(envelope.Title), installationDirName(installationID))
	deferrer := func() {
		if err := os.RemoveAll(installationDir); err != nil {
			vlog.Errorf("RemoveAll(%v) failed: %v", installationDir, err)
		}
	}
	defer func() {
		if deferrer != nil {
			deferrer()
		}
	}()
	if _, err := newVersion(installationDir, envelope, ""); err != nil {
		return "", err
	}
	if err := saveOrigin(installationDir, applicationVON); err != nil {
		return "", err
	}
	if err := initializeInstallation(installationDir, active); err != nil {
		return "", err
	}
	deferrer = nil
	return naming.Join(envelope.Title, installationID), nil
}

func (*appInvoker) Refresh(ipc.ServerContext) error {
	// TODO(jsimsa): Implement.
	return nil
}

func (*appInvoker) Restart(ipc.ServerContext) error {
	// TODO(jsimsa): Implement.
	return nil
}

func openWriteFile(path string) (*os.File, error) {
	perm := os.FileMode(0600)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		vlog.Errorf("OpenFile(%v) failed: %v", path, err)
		return nil, errOperationFailed
	}
	return file, nil
}

// installationDir returns the path to the directory containing the app
// installation referred to by the invoker's suffix.  Returns an error if the
// suffix does not name an installation or if the named installation does not
// exist.
func (i *appInvoker) installationDir() (string, error) {
	components := i.suffix
	if nComponents := len(components); nComponents != 2 {
		return "", errInvalidSuffix
	}
	app, installation := components[0], components[1]
	installationDir := filepath.Join(i.config.Root, applicationDirName(app), installationDirName(installation))
	if _, err := os.Stat(installationDir); err != nil {
		if os.IsNotExist(err) {
			return "", errNotExist
		}
		vlog.Errorf("Stat(%v) failed: %v", installationDir, err)
		return "", errOperationFailed
	}
	return installationDir, nil
}

// newInstance sets up the directory for a new application instance.
func (i *appInvoker) newInstance() (string, string, error) {
	installationDir, err := i.installationDir()
	if err != nil {
		return "", "", err
	}
	if !installationStateIs(installationDir, active) {
		return "", "", errInvalidOperation
	}
	instanceID := generateID()
	instanceDir := filepath.Join(installationDir, "instances", instanceDirName(instanceID))
	if mkdir(instanceDir) != nil {
		return "", instanceID, errOperationFailed
	}
	currLink := filepath.Join(installationDir, "current")
	versionDir, err := filepath.EvalSymlinks(currLink)
	if err != nil {
		vlog.Errorf("EvalSymlinks(%v) failed: %v", currLink, err)
		return instanceDir, instanceID, errOperationFailed
	}
	versionLink := filepath.Join(instanceDir, "version")
	if err := os.Symlink(versionDir, versionLink); err != nil {
		vlog.Errorf("Symlink(%v, %v) failed: %v", versionDir, versionLink, err)
		return instanceDir, instanceID, errOperationFailed
	}
	if err := initializeInstance(instanceDir, suspended); err != nil {
		return instanceDir, instanceID, err
	}
	return instanceDir, instanceID, nil
}

func genCmd(instanceDir string) (*exec.Cmd, error) {
	versionLink := filepath.Join(instanceDir, "version")
	versionDir, err := filepath.EvalSymlinks(versionLink)
	if err != nil {
		vlog.Errorf("EvalSymlinks(%v) failed: %v", versionLink, err)
		return nil, errOperationFailed
	}
	envelope, err := loadEnvelope(versionDir)
	if err != nil {
		return nil, err
	}
	binPath := filepath.Join(versionDir, "bin")
	if _, err := os.Stat(binPath); err != nil {
		vlog.Errorf("Stat(%v) failed: %v", binPath, err)
		return nil, errOperationFailed
	}
	// TODO(caprita): For the purpose of isolating apps, we should run them
	// as different users.  We'll need to either use the root process or a
	// suid script to be able to do it.
	cmd := exec.Command(binPath)
	// TODO(caprita): Also pass in configuration info like NAMESPACE_ROOT to
	// the app (to point to the device mounttable).
	cmd.Env = envelope.Env
	rootDir := filepath.Join(instanceDir, "root")
	if err := mkdir(rootDir); err != nil {
		return nil, err
	}
	cmd.Dir = rootDir
	logDir := filepath.Join(instanceDir, "logs")
	if err := mkdir(logDir); err != nil {
		return nil, err
	}
	timestamp := time.Now().UnixNano()
	if cmd.Stdout, err = openWriteFile(filepath.Join(logDir, fmt.Sprintf("STDOUT-%d", timestamp))); err != nil {
		return nil, err
	}
	if cmd.Stderr, err = openWriteFile(filepath.Join(logDir, fmt.Sprintf("STDERR-%d", timestamp))); err != nil {
		return nil, err
	}
	// Set up args and env.
	cmd.Args = append(cmd.Args, "--log_dir=../logs")
	cmd.Args = append(cmd.Args, envelope.Args...)
	return cmd, nil
}

func (i *appInvoker) startCmd(instanceDir string, cmd *exec.Cmd) error {
	// Setup up the child process callback.
	callbackState := i.callback
	listener := callbackState.listenFor(mgmt.AppCycleManagerConfigKey)
	defer listener.cleanup()
	cfg := config.New()
	cfg.Set(mgmt.ParentNodeManagerConfigKey, listener.name())
	handle := vexec.NewParentHandle(cmd, vexec.ConfigOpt{cfg})
	defer func() {
		if handle != nil {
			if err := handle.Clean(); err != nil {
				vlog.Errorf("Clean() failed: %v", err)
			}
		}
	}()
	// Start the child process.
	if err := handle.Start(); err != nil {
		vlog.Errorf("Start() failed: %v", err)
		return errOperationFailed
	}
	// Wait for the child process to start.
	timeout := 10 * time.Second
	if err := handle.WaitForReady(timeout); err != nil {
		vlog.Errorf("WaitForReady(%v) failed: %v", timeout, err)
		return errOperationFailed
	}
	childName, err := listener.waitForValue(timeout)
	if err != nil {
		return errOperationFailed
	}
	instanceInfo := &instanceInfo{
		AppCycleMgrName: childName,
		Pid:             handle.Pid(),
	}
	if err := saveInstanceInfo(instanceDir, instanceInfo); err != nil {
		return err
	}
	// TODO(caprita): Spin up a goroutine to reap child status upon exit and
	// transition it to suspended state if it exits on its own.
	handle = nil
	return nil
}

func (i *appInvoker) run(instanceDir string) error {
	if err := transitionInstance(instanceDir, suspended, starting); err != nil {
		return err
	}
	cmd, err := genCmd(instanceDir)
	if err == nil {
		err = i.startCmd(instanceDir, cmd)
	}
	if err != nil {
		transitionInstance(instanceDir, starting, suspended)
		return err
	}
	return transitionInstance(instanceDir, starting, started)
}

func (i *appInvoker) Start(ipc.ServerContext) ([]string, error) {
	instanceDir, instanceID, err := i.newInstance()
	if err == nil {
		err = i.run(instanceDir)
	}
	if err != nil {
		if err := os.RemoveAll(instanceDir); err != nil {
			vlog.Errorf("RemoveAll(%v) failed: %v", instanceDir, err)
		}
		return nil, err
	}
	return []string{instanceID}, nil
}

// instanceDir returns the path to the directory containing the app instance
// referred to by the invoker's suffix, as well as the corresponding stopped
// instance dir.  Returns an error if the suffix does not name an instance.
func (i *appInvoker) instanceDir() (string, error) {
	components := i.suffix
	if nComponents := len(components); nComponents != 3 {
		return "", errInvalidSuffix
	}
	app, installation, instance := components[0], components[1], components[2]
	instancesDir := filepath.Join(i.config.Root, applicationDirName(app), installationDirName(installation), "instances")
	instanceDir := filepath.Join(instancesDir, instanceDirName(instance))
	return instanceDir, nil
}

func (i *appInvoker) Resume(ipc.ServerContext) error {
	instanceDir, err := i.instanceDir()
	if err != nil {
		return err
	}
	return i.run(instanceDir)
}

func stopAppRemotely(appVON string) error {
	appStub, err := appcycle.BindAppCycle(appVON)
	if err != nil {
		vlog.Errorf("BindAppCycle(%v) failed: %v", appVON, err)
		return errOperationFailed
	}
	ctx, cancel := rt.R().NewContext().WithTimeout(time.Minute)
	defer cancel()
	stream, err := appStub.Stop(ctx)
	if err != nil {
		vlog.Errorf("%v.Stop() failed: %v", appVON, err)
		return errOperationFailed
	}
	rstream := stream.RecvStream()
	for rstream.Advance() {
		vlog.VI(2).Infof("%v.Stop() task update: %v", appVON, rstream.Value())
	}
	if err := rstream.Err(); err != nil {
		vlog.Errorf("Advance() failed: %v", err)
		return errOperationFailed
	}
	if err := stream.Finish(); err != nil {
		vlog.Errorf("Finish() failed: %v", err)
		return errOperationFailed
	}
	return nil
}

func stop(instanceDir string) error {
	info, err := loadInstanceInfo(instanceDir)
	if err != nil {
		return err
	}
	return stopAppRemotely(info.AppCycleMgrName)
}

// TODO(caprita): implement deadline for Stop.

func (i *appInvoker) Stop(_ ipc.ServerContext, deadline uint32) error {
	instanceDir, err := i.instanceDir()
	if err != nil {
		return err
	}
	if err := transitionInstance(instanceDir, suspended, stopped); err == errOperationFailed || err == nil {
		return err
	}
	if err := transitionInstance(instanceDir, started, stopping); err != nil {
		return err
	}
	if err := stop(instanceDir); err != nil {
		transitionInstance(instanceDir, stopping, started)
		return err
	}
	return transitionInstance(instanceDir, stopping, stopped)
}

func (i *appInvoker) Suspend(ipc.ServerContext) error {
	instanceDir, err := i.instanceDir()
	if err != nil {
		return err
	}
	if err := transitionInstance(instanceDir, started, suspending); err != nil {
		return err
	}
	if err := stop(instanceDir); err != nil {
		transitionInstance(instanceDir, suspending, started)
		return err
	}
	return transitionInstance(instanceDir, suspending, suspended)
}

func (i *appInvoker) Uninstall(ipc.ServerContext) error {
	installationDir, err := i.installationDir()
	if err != nil {
		return err
	}
	return transitionInstallation(installationDir, active, uninstalled)
}

func (i *appInvoker) Update(ipc.ServerContext) error {
	installationDir, err := i.installationDir()
	if err != nil {
		return err
	}
	if !installationStateIs(installationDir, active) {
		return errInvalidOperation
	}
	originVON, err := loadOrigin(installationDir)
	if err != nil {
		return err
	}
	ctx, cancel := rt.R().NewContext().WithTimeout(time.Minute)
	defer cancel()
	newEnvelope, err := fetchAppEnvelope(ctx, originVON)
	if err != nil {
		return err
	}
	currLink := filepath.Join(installationDir, "current")
	oldVersionDir, err := filepath.EvalSymlinks(currLink)
	if err != nil {
		vlog.Errorf("EvalSymlinks(%v) failed: %v", currLink, err)
		return errOperationFailed
	}
	// NOTE(caprita): A race can occur between two competing updates, where
	// both use the old version as their baseline.  This can result in both
	// updates succeeding even if they are updating the app installation to
	// the same new envelope.  This will result in one of the updates
	// becoming the new 'current'.  Both versions will point their
	// 'previous' link to the old version.  This doesn't appear to be of
	// practical concern, so we avoid the complexity of synchronizing
	// updates.
	oldEnvelope, err := loadEnvelope(oldVersionDir)
	if err != nil {
		return err
	}
	if oldEnvelope.Title != newEnvelope.Title {
		return errIncompatibleUpdate
	}
	if reflect.DeepEqual(oldEnvelope, newEnvelope) {
		return errUpdateNoOp
	}
	versionDir, err := newVersion(installationDir, newEnvelope, oldVersionDir)
	if err != nil {
		if err := os.RemoveAll(versionDir); err != nil {
			vlog.Errorf("RemoveAll(%v) failed: %v", versionDir, err)
		}
		return err
	}
	return nil
}

func (*appInvoker) UpdateTo(_ ipc.ServerContext, von string) error {
	// TODO(jsimsa): Implement.
	return nil
}

func (i *appInvoker) Revert(ipc.ServerContext) error {
	installationDir, err := i.installationDir()
	if err != nil {
		return err
	}
	if !installationStateIs(installationDir, active) {
		return errInvalidOperation
	}
	// NOTE(caprita): A race can occur between an update and a revert, where
	// both use the same current version as their starting point.  This will
	// render the update inconsequential.  This doesn't appear to be of
	// practical concern, so we avoid the complexity of synchronizing
	// updates and revert operations.
	currLink := filepath.Join(installationDir, "current")
	currVersionDir, err := filepath.EvalSymlinks(currLink)
	if err != nil {
		vlog.Errorf("EvalSymlinks(%v) failed: %v", currLink, err)
		return errOperationFailed
	}
	previousLink := filepath.Join(currVersionDir, "previous")
	if _, err := os.Lstat(previousLink); err != nil {
		if os.IsNotExist(err) {
			// No 'previous' link -- must be the first version.
			return errUpdateNoOp
		}
		vlog.Errorf("Lstat(%v) failed: %v", previousLink, err)
		return errOperationFailed
	}
	prevVersionDir, err := filepath.EvalSymlinks(previousLink)
	if err != nil {
		vlog.Errorf("EvalSymlinks(%v) failed: %v", previousLink, err)
		return errOperationFailed
	}
	return updateLink(prevVersionDir, currLink)
}
