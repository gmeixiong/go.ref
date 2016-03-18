// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: testdata

package testdata

import (
	"fmt"
	"v.io/v23/discovery"
	"v.io/v23/vdl"
	discovery_2 "v.io/x/ref/lib/discovery"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// AdConversionTestCase represents a test case for converting between
// the advertisement and the Gatt characteristics.
type AdConversionTestCase struct {
	AdInfo discovery_2.AdInfo
	// GattAttrs is a map from uuid to the byte data.
	GattAttrs map[string][]byte
}

func (AdConversionTestCase) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery/plugins/ble/testdata.AdConversionTestCase"`
}) {
}

func (m *AdConversionTestCase) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("AdInfo")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.AdInfo.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("GattAttrs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		mapTarget6, err := fieldTarget5.StartMap(tt.NonOptional().Field(1).Type, len(m.GattAttrs))
		if err != nil {
			return err
		}
		for key8, value10 := range m.GattAttrs {
			keyTarget7, err := mapTarget6.StartKey()
			if err != nil {
				return err
			}
			if err := keyTarget7.FromString(string(key8), tt.NonOptional().Field(1).Type.Key()); err != nil {
				return err
			}
			valueTarget9, err := mapTarget6.FinishKeyStartField(keyTarget7)
			if err != nil {
				return err
			}

			if err := valueTarget9.FromBytes([]byte(value10), tt.NonOptional().Field(1).Type.Elem()); err != nil {
				return err
			}
			if err := mapTarget6.FinishField(keyTarget7, valueTarget9); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishMap(mapTarget6); err != nil {
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

func (m *AdConversionTestCase) MakeVDLTarget() vdl.Target {
	return &AdConversionTestCaseTarget{Value: m}
}

type AdConversionTestCaseTarget struct {
	Value           *AdConversionTestCase
	adInfoTarget    discovery_2.AdInfoTarget
	gattAttrsTarget __VDLTarget1_map
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *AdConversionTestCaseTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*AdConversionTestCase)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *AdConversionTestCaseTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "AdInfo":
		t.adInfoTarget.Value = &t.Value.AdInfo
		target, err := &t.adInfoTarget, error(nil)
		return nil, target, err
	case "GattAttrs":
		t.gattAttrsTarget.Value = &t.Value.GattAttrs
		target, err := &t.gattAttrsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/lib/discovery/plugins/ble/testdata.AdConversionTestCase", name)
	}
}
func (t *AdConversionTestCaseTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *AdConversionTestCaseTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[string][]byte
type __VDLTarget1_map struct {
	Value      *map[string][]byte
	currKey    string
	currElem   []byte
	keyTarget  vdl.StringTarget
	elemTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *__VDLTarget1_map) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*map[string][]byte)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[string][]byte)
	return t, nil
}
func (t *__VDLTarget1_map) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_map) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = []byte(nil)
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_map) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *__VDLTarget1_map) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// Create zero values for each type.
var (
	__VDLZeroAdConversionTestCase = AdConversionTestCase{}
)

//////////////////////////////////////////////////
// Const definitions

