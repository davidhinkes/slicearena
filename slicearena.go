// Package slicearena provides a library for efficently managing creating lots
// of slices without incurring heavy data allocation and GC penalties.
package slicearena

import (
	"reflect"
)

// T is the slice arena.  Create a *T via New().
type T struct {
	slice reflect.Value
	ty    reflect.Type
}

// New creates a T.  The zero argument is used to deduce the type of slice.
func New(zero interface{}) *T {
	ty := reflect.TypeOf(zero)
	slice := reflect.MakeSlice(reflect.SliceOf(ty), 0, 0)
	return &T{
		slice: slice,
		ty:    ty,
	}
}

// Reset should be called when the slices created via T.MakeSlice are no
// longer needed.
func (t *T) Reset() {
	t.slice = t.slice.Slice(0, 0)
}

func (t *T) maybeGrow(cap int) {
	if t.slice.Cap() >= cap {
		return
	}
	biggerSlice := reflect.MakeSlice(reflect.SliceOf(t.ty), t.slice.Len(), cap)
	reflect.Copy(biggerSlice, t.slice)
	t.slice = biggerSlice
}

func (t *T) MakeSlice(size int) interface{} {
	i := t.slice.Len()
	j := i + size
	t.maybeGrow(j)
	ret := t.slice.Slice(i, j).Interface()
	t.slice = t.slice.Slice(0, j)
	return ret
}

func (t *T) len() int {
	return t.slice.Len()
}

func (t *T) cap() int {
	return t.slice.Cap()
}
