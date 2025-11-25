package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/ftypes"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiStreamNotBetweenMethod struct {
	/**
	const(1).not_between(1, 1, 1, 1, 1): TypeError: <'not_between function'> takes exactly 4 arguments (5 given)
	const(1).not_between()             : <'not_between function'> missing required argument 'low_limit'

	const(1).not_between(None, 10, True, True)                  : <'not_between function'> unsupported type 'none' object for argument 'low_limit'
	const(1).not_between(1, 10, True, True)                     :
	const(1).not_between(1.1, 10, True, True)                   :
	const(1).not_between('1', 10, True, True)                   : <'not_between function'> unsupported type 'string' object for argument 'low_limit'
	const(1).not_between(const(1), 10, True, True)              : <'not_between function'> unsupported type <stream of DOUBLE> for argument 'low_limit'
	const(1).not_between(data('x'), 10, True, True)             : <'not_between function'> unsupported type <stream of DOUBLE> for argument 'low_limit'
	const(1).not_between(duration('1m'), 10, True, True)        : <'not_between function'> unsupported type 'duration' object for argument 'low_limit'
	const(1).not_between(True, 10, True, True)                  : <'not_between function'> unsupported type 'boolean' object for argument 'low_limit'
	const(1).not_between(False, 10, True, True)                 : <'not_between function'> unsupported type 'boolean' object for argument 'low_limit'
	const(1).not_between([], 10, True, True)                    : <'not_between function'> unsupported type 'list' object for argument 'low_limit'
	const(1).not_between(['1'], 10, True, True)                 : <'not_between function'> unsupported type 'list' object for argument 'low_limit'
	const(1).not_between((), 10, True, True)                    : <'not_between function'> unsupported type 'tuple' object for argument 'low_limit'
	const(1).not_between(('1',), 10, True, True)                : <'not_between function'> unsupported type 'tuple' object for argument 'low_limit'
	const(1).not_between({}, 10, True, True)                    : <'not_between function'> unsupported type 'dict' object for argument 'low_limit'
	const(1).not_between({'x': 'x'}, 10, True, True)            : <'not_between function'> unsupported type 'dict' object for argument 'low_limit'
	const(1).not_between({'x': ['x']}, 10, True, True)          : <'not_between function'> unsupported type 'dict' object for argument 'low_limit'
	const(1).not_between({'x': 1}, 10, True, True)              : <'not_between function'> unsupported type 'dict' object for argument 'low_limit'
	const(1).not_between({'x': [1]}, 10, True, True)            : <'not_between function'> unsupported type 'dict' object for argument 'low_limit'

	const(1).not_between(1)                                     : <'not_between function'> missing required argument 'high_limit'
	const(1).not_between(1.1)                                   : <'not_between function'> missing required argument 'high_limit'

	const(1).not_between(0, None)                               : <'not_between function'> unsupported type 'none' object for argument 'high_limit'
	const(1).not_between(0, 1)                                  :
	const(1).not_between(0, 1.1)                                :
	const(1).not_between(0, '1')                                : <'not_between function'> unsupported type 'string' object for argument 'high_limit'
	const(1).not_between(0, const(1))                           : <'not_between function'> unsupported type <stream of DOUBLE> for argument 'high_limit'
	const(1).not_between(0, data('x'))                          : <'not_between function'> unsupported type <stream of DOUBLE> for argument 'high_limit'
	const(1).not_between(0, duration('1m'))                     : <'not_between function'> unsupported type 'duration' object for argument 'high_limit'
	const(1).not_between(0, True)                               : <'not_between function'> unsupported type 'boolean' object for argument 'high_limit'
	const(1).not_between(0, False)                              : <'not_between function'> unsupported type 'boolean' object for argument 'high_limit'
	const(1).not_between(0, [])                                 : <'not_between function'> unsupported type 'list' object for argument 'high_limit'
	const(1).not_between(0, ['1'])                              : <'not_between function'> unsupported type 'list' object for argument 'high_limit'
	const(1).not_between(0, ())                                 : <'not_between function'> unsupported type 'tuple' object for argument 'high_limit'
	const(1).not_between(0, ('1',))                             : <'not_between function'> unsupported type 'tuple' object for argument 'high_limit'
	const(1).not_between(0, {})                                 : <'not_between function'> unsupported type 'dict' object for argument 'high_limit'
	const(1).not_between(0, {'x': 'x'})                         : <'not_between function'> unsupported type 'dict' object for argument 'high_limit'
	const(1).not_between(0, {'x': ['x']})                       : <'not_between function'> unsupported type 'dict' object for argument 'high_limit'
	const(1).not_between(0, {'x': 1})                           : <'not_between function'> unsupported type 'dict' object for argument 'high_limit'
	const(1).not_between(0, {'x': [1]})                         : <'not_between function'> unsupported type 'dict' object for argument 'high_limit'

	const(1).not_between(0, 10, None)                           :
	const(1).not_between(0, 10, 1)                              : Error executing SignalFlow program (ref: G6m0ewpAwAA): [error parsing program const(1).not_between(0, 10, 1): expected boolean argument, got 1L]
	const(1).not_between(0, 10, 1.1)                            : Error executing SignalFlow program (ref: G6m0e0uA0AA): [error parsing program const(1).not_between(0, 10, 1.1): expected boolean argument, got 1.100000d]
	const(1).not_between(0, 10, '1')                            : Error executing SignalFlow program (ref: G6m0hDVA4AA): [error parsing program const(1).not_between(0, 10, '1'): expected boolean argument, got "1"]
	const(1).not_between(0, 10, const(1))                       : <'not_between function'> unsupported type <stream of DOUBLE> for argument 'low_inclusive'
	const(1).not_between(0, 10, data('x'))                      : <'not_between function'> unsupported type <stream of DOUBLE> for argument 'low_inclusive'
	const(1).not_between(0, 10, duration('1m'))                 : Error executing SignalFlow program (ref: G6m0hs0A4AA): [error parsing program const(1).not_between(0, 10, duration('1m')): expected boolean argument, got 1m]
	const(1).not_between(0, 10, True)                           :
	const(1).not_between(0, 10, False)                          :
	const(1).not_between(0, 10, [])                             : <'not_between function'> unsupported type 'list' object for argument 'low_inclusive'
	const(1).not_between(0, 10, ['1'])                          : <'not_between function'> unsupported type 'list' object for argument 'low_inclusive'
	const(1).not_between(0, 10, ())                             : <'not_between function'> unsupported type 'tuple' object for argument 'low_inclusive'
	const(1).not_between(0, 10, ('1',))                         : <'not_between function'> unsupported type 'tuple' object for argument 'low_inclusive'
	const(1).not_between(0, 10, {})                             : <'not_between function'> unsupported type 'dict' object for argument 'low_inclusive'
	const(1).not_between(0, 10, {'x': 'x'})                     : <'not_between function'> unsupported type 'dict' object for argument 'low_inclusive'
	const(1).not_between(0, 10, {'x': ['x']})                   : <'not_between function'> unsupported type 'dict' object for argument 'low_inclusive'
	const(1).not_between(0, 10, {'x': 1})                       : <'not_between function'> unsupported type 'dict' object for argument 'low_inclusive'
	const(1).not_between(0, 10, {'x': [1]})                     : <'not_between function'> unsupported type 'dict' object for argument 'low_inclusive'

	Notes: - 4 parameters
	       - first 2 are low_limit/high_limit and must be int or double
	       - second 2 are low_inclusive/high_inclusive and may be bool or None.  Default is presumably None, which resolves to actual default.
	*/

	Self          Stream                                    `ffi:"self"`
	LowLimit      ftypes.IntOrDouble                        `ffi:"low_limit"`
	HighLimit     ftypes.IntOrDouble                        `ffi:"high_limit"`
	LowInclusive  ftypes.ThingOrNone[*primitive.ObjectBool] `ffi:"low_inclusive"`
	HighInclusive ftypes.ThingOrNone[*primitive.ObjectBool] `ffi:"high_inclusive"`
}

func NewFFIStreamNotBetweenMethod() itypes.FlowCall {
	return ffi.NewFFI(ffiStreamNotBetweenMethod{
		LowInclusive:  ftypes.NewThingOrNoneNone[*primitive.ObjectBool](),
		HighInclusive: ftypes.NewThingOrNoneNone[*primitive.ObjectBool](),
	})
}

func (f ffiStreamNotBetweenMethod) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewStreamMethodNotBetween(
		prototypeStreamDouble,
		f.Self,
		f.LowLimit.AsDouble(),
		f.HighLimit.AsDouble(),
		f.LowInclusive.Thing != nil && f.LowInclusive.Thing.Value,
		f.HighInclusive.Thing != nil && f.HighInclusive.Thing.Value,
	), nil
}
