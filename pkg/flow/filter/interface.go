package filter

import (
	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
)

type Filter interface {
	itypes.Object
}

var prototypeFilter = itypes.NewObject(map[string]itypes.Object{
	"__binary_and__":       ffi.NewFFI(ffiFilterBinaryOp{op: 0}),
	"__binary_or__":        ffi.NewFFI(ffiFilterBinaryOp{op: 1}),
	"__unary_binary_not__": ffi.NewFFI(ffiFilterUnaryBinaryNot{}),

	// eq/is and ne/isnot have the same behavior
	"__eq__":    ffi.NewFFI(ffiFilterRelOp{invert: false}),
	"__ne__":    ffi.NewFFI(ffiFilterRelOp{invert: true}),
	"__is__":    ffi.NewFFI(ffiFilterRelOp{invert: false}),
	"__isnot__": ffi.NewFFI(ffiFilterRelOp{invert: true}),
})
