package ffi

import (
	"guppy/pkg/interpreter"
)

type FFIStr struct {
	Value interpreter.FlowStringable `ffi:"value"`
}

func NewFFIStr() interpreter.FlowCall {
	return interpreter.NewFFI(FFIStr{})
}

func (f FFIStr) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if s, err := f.Value.String(i); err != nil {
		return nil, err
	} else {
		return interpreter.NewObjectString(s), nil
	}
}
