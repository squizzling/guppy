package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/ftypes"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

// TODO: Probably explore the space, for now it's just MVP.

type ffiStreamBelow struct {
	Self      Stream                `ffi:"self"`
	Limit     ftypes.IntOrDouble    `ffi:"limit"`
	Inclusive *primitive.ObjectBool `ffi:"inclusive"`
	Clamp     *primitive.ObjectBool `ffi:"clamp"`
}

func NewFFIStreamBelow() itypes.FlowCall {
	return ffi.NewFFI(ffiStreamBelow{
		Inclusive: primitive.NewObjectBool(false),
		Clamp:     primitive.NewObjectBool(false),
	})
}

func (f ffiStreamBelow) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewStreamMethodBelow(prototypeStreamDouble, f.Self, f.Limit.AsDouble(), f.Inclusive.Value, f.Clamp.Value), nil
}
