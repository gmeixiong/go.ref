// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: controller.vdl

package app

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"

	// VDL user imports
	"v.io/v23/security"
	"v.io/v23/vdlroot/signature"
	"v.io/x/ref/services/wspr/internal/principal"
)

// ControllerClientMethods is the client interface
// containing Controller methods.
type ControllerClientMethods interface {
	// Serve instructs WSPR to start listening for calls on behalf
	// of a javascript server.
	Serve(ctx *context.T, name string, serverId uint32, opts ...rpc.CallOpt) error
	// Stop instructs WSPR to stop listening for calls for the
	// given javascript server.
	Stop(ctx *context.T, serverId uint32, opts ...rpc.CallOpt) error
	// AddName adds a published name to an existing server.
	AddName(ctx *context.T, serverId uint32, name string, opts ...rpc.CallOpt) error
	// RemoveName removes a published name from an existing server.
	RemoveName(ctx *context.T, serverId uint32, name string, opts ...rpc.CallOpt) error
	// UnlinkBlessings removes the given blessings from the blessings store.
	UnlinkBlessings(ctx *context.T, handle principal.BlessingsHandle, opts ...rpc.CallOpt) error
	// Bless binds extensions of blessings held by this principal to
	// another principal (represented by its public key).
	Bless(ctx *context.T, publicKey string, blessingHandle principal.BlessingsHandle, extension string, caveat []security.Caveat, opts ...rpc.CallOpt) (string, principal.BlessingsHandle, error)
	// BlessSelf creates a blessing with the provided name for this principal.
	BlessSelf(ctx *context.T, name string, caveats []security.Caveat, opts ...rpc.CallOpt) (string, principal.BlessingsHandle, error)
	// RemoteBlessings fetches the remote blessings for a given name and method.
	RemoteBlessings(ctx *context.T, name string, method string, opts ...rpc.CallOpt) ([]string, error)
	// Signature fetches the signature for a given name.
	Signature(ctx *context.T, name string, opts ...rpc.CallOpt) ([]signature.Interface, error)
}

// ControllerClientStub adds universal methods to ControllerClientMethods.
type ControllerClientStub interface {
	ControllerClientMethods
	rpc.UniversalServiceMethods
}

// ControllerClient returns a client stub for Controller.
func ControllerClient(name string) ControllerClientStub {
	return implControllerClientStub{name}
}

type implControllerClientStub struct {
	name string
}

func (c implControllerClientStub) Serve(ctx *context.T, i0 string, i1 uint32, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Serve", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) Stop(ctx *context.T, i0 uint32, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Stop", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) AddName(ctx *context.T, i0 uint32, i1 string, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "AddName", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) RemoveName(ctx *context.T, i0 uint32, i1 string, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "RemoveName", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) UnlinkBlessings(ctx *context.T, i0 principal.BlessingsHandle, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "UnlinkBlessings", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implControllerClientStub) Bless(ctx *context.T, i0 string, i1 principal.BlessingsHandle, i2 string, i3 []security.Caveat, opts ...rpc.CallOpt) (o0 string, o1 principal.BlessingsHandle, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Bless", []interface{}{i0, i1, i2, i3}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0, &o1)
	return
}

func (c implControllerClientStub) BlessSelf(ctx *context.T, i0 string, i1 []security.Caveat, opts ...rpc.CallOpt) (o0 string, o1 principal.BlessingsHandle, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "BlessSelf", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0, &o1)
	return
}

