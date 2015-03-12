// Package v23tests provides support for writing end-to-end style integration
// tests. In particular, support is provided for building binaries, running
// processes, making assertions about their output/state and ensuring that
// no processes or files are left behind on exit. Since such tests are often
// difficult to debug facilities are provided to help do so.
//
// The preferred usage of this integration test framework is via the v23
// tool which generates supporting code. The primary reason for doing so is
// to cleanly separate integration tests, which can be very expensive to run,
// from normal unit tests which are intended to be fast and used constantly.
// However, it still beneficial to be able to always compile the integration
// test code with the normal test code, just not to run it. Similarly, it
// is beneficial to share as much of the existing go test infrastructure as
// possible, so the generated code uses a flag and a naming convention to
// separate the tests. Integration tests may be run in addition to unit tests
// by supplying the --v23.tests flag; the -run flag can be used
// to avoid running unit tests by specifying a prefix of TestV23 since
// the generate test functions always. Thus:
//
// v23 go test -v <pkgs> --v23.test  // runs both unit and integration tests
// v23 go test -v -run=TestV23 <pkgs> --v23.test // runs just integration tests
//
// The go generate mechanism is used to generate the test code, thus the
// comment:
//
// //go:generate v23 integration generate
//
// will generate the files v23_test.go and internal_v23_test.go for the
// package in which it occurs. Run v23 integration generate help for full
// details and options. In short, any function in an external
// (i.e. <pgk>_test) test package of the following form:
//
// V23Test<x>(t *v23tests.T)
//
// will be invoked as integration test if the --v23.tests flag is used.
//
// The generated code makes use of the RunTest function, documented below.
//
// The test environment is implemented by an instance of the interface T.
// It is constructed with an instance of another interface Test, which is
// generally implemented by testing.T. Thus, the integration test environment
// directly as follows:
//
//   func TestFoo(t *testing.T) {
//     env := v23tests.New(t)
//     defer env.Cleanup()
//
//     ...
//   }
//
// The methods in this API typically do not return error in the case of
// failure. Instead, the current test will fail with an appropriate error
// message. This avoids the need to handle errors inline the test itself.
//
// The test environment manages all built packages, subprocesses and a
// set of environment variables that are passed to subprocesses.
//
// Debugging is supported as follows:
// 1. The DebugShell method creates an interative shell at that point in
//    the tests execution that has access to all of the running processes
//    and environment of those processes. The developer can interact with
//    those processes to determine the state of the test.
// 2. Calls to methods on Test (e.g. FailNow, Fatalf) that fail the test
//    cause the Cleanup method to print out the status of all invocations.
// 3. Similarly, if the --v23.tests.shell-on-error flag is set then the
//    cleanup method will invoke a DebugShell on a test failure allowing
//    the developer to inspect the state of the test.
// 4. The implementation of this package uses filenames that start with v23test
//    to allow for easy tracing with --vmodule=v23test*=2 for example.
//
package v23tests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"testing"
	"time"

	"v.io/v23"
	"v.io/v23/security"
	"v.io/x/lib/vlog"

	"v.io/x/ref/lib/modules"
	"v.io/x/ref/lib/testutil"
	tsecurity "v.io/x/ref/lib/testutil/security"
	"v.io/x/ref/security/agent"
)

// TB is an exact mirror of testing.TB. It is provided to allow for testing
// of this package using a mock implementation. As per testing.TB, it is not
// intended to be implemented outside of this package.
type TB interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

// T represents an integration test environment.
type T struct {
	// The embedded TB
	TB

	// The function to shutdown the context used to create the environment.
	shutdown v23.Shutdown

	// The shell to use to start commands.
	shell *modules.Shell

	// The environment's root security principal.
	principal security.Principal

	// Maps path to Binary.
	builtBinaries map[string]*Binary

	tempFiles            []*os.File
	tempDirs             []string
	binDir, cachedBinDir string
	dirStack             []string

	invocations []*Invocation
}

var errNotShutdown = errors.New("has not been shutdown")

