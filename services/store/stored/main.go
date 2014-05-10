// stored is a storage server.
//
// Usage:
//
//     stored [--name=<mount>] [--db=<dbName>]
//
//     - <name> is the Veyron mount point name, default /global/vstore/<hostname>/<username>.
//     - <dbName> is the filename in which to store the data.
//
// The Store service has Veyron name, <name>/.store.  Individual values with
// path <path> have name <name>/<path>.
package main

import (
	"flag"
	"log"
	"os"
	"os/user"

	"veyron/services/store/server"

	"veyron2"
	"veyron2/rt"
	"veyron2/security"
)

var (
	mountName string
	dbName    = flag.String("db", "/var/tmp/veyron_store.db", "Metadata database")

	// TODO(jyh): Figure out how to get a real public ID.
	rootPublicID    security.PublicID = security.FakePublicID("root")
	serverPrivateID                   = security.FakePrivateID("store")
)

func init() {
	username := "unknown"
	if u, err := user.Current(); err == nil {
		username = u.Username
	}
	hostname := "unknown"
	if h, err := os.Hostname(); err == nil {
		hostname = h
	}
	dir := "global/vstore/" + hostname + "/" + username
	flag.StringVar(&mountName, "name", dir, "Mount point for media")
}

// Main starts the content service, taking arguments from the command line
// flags.
func main() {
	r := rt.Init()

	// Create a new server instance.
	s, err := r.NewServer(veyron2.LocalID(serverPrivateID))
	if err != nil {
		log.Fatal("r.NewServer() failed: ", err)
	}

	// Create a new StoreService.
	storeService, err := server.New(server.ServerConfig{Admin: rootPublicID, DBName: *dbName})
	if err != nil {
		log.Fatal("server.New() failed: ", err)
	}
	defer storeService.Close()

	// Register the services.
	storeDisp := server.NewStoreDispatcher(storeService)
	objectDisp := server.NewObjectDispatcher(storeService)
	if err := s.Register(".store", storeDisp); err != nil {
		log.Fatal("s.Register(storeDisp) failed: ", err)
	}
	if err := s.Register("", objectDisp); err != nil {
		log.Fatal("s.Register(objectDisp) failed: ", err)
	}

	// Create an endpoint and start listening.
	ep, err := s.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("s.Listen() failed: ", err)
	}

	// Publish the service in the mount table.
	log.Printf("Mounting store on %s, endpoint /%s", mountName, ep)
	if err := s.Publish(mountName); err != nil {
		log.Fatal("s.Publish() failed: ", err)
	}

	// Wait forever.
	done := make(chan struct{})
	<-done
}
