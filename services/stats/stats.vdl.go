// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: stats

// Packages stats defines the non-native types exported by the stats service.
package stats

import (
	"fmt"
	"v.io/v23/vdl"
)

// HistogramValue is the value of Histogram objects.
type HistogramValue struct {
	// Count is the total number of values added to the histogram.
	Count int64
	// Sum is the sum of all the values added to the histogram.
	Sum int64
	// Min is the minimum of all the values added to the histogram.
	Min int64
	// Max is the maximum of all the values added to the histogram.
	Max int64
	// Buckets contains all the buckets of the histogram.
	Buckets []HistogramBucket
}

func (HistogramValue) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/stats.HistogramValue"`
}) {
}

func (m *HistogramValue) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_stats_HistogramValue == nil || __VDLType0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Count")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromInt(int64(m.Count), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Sum")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.Sum), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Min")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget7.FromInt(int64(m.Min), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Max")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget9.FromInt(int64(m.Max), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	keyTarget10, fieldTarget11, err := fieldsTarget1.StartField("Buckets")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget12, err := fieldTarget11.StartList(__VDLType1, len(m.Buckets))
		if err != nil {
			return err
		}
		for i, elem14 := range m.Buckets {
			elemTarget13, err := listTarget12.StartElem(i)
			if err != nil {
				return err
			}

			if err := elem14.FillVDLTarget(elemTarget13, __VDLType_v_io_x_ref_services_stats_HistogramBucket); err != nil {
				return err
			}
			if err := listTarget12.FinishElem(elemTarget13); err != nil {
				return err
			}
		}
		if err := fieldTarget11.FinishList(listTarget12); err != nil {
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

func (m *HistogramValue) MakeVDLTarget() vdl.Target {
	return &HistogramValueTarget{Value: m}
}

type HistogramValueTarget struct {
	Value         *HistogramValue
	countTarget   vdl.Int64Target
	sumTarget     vdl.Int64Target
	minTarget     vdl.Int64Target
	maxTarget     vdl.Int64Target
	bucketsTarget unnamed_5b5d762e696f2f782f7265662f73657276696365732f73746174732e486973746f6772616d4275636b6574207374727563747b4c6f77426f756e6420696e7436343b436f756e7420696e7436347dTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *HistogramValueTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_stats_HistogramValue) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_stats_HistogramValue)
	}
	return t, nil
}
func (t *HistogramValueTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Count":
		t.countTarget.Value = &t.Value.Count
		target, err := &t.countTarget, error(nil)
		return nil, target, err
	case "Sum":
		t.sumTarget.Value = &t.Value.Sum
		target, err := &t.sumTarget, error(nil)
		return nil, target, err
	case "Min":
		t.minTarget.Value = &t.Value.Min
		target, err := &t.minTarget, error(nil)
		return nil, target, err
	case "Max":
		t.maxTarget.Value = &t.Value.Max
		target, err := &t.maxTarget, error(nil)
		return nil, target, err
	case "Buckets":
		t.bucketsTarget.Value = &t.Value.Buckets
		target, err := &t.bucketsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_stats_HistogramValue)
	}
}
func (t *HistogramValueTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *HistogramValueTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// []HistogramBucket
type unnamed_5b5d762e696f2f782f7265662f73657276696365732f73746174732e486973746f6772616d4275636b6574207374727563747b4c6f77426f756e6420696e7436343b436f756e7420696e7436347dTarget struct {
	Value      *[]HistogramBucket
	elemTarget HistogramBucketTarget
	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *unnamed_5b5d762e696f2f782f7265662f73657276696365732f73746174732e486973746f6772616d4275636b6574207374727563747b4c6f77426f756e6420696e7436343b436f756e7420696e7436347dTarget) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if !vdl.Compatible(tt, __VDLType1) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType1)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]HistogramBucket, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *unnamed_5b5d762e696f2f782f7265662f73657276696365732f73746174732e486973746f6772616d4275636b6574207374727563747b4c6f77426f756e6420696e7436343b436f756e7420696e7436347dTarget) StartElem(index int) (elem vdl.Target, _ error) {
	t.elemTarget.Value = &(*t.Value)[index]
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *unnamed_5b5d762e696f2f782f7265662f73657276696365732f73746174732e486973746f6772616d4275636b6574207374727563747b4c6f77426f756e6420696e7436343b436f756e7420696e7436347dTarget) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *unnamed_5b5d762e696f2f782f7265662f73657276696365732f73746174732e486973746f6772616d4275636b6574207374727563747b4c6f77426f756e6420696e7436343b436f756e7420696e7436347dTarget) FinishList(elem vdl.ListTarget) error {

	return nil
}

type HistogramBucketTarget struct {
	Value          *HistogramBucket
	lowBoundTarget vdl.Int64Target
	countTarget    vdl.Int64Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *HistogramBucketTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_stats_HistogramBucket) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_stats_HistogramBucket)
	}
	return t, nil
}
func (t *HistogramBucketTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "LowBound":
		t.lowBoundTarget.Value = &t.Value.LowBound
		target, err := &t.lowBoundTarget, error(nil)
		return nil, target, err
	case "Count":
		t.countTarget.Value = &t.Value.Count
		target, err := &t.countTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_stats_HistogramBucket)
	}
}
func (t *HistogramBucketTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *HistogramBucketTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// HistogramBucket is one histogram bucket.
type HistogramBucket struct {
	// LowBound is the lower bound of the bucket.
	LowBound int64
	// Count is the number of values in the bucket.
	Count int64
}

func (HistogramBucket) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/stats.HistogramBucket"`
}) {
}

func (m *HistogramBucket) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_stats_HistogramBucket == nil || __VDLType2 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("LowBound")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromInt(int64(m.LowBound), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Count")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.Count), vdl.Int64Type); err != nil {
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

func (m *HistogramBucket) MakeVDLTarget() vdl.Target {
	return &HistogramBucketTarget{Value: m}
}

func init() {
	vdl.Register((*HistogramValue)(nil))
	vdl.Register((*HistogramBucket)(nil))
}

var __VDLType2 *vdl.Type = vdl.TypeOf((*HistogramBucket)(nil))
var __VDLType0 *vdl.Type = vdl.TypeOf((*HistogramValue)(nil))
var __VDLType1 *vdl.Type = vdl.TypeOf([]HistogramBucket(nil))
var __VDLType_v_io_x_ref_services_stats_HistogramBucket *vdl.Type = vdl.TypeOf(HistogramBucket{})
var __VDLType_v_io_x_ref_services_stats_HistogramValue *vdl.Type = vdl.TypeOf(HistogramValue{})

func __VDLEnsureNativeBuilt() {
}