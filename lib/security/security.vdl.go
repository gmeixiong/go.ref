// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: security

package security

import (
	"fmt"
	"reflect"
	"time"
	"v.io/v23/security"
	"v.io/v23/vdl"
	time_2 "v.io/v23/vdlroot/time"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

type blessingRootsState map[string][]security.BlessingPattern

func (blessingRootsState) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security.blessingRootsState"`
}) {
}

func (m *blessingRootsState) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
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

		listTarget6, err := valueTarget4.StartList(tt.NonOptional().Elem(), len(value5))
		if err != nil {
			return err
		}
		for i, elem8 := range value5 {
			elemTarget7, err := listTarget6.StartElem(i)
			if err != nil {
				return err
			}

			if err := elem8.FillVDLTarget(elemTarget7, tt.NonOptional().Elem().Elem()); err != nil {
				return err
			}
			if err := listTarget6.FinishElem(elemTarget7); err != nil {
				return err
			}
		}
		if err := valueTarget4.FinishList(listTarget6); err != nil {
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

func (m *blessingRootsState) MakeVDLTarget() vdl.Target {
	return &blessingRootsStateTarget{Value: m}
}

type blessingRootsStateTarget struct {
	Value      *blessingRootsState
	currKey    string
	currElem   []security.BlessingPattern
	keyTarget  vdl.StringTarget
	elemTarget __VDLTarget1_list
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *blessingRootsStateTarget) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*blessingRootsState)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(blessingRootsState)
	return t, nil
}
func (t *blessingRootsStateTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *blessingRootsStateTarget) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = []security.BlessingPattern(nil)
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *blessingRootsStateTarget) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *blessingRootsStateTarget) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// []security.BlessingPattern
type __VDLTarget1_list struct {
	Value      *[]security.BlessingPattern
	elemTarget security.BlessingPatternTarget
	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *__VDLTarget1_list) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if ttWant := vdl.TypeOf((*[]security.BlessingPattern)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]security.BlessingPattern, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *__VDLTarget1_list) StartElem(index int) (elem vdl.Target, _ error) {
	t.elemTarget.Value = &(*t.Value)[index]
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_list) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *__VDLTarget1_list) FinishList(elem vdl.ListTarget) error {

	return nil
}

type dischargeCacheKey [32]byte

func (dischargeCacheKey) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security.dischargeCacheKey"`
}) {
}

func (m *dischargeCacheKey) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromBytes([]byte((*m)[:]), tt); err != nil {
		return err
	}
	return nil
}

func (m *dischargeCacheKey) MakeVDLTarget() vdl.Target {
	return &dischargeCacheKeyTarget{Value: m}
}

type dischargeCacheKeyTarget struct {
	Value *dischargeCacheKey
	vdl.TargetBase
}

func (t *dischargeCacheKeyTarget) FromBytes(src []byte, tt *vdl.Type) error {

	if ttWant := vdl.TypeOf((*dischargeCacheKey)(nil)); !vdl.Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	copy((*t.Value)[:], src)

	return nil
}

type CachedDischarge struct {
	Discharge security.Discharge
	// CacheTime is the time at which the discharge was first cached.
	CacheTime time.Time
}

func (CachedDischarge) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security.CachedDischarge"`
}) {
}

func (m *CachedDischarge) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var wireValue2 security.WireDischarge
	if err := security.WireDischargeFromNative(&wireValue2, m.Discharge); err != nil {
		return err
	}

	keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("Discharge")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		unionValue5 := wireValue2
		if unionValue5 == nil {
			unionValue5 = security.WireDischargePublicKey{}
		}
		if err := unionValue5.FillVDLTarget(fieldTarget4, tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
			return err
		}
	}
	var wireValue6 time_2.Time
	if err := time_2.TimeFromNative(&wireValue6, m.CacheTime); err != nil {
		return err
	}

	keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("CacheTime")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue6.FillVDLTarget(fieldTarget8, tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget7, fieldTarget8); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *CachedDischarge) MakeVDLTarget() vdl.Target {
	return &CachedDischargeTarget{Value: m}
}

