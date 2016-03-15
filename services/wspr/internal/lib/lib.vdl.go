// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: lib

package lib

import (
	"fmt"
	"reflect"
	"v.io/v23/vdl"
	"v.io/v23/verror"
	"v.io/v23/vom"
	"v.io/v23/vtrace"
)

// The response from the javascript server to the proxy.
type ServerRpcReply struct {
	Results       []*vom.RawBytes
	Err           error
	TraceResponse vtrace.Response
}

func (ServerRpcReply) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/lib.ServerRpcReply"`
}) {
}

func (m *ServerRpcReply) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	__VDLEnsureNativeBuilt()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Results")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget4, err := fieldTarget3.StartList(__VDLType1, len(m.Results))
		if err != nil {
			return err
		}
		for i, elem6 := range m.Results {
			elemTarget5, err := listTarget4.StartElem(i)
			if err != nil {
				return err
			}

			if elem6 == nil {
				if err := elemTarget5.FromNil(vdl.AnyType); err != nil {
					return err
				}
			} else {
				if err := elem6.FillVDLTarget(elemTarget5, vdl.AnyType); err != nil {
					return err
				}
			}
			if err := listTarget4.FinishElem(elemTarget5); err != nil {
				return err
			}
		}
		if err := fieldTarget3.FinishList(listTarget4); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("Err")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if m.Err == nil {
			if err := fieldTarget8.FromNil(vdl.ErrorType); err != nil {
				return err
			}
		} else {
			var wireError9 vdl.WireError
			if err := verror.WireFromNative(&wireError9, m.Err); err != nil {
				return err
			}
			if err := wireError9.FillVDLTarget(fieldTarget8, vdl.ErrorType); err != nil {
				return err
			}

		}
		if err := fieldsTarget1.FinishField(keyTarget7, fieldTarget8); err != nil {
			return err
		}
	}
	keyTarget10, fieldTarget11, err := fieldsTarget1.StartField("TraceResponse")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.TraceResponse.FillVDLTarget(fieldTarget11, __VDLType_v_io_v23_vtrace_Response); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget10, fieldTarget11); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *ServerRpcReply) MakeVDLTarget() vdl.Target {
	return &ServerRpcReplyTarget{Value: m}
}

type ServerRpcReplyTarget struct {
	Value               *ServerRpcReply
	resultsTarget       unnamed_5b5d616e79Target
	errTarget           verror.ErrorTarget
	traceResponseTarget vtrace.ResponseTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *ServerRpcReplyTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply)
	}
	return t, nil
}
func (t *ServerRpcReplyTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Results":
		t.resultsTarget.Value = &t.Value.Results
		target, err := &t.resultsTarget, error(nil)
		return nil, target, err
	case "Err":
		t.errTarget.Value = &t.Value.Err
		target, err := &t.errTarget, error(nil)
		return nil, target, err
	case "TraceResponse":
		t.traceResponseTarget.Value = &t.Value.TraceResponse
		target, err := &t.traceResponseTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply)
	}
}
func (t *ServerRpcReplyTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ServerRpcReplyTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// []*vom.RawBytes
type unnamed_5b5d616e79Target struct {
	Value *[]*vom.RawBytes

	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *unnamed_5b5d616e79Target) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if !vdl.Compatible(tt, __VDLType1) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType1)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]*vom.RawBytes, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *unnamed_5b5d616e79Target) StartElem(index int) (elem vdl.Target, _ error) {
	target, err := vdl.ReflectTarget(reflect.ValueOf(&(*t.Value)[index]))
	return target, err
}
func (t *unnamed_5b5d616e79Target) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *unnamed_5b5d616e79Target) FinishList(elem vdl.ListTarget) error {

	return nil
}

type LogLevel int

const (
	LogLevelInfo LogLevel = iota
	LogLevelError
)

// LogLevelAll holds all labels for LogLevel.
var LogLevelAll = [...]LogLevel{LogLevelInfo, LogLevelError}

