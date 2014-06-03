package main

import (
	"flag"

	"veyron/lib/signals"
	vflag "veyron/security/flag"
	"veyron/services/mgmt/node/impl"

	"veyron2/rt"
	"veyron2/services/mgmt/application"
	"veyron2/vlog"
)

func main() {
	// TODO(rthellend): Remove the address and protocol flags when the config manager is working.
	var address, protocol, name, origin string
	flag.StringVar(&address, "address", "localhost:0", "network address to listen on")
	flag.StringVar(&name, "name", "", "name to publish the node manager at")
	flag.StringVar(&protocol, "protocol", "tcp", "network type to listen on")
	flag.StringVar(&origin, "origin", "", "node manager application repository")
	flag.Parse()
	if origin == "" {
		vlog.Fatalf("Specify an origin using --origin=<name>")
	}
	runtime := rt.Init()
	defer runtime.Shutdown()
	server, err := runtime.NewServer()
	if err != nil {
		vlog.Fatalf("NewServer() failed: %v", err)
	}
	defer server.Stop()
	envelope := &application.Envelope{}
	dispatcher := impl.NewDispatcher(envelope, origin, vflag.NewAuthorizerOrDie())
	suffix := ""
	if err := server.Register(suffix, dispatcher); err != nil {
		vlog.Fatalf("Register(%v, %v) failed: %v", suffix, dispatcher, err)
	}
	endpoint, err := server.Listen(protocol, address)
	if err != nil {
		vlog.Fatalf("Listen(%v, %v) failed: %v", protocol, address, err)
	}
	vlog.VI(0).Infof("Listening on %v", endpoint)
	if len(name) > 0 {
		if err := server.Publish(name); err != nil {
			vlog.Fatalf("Publish(%v) failed: %v", name, err)
		}
	}
	// Wait until shutdown.
	<-signals.ShutdownOnSignals()
}
