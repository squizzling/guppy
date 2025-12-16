package stream

import (
	"math"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiSqrt struct {
	/**
	sqrt(None)               : unexpected argument 1 to function 'sqrt', expected numeric object, got 'none' object
	sqrt(1)                  :
	sqrt(1.1)                :
	sqrt('1')                : unexpected argument 1 to function 'sqrt', expected numeric object, got 'string' object
	sqrt(const(1))           :
	sqrt(const(1) > const(2)):
	sqrt(data('x'))          :
	sqrt(duration('1m'))     : invalid argument 'sf.analytics.program.parsing.types.DurationType@65a7ee86' to function
	sqrt(True)               : unexpected argument 1 to function 'sqrt', expected numeric object, got 'boolean' object
	sqrt(False)              : unexpected argument 1 to function 'sqrt', expected numeric object, got 'boolean' object
	sqrt([])                 : invalid argument 'sf.analytics.program.parsing.types.ListType@2cf5a1b6' to function
	sqrt(['1'])              : invalid argument 'sf.analytics.program.parsing.types.ListType@580d6e73' to function
	sqrt([1])                : invalid argument 'sf.analytics.program.parsing.types.ListType@197cb0bf' to function
	sqrt(())                 : invalid argument 'sf.analytics.program.parsing.types.TupleType@5029eeb6' to function
	sqrt(('1',))             : invalid argument 'sf.analytics.program.parsing.types.TupleType@75c8eec5' to function
	sqrt(('1',1))            : invalid argument 'sf.analytics.program.parsing.types.TupleType@727095a0' to function
	sqrt({})                 : invalid argument 'sf.analytics.program.parsing.types.DictionaryType@7b5f79f5' to function
	sqrt({'x': 'x'})         : invalid argument 'sf.analytics.program.parsing.types.DictionaryType@59ea4a44' to function
	sqrt({'x': ['x']})       : invalid argument 'sf.analytics.program.parsing.types.DictionaryType@1610183b' to function
	sqrt({'x': 1})           : invalid argument 'sf.analytics.program.parsing.types.DictionaryType@cba446e' to function
	sqrt({'x': [1]})         : invalid argument 'sf.analytics.program.parsing.types.DictionaryType@6e1e84bf' to function
	sqrt(lambda x: x)        : invalid argument '<lambda function>' to function

	Notes:
	- We're going to be lazy and assume this is it.
	*/
	Input struct {
		Stream Stream
		Int    *primitive.ObjectInt
		Double *primitive.ObjectDouble
	} `ffi:"input"`
}

func NewFFISqrt() itypes.FlowCall {
	return ffi.NewFFI(ffiSqrt{})
}

func (f ffiSqrt) Call(i itypes.Interpreter) (itypes.Object, error) {
	if f.Input.Int != nil {
		return f.Input.Int, nil
	} else if f.Input.Double != nil {
		return primitive.NewObjectDouble(math.Sqrt(f.Input.Double.Value)), nil
	} else {
		return NewStreamFuncSqrt(prototypeStreamDouble, f.Input.Stream), nil
	}
}
