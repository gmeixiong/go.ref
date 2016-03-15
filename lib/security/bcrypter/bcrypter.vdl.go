// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: bcrypter

package bcrypter

import (
	"fmt"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/security"
	"v.io/v23/vdl"
	"v.io/v23/verror"
)

// WireCiphertext represents the wire format of the ciphertext
// generated by a Crypter.
type WireCiphertext struct {
	// PatternId is an identifier of the blessing pattern that this
	// ciphertext is for. It is represented by a 16 byte truncated
	// SHA256 hash of the pattern.
	PatternId string
	// Bytes is a map from an identifier of the public IBE params to
	// the ciphertext bytes that were generated using those params.
	//
	// The params identifier is a 16 byte truncated SHA256 hash
	// of the marshaled form of the IBE params.
	Bytes map[string][]byte
}

func (WireCiphertext) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WireCiphertext"`
}) {
}

func (m *WireCiphertext) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_lib_security_bcrypter_WireCiphertext == nil || __VDLType0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PatternId")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.PatternId), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Bytes")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		mapTarget6, err := fieldTarget5.StartMap(__VDLType1, len(m.Bytes))
		if err != nil {
			return err
		}
		for key8, value10 := range m.Bytes {
			keyTarget7, err := mapTarget6.StartKey()
			if err != nil {
				return err
			}
			if err := keyTarget7.FromString(string(key8), vdl.StringType); err != nil {
				return err
			}
			valueTarget9, err := mapTarget6.FinishKeyStartField(keyTarget7)
			if err != nil {
				return err
			}

			if err := valueTarget9.FromBytes([]byte(value10), __VDLType2); err != nil {
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

func (m *WireCiphertext) MakeVDLTarget() vdl.Target {
	return &WireCiphertextTarget{Value: m}
}

type WireCiphertextTarget struct {
	Value           *WireCiphertext
	patternIdTarget vdl.StringTarget
	bytesTarget     unnamed_6d61705b737472696e675d5b5d62797465Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WireCiphertextTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_lib_security_bcrypter_WireCiphertext) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_lib_security_bcrypter_WireCiphertext)
	}
	return t, nil
}
func (t *WireCiphertextTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "PatternId":
		t.patternIdTarget.Value = &t.Value.PatternId
		target, err := &t.patternIdTarget, error(nil)
		return nil, target, err
	case "Bytes":
		t.bytesTarget.Value = &t.Value.Bytes
		target, err := &t.bytesTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_lib_security_bcrypter_WireCiphertext)
	}
}
func (t *WireCiphertextTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WireCiphertextTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[string][]byte
type unnamed_6d61705b737472696e675d5b5d62797465Target struct {
	Value      *map[string][]byte
	currKey    string
	currElem   []byte
	keyTarget  vdl.StringTarget
	elemTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *unnamed_6d61705b737472696e675d5b5d62797465Target) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if !vdl.Compatible(tt, __VDLType1) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType1)
	}
	*t.Value = make(map[string][]byte)
	return t, nil
}
func (t *unnamed_6d61705b737472696e675d5b5d62797465Target) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *unnamed_6d61705b737472696e675d5b5d62797465Target) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = []byte(nil)
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *unnamed_6d61705b737472696e675d5b5d62797465Target) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *unnamed_6d61705b737472696e675d5b5d62797465Target) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// WireParams represents the wire format of the public parameters
// of an identity provider (aka Root).
type WireParams struct {
	// Blessing is the blessing name of the identity provider. The identity
	// provider  can extract private keys for blessings that are extensions
	// of this blessing name.
	Blessing string
	// Params is the marshaled form of the public IBE params of the
	// the identity provider.
	Params []byte
}

func (WireParams) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WireParams"`
}) {
}

func (m *WireParams) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_lib_security_bcrypter_WireParams == nil || __VDLType3 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Blessing")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Blessing), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Params")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget5.FromBytes([]byte(m.Params), __VDLType2); err != nil {
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

func (m *WireParams) MakeVDLTarget() vdl.Target {
	return &WireParamsTarget{Value: m}
}

type WireParamsTarget struct {
	Value          *WireParams
	blessingTarget vdl.StringTarget
	paramsTarget   vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WireParamsTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_lib_security_bcrypter_WireParams) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_lib_security_bcrypter_WireParams)
	}
	return t, nil
}
func (t *WireParamsTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Blessing":
		t.blessingTarget.Value = &t.Value.Blessing
		target, err := &t.blessingTarget, error(nil)
		return nil, target, err
	case "Params":
		t.paramsTarget.Value = &t.Value.Params
		target, err := &t.paramsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_lib_security_bcrypter_WireParams)
	}
}
func (t *WireParamsTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WireParamsTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// WirePrivateKey represents the wire format of the private key corresponding
// to a blessing.
type WirePrivateKey struct {
	// Blessing is the blessing for which this private key was extracted for.
	Blessing string
	// Params are the public parameters of the identity provider that extracted
	// this private key.
	Params WireParams
	// Keys contain the extracted IBE private keys for each pattern that is
	// matched by the blessing and is an extension of the identity provider's
	// name. The keys are enumerated in increasing order of the lengths of the
	// corresponding patterns.
	//
	// For example, if the blessing is "google:u:alice:phone" and the identity
	// provider's name is "google:u" then the keys are extracted for the patterns
	// - "google:u"
	// - "google:u:alice"
	// - "google:u:alice:phone"
	// - "google:u:alice:phone:$"
	//
	// The private keys are listed in increasing order of the lengths of the
	// corresponding patterns.
	Keys [][]byte
}

func (WirePrivateKey) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WirePrivateKey"`
}) {
}

