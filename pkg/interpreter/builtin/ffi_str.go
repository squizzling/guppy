package builtin

import (
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
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
