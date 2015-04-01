// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/verror"

	_ "v.io/x/ref/profiles"
	vsecurity "v.io/x/ref/security"
	isecurity "v.io/x/ref/services/security"
	irole "v.io/x/ref/services/security/roled/internal"
	"v.io/x/ref/test/testutil"
)

func TestSeekBlessings(t *testing.T) {
	ctx, shutdown := v23.Init()
	defer shutdown()

	workdir, err := ioutil.TempDir("", "test-role-server-")
	if err != nil {
		t.Fatal("ioutil.TempDir failed: %v", err)
	}
	defer os.RemoveAll(workdir)

	// Role A is a restricted role, i.e. it can be used in sensitive ACLs.
	roleAConf := irole.Config{
		Members: []security.BlessingPattern{
			"root/users/user1/_role",
			"root/users/user2/_role",
			"root/users/user3", // _role/A implied
		},
		Extend: true,
	}
	writeConfig(t, roleAConf, filepath.Join(workdir, "A.conf"))

	// Role B is an unrestricted role.
	roleBConf := irole.Config{
		Members: []security.BlessingPattern{
			"root/users/user1/_role",
			"root/users/user3/_role",
		},
		Audit:  true,
		Extend: false,
	}
	writeConfig(t, roleBConf, filepath.Join(workdir, "B.conf"))

	root := testutil.NewIDProvider("root")

	var (
		user1  = newPrincipalContext(t, ctx, root, "users/user1")
		user1R = newPrincipalContext(t, ctx, root, "users/user1/_role")
		user2  = newPrincipalContext(t, ctx, root, "users/user2")
		user2R = newPrincipalContext(t, ctx, root, "users/user2/_role")
		user3  = newPrincipalContext(t, ctx, root, "users/user3")
		user3R = newPrincipalContext(t, ctx, root, "users/user3", "users/user3/_role/foo", "users/user3/_role/bar")
	)

	testServerCtx := newPrincipalContext(t, ctx, root, "testserver")
	server, testAddr := newServer(t, testServerCtx)
	tDisp := &testDispatcher{}
	if err := server.ServeDispatcher("", tDisp); err != nil {
		t.Fatalf("server.ServeDispatcher failed: %v", err)
	}

	const noErr = ""
	testcases := []struct {
		ctx       *context.T
		role      string
		errID     verror.ID
		blessings []string
	}{
		{user1, "", verror.ErrNoExist.ID, nil},
		{user1, "unknown", verror.ErrNoAccess.ID, nil},
		{user2, "unknown", verror.ErrNoAccess.ID, nil},
		{user3, "unknown", verror.ErrNoAccess.ID, nil},

		{user1, "A", verror.ErrNoAccess.ID, nil},
		{user1R, "A", noErr, []string{"root/roles/A/root/users/user1/_role"}},
		{user2, "A", verror.ErrNoAccess.ID, nil},
		{user2R, "A", noErr, []string{"root/roles/A/root/users/user2/_role"}},
		{user3, "A", verror.ErrNoAccess.ID, nil},
		{user3R, "A", noErr, []string{"root/roles/A/root/users/user3/_role/bar", "root/roles/A/root/users/user3/_role/foo"}},

		{user1, "B", verror.ErrNoAccess.ID, nil},
		{user1R, "B", noErr, []string{"root/roles/B"}},
		{user2, "B", verror.ErrNoAccess.ID, nil},
		{user2R, "B", verror.ErrNoAccess.ID, nil},
		{user3, "B", verror.ErrNoAccess.ID, nil},
		{user3R, "B", noErr, []string{"root/roles/B"}},
	}
	addr := newRoleServer(t, newPrincipalContext(t, ctx, root, "roles"), workdir)
	for _, tc := range testcases {
		user := v23.GetPrincipal(tc.ctx).BlessingStore().Default()
		c := isecurity.RoleClient(naming.Join(addr, tc.role))
		blessings, err := c.SeekBlessings(tc.ctx)
		if verror.ErrorID(err) != tc.errID {
			t.Errorf("unexpected error ID for (%q, %q). Got %#v, expected %#v", user, tc.role, verror.ErrorID(err), tc.errID)
		}
		if err == nil {
			previousBlessings, _ := v23.GetPrincipal(tc.ctx).BlessingStore().Set(blessings, security.AllPrincipals)
			blessingNames, rejected := callTest(t, tc.ctx, testAddr)
			if !reflect.DeepEqual(blessingNames, tc.blessings) {
				t.Errorf("unexpected blessings for (%q, %q). Got %q, expected %q", user, tc.role, blessingNames, tc.blessings)
			}
			if len(rejected) != 0 {
				t.Errorf("unexpected rejected blessings for (%q, %q): %q", user, tc.role, rejected)
			}
			v23.GetPrincipal(tc.ctx).BlessingStore().Set(previousBlessings, security.AllPrincipals)
		}
	}
}