// LogLevelFromString creates a LogLevel from a string label.
func LogLevelFromString(label string) (x LogLevel, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *LogLevel) Set(label string) error {
	switch label {
	case "Info", "info":
		*x = LogLevelInfo
		return nil
	case "Error", "error":
		*x = LogLevelError
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in lib.LogLevel", label)
}

// String returns the string label of x.
func (x LogLevel) String() string {
	switch x {
	case LogLevelInfo:
		return "Info"
	case LogLevelError:
		return "Error"
	}
	return ""
}

func (LogLevel) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/lib.LogLevel"`
	Enum struct{ Info, Error string }
}) {
}

func (m *LogLevel) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromEnumLabel((*m).String(), __VDLType_v_io_x_ref_services_wspr_internal_lib_LogLevel); err != nil {
		return err
	}
	return nil
}

func (m *LogLevel) MakeVDLTarget() vdl.Target {
	return &LogLevelTarget{Value: m}
}

type LogLevelTarget struct {
	Value *LogLevel
	vdl.TargetBase
}

func (t *LogLevelTarget) FromEnumLabel(src string, tt *vdl.Type) error {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_wspr_internal_lib_LogLevel) {
		return fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_wspr_internal_lib_LogLevel)
	}
	switch src {
	case "Info":
		*t.Value = 0
	case "Error":
		*t.Value = 1
	default:
		return fmt.Errorf("label %s not in enum %v", src, __VDLType_v_io_x_ref_services_wspr_internal_lib_LogLevel)
	}

	return nil
}

type LogMessage struct {
	Level   LogLevel
	Message string
}

func (LogMessage) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/lib.LogMessage"`
}) {
}

func (m *LogMessage) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_wspr_internal_lib_LogMessage == nil || __VDLType2 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Level")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Level.FillVDLTarget(fieldTarget3, __VDLType_v_io_x_ref_services_wspr_internal_lib_LogLevel); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Message")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.Message), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *LogMessage) MakeVDLTarget() vdl.Target {
	return &LogMessageTarget{Value: m}
}