// Caller returns a string of the form <filename>:<lineno> for the
// caller specified by skip, where skip is as per runtime.Caller.
func Caller(skip int) string {
	_, file, line, _ := runtime.Caller(skip + 1)
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

// Run constructs a Binary for path and invokes Run on it.
func (t *T) Run(path string, args ...string) string {
	return t.BinaryFromPath(path).run(args...)
}

// Run constructs a Binary for path and invokes Run on it using
// the specified StartOpts
func (t *T) RunWithOpts(opts modules.StartOpts, path string, args ...string) string {
	b := t.BinaryFromPath(path)
	return b.WithStartOpts(opts).run(args...)
}

// WaitFunc is the type of the functions to be used in conjunction
// with WaitFor and WaitForAsync. It should return a value or an error
// when it wants those functions to terminate, returning a nil value
// and nil error will result in it being called again after the specified
// delay time specified in the calls to WaitFor and WaitForAsync.
type WaitFunc func() (interface{}, error)

// WaitFor calls fn at least once with the specified delay value
// between iterations until the first of the following is encountered:
// 1. fn returns a non-nil value.
// 2. fn returns an error value
// 3. fn is executed at least once and the specified timeout is exceeded.
//
// WaitFor returns the non-nil value for the first case and calls e.Fatalf for
// the other two cases.
// WaitFor will always run fn at least once to completion and hence it will
// hang if that first iteration of fn hangs. If this behaviour is not
// appropriate, then WaitForAsync should be used.
func (t *T) WaitFor(fn WaitFunc, delay, timeout time.Duration) interface{} {
	deadline := time.Now().Add(timeout)
	for {
		val, err := fn()
		if val != nil {
			return val
		}
		if err != nil {
			t.Fatalf("%s: the WaitFunc returned an error: %v", Caller(1), err)
		}
		if time.Now().After(deadline) {
			t.Fatalf("%s: timed out after %s", Caller(1), timeout)
		}
		time.Sleep(delay)
	}
}

// WaitForAsync is like WaitFor except that it calls fn in a goroutine
// and can timeout during the execution of fn.
func (t *T) WaitForAsync(fn WaitFunc, delay, timeout time.Duration) interface{} {
	resultCh := make(chan interface{})
	errCh := make(chan interface{})
	go func() {
		for {
			val, err := fn()
			if val != nil {
				resultCh <- val
				return
			}
			if err != nil {
				errCh <- err
				return
			}
			time.Sleep(delay)
		}
	}()
	select {
	case err := <-errCh:
		t.Fatalf("%s: the WaitFunc returned error: %v", Caller(1), err)
	case result := <-resultCh:
		return result
	case <-time.After(timeout):
		t.Fatalf("%s: timed out after %s", Caller(1), timeout)
	}
	return nil
}

// Pushd pushes the current working directory to the stack of
// directories, returning it as its result, and changes the working
// directory to dir.
func (t *T) Pushd(dir string) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("%s: Getwd failed: %s", Caller(1), err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("%s: Chdir failed: %s", Caller(1), err)
	}
	vlog.VI(1).Infof("Pushd: %s -> %s", cwd, dir)
	t.dirStack = append(t.dirStack, cwd)
	return cwd
}

// Popd pops the most recent entry from the directory stack and changes
// the working directory to that directory. It returns the new working
// directory as its result.
func (t *T) Popd() string {
	if len(t.dirStack) == 0 {
		t.Fatalf("%s: directory stack empty", Caller(1))
	}
	dir := t.dirStack[len(t.dirStack)-1]
	t.dirStack = t.dirStack[:len(t.dirStack)-1]
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("%s: Chdir failed: %s", Caller(1), err)
	}
	vlog.VI(1).Infof("Popd: -> %s", dir)
	return dir
}

// Caller returns a string of the form <filename>:<lineno> for the
// caller specified by skip, where skip is as per runtime.Caller.
func (t *T) Caller(skip int) string {
	return Caller(skip + 1)
}

// Principal returns the security principal of this environment.
func (t *T) Principal() security.Principal {
	return t.principal
}

