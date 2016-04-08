// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: stress

package stress

import (
	"fmt"
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

type SumArg struct {
	ABool        bool
	AInt64       int64
	AListOfBytes []byte
}

func (SumArg) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/runtime/internal/rpc/stress.SumArg"`
}) {
}

func (m *SumArg) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.ABool == false)
	if var4 {
		if err := fieldsTarget1.ZeroField("ABool"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("ABool")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromBool(bool(m.ABool), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.AInt64 == int64(0))
	if var7 {
		if err := fieldsTarget1.ZeroField("AInt64"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("AInt64")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget6.FromInt(int64(m.AInt64), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var var10 bool
	if len(m.AListOfBytes) == 0 {
		var10 = true
	}
	if var10 {
		if err := fieldsTarget1.ZeroField("AListOfBytes"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("AListOfBytes")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := fieldTarget9.FromBytes([]byte(m.AListOfBytes), tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SumArg) MakeVDLTarget() vdl.Target {
	return &SumArgTarget{Value: m}
}

type SumArgTarget struct {
	Value              *SumArg
	aBoolTarget        vdl.BoolTarget
	aInt64Target       vdl.Int64Target
	aListOfBytesTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *SumArgTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*SumArg)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *SumArgTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "ABool":
		t.aBoolTarget.Value = &t.Value.ABool
		target, err := &t.aBoolTarget, error(nil)
		return nil, target, err
	case "AInt64":
		t.aInt64Target.Value = &t.Value.AInt64
		target, err := &t.aInt64Target, error(nil)
		return nil, target, err
	case "AListOfBytes":
		t.aListOfBytesTarget.Value = &t.Value.AListOfBytes
		target, err := &t.aListOfBytesTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/runtime/internal/rpc/stress.SumArg", name)
	}
}
func (t *SumArgTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *SumArgTarget) ZeroField(name string) error {
	switch name {
	case "ABool":
		t.Value.ABool = false
		return nil
	case "AInt64":
		t.Value.AInt64 = int64(0)
		return nil
	case "AListOfBytes":
		t.Value.AListOfBytes = []byte(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/runtime/internal/rpc/stress.SumArg", name)
	}
}
func (t *SumArgTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x *SumArg) VDLRead(dec vdl.Decoder) error {
	*x = SumArg{}
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if dec.Type().Kind() != vdl.Struct {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	match := 0
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			if match == 0 && dec.Type().NumField() > 0 {
				return fmt.Errorf("no matching fields in struct %T, from %v", *x, dec.Type())
			}
			return dec.FinishValue()
		case "ABool":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.ABool, err = dec.DecodeBool(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "AInt64":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.AInt64, err = dec.DecodeInt(64); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "AListOfBytes":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if err = dec.DecodeBytes(-1, &x.AListOfBytes); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		default:
			if err = dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

type SumStats struct {
	SumCount       uint64
	SumStreamCount uint64
	BytesRecv      uint64
	BytesSent      uint64
}

func (SumStats) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/runtime/internal/rpc/stress.SumStats"`
}) {
}

func (m *SumStats) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.SumCount == uint64(0))
	if var4 {
		if err := fieldsTarget1.ZeroField("SumCount"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("SumCount")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromUint(uint64(m.SumCount), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.SumStreamCount == uint64(0))
	if var7 {
		if err := fieldsTarget1.ZeroField("SumStreamCount"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("SumStreamCount")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget6.FromUint(uint64(m.SumStreamCount), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var10 := (m.BytesRecv == uint64(0))
	if var10 {
		if err := fieldsTarget1.ZeroField("BytesRecv"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("BytesRecv")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget9.FromUint(uint64(m.BytesRecv), tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	var13 := (m.BytesSent == uint64(0))
	if var13 {
		if err := fieldsTarget1.ZeroField("BytesSent"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("BytesSent")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget12.FromUint(uint64(m.BytesSent), tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SumStats) MakeVDLTarget() vdl.Target {
	return &SumStatsTarget{Value: m}
}

type SumStatsTarget struct {
	Value                *SumStats
	sumCountTarget       vdl.Uint64Target
	sumStreamCountTarget vdl.Uint64Target
	bytesRecvTarget      vdl.Uint64Target
	bytesSentTarget      vdl.Uint64Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *SumStatsTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*SumStats)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *SumStatsTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "SumCount":
		t.sumCountTarget.Value = &t.Value.SumCount
		target, err := &t.sumCountTarget, error(nil)
		return nil, target, err
	case "SumStreamCount":
		t.sumStreamCountTarget.Value = &t.Value.SumStreamCount
		target, err := &t.sumStreamCountTarget, error(nil)
		return nil, target, err
	case "BytesRecv":
		t.bytesRecvTarget.Value = &t.Value.BytesRecv
		target, err := &t.bytesRecvTarget, error(nil)
		return nil, target, err
	case "BytesSent":
		t.bytesSentTarget.Value = &t.Value.BytesSent
		target, err := &t.bytesSentTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/runtime/internal/rpc/stress.SumStats", name)
	}
}
func (t *SumStatsTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *SumStatsTarget) ZeroField(name string) error {
	switch name {
	case "SumCount":
		t.Value.SumCount = uint64(0)
		return nil
	case "SumStreamCount":
		t.Value.SumStreamCount = uint64(0)
		return nil
	case "BytesRecv":
		t.Value.BytesRecv = uint64(0)
		return nil
	case "BytesSent":
		t.Value.BytesSent = uint64(0)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/runtime/internal/rpc/stress.SumStats", name)
	}
}
func (t *SumStatsTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x *SumStats) VDLRead(dec vdl.Decoder) error {
	*x = SumStats{}
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if dec.Type().Kind() != vdl.Struct {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	match := 0
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			if match == 0 && dec.Type().NumField() > 0 {
				return fmt.Errorf("no matching fields in struct %T, from %v", *x, dec.Type())
			}
			return dec.FinishValue()
		case "SumCount":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.SumCount, err = dec.DecodeUint(64); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "SumStreamCount":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.SumStreamCount, err = dec.DecodeUint(64); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "BytesRecv":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.BytesRecv, err = dec.DecodeUint(64); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "BytesSent":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.BytesSent, err = dec.DecodeUint(64); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		default:
			if err = dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

//////////////////////////////////////////////////
// Interface definitions

// StressClientMethods is the client interface
// containing Stress methods.
type StressClientMethods interface {
	// Echo returns the payload that it receives.
	Echo(_ *context.T, Payload []byte, _ ...rpc.CallOpt) ([]byte, error)
	// Do returns the checksum of the payload that it receives.
	Sum(_ *context.T, arg SumArg, _ ...rpc.CallOpt) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(*context.T, ...rpc.CallOpt) (StressSumStreamClientCall, error)
	// GetSumStats returns the stats on the Sum calls that the server received.
	GetSumStats(*context.T, ...rpc.CallOpt) (SumStats, error)
	// Stop stops the server.
	Stop(*context.T, ...rpc.CallOpt) error
}

// StressClientStub adds universal methods to StressClientMethods.
type StressClientStub interface {
	StressClientMethods
	rpc.UniversalServiceMethods
}

// StressClient returns a client stub for Stress.
func StressClient(name string) StressClientStub {
	return implStressClientStub{name}
}

type implStressClientStub struct {
	name string
}

func (c implStressClientStub) Echo(ctx *context.T, i0 []byte, opts ...rpc.CallOpt) (o0 []byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Echo", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implStressClientStub) Sum(ctx *context.T, i0 SumArg, opts ...rpc.CallOpt) (o0 []byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Sum", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implStressClientStub) SumStream(ctx *context.T, opts ...rpc.CallOpt) (ocall StressSumStreamClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "SumStream", nil, opts...); err != nil {
		return
	}
	ocall = &implStressSumStreamClientCall{ClientCall: call}
	return
}

func (c implStressClientStub) GetSumStats(ctx *context.T, opts ...rpc.CallOpt) (o0 SumStats, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "GetSumStats", nil, []interface{}{&o0}, opts...)
	return
}

func (c implStressClientStub) Stop(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Stop", nil, nil, opts...)
	return
}

// StressSumStreamClientStream is the client stream for Stress.SumStream.
type StressSumStreamClientStream interface {
	// RecvStream returns the receiver side of the Stress.SumStream client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Stress.SumStream client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item SumArg) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// StressSumStreamClientCall represents the call returned from Stress.SumStream.
type StressSumStreamClientCall interface {
	StressSumStreamClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implStressSumStreamClientCall struct {
	rpc.ClientCall
	valRecv []byte
	errRecv error
}

func (c *implStressSumStreamClientCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implStressSumStreamClientCallRecv{c}
}

type implStressSumStreamClientCallRecv struct {
	c *implStressSumStreamClientCall
}

func (c implStressSumStreamClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implStressSumStreamClientCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implStressSumStreamClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implStressSumStreamClientCall) SendStream() interface {
	Send(item SumArg) error
	Close() error
} {
	return implStressSumStreamClientCallSend{c}
}

type implStressSumStreamClientCallSend struct {
	c *implStressSumStreamClientCall
}

func (c implStressSumStreamClientCallSend) Send(item SumArg) error {
	return c.c.Send(item)
}
func (c implStressSumStreamClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implStressSumStreamClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// StressServerMethods is the interface a server writer
// implements for Stress.
type StressServerMethods interface {
	// Echo returns the payload that it receives.
	Echo(_ *context.T, _ rpc.ServerCall, Payload []byte) ([]byte, error)
	// Do returns the checksum of the payload that it receives.
	Sum(_ *context.T, _ rpc.ServerCall, arg SumArg) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(*context.T, StressSumStreamServerCall) error
	// GetSumStats returns the stats on the Sum calls that the server received.
	GetSumStats(*context.T, rpc.ServerCall) (SumStats, error)
	// Stop stops the server.
	Stop(*context.T, rpc.ServerCall) error
}

// StressServerStubMethods is the server interface containing
// Stress methods, as expected by rpc.Server.
// The only difference between this interface and StressServerMethods
// is the streaming methods.
type StressServerStubMethods interface {
	// Echo returns the payload that it receives.
	Echo(_ *context.T, _ rpc.ServerCall, Payload []byte) ([]byte, error)
	// Do returns the checksum of the payload that it receives.
	Sum(_ *context.T, _ rpc.ServerCall, arg SumArg) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(*context.T, *StressSumStreamServerCallStub) error
	// GetSumStats returns the stats on the Sum calls that the server received.
	GetSumStats(*context.T, rpc.ServerCall) (SumStats, error)
	// Stop stops the server.
	Stop(*context.T, rpc.ServerCall) error
}

// StressServerStub adds universal methods to StressServerStubMethods.
type StressServerStub interface {
	StressServerStubMethods
	// Describe the Stress interfaces.
	Describe__() []rpc.InterfaceDesc
}

// StressServer returns a server stub for Stress.
// It converts an implementation of StressServerMethods into
// an object that may be used by rpc.Server.
func StressServer(impl StressServerMethods) StressServerStub {
	stub := implStressServerStub{
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

type implStressServerStub struct {
	impl StressServerMethods
	gs   *rpc.GlobState
}

func (s implStressServerStub) Echo(ctx *context.T, call rpc.ServerCall, i0 []byte) ([]byte, error) {
	return s.impl.Echo(ctx, call, i0)
}

func (s implStressServerStub) Sum(ctx *context.T, call rpc.ServerCall, i0 SumArg) ([]byte, error) {
	return s.impl.Sum(ctx, call, i0)
}

func (s implStressServerStub) SumStream(ctx *context.T, call *StressSumStreamServerCallStub) error {
	return s.impl.SumStream(ctx, call)
}

func (s implStressServerStub) GetSumStats(ctx *context.T, call rpc.ServerCall) (SumStats, error) {
	return s.impl.GetSumStats(ctx, call)
}

func (s implStressServerStub) Stop(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.Stop(ctx, call)
}

func (s implStressServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implStressServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{StressDesc}
}

// StressDesc describes the Stress interface.
var StressDesc rpc.InterfaceDesc = descStress

// descStress hides the desc to keep godoc clean.
var descStress = rpc.InterfaceDesc{
	Name:    "Stress",
	PkgPath: "v.io/x/ref/runtime/internal/rpc/stress",
	Methods: []rpc.MethodDesc{
		{
			Name: "Echo",
			Doc:  "// Echo returns the payload that it receives.",
			InArgs: []rpc.ArgDesc{
				{"Payload", ``}, // []byte
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "Sum",
			Doc:  "// Do returns the checksum of the payload that it receives.",
			InArgs: []rpc.ArgDesc{
				{"arg", ``}, // SumArg
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "SumStream",
			Doc:  "// DoStream returns the checksum of the payload that it receives via the stream.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "GetSumStats",
			Doc:  "// GetSumStats returns the stats on the Sum calls that the server received.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // SumStats
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "Stop",
			Doc:  "// Stop stops the server.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Admin"))},
		},
	},
}

// StressSumStreamServerStream is the server stream for Stress.SumStream.
type StressSumStreamServerStream interface {
	// RecvStream returns the receiver side of the Stress.SumStream server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() SumArg
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Stress.SumStream server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// StressSumStreamServerCall represents the context passed to Stress.SumStream.
type StressSumStreamServerCall interface {
	rpc.ServerCall
	StressSumStreamServerStream
}

// StressSumStreamServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements StressSumStreamServerCall.
type StressSumStreamServerCallStub struct {
	rpc.StreamServerCall
	valRecv SumArg
	errRecv error
}

// Init initializes StressSumStreamServerCallStub from rpc.StreamServerCall.
func (s *StressSumStreamServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Stress.SumStream server stream.
func (s *StressSumStreamServerCallStub) RecvStream() interface {
	Advance() bool
	Value() SumArg
	Err() error
} {
	return implStressSumStreamServerCallRecv{s}
}

type implStressSumStreamServerCallRecv struct {
	s *StressSumStreamServerCallStub
}

func (s implStressSumStreamServerCallRecv) Advance() bool {
	s.s.valRecv = SumArg{}
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implStressSumStreamServerCallRecv) Value() SumArg {
	return s.s.valRecv
}
func (s implStressSumStreamServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Stress.SumStream server stream.
func (s *StressSumStreamServerCallStub) SendStream() interface {
	Send(item []byte) error
} {
	return implStressSumStreamServerCallSend{s}
}

type implStressSumStreamServerCallSend struct {
	s *StressSumStreamServerCallStub
}

func (s implStressSumStreamServerCallSend) Send(item []byte) error {
	return s.s.Send(item)
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*SumArg)(nil))
	vdl.Register((*SumStats)(nil))

	return struct{}{}
}
