// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: server

package server

import (
	"fmt"
	"v.io/v23/security/access"
	"v.io/v23/services/syncbase"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// ServiceData represents the persistent state of a Service.
type ServiceData struct {
	Version uint64 // covers the fields below
	Perms   access.Permissions
}

func (ServiceData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/server.ServiceData"`
}) {
}

func (m *ServiceData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Version == uint64(0))
	if var4 {
		if err := fieldsTarget1.ZeroField("Version"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Version")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromUint(uint64(m.Version), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Perms) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Perms"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Perms")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Perms.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *ServiceData) MakeVDLTarget() vdl.Target {
	return &ServiceDataTarget{Value: m}
}

type ServiceDataTarget struct {
	Value         *ServiceData
	versionTarget vdl.Uint64Target
	permsTarget   access.PermissionsTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *ServiceDataTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*ServiceData)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *ServiceDataTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Version":
		t.versionTarget.Value = &t.Value.Version
		target, err := &t.versionTarget, error(nil)
		return nil, target, err
	case "Perms":
		t.permsTarget.Value = &t.Value.Perms
		target, err := &t.permsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/server.ServiceData", name)
	}
}
func (t *ServiceDataTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ServiceDataTarget) ZeroField(name string) error {
	switch name {
	case "Version":
		t.Value.Version = uint64(0)
		return nil
	case "Perms":
		t.Value.Perms = access.Permissions(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/server.ServiceData", name)
	}
}
func (t *ServiceDataTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x *ServiceData) VDLRead(dec vdl.Decoder) error {
	*x = ServiceData{}
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
		case "Version":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.Version, err = dec.DecodeUint(64); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "Perms":
			match++
			if err = x.Perms.VDLRead(dec); err != nil {
				return err
			}
		default:
			if err = dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

// DbInfo contains information about a single Database, stored in the
// service-level storage engine.
type DbInfo struct {
	Id syncbase.Id
	// Select fields from DatabaseOptions, needed in order to open storage engine
	// on restart.
	RootDir string // interpreted by storage engine
	Engine  string // name of storage engine, e.g. "leveldb"
}

func (DbInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/server.DbInfo"`
}) {
}

func (m *DbInfo) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Id == syncbase.Id{})
	if var4 {
		if err := fieldsTarget1.ZeroField("Id"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Id.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.RootDir == "")
	if var7 {
		if err := fieldsTarget1.ZeroField("RootDir"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("RootDir")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget6.FromString(string(m.RootDir), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var10 := (m.Engine == "")
	if var10 {
		if err := fieldsTarget1.ZeroField("Engine"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Engine")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget9.FromString(string(m.Engine), tt.NonOptional().Field(2).Type); err != nil {
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

func (m *DbInfo) MakeVDLTarget() vdl.Target {
	return &DbInfoTarget{Value: m}
}

type DbInfoTarget struct {
	Value         *DbInfo
	idTarget      syncbase.IdTarget
	rootDirTarget vdl.StringTarget
	engineTarget  vdl.StringTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DbInfoTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*DbInfo)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *DbInfoTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		t.idTarget.Value = &t.Value.Id
		target, err := &t.idTarget, error(nil)
		return nil, target, err
	case "RootDir":
		t.rootDirTarget.Value = &t.Value.RootDir
		target, err := &t.rootDirTarget, error(nil)
		return nil, target, err
	case "Engine":
		t.engineTarget.Value = &t.Value.Engine
		target, err := &t.engineTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/server.DbInfo", name)
	}
}
func (t *DbInfoTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DbInfoTarget) ZeroField(name string) error {
	switch name {
	case "Id":
		t.Value.Id = syncbase.Id{}
		return nil
	case "RootDir":
		t.Value.RootDir = ""
		return nil
	case "Engine":
		t.Value.Engine = ""
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/server.DbInfo", name)
	}
}
func (t *DbInfoTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x *DbInfo) VDLRead(dec vdl.Decoder) error {
	*x = DbInfo{}
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
		case "Id":
			match++
			if err = x.Id.VDLRead(dec); err != nil {
				return err
			}
		case "RootDir":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.RootDir, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "Engine":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.Engine, err = dec.DecodeString(); err != nil {
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

// DatabaseData represents the persistent state of a Database, stored in the
// per-database storage engine.
type DatabaseData struct {
	Id             syncbase.Id
	Version        uint64 // covers the Perms field below
	Perms          access.Permissions
	SchemaMetadata *syncbase.SchemaMetadata
}

func (DatabaseData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/server.DatabaseData"`
}) {
}

func (m *DatabaseData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Id == syncbase.Id{})
	if var4 {
		if err := fieldsTarget1.ZeroField("Id"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Id.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.Version == uint64(0))
	if var7 {
		if err := fieldsTarget1.ZeroField("Version"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Version")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget6.FromUint(uint64(m.Version), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var var10 bool
	if len(m.Perms) == 0 {
		var10 = true
	}
	if var10 {
		if err := fieldsTarget1.ZeroField("Perms"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Perms")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Perms.FillVDLTarget(fieldTarget9, tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	var13 := (m.SchemaMetadata == (*syncbase.SchemaMetadata)(nil))
	if var13 {
		if err := fieldsTarget1.ZeroField("SchemaMetadata"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("SchemaMetadata")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.SchemaMetadata.FillVDLTarget(fieldTarget12, tt.NonOptional().Field(3).Type); err != nil {
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

func (m *DatabaseData) MakeVDLTarget() vdl.Target {
	return &DatabaseDataTarget{Value: m}
}

type DatabaseDataTarget struct {
	Value                *DatabaseData
	idTarget             syncbase.IdTarget
	versionTarget        vdl.Uint64Target
	permsTarget          access.PermissionsTarget
	schemaMetadataTarget __VDLTarget1_optional
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DatabaseDataTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*DatabaseData)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *DatabaseDataTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		t.idTarget.Value = &t.Value.Id
		target, err := &t.idTarget, error(nil)
		return nil, target, err
	case "Version":
		t.versionTarget.Value = &t.Value.Version
		target, err := &t.versionTarget, error(nil)
		return nil, target, err
	case "Perms":
		t.permsTarget.Value = &t.Value.Perms
		target, err := &t.permsTarget, error(nil)
		return nil, target, err
	case "SchemaMetadata":
		t.schemaMetadataTarget.Value = &t.Value.SchemaMetadata
		target, err := &t.schemaMetadataTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/server.DatabaseData", name)
	}
}
func (t *DatabaseDataTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DatabaseDataTarget) ZeroField(name string) error {
	switch name {
	case "Id":
		t.Value.Id = syncbase.Id{}
		return nil
	case "Version":
		t.Value.Version = uint64(0)
		return nil
	case "Perms":
		t.Value.Perms = access.Permissions(nil)
		return nil
	case "SchemaMetadata":
		t.Value.SchemaMetadata = (*syncbase.SchemaMetadata)(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/server.DatabaseData", name)
	}
}
func (t *DatabaseDataTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// Optional syncbase.SchemaMetadata
type __VDLTarget1_optional struct {
	Value      **syncbase.SchemaMetadata
	elemTarget syncbase.SchemaMetadataTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *__VDLTarget1_optional) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if *t.Value == nil {
		*t.Value = &syncbase.SchemaMetadata{}
	}
	t.elemTarget.Value = *t.Value
	target, err := &t.elemTarget, error(nil)
	if err != nil {
		return nil, err
	}
	return target.StartFields(tt)
}
func (t *__VDLTarget1_optional) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}
func (t *__VDLTarget1_optional) FromNil(tt *vdl.Type) error {
	*t.Value = (*syncbase.SchemaMetadata)(nil)
	return nil
}

func (x *DatabaseData) VDLRead(dec vdl.Decoder) error {
	*x = DatabaseData{}
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
		case "Id":
			match++
			if err = x.Id.VDLRead(dec); err != nil {
				return err
			}
		case "Version":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.Version, err = dec.DecodeUint(64); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "Perms":
			match++
			if err = x.Perms.VDLRead(dec); err != nil {
				return err
			}
		case "SchemaMetadata":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if dec.IsNil() {
				if !vdl.Compatible(dec.Type(), vdl.TypeOf(x.SchemaMetadata)) {
					return fmt.Errorf("incompatible optional %T, from %v", x.SchemaMetadata, dec.Type())
				}
				x.SchemaMetadata = nil
				if err = dec.FinishValue(); err != nil {
					return err
				}
			} else {
				x.SchemaMetadata = new(syncbase.SchemaMetadata)
				dec.IgnoreNextStartValue()
				if err = x.SchemaMetadata.VDLRead(dec); err != nil {
					return err
				}
			}
		default:
			if err = dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

// CollectionPerms represent the persistent, synced permissions of a Collection.
// Existence of CollectionPerms in the store determines existence of the
// Collection.
// Note: Since CollectionPerms is synced and conflict resolved, the sync
// protocol needs to be aware of it. Any potential additions to synced
// Collection metadata should be written to a separate, synced key prefix,
// written in the same transaction with CollectionPerms and incorporated into
// the sync protocol. All persistent Collection metadata should be synced;
// local-only metadata is acceptable only if optional (e.g. stats).
type CollectionPerms access.Permissions

func (CollectionPerms) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/server.CollectionPerms"`
}) {
}

func (m *CollectionPerms) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	mapTarget1, err := t.StartMap(tt, len((*m)))
	if err != nil {
		return err
	}
	for key3, value5 := range *m {
		keyTarget2, err := mapTarget1.StartKey()
		if err != nil {
			return err
		}
		if err := keyTarget2.FromString(string(key3), tt.NonOptional().Key()); err != nil {
			return err
		}
		valueTarget4, err := mapTarget1.FinishKeyStartField(keyTarget2)
		if err != nil {
			return err
		}

		if err := value5.FillVDLTarget(valueTarget4, tt.NonOptional().Elem()); err != nil {
			return err
		}
		if err := mapTarget1.FinishField(keyTarget2, valueTarget4); err != nil {
			return err
		}
	}
	if err := t.FinishMap(mapTarget1); err != nil {
		return err
	}
	return nil
}

func (m *CollectionPerms) MakeVDLTarget() vdl.Target {
	return &CollectionPermsTarget{Value: m}
}

type CollectionPermsTarget struct {
	Value      *CollectionPerms
	currKey    string
	currElem   access.AccessList
	keyTarget  vdl.StringTarget
	elemTarget access.AccessListTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *CollectionPermsTarget) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*CollectionPerms)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(CollectionPerms)
	return t, nil
}
func (t *CollectionPermsTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *CollectionPermsTarget) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = access.AccessList{}
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *CollectionPermsTarget) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *CollectionPermsTarget) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

func (x *CollectionPerms) VDLRead(dec vdl.Decoder) error {
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if k := dec.Type().Kind(); k != vdl.Map {
		return fmt.Errorf("incompatible map %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len == 0:
		*x = nil
		return dec.FinishValue()
	case len > 0:
		*x = make(CollectionPerms, len)
	default:
		*x = make(CollectionPerms)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var key string
		{
			if err = dec.StartValue(); err != nil {
				return err
			}
			if key, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		}
		var elem access.AccessList
		{
			if err = elem.VDLRead(dec); err != nil {
				return err
			}
		}
		(*x)[key] = elem
	}
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
	vdl.Register((*ServiceData)(nil))
	vdl.Register((*DbInfo)(nil))
	vdl.Register((*DatabaseData)(nil))
	vdl.Register((*CollectionPerms)(nil))

	return struct{}{}
}