func (m *WirePrivateKey) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_lib_security_bcrypter_WirePrivateKey == nil || __VDLType4 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Blessing")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Blessing), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Params")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Params.FillVDLTarget(fieldTarget5, __VDLType_v_io_x_ref_lib_security_bcrypter_WireParams); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Keys")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget8, err := fieldTarget7.StartList(__VDLType5, len(m.Keys))
		if err != nil {
			return err
		}
		for i, elem10 := range m.Keys {
			elemTarget9, err := listTarget8.StartElem(i)
			if err != nil {
				return err
			}

			if err := elemTarget9.FromBytes([]byte(elem10), __VDLType2); err != nil {
				return err
			}
			if err := listTarget8.FinishElem(elemTarget9); err != nil {
				return err
			}
		}
		if err := fieldTarget7.FinishList(listTarget8); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WirePrivateKey) MakeVDLTarget() vdl.Target {
	return &WirePrivateKeyTarget{Value: m}
}

type WirePrivateKeyTarget struct {
	Value          *WirePrivateKey
	blessingTarget vdl.StringTarget
	paramsTarget   WireParamsTarget
	keysTarget     unnamed_5b5d5b5d62797465Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WirePrivateKeyTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_lib_security_bcrypter_WirePrivateKey) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_lib_security_bcrypter_WirePrivateKey)
	}
	return t, nil
}
func (t *WirePrivateKeyTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Blessing":
		t.blessingTarget.Value = &t.Value.Blessing
		target, err := &t.blessingTarget, error(nil)
		return nil, target, err
	case "Params":
		t.paramsTarget.Value = &t.Value.Params
		target, err := &t.paramsTarget, error(nil)
		return nil, target, err
	case "Keys":
		t.keysTarget.Value = &t.Value.Keys
		target, err := &t.keysTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_lib_security_bcrypter_WirePrivateKey)
	}
}
func (t *WirePrivateKeyTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WirePrivateKeyTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// [][]byte
type unnamed_5b5d5b5d62797465Target struct {
	Value      *[][]byte
	elemTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *unnamed_5b5d5b5d62797465Target) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if !vdl.Compatible(tt, __VDLType5) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType5)
	}
	if cap(*t.Value) < len {
		*t.Value = make([][]byte, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *unnamed_5b5d5b5d62797465Target) StartElem(index int) (elem vdl.Target, _ error) {
	t.elemTarget.Value = &(*t.Value)[index]
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *unnamed_5b5d5b5d62797465Target) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *unnamed_5b5d5b5d62797465Target) FinishList(elem vdl.ListTarget) error {

	return nil
}

func init() {
	vdl.Register((*WireCiphertext)(nil))
	vdl.Register((*WireParams)(nil))
	vdl.Register((*WirePrivateKey)(nil))
}

var __VDLType0 *vdl.Type = vdl.TypeOf((*WireCiphertext)(nil))
var __VDLType3 *vdl.Type = vdl.TypeOf((*WireParams)(nil))
var __VDLType4 *vdl.Type = vdl.TypeOf((*WirePrivateKey)(nil))
var __VDLType5 *vdl.Type = vdl.TypeOf([][]byte(nil))
var __VDLType2 *vdl.Type = vdl.TypeOf([]byte(nil))
var __VDLType1 *vdl.Type = vdl.TypeOf(map[string][]byte(nil))
var __VDLType_v_io_x_ref_lib_security_bcrypter_WireCiphertext *vdl.Type = vdl.TypeOf(WireCiphertext{})
var __VDLType_v_io_x_ref_lib_security_bcrypter_WireParams *vdl.Type = vdl.TypeOf(WireParams{})
var __VDLType_v_io_x_ref_lib_security_bcrypter_WirePrivateKey *vdl.Type = vdl.TypeOf(WirePrivateKey{})

func __VDLEnsureNativeBuilt() {
}

var (
	ErrInternal           = verror.Register("v.io/x/ref/lib/security/bcrypter.Internal", verror.NoRetry, "{1:}{2:} internal error: {3}")
	ErrNoParams           = verror.Register("v.io/x/ref/lib/security/bcrypter.NoParams", verror.NoRetry, "{1:}{2:} no public parameters available for encrypting for pattern: {3}")
	ErrPrivateKeyNotFound = verror.Register("v.io/x/ref/lib/security/bcrypter.PrivateKeyNotFound", verror.NoRetry, "{1:}{2:} no private key found for decrypting ciphertext")
	ErrInvalidPrivateKey  = verror.Register("v.io/x/ref/lib/security/bcrypter.InvalidPrivateKey", verror.NoRetry, "{1:}{2:} private key is invalid: {3}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInternal.ID), "{1:}{2:} internal error: {3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoParams.ID), "{1:}{2:} no public parameters available for encrypting for pattern: {3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrPrivateKeyNotFound.ID), "{1:}{2:} no private key found for decrypting ciphertext")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidPrivateKey.ID), "{1:}{2:} private key is invalid: {3}")
}

// NewErrInternal returns an error with the ErrInternal ID.
func NewErrInternal(ctx *context.T, err error) error {
	return verror.New(ErrInternal, ctx, err)
}

// NewErrNoParams returns an error with the ErrNoParams ID.
func NewErrNoParams(ctx *context.T, pattern security.BlessingPattern) error {
	return verror.New(ErrNoParams, ctx, pattern)
}

// NewErrPrivateKeyNotFound returns an error with the ErrPrivateKeyNotFound ID.
func NewErrPrivateKeyNotFound(ctx *context.T) error {
	return verror.New(ErrPrivateKeyNotFound, ctx)
}

// NewErrInvalidPrivateKey returns an error with the ErrInvalidPrivateKey ID.
func NewErrInvalidPrivateKey(ctx *context.T, err error) error {
	return verror.New(ErrInvalidPrivateKey, ctx, err)
}