type CachedDischargeTarget struct {
	Value *CachedDischarge

	cacheTimeTarget time_2.TimeTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *CachedDischargeTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*CachedDischarge)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *CachedDischargeTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Discharge":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Discharge))
		return nil, target, err
	case "CacheTime":
		t.cacheTimeTarget.Value = &t.Value.CacheTime
		target, err := &t.cacheTimeTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/lib/security.CachedDischarge", name)
	}
}
func (t *CachedDischargeTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *CachedDischargeTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

type blessingStoreState struct {
	// PeerBlessings maps BlessingPatterns to the Blessings object that is to
	// be shared with peers which present blessings of their own that match the
	// pattern.
	//
	// All blessings bind to the same public key.
	PeerBlessings map[security.BlessingPattern]security.Blessings
	// DefaultBlessings is the default Blessings to be shared with peers for which
	// no other information is available to select blessings.
	DefaultBlessings security.Blessings
	// DischargeCache is the cache of discharges.
	// TODO(mattr): This map is deprecated in favor of the Discharges map below.
	DischargeCache map[dischargeCacheKey]security.Discharge
	// DischargeCache is the cache of discharges.
	Discharges map[dischargeCacheKey]CachedDischarge
	// CacheKeyFormat is the dischargeCacheKey format version. It should incremented
	// any time the format of the dischargeCacheKey is changed.
	CacheKeyFormat uint32
}

func (blessingStoreState) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security.blessingStoreState"`
}) {
}

func (m *blessingStoreState) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PeerBlessings")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		mapTarget4, err := fieldTarget3.StartMap(tt.NonOptional().Field(0).Type, len(m.PeerBlessings))
		if err != nil {
			return err
		}
		for key6, value8 := range m.PeerBlessings {
			keyTarget5, err := mapTarget4.StartKey()
			if err != nil {
				return err
			}

			if err := key6.FillVDLTarget(keyTarget5, tt.NonOptional().Field(0).Type.Key()); err != nil {
				return err
			}
			valueTarget7, err := mapTarget4.FinishKeyStartField(keyTarget5)
			if err != nil {
				return err
			}

			var wireValue9 security.WireBlessings
			if err := security.WireBlessingsFromNative(&wireValue9, value8); err != nil {
				return err
			}

			if err := wireValue9.FillVDLTarget(valueTarget7, tt.NonOptional().Field(0).Type.Elem()); err != nil {
				return err
			}
			if err := mapTarget4.FinishField(keyTarget5, valueTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget3.FinishMap(mapTarget4); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	var wireValue10 security.WireBlessings
	if err := security.WireBlessingsFromNative(&wireValue10, m.DefaultBlessings); err != nil {
		return err
	}

	keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("DefaultBlessings")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue10.FillVDLTarget(fieldTarget12, tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
			return err
		}
	}
	keyTarget13, fieldTarget14, err := fieldsTarget1.StartField("DischargeCache")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		mapTarget15, err := fieldTarget14.StartMap(tt.NonOptional().Field(2).Type, len(m.DischargeCache))
		if err != nil {
			return err
		}
		for key17, value19 := range m.DischargeCache {
			keyTarget16, err := mapTarget15.StartKey()
			if err != nil {
				return err
			}

			if err := key17.FillVDLTarget(keyTarget16, tt.NonOptional().Field(2).Type.Key()); err != nil {
				return err
			}
			valueTarget18, err := mapTarget15.FinishKeyStartField(keyTarget16)
			if err != nil {
				return err
			}

			var wireValue20 security.WireDischarge
			if err := security.WireDischargeFromNative(&wireValue20, value19); err != nil {
				return err
			}

			unionValue21 := wireValue20
			if unionValue21 == nil {
				unionValue21 = security.WireDischargePublicKey{}
			}
			if err := unionValue21.FillVDLTarget(valueTarget18, tt.NonOptional().Field(2).Type.Elem()); err != nil {
				return err
			}
			if err := mapTarget15.FinishField(keyTarget16, valueTarget18); err != nil {
				return err
			}
		}
		if err := fieldTarget14.FinishMap(mapTarget15); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget13, fieldTarget14); err != nil {
			return err
		}
	}
	keyTarget22, fieldTarget23, err := fieldsTarget1.StartField("Discharges")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		mapTarget24, err := fieldTarget23.StartMap(tt.NonOptional().Field(3).Type, len(m.Discharges))
		if err != nil {
			return err
		}
		for key26, value28 := range m.Discharges {
			keyTarget25, err := mapTarget24.StartKey()
			if err != nil {
				return err
			}

			if err := key26.FillVDLTarget(keyTarget25, tt.NonOptional().Field(3).Type.Key()); err != nil {
				return err
			}
			valueTarget27, err := mapTarget24.FinishKeyStartField(keyTarget25)
			if err != nil {
				return err
			}

			if err := value28.FillVDLTarget(valueTarget27, tt.NonOptional().Field(3).Type.Elem()); err != nil {
				return err
			}
			if err := mapTarget24.FinishField(keyTarget25, valueTarget27); err != nil {
				return err
			}
		}
		if err := fieldTarget23.FinishMap(mapTarget24); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget22, fieldTarget23); err != nil {
			return err
		}
	}
	keyTarget29, fieldTarget30, err := fieldsTarget1.StartField("CacheKeyFormat")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget30.FromUint(uint64(m.CacheKeyFormat), tt.NonOptional().Field(4).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget29, fieldTarget30); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *blessingStoreState) MakeVDLTarget() vdl.Target {
	return &blessingStoreStateTarget{Value: m}
}

