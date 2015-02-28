// This file was auto-generated by the veyron vdl tool.
// Source: types.vdl

// Packages stats defines the non-native types exported by the stats service.
package stats

import (
	// VDL system imports
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
	Name string "v.io/x/ref/services/mgmt/stats.HistogramValue"
}) {
}

// HistogramBucket is one histogram bucket.
type HistogramBucket struct {
	// LowBound is the lower bound of the bucket.
	LowBound int64
	// Count is the number of values in the bucket.
	Count int64
}

func (HistogramBucket) __VDLReflect(struct {
	Name string "v.io/x/ref/services/mgmt/stats.HistogramBucket"
}) {
}

func init() {
	vdl.Register((*HistogramValue)(nil))
	vdl.Register((*HistogramBucket)(nil))
}
