// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysinit

// TODO(cnicolaou): will need to figure out a simple of way of handling the
// different init systems supported by various versions of linux. One simple
// option is to just include them in the name when installing - e.g.
// simplevns-upstart.

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"v.io/x/ref"
)

// action is a var so we can override it for testing.
var action = func(command, action, service string) error {
	cmd := exec.Command(command, action, service)
	if os.Geteuid() == 0 && os.Getuid() > 0 {
		// Set uid to root (e.g. when running from a suid binary),
		// otherwise initctl doesn't work.
		sysProcAttr := new(syscall.SysProcAttr)
		sysProcAttr.Credential = new(syscall.Credential)
		sysProcAttr.Credential.Gid = uint32(0)
		sysProcAttr.Credential.Uid = uint32(0)
		cmd.SysProcAttr = sysProcAttr
	}
	// Clear env.  In particular, initctl doesn't like USER being set to
	// something other than root.
	cmd.Env = []string{ref.RPCTransitionStateVar + "=" + os.Getenv(ref.RPCTransitionStateVar)}
	output, err := cmd.CombinedOutput()
	fmt.Fprintf(os.Stderr, "%s output: for %s %s: %s\n", command, action, service, output)
	return err
}

var (
	upstartDir        = "/etc/init"
	upstartBin        = "/sbin/initctl"
	systemdDir        = "/lib/systemd/system" // This works for both rpi and edison (/usr/lib does not)
	systemdTmpFileDir = "/usr/lib/tmpfiles.d"
	dockerDir         = "/home/veyron/init"
)

// InitSystem attempts to determine what kind of init system is in use on
// the platform that it is run on. It recognises upstart and systemd by
// testing for the presence of the initctl and systemctl commands. upstart
// is tested for first and hence is preferred in the unlikely case that both
// are installed. Docker containers do not support upstart and systemd and
// for them we have our own init system that uses the daemon command to
// start/stop/respawn jobs.
func InitSystem() string {
	// NOTE(spetrovic): This check is kind of a hack. Ideally, we would
	// detect a docker system by looking at the "container=lxc" environment
	// variable.  However, we run sysinit during image creation, at which
	// point we're on a native system and this variable isn't set.
	if fi, err := os.Stat("/home/veyron/init"); err == nil && fi.Mode().IsDir() {
		return "docker"
	}
	if fi, err := os.Stat(upstartBin); err == nil {
		if (fi.Mode() & 0100) != 0 {
			return "upstart"
		}
	}

	if findSystemdSystemCtl() != "" {
		return "systemd"
	}

	return ""
}

