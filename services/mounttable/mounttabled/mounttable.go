// mounttabled is a simple mount table daemon.
package main

import (
	"flag"
	"fmt"
	"os"

	"veyron2/rt"
	"veyron2/vlog"

	"veyron/lib/signals"

	"veyron/services/mounttable/lib"
)

var (
	mountName = flag.String("name", "", "Name to mount this mountable as.  Empty means don't mount.")
	address   = flag.String("address", ":0", "Address to listen on.  Default is to use a randomly assigned port")
)

const usage = `%s is a simple mount table daemon.

Usage:

  %s [--name=<name>]

  <name>, if provided, causes the mount table to mount itself under that name.
  The name may be absolute for a remote mount table service (e.g., "/<remote mt
  address>//some/suffix") or could be relative to this process' default mount
  table (e.g., "some/suffix").
`

func Usage() {
	fmt.Fprintf(os.Stderr, usage, os.Args[0], os.Args[0])
}

func main() {
	// TODO(cnicolaou): fix Usage so that it includes the flags defined by
	// the runtime
	flag.Usage = Usage
	r := rt.Init()
	defer r.Shutdown()

	server, err := r.NewServer()
	if err != nil {
		vlog.Errorf("r.NewServer failed: %v", err)
		return
	}
	defer server.Stop()
	mtPrefix := "mt"
	if err := server.Register(mtPrefix, mounttable.NewMountTable()); err != nil {
		vlog.Errorf("server.Register failed to register mount table: %v", err)
		return
	}
	endpoint, err := server.Listen("tcp", *address)
	if err != nil {
		vlog.Errorf("server.Listen failed: %v", err)
		return
	}
	if name := *mountName; len(name) > 0 {
		if err := server.Publish(name); err != nil {
			vlog.Errorf("Publish(%v) failed: %v", name, err)
			return
		}
		vlog.Infof("Mount table service at: %v/%v (/%v/%v)", name, mtPrefix, endpoint, mtPrefix)

	} else {
		vlog.Infof("Mount table at: /%v/%v", endpoint, mtPrefix)
	}

	// Wait until signal is received.
	vlog.Info("Received signal ", <-signals.ShutdownOnSignals())
}