type LogMessageTarget struct {
	Value         *LogMessage
	levelTarget   LogLevelTarget
	messageTarget vdl.StringTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *LogMessageTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_wspr_internal_lib_LogMessage) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_wspr_internal_lib_LogMessage)
	}
	return t, nil
}
func (t *LogMessageTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Level":
		t.levelTarget.Value = &t.Value.Level
		target, err := &t.levelTarget, error(nil)
		return nil, target, err
	case "Message":
		t.messageTarget.Value = &t.Value.Message
		target, err := &t.messageTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_wspr_internal_lib_LogMessage)
	}
}
func (t *LogMessageTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *LogMessageTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func init() {
	vdl.Register((*ServerRpcReply)(nil))
	vdl.Register((*LogLevel)(nil))
	vdl.Register((*LogMessage)(nil))
}

var __VDLType2 *vdl.Type = vdl.TypeOf((*LogMessage)(nil))
var __VDLType0 *vdl.Type

func __VDLType0_gen() *vdl.Type {
	__VDLType0Builder := vdl.TypeBuilder{}

	__VDLType01 := __VDLType0Builder.Optional()
	__VDLType02 := __VDLType0Builder.Struct()
	__VDLType03 := __VDLType0Builder.Named("v.io/x/ref/services/wspr/internal/lib.ServerRpcReply").AssignBase(__VDLType02)
	__VDLType04 := __VDLType0Builder.List()
	__VDLType05 := vdl.AnyType
	__VDLType04.AssignElem(__VDLType05)
	__VDLType02.AppendField("Results", __VDLType04)
	__VDLType06 := __VDLType0Builder.Optional()
	__VDLType07 := __VDLType0Builder.Struct()
	__VDLType08 := __VDLType0Builder.Named("error").AssignBase(__VDLType07)
	__VDLType09 := vdl.StringType
	__VDLType07.AppendField("Id", __VDLType09)
	__VDLType010 := __VDLType0Builder.Enum()
	__VDLType010.AppendLabel("NoRetry")
	__VDLType010.AppendLabel("RetryConnection")
	__VDLType010.AppendLabel("RetryRefetch")
	__VDLType010.AppendLabel("RetryBackoff")
	__VDLType07.AppendField("RetryCode", __VDLType010)
	__VDLType07.AppendField("Msg", __VDLType09)
	__VDLType07.AppendField("ParamList", __VDLType04)
	__VDLType06.AssignElem(__VDLType08)
	__VDLType02.AppendField("Err", __VDLType06)
	__VDLType011 := __VDLType0Builder.Struct()
	__VDLType012 := __VDLType0Builder.Named("v.io/v23/vtrace.Response").AssignBase(__VDLType011)
	__VDLType013 := vdl.Int32Type
	__VDLType014 := __VDLType0Builder.Named("v.io/v23/vtrace.TraceFlags").AssignBase(__VDLType013)
	__VDLType011.AppendField("Flags", __VDLType014)
	__VDLType015 := __VDLType0Builder.Struct()
	__VDLType016 := __VDLType0Builder.Named("v.io/v23/vtrace.TraceRecord").AssignBase(__VDLType015)
	__VDLType017 := __VDLType0Builder.Array()
	__VDLType018 := __VDLType0Builder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLType017)
	__VDLType019 := vdl.ByteType
	__VDLType017.AssignElem(__VDLType019)
	__VDLType017.AssignLen(16)
	__VDLType015.AppendField("Id", __VDLType018)
	__VDLType020 := __VDLType0Builder.List()
	__VDLType021 := __VDLType0Builder.Struct()
	__VDLType022 := __VDLType0Builder.Named("v.io/v23/vtrace.SpanRecord").AssignBase(__VDLType021)
	__VDLType021.AppendField("Id", __VDLType018)
	__VDLType021.AppendField("Parent", __VDLType018)
	__VDLType021.AppendField("Name", __VDLType09)
	__VDLType023 := __VDLType0Builder.Struct()
	__VDLType024 := __VDLType0Builder.Named("time.Time").AssignBase(__VDLType023)
	__VDLType025 := vdl.Int64Type
	__VDLType023.AppendField("Seconds", __VDLType025)
	__VDLType026 := vdl.Int32Type
	__VDLType023.AppendField("Nanos", __VDLType026)
	__VDLType021.AppendField("Start", __VDLType024)
	__VDLType021.AppendField("End", __VDLType024)
	__VDLType027 := __VDLType0Builder.List()
	__VDLType028 := __VDLType0Builder.Struct()
	__VDLType029 := __VDLType0Builder.Named("v.io/v23/vtrace.Annotation").AssignBase(__VDLType028)
	__VDLType028.AppendField("When", __VDLType024)
	__VDLType028.AppendField("Message", __VDLType09)
	__VDLType027.AssignElem(__VDLType029)
	__VDLType021.AppendField("Annotations", __VDLType027)
	__VDLType020.AssignElem(__VDLType022)
	__VDLType015.AppendField("Spans", __VDLType020)
	__VDLType011.AppendField("Trace", __VDLType016)
	__VDLType02.AppendField("TraceResponse", __VDLType012)
	__VDLType01.AssignElem(__VDLType03)
	__VDLType0Builder.Build()
	__VDLType0v, err := __VDLType01.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType0v
}
func init() {
	__VDLType0 = __VDLType0_gen()
}

var __VDLType1 *vdl.Type = vdl.TypeOf([]*vom.RawBytes(nil))
var __VDLType_v_io_v23_vtrace_Response *vdl.Type

