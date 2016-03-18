// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: blobtestsvdl

package blobtestsvdl

import (
	"fmt"
	"reflect"
	"v.io/v23/services/syncbase/nosql"
	"v.io/v23/vdl"
	"v.io/v23/vom"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

type BlobInfo struct {
	Info string
	Br   nosql.BlobRef
}

func (BlobInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobInfo"`
}) {
}

func (m *BlobInfo) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Br")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Br.FillVDLTarget(fieldTarget5, tt.NonOptional().Field(1).Type); err != nil {
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

func (m *BlobInfo) MakeVDLTarget() vdl.Target {
	return &BlobInfoTarget{Value: m}
}

type BlobInfoTarget struct {
	Value      *BlobInfo
	infoTarget vdl.StringTarget
	brTarget   nosql.BlobRefTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobInfoTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobInfo)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobInfoTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Br":
		t.brTarget.Value = &t.Value.Br
		target, err := &t.brTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync/testdata.BlobInfo", name)
	}
}
func (t *BlobInfoTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobInfoTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

type (
	// BlobUnion represents any single field of the BlobUnion union type.
	BlobUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the BlobUnion union type.
		__VDLReflect(__BlobUnionReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
	}
	// BlobUnionNum represents field Num of the BlobUnion union type.
	BlobUnionNum struct{ Value int32 }
	// BlobUnionBi represents field Bi of the BlobUnion union type.
	BlobUnionBi struct{ Value BlobInfo }
	// __BlobUnionReflect describes the BlobUnion union type.
	__BlobUnionReflect struct {
		Name  string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobUnion"`
		Type  BlobUnion
		Union struct {
			Num BlobUnionNum
			Bi  BlobUnionBi
		}
	}
)

func (x BlobUnionNum) Index() int                      { return 0 }
func (x BlobUnionNum) Interface() interface{}          { return x.Value }
func (x BlobUnionNum) Name() string                    { return "Num" }
func (x BlobUnionNum) __VDLReflect(__BlobUnionReflect) {}

func (m BlobUnionNum) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Num")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromInt(int64(m.Value), tt.NonOptional().Field(0).Type); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlobUnionNum) MakeVDLTarget() vdl.Target {
	return nil
}

func (x BlobUnionBi) Index() int                      { return 1 }
func (x BlobUnionBi) Interface() interface{}          { return x.Value }
func (x BlobUnionBi) Name() string                    { return "Bi" }
func (x BlobUnionBi) __VDLReflect(__BlobUnionReflect) {}

func (m BlobUnionBi) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Bi")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(1).Type); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlobUnionBi) MakeVDLTarget() vdl.Target {
	return nil
}

type BlobSet struct {
	Info string
	Bs   map[nosql.BlobRef]struct{}
}

func (BlobSet) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobSet"`
}) {
}

func (m *BlobSet) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Bs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		setTarget6, err := fieldTarget5.StartSet(tt.NonOptional().Field(1).Type, len(m.Bs))
		if err != nil {
			return err
		}
		for key8 := range m.Bs {
			keyTarget7, err := setTarget6.StartKey()
			if err != nil {
				return err
			}

			if err := key8.FillVDLTarget(keyTarget7, tt.NonOptional().Field(1).Type.Key()); err != nil {
				return err
			}
			if err := setTarget6.FinishKey(keyTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishSet(setTarget6); err != nil {
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

func (m *BlobSet) MakeVDLTarget() vdl.Target {
	return &BlobSetTarget{Value: m}
}

type BlobSetTarget struct {
	Value      *BlobSet
	infoTarget vdl.StringTarget
	bsTarget   __VDLTarget1_set
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobSetTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobSet)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobSetTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Bs":
		t.bsTarget.Value = &t.Value.Bs
		target, err := &t.bsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync/testdata.BlobSet", name)
	}
}
func (t *BlobSetTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobSetTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[nosql.BlobRef]struct{}
type __VDLTarget1_set struct {
	Value     *map[nosql.BlobRef]struct{}
	currKey   nosql.BlobRef
	keyTarget nosql.BlobRefTarget
	vdl.TargetBase
	vdl.SetTargetBase
}

func (t *__VDLTarget1_set) StartSet(tt *vdl.Type, len int) (vdl.SetTarget, error) {

	if ttWant := vdl.TypeOf((*map[nosql.BlobRef]struct{})(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[nosql.BlobRef]struct{})
	return t, nil
}
func (t *__VDLTarget1_set) StartKey() (key vdl.Target, _ error) {
	t.currKey = nosql.BlobRef("")
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_set) FinishKey(key vdl.Target) error {
	(*t.Value)[t.currKey] = struct{}{}
	return nil
}
func (t *__VDLTarget1_set) FinishSet(list vdl.SetTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

type BlobAny struct {
	Info string
	Baa  []*vom.RawBytes
}

func (BlobAny) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobAny"`
}) {
}

func (m *BlobAny) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Baa")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget6, err := fieldTarget5.StartList(tt.NonOptional().Field(1).Type, len(m.Baa))
		if err != nil {
			return err
		}
		for i, elem8 := range m.Baa {
			elemTarget7, err := listTarget6.StartElem(i)
			if err != nil {
				return err
			}

			if elem8 == nil {
				if err := elemTarget7.FromNil(tt.NonOptional().Field(1).Type.Elem()); err != nil {
					return err
				}
			} else {
				if err := elem8.FillVDLTarget(elemTarget7, tt.NonOptional().Field(1).Type.Elem()); err != nil {
					return err
				}
			}
			if err := listTarget6.FinishElem(elemTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishList(listTarget6); err != nil {
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

func (m *BlobAny) MakeVDLTarget() vdl.Target {
	return &BlobAnyTarget{Value: m}
}

type BlobAnyTarget struct {
	Value      *BlobAny
	infoTarget vdl.StringTarget
	baaTarget  __VDLTarget2_list
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobAnyTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobAny)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobAnyTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Baa":
		t.baaTarget.Value = &t.Value.Baa
		target, err := &t.baaTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync/testdata.BlobAny", name)
	}
}
func (t *BlobAnyTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobAnyTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// []*vom.RawBytes
type __VDLTarget2_list struct {
	Value *[]*vom.RawBytes

	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *__VDLTarget2_list) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if ttWant := vdl.TypeOf((*[]*vom.RawBytes)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]*vom.RawBytes, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *__VDLTarget2_list) StartElem(index int) (elem vdl.Target, _ error) {
	target, err := vdl.ReflectTarget(reflect.ValueOf(&(*t.Value)[index]))
	return target, err
}
func (t *__VDLTarget2_list) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *__VDLTarget2_list) FinishList(elem vdl.ListTarget) error {

	return nil
}

type NonBlobSet struct {
	Info string
	S    map[string]struct{}
}

func (NonBlobSet) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.NonBlobSet"`
}) {
}

