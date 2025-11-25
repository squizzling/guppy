package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiThreshold struct {
	/**
	threshold(None)          : <'threshold function'> unsupported type 'none' object for argument 'object'
	threshold(1)             :
	threshold(1.1)           :
	threshold('1')           : <'threshold function'> unsupported type 'string' object for argument 'object'
	threshold(const(1))      :
	threshold(data('x'))     : Threshold expressions cannot contain stream operators
	threshold(duration('1m')): <'threshold function'> unsupported type 'duration' object for argument 'object'
	threshold(True)          : <'threshold function'> unsupported type 'boolean' object for argument 'object'
	threshold(False)         : <'threshold function'> unsupported type 'boolean' object for argument 'object'
	threshold([])            : <'threshold function'> unsupported type 'list' object for argument 'object'
	threshold(['1'])         : <'threshold function'> unsupported type 'list' object for argument 'object'
	threshold(())            : <'threshold function'> unsupported type 'tuple' object for argument 'object'
	threshold(('1',))        : <'threshold function'> unsupported type 'tuple' object for argument 'object'
	threshold({})            : <'threshold function'> unsupported type 'dict' object for argument 'object'
	threshold({'x': 'x'})    : <'threshold function'> unsupported type 'dict' object for argument 'object'
	threshold({'x': ['x']})  : <'threshold function'> unsupported type 'dict' object for argument 'object'
	threshold({'x': 1})      : <'threshold function'> unsupported type 'dict' object for argument 'object'
	threshold({'x': [1]})    : <'threshold function'> unsupported type 'dict' object for argument 'object'

	A=data('x');threshold(A) :
	A=threshold(1);A.sum()   :
	threshold(1).sum()       : <threshold stream of DOUBLE> has no attribute 'sum'

	threshold(None, None)    : <'threshold function'> takes exactly 1 arguments (2 given)

	Supported: int, double, sort of stream
	Notes: There is a corner case where threshold can take a Stream, but only if the Stream was assigned to
	       a variable first.  There is also no Stream methods on a threshold stream, until after the
	       ThresholdStream is assigned to a variable.  We will ignore all this, accept a Stream, and
	       treat it as a Stream of double regardless of input type.
	*/

	Object struct {
		Double *primitive.ObjectDouble
		Int    *primitive.ObjectInt
		Stream Stream
	} `ffi:"object"`
}

func NewFFIThreshold() itypes.FlowCall {
	return ffi.NewFFI(ffiThreshold{})
}

func (f ffiThreshold) Call(i itypes.Interpreter) (itypes.Object, error) {
	if f.Object.Int != nil {
		return NewStreamFuncThresholdDouble(prototypeStreamDouble, float64(f.Object.Int.Value)), nil
	} else if f.Object.Double != nil {
		return NewStreamFuncThresholdDouble(prototypeStreamDouble, f.Object.Double.Value), nil
	} else {
		return NewStreamFuncThresholdStream(prototypeStreamDouble, f.Object.Stream), nil
	}
}