// Cleanup cleans up the environment, deletes all its artifacts and
// kills all subprocesses. It will kill subprocesses in LIFO order.
// Cleanup checks to see if the test has failed and logs information
// as to the state of the processes it was asked to invoke up to that
// point and optionally, if the --v23.tests.shell-on-fail flag is set
// then it will run a debug shell before cleaning up its state.
func (t *T) Cleanup() {
	if t.Failed() {
		if testutil.IntegrationTestsDebugShellOnError {
			t.DebugSystemShell()
		}
		// Print out a summary of the invocations and their status.
		for i, inv := range t.invocations {
			if inv.hasShutdown && inv.Exists() {
				m := fmt.Sprintf("%d: %s has been shutdown but still exists: %v", i, inv.path, inv.shutdownErr)
				t.Log(m)
				vlog.VI(1).Info(m)
				vlog.VI(2).Infof("%d: %s %v", i, inv.path, inv.args)
				continue
			}
			if inv.shutdownErr != nil {
				m := fmt.Sprintf("%d: %s: shutdown status: %v", i, inv.path, inv.shutdownErr)
				t.Log(m)
				vlog.VI(1).Info(m)
				vlog.VI(2).Infof("%d: %s %v", i, inv.path, inv.args)
			}
		}
	}

	vlog.VI(1).Infof("V23Test.Cleanup")
	// Shut down all processes in LIFO order before attempting to delete any
	// files/directories to avoid potential 'file system busy' problems
	// on non-unix systems.
	for i := len(t.invocations); i > 0; i-- {
		inv := t.invocations[i-1]
		if inv.hasShutdown {
			vlog.VI(1).Infof("V23Test.Cleanup: %q has been shutdown", inv.Path())
			continue
		}
		vlog.VI(1).Infof("V23Test.Cleanup: Kill: %q", inv.Path())
		err := inv.Kill(syscall.SIGTERM)
		inv.Wait(os.Stdout, os.Stderr)
		vlog.VI(1).Infof("V23Test.Cleanup: Killed: %q: %v", inv.Path(), err)
	}
	vlog.VI(1).Infof("V23Test.Cleanup: all invocations taken care of.")

	if err := t.shell.Cleanup(os.Stdout, os.Stderr); err != nil {
		t.Fatalf("WARNING: could not clean up shell (%v)", err)
	}

	vlog.VI(1).Infof("V23Test.Cleanup: cleaning up binaries & files")

	for _, tempFile := range t.tempFiles {
		vlog.VI(1).Infof("V23Test.Cleanup: cleaning up %s", tempFile.Name())
		if err := tempFile.Close(); err != nil {
			vlog.Errorf("WARNING: Close(%q) failed: %v", tempFile.Name(), err)
		}
		if err := os.RemoveAll(tempFile.Name()); err != nil {
			vlog.Errorf("WARNING: RemoveAll(%q) failed: %v", tempFile.Name(), err)
		}
	}

	for _, tempDir := range t.tempDirs {
		vlog.VI(1).Infof("V23Test.Cleanup: cleaning up %s", tempDir)
		if err := os.RemoveAll(tempDir); err != nil {
			vlog.Errorf("WARNING: RemoveAll(%q) failed: %v", tempDir, err)
		}
	}

	// shutdown the runtime
	t.shutdown()
}

// GetVar returns the variable associated with the specified key
// and an indication of whether it is defined or not.
func (t *T) GetVar(key string) (string, bool) {
	return t.shell.GetVar(key)
}

// SetVar sets the value to be associated with key.
func (t *T) SetVar(key, value string) {
	t.shell.SetVar(key, value)
}

// ClearVar removes the speficied variable from the Shell's environment
func (t *T) ClearVar(key string) {
	t.shell.ClearVar(key)
}

func writeStringOrDie(t *T, f *os.File, s string) {
	if _, err := f.WriteString(s); err != nil {
		t.Fatalf("Write() failed: %v", err)
	}
}

// DebugSystemShell drops the user into a debug system shell (e.g. bash)
// with any environment variables specified in env... (in VAR=VAL format)
// available to it.
// If there is no controlling TTY, DebugSystemShell will emit a warning message
// and take no futher action. The DebugSystemShell also sets some environment
// variables that relate to the running test:
// - V23_TMP_DIR<#> contains the name of each temp directory created.
// - V23_BIN_DIR contains the name of the directory containing binaries.
func (t *T) DebugSystemShell(env ...string) {
	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd() failed: %v", err)
	}

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	dev := "/dev/tty"
	fd, err := syscall.Open(dev, syscall.O_RDWR, 0)
	if err != nil {
		vlog.Errorf("WARNING: Open(%v) failed, was asked to create a debug shell but cannot: %v", dev, err)
		return
	}

	var agentFile *os.File
	if creds, err := t.shell.NewChildCredentials("debug"); err == nil {
		if agentFile, err = creds.File(); err != nil {
			vlog.Errorf("WARNING: failed to obtain credentials for the debug shell: %v", err)
		}
	} else {
		vlog.Errorf("WARNING: failed to obtain credentials for the debug shell: %v", err)
	}

	file := os.NewFile(uintptr(fd), dev)
	attr := os.ProcAttr{
		Files: []*os.File{file, file, file},
		Dir:   cwd,
	}
	// Set up agent for Child.
	attr.Files = append(attr.Files, agentFile)
	attr.Env = append(attr.Env, fmt.Sprintf("%s=%d", agent.FdVarName, len(attr.Files)-1))

	// Set up environment for Child.
	for _, v := range t.shell.Env() {
		attr.Env = append(attr.Env, v)
	}

	for i, td := range t.tempDirs {
		attr.Env = append(attr.Env, fmt.Sprintf("V23_TMP_DIR%d=%s", i, td))
	}

	if len(t.cachedBinDir) > 0 {
		attr.Env = append(attr.Env, "V23_BIN_DIR="+t.BinDir())
	}
	attr.Env = append(attr.Env, env...)

	// Start up a new shell.
	writeStringOrDie(t, file, ">> Starting a new interactive shell\n")
	writeStringOrDie(t, file, "Hit CTRL-D to resume the test\n")
	if len(t.builtBinaries) > 0 {
		writeStringOrDie(t, file, "Built binaries:\n")
		for _, value := range t.builtBinaries {
			writeStringOrDie(t, file, "\t"+value.Path()+"\n")
		}
	}
	if len(t.cachedBinDir) > 0 {
		writeStringOrDie(t, file, fmt.Sprintf("Binaries are cached in %q\n", t.cachedBinDir))
	} else {
		writeStringOrDie(t, file, fmt.Sprintf("Caching of binaries was not enabled, being written to %q\n", t.binDir))
	}

	shellPath := "/bin/sh"
	if shellPathFromEnv := os.Getenv("SHELL"); shellPathFromEnv != "" {
		shellPath = shellPathFromEnv
	}
	proc, err := os.StartProcess(shellPath, []string{}, &attr)
	if err != nil {
		t.Fatalf("StartProcess(%q) failed: %v", shellPath, err)
	}

	// Wait until user exits the shell
	state, err := proc.Wait()
	if err != nil {
		t.Fatalf("Wait(%v) failed: %v", shellPath, err)
	}

	writeStringOrDie(t, file, fmt.Sprintf("<< Exited shell: %s\n", state.String()))
}

