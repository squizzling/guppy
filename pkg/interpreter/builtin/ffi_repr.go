package builtin

import (
	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type FFIRepr struct {
	Value itypes.Reprable `ffi:"value"`
}

func NewFFIRepr() itypes.FlowCall {
	return ffi.NewFFI(FFIRepr{})
}

func (f FFIRepr) Call(i itypes.Interpreter) (itypes.Object, error) {
	return primitive.NewObjectString(f.Value.Repr()), nil
}