// New returns the appropriate implementation of InstallSystemInit for the
// underlying system.
func New(system string, sd *ServiceDescription) InstallSystemInit {
	switch system {
	case "docker":
		return (*DockerService)(sd)
	case "upstart":
		return (*UpstartService)(sd)
	case "systemd":
		return (*SystemdService)(sd)
	default:
		return nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// Upstart support

// See http://upstart.ubuntu.com/cookbook/ for info on upstart

// UpstartService is the implementation of InstallSystemInit interfacing with
// the Upstart system.
type UpstartService ServiceDescription

var upstartTemplate = `# This file was auto-generated by the Vanadium SysInit tool.
# Date: {{.Date}}
#
# {{.Service}} - {{.Description}}
#
# Upstart config for Ubuntu-GCE

description	"{{.Description}}"

start on runlevel [2345]
stop on runlevel [!2345]
{{if .Environment}}
# Environment variables
{{range $var, $value := .Environment}}
env {{$var}}={{$value}}{{end}}
{{end}}
respawn
respawn limit 10 5
umask 022

pre-start script
    test -x {{.Binary}} || { stop; exit 0; }
    mkdir -p -m0755 /var/log/veyron
    chown -R {{.User}} /var/log/veyron
end script

script
  set -e
  echo '{{.Service}} starting'
  exec sudo -u {{.User}} {{range $cmd := .Command}} {{$cmd}}{{end}}
end script
`

// Implements the InstallSystemInit method.
func (u *UpstartService) Install() error {
	file := fmt.Sprintf("%s/%s.conf", upstartDir, u.Service)
	return (*ServiceDescription)(u).writeTemplate(upstartTemplate, file)
}

// Print implements the InstallSystemInit method.
func (u *UpstartService) Print() error {
	return (*ServiceDescription)(u).writeTemplate(upstartTemplate, "")
}

// Implements the InstallSystemInit method.
func (u *UpstartService) Uninstall() error {
	// For now, ignore any errors returned by Stop, since Stop complains
	// when there is no instance to stop.
	// TODO(caprita): Only call Stop if there are running instances.
	u.Stop()
	file := fmt.Sprintf("%s/%s.conf", upstartDir, u.Service)
	return os.Remove(file)
}

// Start implements the InstallSystemInit method.
func (u *UpstartService) Start() error {
	return action(upstartBin, "start", u.Service)
}

// Stop implements the InstallSystemInit method.
func (u *UpstartService) Stop() error {
	return action(upstartBin, "stop", u.Service)
}

///////////////////////////////////////////////////////////////////////////////
// Systemd support

// SystemdService is the implementation of InstallSystemInit interfacing with
// the Systemd system.
type SystemdService ServiceDescription

const systemdTemplate = `# This file was auto-generated by the Vanadium SysInit tool.
# Date: {{.Date}}
#
# {{.Service}} - {{.Description}}
#
[Unit]
Description={{.Description}}
After=openntpd.service

[Service]
User={{.User}}{{if .Environment}}{{println}}Environment={{range $var, $value := .Environment}}"{{$var}}={{$value}}" {{end}}{{end}}
ExecStart={{range $cmd := .Command}}{{$cmd}} {{end}}
ExecReload=/bin/kill -HUP $MAINPID
KillMode=process
Restart=always
RestartSecs=10
StandardOutput=syslog


[Install]
WantedBy=multi-user.target
`

// Install implements the InstallSystemInit method.
func (s *SystemdService) Install() error {
	file := fmt.Sprintf("%s/%s.service", systemdDir, s.Service)
	if err := (*ServiceDescription)(s).writeTemplate(systemdTemplate, file); err != nil {
		return fmt.Errorf("failed to write template (uid= %d, euid= %d): %v", os.Getuid(), os.Geteuid(), err)
	}
	file = fmt.Sprintf("%s/veyron.conf", systemdTmpFileDir)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	f.WriteString("d /var/log/veyron 0755 veyron veyron\n")
	f.Close()
	err = action("systemd-tmpfiles", "--create", file)
	if err != nil {
		return err
	}

	// First call disable to get rid of any symlink lingering around from a previous install
	// We don't care about the return status on the disable action.
	action(findSystemdSystemCtl(), "disable", s.Service)
	return action(findSystemdSystemCtl(), "enable", s.Service)
}

// Print implements the InstallSystemInit method.
func (s *SystemdService) Print() error {
	return (*ServiceDescription)(s).writeTemplate(systemdTemplate, "")
}

// Uninstall implements the InstallSystemInit method.
func (s *SystemdService) Uninstall() error {
	if err := s.Stop(); err != nil {
		return err
	}
	if err := action(findSystemdSystemCtl(), "disable", s.Service); err != nil {
		return err
	}
	file := fmt.Sprintf("%s/%s.service", systemdDir, s.Service)
	return os.Remove(file)
}

// Start implements the InstallSystemInit method.
func (s *SystemdService) Start() error {
	return action(findSystemdSystemCtl(), "start", s.Service)
}

// Stop implements the InstallSystemInit method.
func (s *SystemdService) Stop() error {
	return action(findSystemdSystemCtl(), "stop", s.Service)
}

// This is a variable so it can be overridden for testing.
var findSystemdSystemCtl = func() string {
	// Systems using systemd may have systemctl in one of several possible places. This finds it.
	paths := []string{"/sbin", "/bin", "/usr/bin", "/usr/sbin"}

	for _, path := range paths {
		testpath := filepath.Join(path, "systemctl")
		if fi, err := os.Stat(testpath); err == nil && (fi.Mode()&0100) != 0 {
			return testpath
		}
	}

	return ""
}

///////////////////////////////////////////////////////////////////////////////
// Docker support

// DockerService is the implementation of InstallSystemInit interfacing with
// Docker.
type DockerService ServiceDescription

const dockerTemplate = `#!/bin/bash
# This file was auto-generated by the Vanadium SysInit tool.
# Date: {{.Date}}
#
# {{.Service}} - {{.Description}}
#
set -e
{{if .Environment}}
# Environment variables
{{range $var, $value := .Environment}}export {{$var}}={{$value}}{{end}}
{{end}}
echo '{{.Service}} setup.'
test -x {{.Binary}} || { stop; exit 0; }
mkdir -p -m0755 /var/log/veyron
chown -R {{.User}} /var/log/veyron

echo '{{.Service}} starting'
exec daemon -n {{.Service}} -r -A 2 -L 10 -M 5 -X '{{range $cmd := .Command}} {{$cmd}}{{end}}' &
`

// Install implements the InstallSystemInit method.
func (s *DockerService) Install() error {
	file := fmt.Sprintf("%s/%s.sh", dockerDir, s.Service)
	if err := (*ServiceDescription)(s).writeTemplate(dockerTemplate, file); err != nil {
		return err
	}
	os.Chmod(file, 0755)
	return nil
}

// Print implements the InstallSystemInit method.
func (s *DockerService) Print() error {
	return (*ServiceDescription)(s).writeTemplate(dockerTemplate, "")
}

// Uninstall implements the InstallSystemInit method.
func (s *DockerService) Uninstall() error {
	if err := s.Stop(); err != nil {
		return err
	}
	file := fmt.Sprintf("%s/%s.sh", dockerDir, s.Service)
	return os.Remove(file)
}

// Start implements the InstallSystemInit method.
func (s *DockerService) Start() error {
	return action(fmt.Sprintf("%s/%s.sh", dockerDir, s.Service), "", s.Service)
}

// Stop implements the InstallSystemInit method.
func (s *DockerService) Stop() error {
	return action("daemon", fmt.Sprintf("-n %s --stop", s.Service), s.Service)
}
