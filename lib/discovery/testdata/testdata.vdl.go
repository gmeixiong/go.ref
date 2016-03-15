// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: testdata

package testdata

import (
	"fmt"
	"v.io/v23/vdl"
	"v.io/x/ref/lib/discovery"
)

// PackAddressTest represents a test case for PackAddress.
type PackAddressTest struct {
	// In is the addresses to pack.
	In []string
	// Packed is the expected packed output.
	Packed []byte
}

func (PackAddressTest) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery/testdata.PackAddressTest"`
}) {
}

func (m *PackAddressTest) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_lib_discovery_testdata_PackAddressTest == nil || __VDLType0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("In")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget4, err := fieldTarget3.StartList(__VDLType1, len(m.In))
		if err != nil {
			return err
		}
		for i, elem6 := range m.In {
			elemTarget5, err := listTarget4.StartElem(i)
			if err != nil {
				return err
			}
			if err := elemTarget5.FromString(string(elem6), vdl.StringType); err != nil {
				return err
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
	keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("Packed")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget8.FromBytes([]byte(m.Packed), __VDLType2); err != nil {
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

func (m *PackAddressTest) MakeVDLTarget() vdl.Target {
	return &PackAddressTestTarget{Value: m}
}

type PackAddressTestTarget struct {
	Value        *PackAddressTest
	inTarget     vdl.StringSliceTarget
	packedTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *PackAddressTestTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_lib_discovery_testdata_PackAddressTest) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_lib_discovery_testdata_PackAddressTest)
	}
	return t, nil
}
func (t *PackAddressTestTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "In":
		t.inTarget.Value = &t.Value.In
		target, err := &t.inTarget, error(nil)
		return nil, target, err
	case "Packed":
		t.packedTarget.Value = &t.Value.Packed
		target, err := &t.packedTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_lib_discovery_testdata_PackAddressTest)
	}
}
func (t *PackAddressTestTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *PackAddressTestTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// PackEncryptionKeysTest represents a test case for PackEncryptionKeys
type PackEncryptionKeysTest struct {
	// Algo is the algorithm that's in use.
	// but that isn't defined in vdl yet.
	Algo discovery.EncryptionAlgorithm
	// Keys are the encryption keys.
	// but that isn't defined in vdl yet.
	Keys []discovery.EncryptionKey
	// Packed is the expected output bytes.
	Packed []byte
}

func (PackEncryptionKeysTest) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery/testdata.PackEncryptionKeysTest"`
}) {
}

func (m *PackEncryptionKeysTest) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_lib_discovery_testdata_PackEncryptionKeysTest == nil || __VDLType3 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Algo")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Algo.FillVDLTarget(fieldTarget3, __VDLType_v_io_x_ref_lib_discovery_EncryptionAlgorithm); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Keys")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget6, err := fieldTarget5.StartList(__VDLType4, len(m.Keys))
		if err != nil {
			return err
		}
		for i, elem8 := range m.Keys {
			elemTarget7, err := listTarget6.StartElem(i)
			if err != nil {
				return err
			}

			if err := elem8.FillVDLTarget(elemTarget7, __VDLType_v_io_x_ref_lib_discovery_EncryptionKey); err != nil {
				return err
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
	keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("Packed")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget10.FromBytes([]byte(m.Packed), __VDLType2); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *PackEncryptionKeysTest) MakeVDLTarget() vdl.Target {
	return &PackEncryptionKeysTestTarget{Value: m}
}

type PackEncryptionKeysTestTarget struct {
	Value        *PackEncryptionKeysTest
	algoTarget   discovery.EncryptionAlgorithmTarget
	keysTarget   unnamed_5b5d762e696f2f782f7265662f6c69622f646973636f766572792e456e6372797074696f6e4b6579205b5d62797465Target
	packedTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *PackEncryptionKeysTestTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_lib_discovery_testdata_PackEncryptionKeysTest) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_lib_discovery_testdata_PackEncryptionKeysTest)
	}
	return t, nil
}
func (t *PackEncryptionKeysTestTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Algo":
		t.algoTarget.Value = &t.Value.Algo
		target, err := &t.algoTarget, error(nil)
		return nil, target, err
	case "Keys":
		t.keysTarget.Value = &t.Value.Keys
		target, err := &t.keysTarget, error(nil)
		return nil, target, err
	case "Packed":
		t.packedTarget.Value = &t.Value.Packed
		target, err := &t.packedTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_lib_discovery_testdata_PackEncryptionKeysTest)
	}
}
func (t *PackEncryptionKeysTestTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *PackEncryptionKeysTestTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// []discovery.EncryptionKey
type unnamed_5b5d762e696f2f782f7265662f6c69622f646973636f766572792e456e6372797074696f6e4b6579205b5d62797465Target struct {
	Value      *[]discovery.EncryptionKey
	elemTarget discovery.EncryptionKeyTarget
	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *unnamed_5b5d762e696f2f782f7265662f6c69622f646973636f766572792e456e6372797074696f6e4b6579205b5d62797465Target) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if !vdl.Compatible(tt, __VDLType4) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType4)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]discovery.EncryptionKey, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *unnamed_5b5d762e696f2f782f7265662f6c69622f646973636f766572792e456e6372797074696f6e4b6579205b5d62797465Target) StartElem(index int) (elem vdl.Target, _ error) {
	t.elemTarget.Value = &(*t.Value)[index]
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *unnamed_5b5d762e696f2f782f7265662f6c69622f646973636f766572792e456e6372797074696f6e4b6579205b5d62797465Target) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *unnamed_5b5d762e696f2f782f7265662f6c69622f646973636f766572792e456e6372797074696f6e4b6579205b5d62797465Target) FinishList(elem vdl.ListTarget) error {

	return nil
}

