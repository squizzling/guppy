package builtin

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type FFILen struct {
	Value struct {
		List   *primitive.ObjectList
		Tuple  *primitive.ObjectTuple
		String *primitive.ObjectString
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
