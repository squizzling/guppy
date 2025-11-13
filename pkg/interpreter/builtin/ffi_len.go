package builtin

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type FFILen struct {
	Value struct {
		List   *interpreter.ObjectList
		Tuple  *interpreter.ObjectTuple
		String *interpreter.ObjectString
	} `ffi:"value"`
}

func NewFFILen() itypes.FlowCall {
	return ffi.NewFFI(FFILen{})
}

func (f FFILen) Call(i itypes.Interpreter) (itypes.Object, error) {
	switch {
	case f.Value.List != nil:
		return primitive.NewObjectInt(len(f.Value.List.Items)), nil
	case f.Value.Tuple != nil:
		return primitive.NewObjectInt(len(f.Value.Tuple.Items)), nil
	case f.Value.String != nil:
		return primitive.NewObjectInt(len(f.Value.String.Value)), nil
	default:
		return nil, fmt.Errorf("FFILen.Call: FFILen.Value is not set")
	}
}