// UuidTestData represents the inputs and outputs for a uuid test.
type UuidTestData struct {
	// In is the input string.
	In string
	// Want is the expected uuid's human-readable string form.
	Want string
}

func (UuidTestData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery/testdata.UuidTestData"`
}) {
}

func (m *UuidTestData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_lib_discovery_testdata_UuidTestData == nil || __VDLType5 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("In")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.In), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Want")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.Want), vdl.StringType); err != nil {
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

func (m *UuidTestData) MakeVDLTarget() vdl.Target {
	return &UuidTestDataTarget{Value: m}
}

type UuidTestDataTarget struct {
	Value      *UuidTestData
	inTarget   vdl.StringTarget
	wantTarget vdl.StringTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *UuidTestDataTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_lib_discovery_testdata_UuidTestData) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_lib_discovery_testdata_UuidTestData)
	}
	return t, nil
}
func (t *UuidTestDataTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "In":
		t.inTarget.Value = &t.Value.In
		target, err := &t.inTarget, error(nil)
		return nil, target, err
	case "Want":
		t.wantTarget.Value = &t.Value.Want
		target, err := &t.wantTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_lib_discovery_testdata_UuidTestData)
	}
}
func (t *UuidTestDataTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *UuidTestDataTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func init() {
	vdl.Register((*PackAddressTest)(nil))
	vdl.Register((*PackEncryptionKeysTest)(nil))
	vdl.Register((*UuidTestData)(nil))
}

var __VDLType0 *vdl.Type = vdl.TypeOf((*PackAddressTest)(nil))
var __VDLType3 *vdl.Type = vdl.TypeOf((*PackEncryptionKeysTest)(nil))
var __VDLType5 *vdl.Type = vdl.TypeOf((*UuidTestData)(nil))
var __VDLType2 *vdl.Type = vdl.TypeOf([]byte(nil))
var __VDLType1 *vdl.Type = vdl.TypeOf([]string(nil))
var __VDLType4 *vdl.Type = vdl.TypeOf([]discovery.EncryptionKey(nil))
var __VDLType_v_io_x_ref_lib_discovery_EncryptionAlgorithm *vdl.Type = vdl.TypeOf(discovery.EncryptionAlgorithm(0))
var __VDLType_v_io_x_ref_lib_discovery_EncryptionKey *vdl.Type = vdl.TypeOf(discovery.EncryptionKey(nil))
var __VDLType_v_io_x_ref_lib_discovery_testdata_PackAddressTest *vdl.Type = vdl.TypeOf(PackAddressTest{})
var __VDLType_v_io_x_ref_lib_discovery_testdata_PackEncryptionKeysTest *vdl.Type = vdl.TypeOf(PackEncryptionKeysTest{})
var __VDLType_v_io_x_ref_lib_discovery_testdata_UuidTestData *vdl.Type = vdl.TypeOf(UuidTestData{})

func __VDLEnsureNativeBuilt() {
}

var PackAddressTestData = []PackAddressTest{
	{
		In: []string{
			"a12345",
		},
		Packed: []byte("\x06a12345"),
	},
	{
		In: []string{
			"a1234",
			"b5678",
			"c9012",
		},
		Packed: []byte("\x05a1234\x05b5678\x05c9012"),
	},
	{},
}

var PackEncryptionKeysTestData = []PackEncryptionKeysTest{
	{
		Algo: 1,
		Keys: []discovery.EncryptionKey{
			discovery.EncryptionKey("0123456789"),
		},
		Packed: []byte("\x01\n0123456789"),
	},
	{
		Algo: 2,
		Keys: []discovery.EncryptionKey{
			discovery.EncryptionKey("012345"),
			discovery.EncryptionKey("123456"),
			discovery.EncryptionKey("234567"),
		},
		Packed: []byte("\x02\x06012345\x06123456\x06234567"),
	},
	{
		Packed: []byte("\x00"),
	},
}

var ServiceUuidTest = []UuidTestData{
	{
		In:   "v.io",
		Want: "2101363c-688d-548a-a600-34d506e1aad0",
	},
	{
		In:   "v.io/v23/abc",
		Want: "6726c4e5-b6eb-5547-9228-b2913f4fad52",
	},
	{
		In:   "v.io/v23/abc/xyz",
		Want: "be8a57d7-931d-5ee4-9243-0bebde0029a5",
	},
}

var AttributeUuidTest = []UuidTestData{
	{
		In:   "name",
		Want: "217a496d-3aae-5748-baf0-a77555f8f4f4",
	},
	{
		In:   "_attr",
		Want: "6c020e4b-9a59-5c7f-92e7-45954a16a402",
	},
	{
		In:   "xyz",
		Want: "c10b25a2-2d4d-5a19-bb7c-1ee1c4972b4c",
	},
}