func __VDLType_v_io_v23_vtrace_Response_gen() *vdl.Type {
	__VDLType_v_io_v23_vtrace_ResponseBuilder := vdl.TypeBuilder{}

	__VDLType_v_io_v23_vtrace_Response1 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Struct()
	__VDLType_v_io_v23_vtrace_Response2 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Named("v.io/v23/vtrace.Response").AssignBase(__VDLType_v_io_v23_vtrace_Response1)
	__VDLType_v_io_v23_vtrace_Response3 := vdl.Int32Type
	__VDLType_v_io_v23_vtrace_Response4 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Named("v.io/v23/vtrace.TraceFlags").AssignBase(__VDLType_v_io_v23_vtrace_Response3)
	__VDLType_v_io_v23_vtrace_Response1.AppendField("Flags", __VDLType_v_io_v23_vtrace_Response4)
	__VDLType_v_io_v23_vtrace_Response5 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Struct()
	__VDLType_v_io_v23_vtrace_Response6 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Named("v.io/v23/vtrace.TraceRecord").AssignBase(__VDLType_v_io_v23_vtrace_Response5)
	__VDLType_v_io_v23_vtrace_Response7 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Array()
	__VDLType_v_io_v23_vtrace_Response8 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLType_v_io_v23_vtrace_Response7)
	__VDLType_v_io_v23_vtrace_Response9 := vdl.ByteType
	__VDLType_v_io_v23_vtrace_Response7.AssignElem(__VDLType_v_io_v23_vtrace_Response9)
	__VDLType_v_io_v23_vtrace_Response7.AssignLen(16)
	__VDLType_v_io_v23_vtrace_Response5.AppendField("Id", __VDLType_v_io_v23_vtrace_Response8)
	__VDLType_v_io_v23_vtrace_Response10 := __VDLType_v_io_v23_vtrace_ResponseBuilder.List()
	__VDLType_v_io_v23_vtrace_Response11 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Struct()
	__VDLType_v_io_v23_vtrace_Response12 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Named("v.io/v23/vtrace.SpanRecord").AssignBase(__VDLType_v_io_v23_vtrace_Response11)
	__VDLType_v_io_v23_vtrace_Response11.AppendField("Id", __VDLType_v_io_v23_vtrace_Response8)
	__VDLType_v_io_v23_vtrace_Response11.AppendField("Parent", __VDLType_v_io_v23_vtrace_Response8)
	__VDLType_v_io_v23_vtrace_Response13 := vdl.StringType
	__VDLType_v_io_v23_vtrace_Response11.AppendField("Name", __VDLType_v_io_v23_vtrace_Response13)
	__VDLType_v_io_v23_vtrace_Response14 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Struct()
	__VDLType_v_io_v23_vtrace_Response15 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Named("time.Time").AssignBase(__VDLType_v_io_v23_vtrace_Response14)
	__VDLType_v_io_v23_vtrace_Response16 := vdl.Int64Type
	__VDLType_v_io_v23_vtrace_Response14.AppendField("Seconds", __VDLType_v_io_v23_vtrace_Response16)
	__VDLType_v_io_v23_vtrace_Response17 := vdl.Int32Type
	__VDLType_v_io_v23_vtrace_Response14.AppendField("Nanos", __VDLType_v_io_v23_vtrace_Response17)
	__VDLType_v_io_v23_vtrace_Response11.AppendField("Start", __VDLType_v_io_v23_vtrace_Response15)
	__VDLType_v_io_v23_vtrace_Response11.AppendField("End", __VDLType_v_io_v23_vtrace_Response15)
	__VDLType_v_io_v23_vtrace_Response18 := __VDLType_v_io_v23_vtrace_ResponseBuilder.List()
	__VDLType_v_io_v23_vtrace_Response19 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Struct()
	__VDLType_v_io_v23_vtrace_Response20 := __VDLType_v_io_v23_vtrace_ResponseBuilder.Named("v.io/v23/vtrace.Annotation").AssignBase(__VDLType_v_io_v23_vtrace_Response19)
	__VDLType_v_io_v23_vtrace_Response19.AppendField("When", __VDLType_v_io_v23_vtrace_Response15)
	__VDLType_v_io_v23_vtrace_Response19.AppendField("Message", __VDLType_v_io_v23_vtrace_Response13)
	__VDLType_v_io_v23_vtrace_Response18.AssignElem(__VDLType_v_io_v23_vtrace_Response20)
	__VDLType_v_io_v23_vtrace_Response11.AppendField("Annotations", __VDLType_v_io_v23_vtrace_Response18)
	__VDLType_v_io_v23_vtrace_Response10.AssignElem(__VDLType_v_io_v23_vtrace_Response12)
	__VDLType_v_io_v23_vtrace_Response5.AppendField("Spans", __VDLType_v_io_v23_vtrace_Response10)
	__VDLType_v_io_v23_vtrace_Response1.AppendField("Trace", __VDLType_v_io_v23_vtrace_Response6)
	__VDLType_v_io_v23_vtrace_ResponseBuilder.Build()
	__VDLType_v_io_v23_vtrace_Responsev, err := __VDLType_v_io_v23_vtrace_Response2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_v_io_v23_vtrace_Responsev
}
func init() {
	__VDLType_v_io_v23_vtrace_Response = __VDLType_v_io_v23_vtrace_Response_gen()
}

