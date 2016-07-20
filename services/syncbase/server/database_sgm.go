// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"v.io/v23/context"
	"v.io/v23/rpc"
	wire "v.io/v23/services/syncbase"
	"v.io/x/ref/services/syncbase/vsync"
)

////////////////////////////////////////
// Syncgroup RPC methods

// Note, existence and access authorization is checked in SyncDatabase methods.

func (d *database) ListSyncgroups(ctx *context.T, call rpc.ServerCall) ([]wire.Id, error) {
	sd := vsync.NewSyncDatabase(d)
	return sd.ListSyncgroups(ctx, call)
}

func (d *database) CreateSyncgroup(ctx *context.T, call rpc.ServerCall, sgId wire.Id, spec wire.SyncgroupSpec, myInfo wire.SyncgroupMemberInfo) error {
	sd := vsync.NewSyncDatabase(d)
	return sd.CreateSyncgroup(ctx, call, sgId, spec, myInfo)
}

func (d *database) JoinSyncgroup(ctx *context.T, call rpc.ServerCall, remoteSyncbaseName string, expectedSyncbaseBlessings []string, sgId wire.Id, myInfo wire.SyncgroupMemberInfo) (wire.SyncgroupSpec, error) {
	sd := vsync.NewSyncDatabase(d)
	return sd.JoinSyncgroup(ctx, call, remoteSyncbaseName, expectedSyncbaseBlessings, sgId, myInfo)
}

func (d *database) LeaveSyncgroup(ctx *context.T, call rpc.ServerCall, sgId wire.Id) error {
	sd := vsync.NewSyncDatabase(d)
	return sd.LeaveSyncgroup(ctx, call, sgId)
}

func (d *database) DestroySyncgroup(ctx *context.T, call rpc.ServerCall, sgId wire.Id) error {
	sd := vsync.NewSyncDatabase(d)
	return sd.DestroySyncgroup(ctx, call, sgId)
}

func (d *database) EjectFromSyncgroup(ctx *context.T, call rpc.ServerCall, sgId wire.Id, member string) error {
	sd := vsync.NewSyncDatabase(d)
	return sd.EjectFromSyncgroup(ctx, call, sgId, member)
}

func (d *database) GetSyncgroupSpec(ctx *context.T, call rpc.ServerCall, sgId wire.Id) (wire.SyncgroupSpec, string, error) {
	sd := vsync.NewSyncDatabase(d)
	return sd.GetSyncgroupSpec(ctx, call, sgId)
}

func (d *database) SetSyncgroupSpec(ctx *context.T, call rpc.ServerCall, sgId wire.Id, spec wire.SyncgroupSpec, version string) error {
	sd := vsync.NewSyncDatabase(d)
	return sd.SetSyncgroupSpec(ctx, call, sgId, spec, version)
}

func (d *database) GetSyncgroupMembers(ctx *context.T, call rpc.ServerCall, sgId wire.Id) (map[string]wire.SyncgroupMemberInfo, error) {
	sd := vsync.NewSyncDatabase(d)
	return sd.GetSyncgroupMembers(ctx, call, sgId)
}