type blessingStoreStateTarget struct {
	Value                  *blessingStoreState
	peerBlessingsTarget    __VDLTarget2_map
	defaultBlessingsTarget security.WireBlessingsTarget
	dischargeCacheTarget   __VDLTarget3_map
	dischargesTarget       __VDLTarget4_map
	cacheKeyFormatTarget   vdl.Uint32Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *blessingStoreStateTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*blessingStoreState)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *blessingStoreStateTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "PeerBlessings":
		t.peerBlessingsTarget.Value = &t.Value.PeerBlessings
		target, err := &t.peerBlessingsTarget, error(nil)
		return nil, target, err
	case "DefaultBlessings":
		t.defaultBlessingsTarget.Value = &t.Value.DefaultBlessings
		target, err := &t.defaultBlessingsTarget, error(nil)
		return nil, target, err
	case "DischargeCache":
		t.dischargeCacheTarget.Value = &t.Value.DischargeCache
		target, err := &t.dischargeCacheTarget, error(nil)
		return nil, target, err
	case "Discharges":
		t.dischargesTarget.Value = &t.Value.Discharges
		target, err := &t.dischargesTarget, error(nil)
		return nil, target, err
	case "CacheKeyFormat":
		t.cacheKeyFormatTarget.Value = &t.Value.CacheKeyFormat
		target, err := &t.cacheKeyFormatTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/lib/security.blessingStoreState", name)
	}
}
func (t *blessingStoreStateTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *blessingStoreStateTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[security.BlessingPattern]security.Blessings
type __VDLTarget2_map struct {
	Value      *map[security.BlessingPattern]security.Blessings
	currKey    security.BlessingPattern
	currElem   security.Blessings
	keyTarget  security.BlessingPatternTarget
	elemTarget security.WireBlessingsTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *__VDLTarget2_map) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*map[security.BlessingPattern]security.Blessings)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[security.BlessingPattern]security.Blessings)
	return t, nil
}
func (t *__VDLTarget2_map) StartKey() (key vdl.Target, _ error) {
	t.currKey = security.BlessingPattern("")
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget2_map) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = reflect.Zero(reflect.TypeOf(t.currElem)).Interface().(security.Blessings)
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget2_map) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *__VDLTarget2_map) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// map[dischargeCacheKey]security.Discharge
type __VDLTarget3_map struct {
	Value     *map[dischargeCacheKey]security.Discharge
	currKey   dischargeCacheKey
	currElem  security.Discharge
	keyTarget dischargeCacheKeyTarget

	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *__VDLTarget3_map) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*map[dischargeCacheKey]security.Discharge)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[dischargeCacheKey]security.Discharge)
	return t, nil
}
func (t *__VDLTarget3_map) StartKey() (key vdl.Target, _ error) {
	t.currKey = dischargeCacheKey{}
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget3_map) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = reflect.Zero(reflect.TypeOf(t.currElem)).Interface().(security.Discharge)
	target, err := vdl.ReflectTarget(reflect.ValueOf(&t.currElem))
	return target, err
}
func (t *__VDLTarget3_map) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *__VDLTarget3_map) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// map[dischargeCacheKey]CachedDischarge
type __VDLTarget4_map struct {
	Value      *map[dischargeCacheKey]CachedDischarge
	currKey    dischargeCacheKey
	currElem   CachedDischarge
	keyTarget  dischargeCacheKeyTarget
	elemTarget CachedDischargeTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *__VDLTarget4_map) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*map[dischargeCacheKey]CachedDischarge)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[dischargeCacheKey]CachedDischarge)
	return t, nil
}
func (t *__VDLTarget4_map) StartKey() (key vdl.Target, _ error) {
	t.currKey = dischargeCacheKey{}
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget4_map) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = reflect.Zero(reflect.TypeOf(t.currElem)).Interface().(CachedDischarge)
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget4_map) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *__VDLTarget4_map) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// Create zero values for each type.
var (
	__VDLZeroblessingRootsState = blessingRootsState(nil)
	__VDLZerodischargeCacheKey  = dischargeCacheKey{}
	__VDLZeroCachedDischarge    = CachedDischarge{
		Discharge: func() security.Discharge {
			var native security.Discharge
			if err := vdl.Convert(&native, security.WireDischarge(security.WireDischargePublicKey{})); err != nil {
				panic(err)
			}
			return native
		}(),
	}
	__VDLZeroblessingStoreState = blessingStoreState{}
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
	vdl.Register((*blessingRootsState)(nil))
	vdl.Register((*dischargeCacheKey)(nil))
	vdl.Register((*CachedDischarge)(nil))
	vdl.Register((*blessingStoreState)(nil))

	return struct{}{}
}