var __VDLType_v_io_x_ref_services_wspr_internal_lib_LogLevel *vdl.Type = vdl.TypeOf(LogLevelInfo)
var __VDLType_v_io_x_ref_services_wspr_internal_lib_LogMessage *vdl.Type = vdl.TypeOf(LogMessage{})
var __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply *vdl.Type

func __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply_gen() *vdl.Type {
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder := vdl.TypeBuilder{}

	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply1 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Struct()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply2 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("v.io/x/ref/services/wspr/internal/lib.ServerRpcReply").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply1)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply3 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.List()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply4 := vdl.AnyType
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply3.AssignElem(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply4)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply1.AppendField("Results", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply3)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply5 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Optional()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply6 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Struct()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply7 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("error").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply6)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply8 := vdl.StringType
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply6.AppendField("Id", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply8)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply9 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Enum()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply9.AppendLabel("NoRetry")
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply9.AppendLabel("RetryConnection")
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply9.AppendLabel("RetryRefetch")
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply9.AppendLabel("RetryBackoff")
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply6.AppendField("RetryCode", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply9)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply6.AppendField("Msg", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply8)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply6.AppendField("ParamList", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply3)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply5.AssignElem(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply7)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply1.AppendField("Err", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply5)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply10 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Struct()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply11 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("v.io/v23/vtrace.Response").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply10)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply12 := vdl.Int32Type
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply13 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("v.io/v23/vtrace.TraceFlags").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply12)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply10.AppendField("Flags", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply13)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply14 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Struct()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply15 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("v.io/v23/vtrace.TraceRecord").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply14)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply16 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Array()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply17 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply16)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply18 := vdl.ByteType
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply16.AssignElem(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply18)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply16.AssignLen(16)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply14.AppendField("Id", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply17)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply19 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.List()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply20 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Struct()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply21 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("v.io/v23/vtrace.SpanRecord").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply20)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply20.AppendField("Id", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply17)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply20.AppendField("Parent", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply17)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply20.AppendField("Name", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply8)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply22 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Struct()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply23 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("time.Time").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply22)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply24 := vdl.Int64Type
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply22.AppendField("Seconds", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply24)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply25 := vdl.Int32Type
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply22.AppendField("Nanos", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply25)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply20.AppendField("Start", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply23)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply20.AppendField("End", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply23)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply26 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.List()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply27 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Struct()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply28 := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Named("v.io/v23/vtrace.Annotation").AssignBase(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply27)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply27.AppendField("When", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply23)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply27.AppendField("Message", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply8)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply26.AssignElem(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply28)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply20.AppendField("Annotations", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply26)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply19.AssignElem(__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply21)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply14.AppendField("Spans", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply19)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply10.AppendField("Trace", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply15)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply1.AppendField("TraceResponse", __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply11)
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyBuilder.Build()
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyv, err := __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReplyv
}
func init() {
	__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply = __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply_gen()
}
func __VDLEnsureNativeBuilt() {
	if __VDLType0 == nil {
		__VDLType0 = __VDLType0_gen()
	}
	if __VDLType_v_io_v23_vtrace_Response == nil {
		__VDLType_v_io_v23_vtrace_Response = __VDLType_v_io_v23_vtrace_Response_gen()
	}
	if __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply == nil {
		__VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply = __VDLType_v_io_x_ref_services_wspr_internal_lib_ServerRpcReply_gen()
	}
}