// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: server.vdl

package server

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/vdl"
	"v.io/v23/verror"

	// VDL user imports
	"v.io/v23/security"
	"v.io/v23/vdlroot/time"
	"v.io/v23/vtrace"
	"v.io/x/ref/services/wspr/internal/principal"
)

type Context struct {
	Language string
}

func (Context) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/rpc/server.Context"`
}) {
}

type SecurityCall struct {
	Method                string
	Suffix                string
	MethodTags            []*vdl.Value
	LocalBlessings        principal.JsBlessings
	LocalBlessingStrings  []string
	RemoteBlessings       principal.JsBlessings
	RemoteBlessingStrings []string
	LocalEndpoint         string
	RemoteEndpoint        string
}

func (SecurityCall) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/rpc/server.SecurityCall"`
}) {
}

type CaveatValidationRequest struct {
	Call    SecurityCall
	Context Context
	Cavs    [][]security.Caveat
}

func (CaveatValidationRequest) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/rpc/server.CaveatValidationRequest"`
}) {
}

type CaveatValidationResponse struct {
	Results []error
}

func (CaveatValidationResponse) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/rpc/server.CaveatValidationResponse"`
}) {
}

type ServerRpcRequestCall struct {
	SecurityCall     SecurityCall
	Deadline         time.Deadline
	Context          Context
	TraceRequest     vtrace.Request
	GrantedBlessings *principal.JsBlessings
}

func (ServerRpcRequestCall) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/rpc/server.ServerRpcRequestCall"`
}) {
}

// A request from the proxy to javascript to handle an RPC
type ServerRpcRequest struct {
	ServerId uint32
	Handle   int32
	Method   string
	Args     []*vdl.Value
	Call     ServerRpcRequestCall
}

func (ServerRpcRequest) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/rpc/server.ServerRpcRequest"`
}) {
}

func init() {
	vdl.Register((*Context)(nil))
	vdl.Register((*SecurityCall)(nil))
	vdl.Register((*CaveatValidationRequest)(nil))
	vdl.Register((*CaveatValidationResponse)(nil))
	vdl.Register((*ServerRpcRequestCall)(nil))
	vdl.Register((*ServerRpcRequest)(nil))
}

var (
	ErrCaveatValidationTimeout                 = verror.Register("v.io/x/ref/services/wspr/internal/rpc/server.CaveatValidationTimeout", verror.NoRetry, "{1:}{2:} Caveat validation has timed out")
	ErrInvalidValidationResponseFromJavascript = verror.Register("v.io/x/ref/services/wspr/internal/rpc/server.InvalidValidationResponseFromJavascript", verror.NoRetry, "{1:}{2:} Invalid validation response from javascript")
	ErrServerStopped                           = verror.Register("v.io/x/ref/services/wspr/internal/rpc/server.ServerStopped", verror.RetryBackoff, "{1:}{2:} Server has been stopped")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCaveatValidationTimeout.ID), "{1:}{2:} Caveat validation has timed out")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidValidationResponseFromJavascript.ID), "{1:}{2:} Invalid validation response from javascript")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrServerStopped.ID), "{1:}{2:} Server has been stopped")
}

// NewErrCaveatValidationTimeout returns an error with the ErrCaveatValidationTimeout ID.
func NewErrCaveatValidationTimeout(ctx *context.T) error {
	return verror.New(ErrCaveatValidationTimeout, ctx)
}

// NewErrInvalidValidationResponseFromJavascript returns an error with the ErrInvalidValidationResponseFromJavascript ID.
func NewErrInvalidValidationResponseFromJavascript(ctx *context.T) error {
	return verror.New(ErrInvalidValidationResponseFromJavascript, ctx)
}

// NewErrServerStopped returns an error with the ErrServerStopped ID.
func NewErrServerStopped(ctx *context.T) error {
	return verror.New(ErrServerStopped, ctx)
}