func (m *NonBlobSet) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("S")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		setTarget6, err := fieldTarget5.StartSet(tt.NonOptional().Field(1).Type, len(m.S))
		if err != nil {
			return err
		}
		for key8 := range m.S {
			keyTarget7, err := setTarget6.StartKey()
			if err != nil {
				return err
			}
			if err := keyTarget7.FromString(string(key8), tt.NonOptional().Field(1).Type.Key()); err != nil {
				return err
			}
			if err := setTarget6.FinishKey(keyTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishSet(setTarget6); err != nil {
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

func (m *NonBlobSet) MakeVDLTarget() vdl.Target {
	return &NonBlobSetTarget{Value: m}
}

type NonBlobSetTarget struct {
	Value      *NonBlobSet
	infoTarget vdl.StringTarget
	sTarget    __VDLTarget3_set
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *NonBlobSetTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*NonBlobSet)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *NonBlobSetTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "S":
		t.sTarget.Value = &t.Value.S
		target, err := &t.sTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync/testdata.NonBlobSet", name)
	}
}
func (t *NonBlobSetTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *NonBlobSetTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[string]struct{}
type __VDLTarget3_set struct {
	Value     *map[string]struct{}
	currKey   string
	keyTarget vdl.StringTarget
	vdl.TargetBase
	vdl.SetTargetBase
}

func (t *__VDLTarget3_set) StartSet(tt *vdl.Type, len int) (vdl.SetTarget, error) {

	if ttWant := vdl.TypeOf((*map[string]struct{})(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[string]struct{})
	return t, nil
}
func (t *__VDLTarget3_set) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget3_set) FinishKey(key vdl.Target) error {
	(*t.Value)[t.currKey] = struct{}{}
	return nil
}
func (t *__VDLTarget3_set) FinishSet(list vdl.SetTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

type BlobOpt struct {
	Info string
	Bo   *BlobInfo
}

func (BlobOpt) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobOpt"`
}) {
}

func (m *BlobOpt) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Bo")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if m.Bo == nil {
			if err := fieldTarget5.FromNil(tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
		} else {
			if err := m.Bo.FillVDLTarget(fieldTarget5, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
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

func (m *BlobOpt) MakeVDLTarget() vdl.Target {
	return &BlobOptTarget{Value: m}
}

type BlobOptTarget struct {
	Value      *BlobOpt
	infoTarget vdl.StringTarget
	boTarget   __VDLTarget4_optional
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobOptTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobOpt)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobOptTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Bo":
		t.boTarget.Value = &t.Value.Bo
		target, err := &t.boTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync/testdata.BlobOpt", name)
	}
}
func (t *BlobOptTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobOptTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// Optional BlobInfo
type __VDLTarget4_optional struct {
	Value      **BlobInfo
	elemTarget BlobInfoTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *__VDLTarget4_optional) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if *t.Value == nil {
		*t.Value = &BlobInfo{}
	}
	t.elemTarget.Value = *t.Value
	target, err := &t.elemTarget, error(nil)
	if err != nil {
		return nil, err
	}
	return target.StartFields(tt)
}
func (t *__VDLTarget4_optional) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}
func (t *__VDLTarget4_optional) FromNil(tt *vdl.Type) error {

	*t.Value = nil

	return nil
}

// Create zero values for each type.
var (
	__VDLZeroBlobInfo   = BlobInfo{}
	__VDLZeroBlobUnion  = BlobUnion(BlobUnionNum{})
	__VDLZeroBlobSet    = BlobSet{}
	__VDLZeroBlobAny    = BlobAny{}
	__VDLZeroNonBlobSet = NonBlobSet{}
	__VDLZeroBlobOpt    = BlobOpt{}
)

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

	// Register types.
	vdl.Register((*BlobInfo)(nil))
	vdl.Register((*BlobUnion)(nil))
	vdl.Register((*BlobSet)(nil))
	vdl.Register((*BlobAny)(nil))
	vdl.Register((*NonBlobSet)(nil))
	vdl.Register((*BlobOpt)(nil))

	return struct{}{}
}
