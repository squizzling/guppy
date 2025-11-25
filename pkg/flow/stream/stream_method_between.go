package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/ftypes"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiStreamBetweenMethod struct {
	/**
	const(1).between(1, 1, 1, 1, 1, 1): TypeError: <'between function'> takes exactly 5 arguments (6 given)
	const(1).between()             : <'between function'> missing required argument 'low_limit'

	const(1).between(None, 10, True, True)                  : <'between function'> unsupported type 'none' object for argument 'low_limit'
	const(1).between(1, 10, True, True)                     :
	const(1).between(1.1, 10, True, True)                   :
	const(1).between('1', 10, True, True)                   : <'between function'> unsupported type 'string' object for argument 'low_limit'
	const(1).between(const(1), 10, True, True)              : <'between function'> unsupported type <stream of DOUBLE> for argument 'low_limit'
	const(1).between(data('x'), 10, True, True)             : <'between function'> unsupported type <stream of DOUBLE> for argument 'low_limit'
	const(1).between(duration('1m'), 10, True, True)        : <'between function'> unsupported type 'duration' object for argument 'low_limit'
	const(1).between(True, 10, True, True)                  : <'between function'> unsupported type 'boolean' object for argument 'low_limit'
	const(1).between(False, 10, True, True)                 : <'between function'> unsupported type 'boolean' object for argument 'low_limit'
	const(1).between([], 10, True, True)                    : <'between function'> unsupported type 'list' object for argument 'low_limit'
	const(1).between(['1'], 10, True, True)                 : <'between function'> unsupported type 'list' object for argument 'low_limit'
	const(1).between((), 10, True, True)                    : <'between function'> unsupported type 'tuple' object for argument 'low_limit'
	const(1).between(('1',), 10, True, True)                : <'between function'> unsupported type 'tuple' object for argument 'low_limit'
	const(1).between({}, 10, True, True)                    : <'between function'> unsupported type 'dict' object for argument 'low_limit'
	const(1).between({'x': 'x'}, 10, True, True)            : <'between function'> unsupported type 'dict' object for argument 'low_limit'
	const(1).between({'x': ['x']}, 10, True, True)          : <'between function'> unsupported type 'dict' object for argument 'low_limit'
	const(1).between({'x': 1}, 10, True, True)              : <'between function'> unsupported type 'dict' object for argument 'low_limit'
	const(1).between({'x': [1]}, 10, True, True)            : <'between function'> unsupported type 'dict' object for argument 'low_limit'

	const(1).between(1)                                     : <'between function'> missing required argument 'high_limit'
	const(1).between(1.1)                                   : <'between function'> missing required argument 'high_limit'

	const(1).between(0, None)                               : <'between function'> unsupported type 'none' object for argument 'high_limit'
	const(1).between(0, 1)                                  :
	const(1).between(0, 1.1)                                :
	const(1).between(0, '1')                                : <'between function'> unsupported type 'string' object for argument 'high_limit'
	const(1).between(0, const(1))                           : <'between function'> unsupported type <stream of DOUBLE> for argument 'high_limit'
	const(1).between(0, data('x'))                          : <'between function'> unsupported type <stream of DOUBLE> for argument 'high_limit'
	const(1).between(0, duration('1m'))                     : <'between function'> unsupported type 'duration' object for argument 'high_limit'
	const(1).between(0, True)                               : <'between function'> unsupported type 'boolean' object for argument 'high_limit'
	const(1).between(0, False)                              : <'between function'> unsupported type 'boolean' object for argument 'high_limit'
	const(1).between(0, [])                                 : <'between function'> unsupported type 'list' object for argument 'high_limit'
	const(1).between(0, ['1'])                              : <'between function'> unsupported type 'list' object for argument 'high_limit'
	const(1).between(0, ())                                 : <'between function'> unsupported type 'tuple' object for argument 'high_limit'
	const(1).between(0, ('1',))                             : <'between function'> unsupported type 'tuple' object for argument 'high_limit'
	const(1).between(0, {})                                 : <'between function'> unsupported type 'dict' object for argument 'high_limit'
	const(1).between(0, {'x': 'x'})                         : <'between function'> unsupported type 'dict' object for argument 'high_limit'
	const(1).between(0, {'x': ['x']})                       : <'between function'> unsupported type 'dict' object for argument 'high_limit'
	const(1).between(0, {'x': 1})                           : <'between function'> unsupported type 'dict' object for argument 'high_limit'
	const(1).between(0, {'x': [1]})                         : <'between function'> unsupported type 'dict' object for argument 'high_limit'

	const(1).between(0, 10, None)                           :
	const(1).between(0, 10, 1)                              : Error executing SignalFlow program (ref: G6m0ewpAwAA): [error parsing program const(1).between(0, 10, 1): expected boolean argument, got 1L]
	const(1).between(0, 10, 1.1)                            : Error executing SignalFlow program (ref: G6m0e0uA0AA): [error parsing program const(1).between(0, 10, 1.1): expected boolean argument, got 1.100000d]
	const(1).between(0, 10, '1')                            : Error executing SignalFlow program (ref: G6m0hDVA4AA): [error parsing program const(1).between(0, 10, '1'): expected boolean argument, got "1"]
	const(1).between(0, 10, const(1))                       : <'between function'> unsupported type <stream of DOUBLE> for argument 'low_inclusive'
	const(1).between(0, 10, data('x'))                      : <'between function'> unsupported type <stream of DOUBLE> for argument 'low_inclusive'
	const(1).between(0, 10, duration('1m'))                 : Error executing SignalFlow program (ref: G6m0hs0A4AA): [error parsing program const(1).between(0, 10, duration('1m')): expected boolean argument, got 1m]
	const(1).between(0, 10, True)                           :
	const(1).between(0, 10, False)                          :
	const(1).between(0, 10, [])                             : <'between function'> unsupported type 'list' object for argument 'low_inclusive'
	const(1).between(0, 10, ['1'])                          : <'between function'> unsupported type 'list' object for argument 'low_inclusive'
	const(1).between(0, 10, ())                             : <'between function'> unsupported type 'tuple' object for argument 'low_inclusive'
	const(1).between(0, 10, ('1',))                         : <'between function'> unsupported type 'tuple' object for argument 'low_inclusive'
	const(1).between(0, 10, {})                             : <'between function'> unsupported type 'dict' object for argument 'low_inclusive'
	const(1).between(0, 10, {'x': 'x'})                     : <'between function'> unsupported type 'dict' object for argument 'low_inclusive'
	const(1).between(0, 10, {'x': ['x']})                   : <'between function'> unsupported type 'dict' object for argument 'low_inclusive'
	const(1).between(0, 10, {'x': 1})                       : <'between function'> unsupported type 'dict' object for argument 'low_inclusive'
	const(1).between(0, 10, {'x': [1]})                     : <'between function'> unsupported type 'dict' object for argument 'low_inclusive'

	const(1).between(0, 10, None, None, None)                   : Error executing SignalFlow program (ref: G6m4Z7DA4AA): [error parsing program const(1).between(0, 10, None, None, None): expected boolean argument, got <null>]
	const(1).between(0, 10, None, None, 1)                      : Error executing SignalFlow program (ref: G6m4ayMAwAE): [error parsing program const(1).between(0, 10, None, None, 1): expected boolean argument, got 1L]
	const(1).between(0, 10, None, None, 1.1)                    : Error executing SignalFlow program (ref: G6m4bWpA0AE): [error parsing program const(1).between(0, 10, None, None, 1.1): expected boolean argument, got 1.100000d]
	const(1).between(0, 10, None, None, '1')                    : Error executing SignalFlow program (ref: G6m4by7A4AE): [error parsing program const(1).between(0, 10, None, None, '1'): expected boolean argument, got "1"]
	const(1).between(0, 10, None, None, const(1))               : <'between function'> unsupported type <stream of DOUBLE> for argument 'clamp'
	const(1).between(0, 10, None, None, data('x'))              : <'between function'> unsupported type <stream of DOUBLE> for argument 'clamp'
	const(1).between(0, 10, None, None, duration('1m'))         : Error executing SignalFlow program (ref: G6m4cS_A4AA): [error parsing program const(1).between(0, 10, None, None, duration('1m')): expected boolean argument, got 1m]
	const(1).between(0, 10, None, None, True)                   :
	const(1).between(0, 10, None, None, False)                  :
	const(1).between(0, 10, None, None, [])                     : <'between function'> unsupported type 'list' object for argument 'clamp'
	const(1).between(0, 10, None, None, ['1'])                  : <'between function'> unsupported type 'list' object for argument 'clamp'
	const(1).between(0, 10, None, None, ())                     : <'between function'> unsupported type 'tuple' object for argument 'clamp'
	const(1).between(0, 10, None, None, ('1',))                 : <'between function'> unsupported type 'tuple' object for argument 'clamp'
	const(1).between(0, 10, None, None, {})                     : <'between function'> unsupported type 'dict' object for argument 'clamp'
	const(1).between(0, 10, None, None, {'x': 'x'})             : <'between function'> unsupported type 'dict' object for argument 'clamp'
	const(1).between(0, 10, None, None, {'x': ['x']})           : <'between function'> unsupported type 'dict' object for argument 'clamp'
	const(1).between(0, 10, None, None, {'x': 1})               : <'between function'> unsupported type 'dict' object for argument 'clamp'
	const(1).between(0, 10, None, None, {'x': [1]})             : <'between function'> unsupported type 'dict' object for argument 'clamp'

	Notes: - 5 parameters
	       - first 2 are low_limit/high_limit and must be int or double
	       - second 2 are low_inclusive/high_inclusive and may be bool or None.  Default is presumably None, which resolves to actual default.
	       - final 1 is clamp, which must be bool, defaults to False
	*/

	Self          Stream                                    `ffi:"self"`
	LowLimit      ftypes.IntOrDouble                        `ffi:"low_limit"`
	HighLimit     ftypes.IntOrDouble                        `ffi:"high_limit"`
	LowInclusive  ftypes.ThingOrNone[*primitive.ObjectBool] `ffi:"low_inclusive"`
	HighInclusive ftypes.ThingOrNone[*primitive.ObjectBool] `ffi:"high_inclusive"`
	Clamp         *primitive.ObjectBool                     `ffi:"clamp"`
}

func NewFFIStreamBetweenMethod() itypes.FlowCall {
	return ffi.NewFFI(ffiStreamBetweenMethod{
		LowInclusive:  ftypes.NewThingOrNoneNone[*primitive.ObjectBool](),
		HighInclusive: ftypes.NewThingOrNoneNone[*primitive.ObjectBool](),
		Clamp:         primitive.NewObjectBool(false),
	})
}

func (f ffiStreamBetweenMethod) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewStreamMethodBetween(
		prototypeStreamDouble,
		f.Self,
		f.LowLimit.AsDouble(),
		f.HighLimit.AsDouble(),
		f.LowInclusive.Thing != nil && f.LowInclusive.Thing.Value,
		f.HighInclusive.Thing != nil && f.HighInclusive.Thing.Value,
		f.Clamp.Value,
	), nil
}
