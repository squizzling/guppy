package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type ffiStreamMap struct {
	/**
	const(1).map(None)                                          : <'map function'> unsupported type 'none' object for argument 'mapfn'
	const(1).map(1)                                             :
	const(1).map(1.1)                                           :
	const(1).map('1')                                           :
	const(1).map(const(1))                                      : <'map function'> unsupported type <stream of DOUBLE> for argument 'mapfn'
	const(1).map(data('x'))                                     : <'map function'> unsupported type <stream of DOUBLE> for argument 'mapfn'
	const(1).map(duration('1m'))                                : <'map function'> unsupported type 'duration' object for argument 'mapfn'
	const(1).map(True)                                          :
	const(1).map(False)                                         :
	const(1).map([])                                            : <'map function'> unsupported type 'list' object for argument 'mapfn'
	const(1).map(['1'])                                         : <'map function'> unsupported type 'list' object for argument 'mapfn'
	const(1).map([1])                                           : <'map function'> unsupported type 'list' object for argument 'mapfn'
	const(1).map(())                                            : <'map function'> unsupported type 'tuple' object for argument 'mapfn'
	const(1).map(('1',))                                        : <'map function'> unsupported type 'tuple' object for argument 'mapfn'
	const(1).map(('1',1))                                       : <'map function'> unsupported type 'tuple' object for argument 'mapfn'
	const(1).map({})                                            : <'map function'> unsupported type 'dict' object for argument 'mapfn'
	const(1).map({'x': 'x'})                                    : <'map function'> unsupported type 'dict' object for argument 'mapfn'
	const(1).map({'x': ['x']})                                  : <'map function'> unsupported type 'dict' object for argument 'mapfn'
	const(1).map({'x': 1})                                      : <'map function'> unsupported type 'dict' object for argument 'mapfn'
	const(1).map({'x': [1]})                                    : <'map function'> unsupported type 'dict' object for argument 'mapfn'
	const(1).map(lambda x: x)                                   :

	- Passing an int or float will yield an int or float.
	- Passing a string will coerce in to a boolean
	- Passing a boolean will return a boolean
	- Passing a lambda will evaluate the lambda as:

	data().map(lambda x: ...)

	Becomes:

	def anon(x):
	  return ...

	data().map(anon)

	Becomes:

	anon(data())

	This is not strictly how it works, but we need some sort of model, and this is the first pass.

	Note: You can't pass an actual function, only a lambda.
	*/

	Self  Stream `ffi:"self"`
	MapFn struct {
		Lambda *interpreter.ObjectLambda
		/*Int    *primitive.ObjectInt
		String *primitive.ObjectString
		Double *primitive.ObjectDouble
		Bool   *primitive.ObjectBool*/
	} `ffi:"mapfn"`
}

func NewFFIStreamMap() itypes.FlowCall {
	return ffi.NewFFI(ffiStreamMap{})
}

func (f ffiStreamMap) Call(i itypes.Interpreter) (itypes.Object, error) {
	if f.MapFn.Lambda != nil {
		i.PushIntrinsicScope()
		defer i.PopScope()

		if err := i.Set(f.MapFn.Lambda.Identifier, f.Self); err != nil {
			return nil, err
		} else if o, err := f.MapFn.Lambda.Expression.Accept(i); err != nil {
			return nil, err
		} else {
			return o.(itypes.Object), err
		}
	} else {
		panic("unreachable")
	}
}
