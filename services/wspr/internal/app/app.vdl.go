// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: app.vdl

// The app package contains the struct that keeps per javascript app state and handles translating
// javascript requests to vanadium requests and vice versa.
package app

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"time"
	"v.io/v23/security"
	time_2 "v.io/v23/vdlroot/time"
	"v.io/v23/vtrace"
	"v.io/x/ref/services/wspr/internal/principal"
	"v.io/x/ref/services/wspr/internal/rpc/server"
)

type RpcRequest struct {
	Name         string
	Method       string
	NumInArgs    int32
	NumOutArgs   int32
	IsStreaming  bool
	Deadline     time_2.Deadline
	TraceRequest vtrace.Request
	Context      server.Context
	CallOptions  []RpcCallOption
}

func (RpcRequest) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/app.RpcRequest"`
}) {
}

type (
	// RpcCallOption represents any single field of the RpcCallOption union type.
	RpcCallOption interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the RpcCallOption union type.
		__VDLReflect(__RpcCallOptionReflect)
	}
	// RpcCallOptionAllowedServersPolicy represents field AllowedServersPolicy of the RpcCallOption union type.
	RpcCallOptionAllowedServersPolicy struct{ Value []security.BlessingPattern }
	// RpcCallOptionRetryTimeout represents field RetryTimeout of the RpcCallOption union type.
	RpcCallOptionRetryTimeout struct{ Value time.Duration }
	// RpcCallOptionGranter represents field Granter of the RpcCallOption union type.
	RpcCallOptionGranter struct{ Value GranterHandle }
	// __RpcCallOptionReflect describes the RpcCallOption union type.
	__RpcCallOptionReflect struct {
		Name  string `vdl:"v.io/x/ref/services/wspr/internal/app.RpcCallOption"`
		Type  RpcCallOption
		Union struct {
			AllowedServersPolicy RpcCallOptionAllowedServersPolicy
			RetryTimeout         RpcCallOptionRetryTimeout
			Granter              RpcCallOptionGranter
		}
	}
)

func (x RpcCallOptionAllowedServersPolicy) Index() int                          { return 0 }
func (x RpcCallOptionAllowedServersPolicy) Interface() interface{}              { return x.Value }
func (x RpcCallOptionAllowedServersPolicy) Name() string                        { return "AllowedServersPolicy" }
func (x RpcCallOptionAllowedServersPolicy) __VDLReflect(__RpcCallOptionReflect) {}

func (x RpcCallOptionRetryTimeout) Index() int                          { return 1 }
func (x RpcCallOptionRetryTimeout) Interface() interface{}              { return x.Value }
func (x RpcCallOptionRetryTimeout) Name() string                        { return "RetryTimeout" }
func (x RpcCallOptionRetryTimeout) __VDLReflect(__RpcCallOptionReflect) {}

func (x RpcCallOptionGranter) Index() int                          { return 2 }
func (x RpcCallOptionGranter) Interface() interface{}              { return x.Value }
func (x RpcCallOptionGranter) Name() string                        { return "Granter" }
func (x RpcCallOptionGranter) __VDLReflect(__RpcCallOptionReflect) {}

type (
	// RpcServerOption represents any single field of the RpcServerOption union type.
	RpcServerOption interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the RpcServerOption union type.
		__VDLReflect(__RpcServerOptionReflect)
	}
	// RpcServerOptionIsLeaf represents field IsLeaf of the RpcServerOption union type.
	RpcServerOptionIsLeaf struct{ Value bool }
	// RpcServerOptionServesMountTable represents field ServesMountTable of the RpcServerOption union type.
	RpcServerOptionServesMountTable struct{ Value bool }
	// __RpcServerOptionReflect describes the RpcServerOption union type.
	__RpcServerOptionReflect struct {
		Name  string `vdl:"v.io/x/ref/services/wspr/internal/app.RpcServerOption"`
		Type  RpcServerOption
		Union struct {
			IsLeaf           RpcServerOptionIsLeaf
			ServesMountTable RpcServerOptionServesMountTable
		}
	}
)

func (x RpcServerOptionIsLeaf) Index() int                            { return 0 }
func (x RpcServerOptionIsLeaf) Interface() interface{}                { return x.Value }
func (x RpcServerOptionIsLeaf) Name() string                          { return "IsLeaf" }
func (x RpcServerOptionIsLeaf) __VDLReflect(__RpcServerOptionReflect) {}

func (x RpcServerOptionServesMountTable) Index() int                            { return 1 }
func (x RpcServerOptionServesMountTable) Interface() interface{}                { return x.Value }
func (x RpcServerOptionServesMountTable) Name() string                          { return "ServesMountTable" }
func (x RpcServerOptionServesMountTable) __VDLReflect(__RpcServerOptionReflect) {}

type RpcResponse struct {
	OutArgs       []*vdl.Value
	TraceResponse vtrace.Response
}

func (RpcResponse) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/app.RpcResponse"`
}) {
}

type GranterHandle int32

func (GranterHandle) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/app.GranterHandle"`
}) {
}

type GranterRequest struct {
	GranterHandle GranterHandle
	Call          server.SecurityCall
}

func (GranterRequest) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/app.GranterRequest"`
}) {
}

type GranterResponse struct {
	Blessings principal.BlessingsHandle
	Err       error
}

func (GranterResponse) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/app.GranterResponse"`
}) {
}

func init() {
	vdl.Register((*RpcRequest)(nil))
	vdl.Register((*RpcCallOption)(nil))
	vdl.Register((*RpcServerOption)(nil))
	vdl.Register((*RpcResponse)(nil))
	vdl.Register((*GranterHandle)(nil))
	vdl.Register((*GranterRequest)(nil))
	vdl.Register((*GranterResponse)(nil))
}