// ConversionTestData contains test cases for conversions between
// the advertisement and the Gatt characteristics
var ConversionTestData = []AdConversionTestCase{
	{
		AdInfo: discovery_2.AdInfo{
			Ad: discovery.Advertisement{
				Id: discovery.AdId{
					1,
					2,
					3,
					4,
					5,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
				},
				InterfaceName: "v.io/x/ref",
				Addresses: []string{
					"localhost:1000",
					"example.com:540",
				},
				Attributes: discovery.Attributes{
					"key1": "value1",
					"key2": "value2",
				},
				Attachments: discovery.Attachments{
					"key1": []byte("\x00\x01\x02\x03\x04"),
					"key3": []byte("\x05\x06\a\b\t"),
				},
			},
			EncryptionAlgorithm: 1,
			EncryptionKeys: []discovery_2.EncryptionKey{
				discovery_2.EncryptionKey("k"),
			},
			Hash: discovery_2.AdHash{
				1,
				3,
				5,
				7,
				9,
				0,
				0,
				0,
			},
			DirAddrs: []string{
				"localhost:1001",
				"example.com:541",
			},
		},
		GattAttrs: map[string][]byte{
			"402cce84-58f4-535b-8289-940365f62c96": []byte("__key1=\x00\x01\x02\x03\x04"),
			"4ce68e8b-173b-597e-9f93-ca453e7bb790": []byte("key1=value1"),
			"6286d80a-adaa-519a-8a06-281a4645a607": []byte("\x01\x01k"),
			"777f349c-d01f-5543-aa31-528e48bb53bd": []byte("key2=value2"),
			"7d8b5c56-0d05-5a7a-a21e-6c0c3c31245e": []byte("\x0elocalhost:1001\x0fexample.com:541"),
			"9c6286f5-aab0-5009-b81b-704d57ed6035": []byte("\x01\x03\x05\a\t\x00\x00\x00"),
			"b2cadfd4-d003-576c-acad-58b8e3a9cbc8": []byte("v.io/x/ref"),
			"bf0a3657-37cb-5aad-8c13-00c1d69a141c": []byte("\x01\x02\x03\x04\x05\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			"f3834b25-501c-566e-8343-d01fc632c922": []byte("__key3=\x05\x06\a\b\t"),
			"fe3fa941-1eda-5265-806f-d5127794a9a9": []byte("\x0elocalhost:1000\x0fexample.com:540"),
		},
	},
	{
		AdInfo: discovery_2.AdInfo{
			Ad: discovery.Advertisement{
				Id: discovery.AdId{
					9,
					8,
					7,
					6,
					5,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
					0,
				},
				InterfaceName: "v.io/x/large",
				Addresses: []string{
					"192.168.100.100:8000",
					"192.168.100.100:8001",
					"192.168.100.100:8002",
					"192.168.100.100:8003",
					"192.168.100.100:8004",
					"192.168.100.100:8005",
					"192.168.100.100:8006",
					"192.168.100.100:8007",
					"192.168.100.100:8008",
					"192.168.100.100:8009",
					"192.168.100.100:8010",
					"192.168.100.100:8011",
					"192.168.100.100:8012",
				},
				Attributes: discovery.Attributes{
					"key1": "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890012345678901234567890012345678901234567890",
					"key2": "9876543210",
				},
				Attachments: discovery.Attachments{
					"key1": []byte("\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00"),
					"key3": []byte("\x00\x01\x02\x03\x04\x05\x06\a\b\t"),
				},
			},
			EncryptionAlgorithm: 1,
			EncryptionKeys: []discovery_2.EncryptionKey{
				discovery_2.EncryptionKey("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"),
			},
			Hash: discovery_2.AdHash{
				2,
				4,
				6,
				8,
				0,
				0,
				0,
				0,
			},
			DirAddrs: []string{
				"192.168.100.100:9000",
				"192.168.100.100:9001",
				"192.168.100.100:9002",
				"192.168.100.100:9003",
				"192.168.100.100:9004",
				"192.168.100.100:9005",
				"192.168.100.100:9006",
				"192.168.100.100:9007",
				"192.168.100.100:9008",
				"192.168.100.100:9009",
				"192.168.100.100:9010",
				"192.168.100.100:9011",
				"192.168.100.100:9012",
			},
		},
		GattAttrs: map[string][]byte{
			"402cce84-58f4-535b-8289-940365f62c96": []byte("__key1=\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00\t\b\a\x06\x05\x04\x03\x02\x01\x00"),
			"4ce68e8b-173b-597e-9f93-ca453e7bb790": []byte("key1=01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890012345678901234567890012345678901234567890"),
			"6286d80a-adaa-519a-8a06-281a4645a607": []byte("\x01\x87\x02kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"),
			"777f349c-d01f-5543-aa31-528e48bb53bd": []byte("key2=9876543210"),
			"7d8b5c56-0d05-5a7a-a21e-6c0c3c31245e": []byte("\x14192.168.100.100:9000\x14192.168.100.100:9001\x14192.168.100.100:9002\x14192.168.100.100:9003\x14192.168.100.100:9004\x14192.168.100.100:9005\x14192.168.100.100:9006\x14192.168.100.100:9007\x14192.168.100.100:9008\x14192.168.100.100:9009\x14192.168.100.100:9010\x14192.168.100.100:9011\x14192.168.100.100:9012"),
			"9c6286f5-aab0-5009-b81b-704d57ed6035": []byte("\x02\x04\x06\b\x00\x00\x00\x00"),
			"b2cadfd4-d003-576c-acad-58b8e3a9cbc8": []byte("v.io/x/large"),
			"bf0a3657-37cb-5aad-8c13-00c1d69a141c": []byte("\t\b\a\x06\x05\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			"f3834b25-501c-566e-8343-d01fc632c922": []byte("__key3=\x00\x01\x02\x03\x04\x05\x06\a\b\t"),
			"fe3fa941-1eda-5265-806f-d5127794a9a9": []byte("\x14192.168.100.100:8000\x14192.168.100.100:8001\x14192.168.100.100:8002\x14192.168.100.100:8003\x14192.168.100.100:8004\x14192.168.100.100:8005\x14192.168.100.100:8006\x14192.168.100.100:8007\x14192.168.100.100:8008\x14192.168.100.100:8009\x14192.168.100.100:8010\x14192.168.100.100:8011\x14192.168.100.100:8012"),
		},
	},
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

	// Register types.
	vdl.Register((*AdConversionTestCase)(nil))

	return struct{}{}
}
