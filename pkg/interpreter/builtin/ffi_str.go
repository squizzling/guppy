package builtin

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type FFIStr struct {
	Value interpreter.FlowStringable `ffi:"value"`
}

func NewFFIStr() itypes.FlowCall {
	return ffi.NewFFI(FFIStr{})
}

func (f FFIStr) Call(i itypes.Interpreter) (itypes.Object, error) {
	if s, err := f.Value.String(i); err != nil {
		return nil, err
	} else {
		return primitive.NewObjectString(s), nil
	}
}