// BinaryFromPath returns a new Binary that, when started, will
// execute the executable or script at the given path.
//
// E.g. env.BinaryFromPath("/bin/bash").Start("-c", "echo hello world").Output() -> "hello world"
func (t *T) BinaryFromPath(path string) *Binary {
	return &Binary{
		env:     t,
		envVars: nil,
		path:    path,
		opts:    t.shell.DefaultStartOpts().NoExecCommand(),
	}
}

// BuildGoPkg expects a Go package path that identifies a "main"
// package and returns a Binary representing the newly built
// binary. This binary does not use the exec protocol defined
// in v.io/x/ref/lib/exec. Use this for command line tools and non
// Vanadium servers.
func (t *T) BuildGoPkg(pkg string) *Binary {
	return t.buildPkg(pkg)
}

// BuildV23 is like BuildGoPkg, but instead assumes that the resulting
// binary is a Vanadium application and does implement the exec protocol
// defined in v.io/x/ref/lib/exec. Use this for Vanadium servers.
func (t *T) BuildV23Pkg(pkg string) *Binary {
	b := t.buildPkg(pkg)
	b.opts = t.shell.DefaultStartOpts().ExternalCommand()
	return b
}

func (t *T) buildPkg(pkg string) *Binary {
	then := time.Now()
	loc := Caller(1)
	cached, built_path, err := buildPkg(t.BinDir(), pkg)
	if err != nil {
		t.Fatalf("%s: buildPkg(%s) failed: %v", loc, pkg, err)
		return nil
	}
	if _, err := os.Stat(built_path); err != nil {
		t.Fatalf("%s: buildPkg(%s) failed to stat %q", loc, pkg, built_path)
	}
	taken := time.Now().Sub(then)
	if cached {
		vlog.Infof("%s: using %s, from %s in %s.", loc, pkg, built_path, taken)
	} else {
		vlog.Infof("%s: built %s, written to %s in %s.", loc, pkg, built_path, taken)
	}
	binary := &Binary{
		env:     t,
		envVars: nil,
		path:    built_path,
		opts:    t.shell.DefaultStartOpts().NoExecCommand(),
	}
	t.builtBinaries[pkg] = binary
	return binary
}

// NewTempFile creates a temporary file. Temporary files will be deleted
// by Cleanup.
func (t *T) NewTempFile() *os.File {
	loc := Caller(1)
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("%s: TempFile() failed: %v", loc, err)
	}
	vlog.Infof("%s: created temporary file at %s", loc, f.Name())
	t.tempFiles = append(t.tempFiles, f)
	return f
}

// NewTempDir creates a temporary directory. Temporary directories and
// their contents will be deleted by Cleanup.
func (t *T) NewTempDir() string {
	loc := Caller(1)
	f, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("%s: TempDir() failed: %v", loc, err)
	}
	vlog.Infof("%s: created temporary directory at %s", loc, f)
	t.tempDirs = append(t.tempDirs, f)
	return f
}

func (t *T) appendInvocation(inv *Invocation) {
	t.invocations = append(t.invocations, inv)
}