func (c implControllerClientStub) RemoteBlessings(ctx *context.T, i0 string, i1 string, opts ...rpc.CallOpt) (o0 []string, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "RemoteBlessings", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implControllerClientStub) Signature(ctx *context.T, i0 string, opts ...rpc.CallOpt) (o0 []signature.Interface, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Signature", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

// ControllerServerMethods is the interface a server writer
// implements for Controller.
type ControllerServerMethods interface {
	// Serve instructs WSPR to start listening for calls on behalf
	// of a javascript server.
	Serve(call rpc.ServerCall, name string, serverId uint32) error
	// Stop instructs WSPR to stop listening for calls for the
	// given javascript server.
	Stop(call rpc.ServerCall, serverId uint32) error
	// AddName adds a published name to an existing server.
	AddName(call rpc.ServerCall, serverId uint32, name string) error
	// RemoveName removes a published name from an existing server.
	RemoveName(call rpc.ServerCall, serverId uint32, name string) error
	// UnlinkBlessings removes the given blessings from the blessings store.
	UnlinkBlessings(call rpc.ServerCall, handle principal.BlessingsHandle) error
	// Bless binds extensions of blessings held by this principal to
	// another principal (represented by its public key).
	Bless(call rpc.ServerCall, publicKey string, blessingHandle principal.BlessingsHandle, extension string, caveat []security.Caveat) (string, principal.BlessingsHandle, error)
	// BlessSelf creates a blessing with the provided name for this principal.
	BlessSelf(call rpc.ServerCall, name string, caveats []security.Caveat) (string, principal.BlessingsHandle, error)
	// RemoteBlessings fetches the remote blessings for a given name and method.
	RemoteBlessings(call rpc.ServerCall, name string, method string) ([]string, error)
	// Signature fetches the signature for a given name.
	Signature(call rpc.ServerCall, name string) ([]signature.Interface, error)
}

// ControllerServerStubMethods is the server interface containing
// Controller methods, as expected by rpc.Server.
// There is no difference between this interface and ControllerServerMethods
// since there are no streaming methods.
type ControllerServerStubMethods ControllerServerMethods

// ControllerServerStub adds universal methods to ControllerServerStubMethods.
type ControllerServerStub interface {
	ControllerServerStubMethods
	// Describe the Controller interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ControllerServer returns a server stub for Controller.
// It converts an implementation of ControllerServerMethods into
// an object that may be used by rpc.Server.
func ControllerServer(impl ControllerServerMethods) ControllerServerStub {
	stub := implControllerServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implControllerServerStub struct {
	impl ControllerServerMethods
	gs   *rpc.GlobState
}

func (s implControllerServerStub) Serve(call rpc.ServerCall, i0 string, i1 uint32) error {
	return s.impl.Serve(call, i0, i1)
}

func (s implControllerServerStub) Stop(call rpc.ServerCall, i0 uint32) error {
	return s.impl.Stop(call, i0)
}

func (s implControllerServerStub) AddName(call rpc.ServerCall, i0 uint32, i1 string) error {
	return s.impl.AddName(call, i0, i1)
}

func (s implControllerServerStub) RemoveName(call rpc.ServerCall, i0 uint32, i1 string) error {
	return s.impl.RemoveName(call, i0, i1)
}

func (s implControllerServerStub) UnlinkBlessings(call rpc.ServerCall, i0 principal.BlessingsHandle) error {
	return s.impl.UnlinkBlessings(call, i0)
}

func (s implControllerServerStub) Bless(call rpc.ServerCall, i0 string, i1 principal.BlessingsHandle, i2 string, i3 []security.Caveat) (string, principal.BlessingsHandle, error) {
	return s.impl.Bless(call, i0, i1, i2, i3)
}

func (s implControllerServerStub) BlessSelf(call rpc.ServerCall, i0 string, i1 []security.Caveat) (string, principal.BlessingsHandle, error) {
	return s.impl.BlessSelf(call, i0, i1)
}

func (s implControllerServerStub) RemoteBlessings(call rpc.ServerCall, i0 string, i1 string) ([]string, error) {
	return s.impl.RemoteBlessings(call, i0, i1)
}

func (s implControllerServerStub) Signature(call rpc.ServerCall, i0 string) ([]signature.Interface, error) {
	return s.impl.Signature(call, i0)
}

func (s implControllerServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implControllerServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ControllerDesc}
}

// ControllerDesc describes the Controller interface.
var ControllerDesc rpc.InterfaceDesc = descController

// descController hides the desc to keep godoc clean.
var descController = rpc.InterfaceDesc{
	Name:    "Controller",
	PkgPath: "v.io/x/ref/services/wspr/internal/app",
	Methods: []rpc.MethodDesc{
		{
			Name: "Serve",
			Doc:  "// Serve instructs WSPR to start listening for calls on behalf\n// of a javascript server.",
			InArgs: []rpc.ArgDesc{
				{"name", ``},     // string
				{"serverId", ``}, // uint32
			},
		},
		{
			Name: "Stop",
			Doc:  "// Stop instructs WSPR to stop listening for calls for the\n// given javascript server.",
			InArgs: []rpc.ArgDesc{
				{"serverId", ``}, // uint32
			},
		},
		{
			Name: "AddName",
			Doc:  "// AddName adds a published name to an existing server.",
			InArgs: []rpc.ArgDesc{
				{"serverId", ``}, // uint32
				{"name", ``},     // string
			},
		},
		{
			Name: "RemoveName",
			Doc:  "// RemoveName removes a published name from an existing server.",
			InArgs: []rpc.ArgDesc{
				{"serverId", ``}, // uint32
				{"name", ``},     // string
			},
		},
		{
			Name: "UnlinkBlessings",
			Doc:  "// UnlinkBlessings removes the given blessings from the blessings store.",
			InArgs: []rpc.ArgDesc{
				{"handle", ``}, // principal.BlessingsHandle
			},
		},
		{
			Name: "Bless",
			Doc:  "// Bless binds extensions of blessings held by this principal to\n// another principal (represented by its public key).",
			InArgs: []rpc.ArgDesc{
				{"publicKey", ``},      // string
				{"blessingHandle", ``}, // principal.BlessingsHandle
				{"extension", ``},      // string
				{"caveat", ``},         // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
				{"", ``}, // principal.BlessingsHandle
			},
		},
		{
			Name: "BlessSelf",
			Doc:  "// BlessSelf creates a blessing with the provided name for this principal.",
			InArgs: []rpc.ArgDesc{
				{"name", ``},    // string
				{"caveats", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
				{"", ``}, // principal.BlessingsHandle
			},
		},
		{
			Name: "RemoteBlessings",
			Doc:  "// RemoteBlessings fetches the remote blessings for a given name and method.",
			InArgs: []rpc.ArgDesc{
				{"name", ``},   // string
				{"method", ``}, // string
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []string
			},
		},
		{
			Name: "Signature",
			Doc:  "// Signature fetches the signature for a given name.",
			InArgs: []rpc.ArgDesc{
				{"name", ``}, // string
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []signature.Interface
			},
		},
	},
}