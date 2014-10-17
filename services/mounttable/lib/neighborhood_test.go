package mounttable

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"veyron.io/veyron/veyron2/naming"
	"veyron.io/veyron/veyron2/services/mounttable"
	"veyron.io/veyron/veyron2/vlog"

	"veyron.io/veyron/veyron/lib/testutil"
	"veyron.io/veyron/veyron/profiles"
)

func init() { testutil.Init() }

func protocolAndAddress(e naming.Endpoint) (string, string, error) {
	addr := e.Addr()
	if addr == nil {
		return "", "", fmt.Errorf("failed to get address")
	}
	return addr.Network(), addr.String(), nil
}

func TestNeighborhood(t *testing.T) {
	vlog.Infof("TestNeighborhood")
	server, err := rootRT.NewServer()
	if err != nil {
		boom(t, "r.NewServer: %s", err)
	}
	defer server.Stop()

	// Start serving on a loopback address.
	e, err := server.ListenX(profiles.LocalListenSpec)
	if err != nil {
		boom(t, "Failed to Listen mount table: %s", err)
	}
	estr := e.String()
	addresses := []string{
		naming.JoinAddressName(estr, ""),
		naming.JoinAddressName(estr, "suffix1"),
		naming.JoinAddressName(estr, "suffix2"),
	}

	// Create a name for the server.
	serverName := fmt.Sprintf("nhtest%d", os.Getpid())

	// Add neighborhood server.
	nhd, err := NewLoopbackNeighborhoodServer(serverName, addresses...)
	if err != nil {
		boom(t, "Failed to create neighborhood server: %s\n", err)
	}
	defer nhd.Stop()
	if err := server.Serve("", nhd); err != nil {
		boom(t, "Failed to register neighborhood server: %s", err)
	}

	// Wait for the mounttable to appear in mdns
L:
	for tries := 1; tries < 2; tries++ {
		names := doGlob(t, naming.JoinAddressName(estr, "//"), "*", rootRT)
		t.Logf("names %v", names)
		for _, n := range names {
			if n == serverName {
				break L
			}
		}
		time.Sleep(1 * time.Second)
	}

	// Make sure we get back a root for the server.
	want, got := []string{""}, doGlob(t, naming.JoinAddressName(estr, "//"+serverName), "", rootRT)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Unexpected Glob result want: %q, got: %q", want, got)
	}

	// Make sure we can resolve through the neighborhood.
	expectedSuffix := "a/b"
	objectPtr, err := mounttable.BindMountTable(naming.JoinAddressName(estr, "//"+serverName+"/"+expectedSuffix), quuxClient(rootRT))
	if err != nil {
		boom(t, "BindMountTable: %s", err)
	}
	servers, suffix, err := objectPtr.ResolveStep(rootRT.NewContext())
	if err != nil {
		boom(t, "resolveStep: %s", err)
	}

	// Resolution returned something.  Make sure its correct.
	if suffix != expectedSuffix {
		boom(t, "resolveStep suffix: expected %s, got %s", expectedSuffix, suffix)
	}
	if len(servers) == 0 {
		boom(t, "resolveStep returns no severs")
	}
L2:
	for _, s := range servers {
		for _, a := range addresses {
			if a == s.Server {
				continue L2
			}
		}
		boom(t, "Unexpected address from resolveStep result: %v", s.Server)
	}
L3:
	for _, a := range addresses {
		for _, s := range servers {
			if a == s.Server {
				continue L3
			}
		}
		boom(t, "Missing address from resolveStep result: %v", a)
	}
}