// Creates a new local testing environment. A local testing environment has a
// a security principle available via Principal().
//
// You should clean up the returned environment using the env.Cleanup() method.
// A typical end-to-end test will begin like:
//
//   func TestFoo(t *testing.T) {
//     env := integration.NewT(t)
//     defer env.Cleanup()
//
//     ...
//   }
func New(t TB) *T {
	ctx, shutdown := v23.Init()

	vlog.Infof("creating root principal")
	principal := tsecurity.NewPrincipal("root")
	ctx, err := v23.SetPrincipal(ctx, principal)
	if err != nil {
		t.Fatalf("failed to set principal: %v", err)
	}

	shell, err := modules.NewShell(ctx, principal)
	if err != nil {
		t.Fatalf("NewShell() failed: %v", err)
	}
	opts := modules.DefaultStartOpts()
	opts.StartTimeout = time.Minute
	opts.ShutdownTimeout = 5 * time.Minute
	shell.SetDefaultStartOpts(opts)

	// The V23_BIN_DIR environment variable can be
	// used to identify a directory that multiple integration
	// tests can use to share binaries. Whoever sets this
	// environment variable is responsible for cleaning up the
	// directory it points to.
	cachedBinDir := os.Getenv("V23_BIN_DIR")
	e := &T{
		TB:            t,
		principal:     principal,
		builtBinaries: make(map[string]*Binary),
		shell:         shell,
		tempFiles:     []*os.File{},
		tempDirs:      []string{},
		cachedBinDir:  cachedBinDir,
		shutdown:      shutdown,
	}
	if len(e.cachedBinDir) == 0 {
		e.binDir = e.NewTempDir()
	}
	return e
}

func (t *T) Shell() *modules.Shell {
	return t.shell
}

// BinDir returns the directory that binarie files are stored in.
func (t *T) BinDir() string {
	if len(t.cachedBinDir) > 0 {
		return t.cachedBinDir
	}
	return t.binDir
}

// BuildPkg returns a path to a directory that contains the built binary for
// the given packages and a function that should be invoked to clean up the
// build artifacts. Note that the clients of this function should not modify
// the contents of this directory directly and instead defer to the cleanup
// function.
func buildPkg(binDir, pkg string) (bool, string, error) {
	binFile := filepath.Join(binDir, path.Base(pkg))
	if _, err := os.Stat(binFile); err != nil {
		if !os.IsNotExist(err) {
			return false, "", err
		}
		cmd := exec.Command("v23", "go", "build", "-o", binFile, pkg)
		if output, err := cmd.CombinedOutput(); err != nil {
			vlog.VI(1).Infof("\n%v:\n%v\n", strings.Join(cmd.Args, " "), string(output))
			return false, "", err
		}
		return false, binFile, nil
	}
	return true, binFile, nil
}

// RunTest runs a single Vanadium 'v23 style' integration test.
func RunTest(t *testing.T, fn func(i *T)) {
	if !testutil.IntegrationTestsEnabled {
		t.Skip()
	}
	i := New(t)
	// defer the Cleanup method so that it will be called even if
	// t.Fatalf/FailNow etc are called and can print out useful information.
	defer i.Cleanup()
	fn(i)
}

// RunRootMT builds and runs a root mount table instance. It populates
// the NAMESPACE_ROOT variable in the test environment so that all subsequent
// invocations will access this root mount table.
func RunRootMT(i *T, args ...string) (*Binary, *Invocation) {
	b := i.BuildV23Pkg("v.io/x/ref/services/mounttable/mounttabled")
	inv := b.start(1, args...)
	name := inv.ExpectVar("NAME")
	inv.Environment().SetVar("NAMESPACE_ROOT", name)
	vlog.Infof("Running root mount table: %q", name)
	return b, inv
}

// UseSharedBinDir ensures that a shared directory is used for binaries
// across multiple instances of the test environment. This is achieved
// by setting the V23_BIN_DIR environment variable if it is not already
// set in the test processes environment (as will typically be the case when
// these tests are run from the v23 tool). It is intended to be called
// from TestMain.
func UseSharedBinDir() func() {
	if v23BinDir := os.Getenv("V23_BIN_DIR"); len(v23BinDir) == 0 {
		v23BinDir, err := ioutil.TempDir("", "bin-")
		if err == nil {
			vlog.Infof("Setting V23_BIN_DIR to %q", v23BinDir)
			os.Setenv("V23_BIN_DIR", v23BinDir)
			return func() { os.RemoveAll(v23BinDir) }
		}
	} else {
		vlog.Infof("Using V23_BIN_DIR %q", v23BinDir)
	}
	return func() {}
}
