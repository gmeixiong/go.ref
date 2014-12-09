package vif_test

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"testing"

	"veyron.io/veyron/veyron/runtimes/google/ipc/stream/vif"
	"veyron.io/veyron/veyron2/naming"
)

func TestSetWithPipes(t *testing.T) {
	var (
		conn1, conn1s = net.Pipe()
		conn2, _      = net.Pipe()
		addr1         = conn1.RemoteAddr()
		addr2         = conn2.RemoteAddr()

		set = vif.NewSet()
	)
	// The server is needed to negotiate with the dialed VIF.  Otherwise, the
	// InternalNewDialedVIF below will block.
	go func() {
		_, err := vif.InternalNewAcceptedVIF(conn1s, naming.FixedRoutingID(2), nil)
		if err != nil {
			panic(fmt.Sprintf("failed to create server: %s", err))
		}
	}()
	vf, err := vif.InternalNewDialedVIF(conn1, naming.FixedRoutingID(1), nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if addr1.Network() != addr2.Network() || addr1.String() != addr2.String() {
		t.Fatalf("This test was intended for distinct connections that have duplicate RemoteAddrs. "+
			"That does not seem to be the case with (%q, %q) and (%q, %q)",
			addr1.Network(), addr1, addr2.Network(), addr2)
	}
	set.Insert(vf)
	if found := set.Find(addr2.Network(), addr2.String()); found != nil {
		t.Fatalf("Got [%v] want nil on Find(%q, %q))", found, addr2.Network(), addr2)
	}
}

func TestSetWithUnixSocket(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestSetWithFileConn")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	sockname := path.Join(dir, "socket")
	ln, err := net.Listen("unix", sockname)
	if err != nil {
		t.Fatal(err)
	}
	// Setup the creation of two separate connections.
	// At the listener, they will end up with two distinct connections
	// with the same "address" (empty string).
	go func() {
		for i := 0; i < 2; i++ {
			conn, err := net.Dial("unix", sockname)
			if err != nil {
				ln.Close()
				panic(fmt.Sprintf("dial failure: %s", err))
			}
			go func() {
				if _, err := vif.InternalNewDialedVIF(conn, naming.FixedRoutingID(1), nil, nil, nil); err != nil {
					panic(fmt.Sprintf("failed to dial VIF: %s", err))
				}
			}()
		}
	}()

	conn1, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	conn2, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	addr1 := conn1.RemoteAddr()
	addr2 := conn2.RemoteAddr()
	if addr1.Network() != addr2.Network() || addr1.String() != addr2.String() {
		t.Fatalf("This test was intended for distinct connections that have duplicate RemoteAddrs. "+
			"That does not seem to be the case with (%q, %q) and (%q, %q)",
			addr1.Network(), addr1, addr2.Network(), addr2)
	}
	vif1, err := vif.InternalNewAcceptedVIF(conn1, naming.FixedRoutingID(1), nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := vif.InternalNewAcceptedVIF(conn2, naming.FixedRoutingID(1), nil); err != nil {
		t.Fatal(err)
	}
	set := vif.NewSet()
	set.Insert(vif1)
	if found := set.Find(addr2.Network(), addr2.String()); found != nil {
		t.Errorf("Got [%v] want nil on Find(%q, %q)", found, addr2.Network(), addr2)
	}
}