func newPrincipalContext(t *testing.T, ctx *context.T, root *testutil.IDProvider, names ...string) *context.T {
	principal := testutil.NewPrincipal()
	var blessings []security.Blessings
	for _, n := range names {
		blessing, err := root.NewBlessings(principal, n)
		if err != nil {
			t.Fatal("root.Bless failed for %q: %v", n, err)
		}
		blessings = append(blessings, blessing)
	}
	bUnion, err := security.UnionOfBlessings(blessings...)
	if err != nil {
		t.Fatal("security.UnionOfBlessings failed: %v", err)
	}
	vsecurity.SetDefaultBlessings(principal, bUnion)
	ctx, err = v23.SetPrincipal(ctx, principal)
	if err != nil {
		t.Fatal("v23.SetPrincipal failed: %v", err)
	}
	return ctx
}

func newRoleServer(t *testing.T, ctx *context.T, dir string) string {
	server, addr := newServer(t, ctx)
	if err := server.ServeDispatcher("", irole.NewDispatcher(dir, addr)); err != nil {
		t.Fatalf("ServeDispatcher failed: %v", err)
	}
	return addr
}

func newServer(t *testing.T, ctx *context.T) (rpc.Server, string) {
	server, err := v23.NewServer(ctx)
	if err != nil {
		t.Fatalf("NewServer() failed: %v", err)
	}
	spec := rpc.ListenSpec{Addrs: rpc.ListenAddrs{{"tcp", "127.0.0.1:0"}}}
	endpoints, err := server.Listen(spec)
	if err != nil {
		t.Fatalf("Listen(%v) failed: %v", spec, err)
	}
	return server, endpoints[0].Name()
}

func writeConfig(t *testing.T, config irole.Config, fileName string) {
	mConf, err := json.Marshal(config)
	if err != nil {
		t.Fatal("json.MarshalIndent failed: %v", err)
	}
	if err := ioutil.WriteFile(fileName, mConf, 0644); err != nil {
		t.Fatal("ioutil.WriteFile(%q, %q) failed: %v", fileName, string(mConf), err)
	}
}

func callTest(t *testing.T, ctx *context.T, addr string) (blessingNames []string, rejected []security.RejectedBlessing) {
	call, err := v23.GetClient(ctx).StartCall(ctx, addr, "Test", nil)
	if err != nil {
		t.Fatalf("StartCall failed: %v", err)
	}
	if err := call.Finish(&blessingNames, &rejected); err != nil {
		t.Fatalf("Finish failed: %v", err)
	}
	return
}

type testDispatcher struct {
}

func (d *testDispatcher) Lookup(suffix string) (interface{}, security.Authorizer, error) {
	return d, d, nil
}

func (d *testDispatcher) Authorize(*context.T) error {
	return nil
}

func (d *testDispatcher) Test(call rpc.ServerCall) ([]string, []security.RejectedBlessing, error) {
	blessings, rejected := security.RemoteBlessingNames(call.Context())
	return blessings, rejected, nil
}