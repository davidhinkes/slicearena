// Package slicearena provides a library for efficently managing creating lots
// of slices without incurring heavy data allocation and GC penalties.
package slicearena

import (
	"reflect"
)

// T is the slice arena.  Create a *T via New().
type T struct {
	slice     reflect.Value
	ty        reflect.Type
	spillover int
}

// New creates a T.  The zero argument is used to deduce the type of slice.
func New(zero interface{}) *T {
	ty := reflect.SliceOf(reflect.TypeOf(zero))
	return &T{
		slice: reflect.MakeSlice(ty, 0, 0),
		ty:    ty,
	}
}

// Reset should be called when the slices created via T.MakeSlice are no
// longer needed.
func (t *T) Reset() {
	if spillover := t.spillover; spillover > 0 {
		newLength := t.slice.Len() + spillover
		t.slice = reflect.MakeSlice(t.ty, 0, newLength)
		t.spillover = 0
	} else {
		t.slice = t.slice.Slice(0, 0)
	}
}

func (t *T) MakeSlice(size int) interface{} {
	i := t.len()
	j := i + size
	if j > t.cap() {
		t.spillover += size
		return reflect.MakeSlice(t.ty, size, size).Interface()
	}
	t.slice = t.slice.Slice(0, j)
	ret := t.slice.Slice(i, j).Interface()
	return ret
}

func (t *T) len() int {
	return t.slice.Len()
}

func (t *T) cap() int {
	return t.slice.Cap()
}
