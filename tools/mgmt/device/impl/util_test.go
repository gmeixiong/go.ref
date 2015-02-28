package impl_test

import (
	"v.io/v23"
	"v.io/v23/context"

	"v.io/x/ref/lib/testutil"
	_ "v.io/x/ref/profiles"
	"v.io/x/ref/tools/mgmt/device/impl"
)

var gctx *context.T

func initTest() v23.Shutdown {
	var shutdown v23.Shutdown
	gctx, shutdown = testutil.InitForTest()
	impl.SetGlobalContext(gctx)
	return func() {
		shutdown()
		impl.SetGlobalContext(nil)
		gctx = nil
	}
}
