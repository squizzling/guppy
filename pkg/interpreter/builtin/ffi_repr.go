package builtin

import (
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type FFIRepr struct {
	Value itypes.Object `ffi:"value"`
}

func NewFFIRepr() itypes.FlowCall {
	return ffi.NewFFI(FFIRepr{})
}

func (f FFIRepr) Call(i itypes.Interpreter) (itypes.Object, error) {
	return primitive.NewObjectString(f.Value.Repr()), nil
}
