package stream

import (
	"math"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiCeil struct {
	/**
	Notes:
	- If there is a single argument, it must be a stream, or an int/double
	- If there are multiple arguments, there must be at least one stream, and any number of int/doubles.

	This is hard to express, so as a first pass, we won't try.  We'll accept a single stream, int, or double, and see
	what blows up.

	TODO: The parameter does not have a name, it behaves like a *args.  Should we allow anonymous parameters in the
	      ffi?  We technically don't require a name.
	*/
	Input struct {
		Stream Stream
		Int    *primitive.ObjectInt
		Double *primitive.ObjectDouble
	} `ffi:"input"`
}

func NewFFICeil() itypes.FlowCall {
	return ffi.NewFFI(ffiCeil{})
}

func (f ffiCeil) Call(i itypes.Interpreter) (itypes.Object, error) {
	if f.Input.Int != nil {
		return f.Input.Int, nil
	} else if f.Input.Double != nil {
		return primitive.NewObjectDouble(math.Ceil(f.Input.Double.Value)), nil
	} else {
		return NewStreamFuncCeil(prototypeStreamDouble, f.Input.Stream), nil
	}
}
