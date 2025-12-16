package stream

import (
	"errors"
	"fmt"
	"time"

	"github.com/squizzling/guppy/pkg/flow/duration"
	"github.com/squizzling/guppy/pkg/flow/filter"
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/ftypes"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

/**
graphite(None)                                              : <'graphite function'> unsupported type 'none' object for argument 'metric'
graphite(1)                                                 : <'graphite function'> unsupported type 'long' object for argument 'metric'
graphite(1.1)                                               : <'graphite function'> unsupported type 'double' object for argument 'metric'
graphite('1')                                               :
graphite(const(1))                                          : <'graphite function'> unsupported type <stream of DOUBLE> for argument 'metric'
graphite(data('x'))                                         : <'graphite function'> unsupported type <stream of DOUBLE> for argument 'metric'
graphite(duration('1m'))                                    : <'graphite function'> unsupported type 'duration' object for argument 'metric'
graphite(True)                                              : <'graphite function'> unsupported type 'boolean' object for argument 'metric'
graphite(False)                                             : <'graphite function'> unsupported type 'boolean' object for argument 'metric'
graphite([])                                                : <'graphite function'> unsupported type 'list' object for argument 'metric'
graphite(['1'])                                             : <'graphite function'> unsupported type 'list' object for argument 'metric'
graphite([1])                                               : <'graphite function'> unsupported type 'list' object for argument 'metric'
graphite(())                                                : <'graphite function'> unsupported type 'tuple' object for argument 'metric'
graphite(('1',))                                            : <'graphite function'> unsupported type 'tuple' object for argument 'metric'
graphite(('1',1))                                           : <'graphite function'> unsupported type 'tuple' object for argument 'metric'
graphite({})                                                : <'graphite function'> unsupported type 'dict' object for argument 'metric'
graphite({'x': 'x'})                                        : <'graphite function'> unsupported type 'dict' object for argument 'metric'
graphite({'x': ['x']})                                      : <'graphite function'> unsupported type 'dict' object for argument 'metric'
graphite({'x': 1})                                          : <'graphite function'> unsupported type 'dict' object for argument 'metric'
graphite({'x': [1]})                                        : <'graphite function'> unsupported type 'dict' object for argument 'metric'
graphite(lambda x: x)                                       : <'graphite function'> unsupported type <lambda function> for argument 'metric'

graphite('metric', None)                                    :
graphite('metric', 1)                                       : <'graphite function'> unsupported type 'long' object for argument 'filter'
graphite('metric', 1.1)                                     : <'graphite function'> unsupported type 'double' object for argument 'filter'
graphite('metric', '1')                                     : <'graphite function'> unsupported type 'string' object for argument 'filter'
graphite('metric', const(1))                                : <'graphite function'> unsupported type <stream of DOUBLE> for argument 'filter'
graphite('metric', data('x'))                               : <'graphite function'> unsupported type <stream of DOUBLE> for argument 'filter'
graphite('metric', duration('1m'))                          : <'graphite function'> unsupported type 'duration' object for argument 'filter'
graphite('metric', True)                                    : <'graphite function'> unsupported type 'boolean' object for argument 'filter'
graphite('metric', False)                                   : <'graphite function'> unsupported type 'boolean' object for argument 'filter'
graphite('metric', [])                                      : <'graphite function'> unsupported type 'list' object for argument 'filter'
graphite('metric', ['1'])                                   : <'graphite function'> unsupported type 'list' object for argument 'filter'
graphite('metric', [1])                                     : <'graphite function'> unsupported type 'list' object for argument 'filter'
graphite('metric', ())                                      : <'graphite function'> unsupported type 'tuple' object for argument 'filter'
graphite('metric', ('1',))                                  : <'graphite function'> unsupported type 'tuple' object for argument 'filter'
graphite('metric', ('1',1))                                 : <'graphite function'> unsupported type 'tuple' object for argument 'filter'
graphite('metric', {})                                      : <'graphite function'> unsupported type 'dict' object for argument 'filter'
graphite('metric', {'x': 'x'})                              : <'graphite function'> unsupported type 'dict' object for argument 'filter'
graphite('metric', {'x': ['x']})                            : <'graphite function'> unsupported type 'dict' object for argument 'filter'
graphite('metric', {'x': 1})                                : <'graphite function'> unsupported type 'dict' object for argument 'filter'
graphite('metric', {'x': [1]})                              : <'graphite function'> unsupported type 'dict' object for argument 'filter'
graphite('metric', lambda x: x)                             : <'graphite function'> unsupported type <lambda function> for argument 'filter'

graphite('metric', None, None)                              : <'graphite function'> unsupported type 'none' object for argument '_sf_delimiter'
graphite('metric', None, 1)                                 : <'graphite function'> unsupported type 'long' object for argument '_sf_delimiter'
graphite('metric', None, 1.1)                               : <'graphite function'> unsupported type 'double' object for argument '_sf_delimiter'
graphite('metric', None, '1')                               : Error executing SignalFlow program (ref: G8OhcAYAwAA): [parameter '_sf_delimiter' assigned multiple values: "1" and "."]
graphite('metric', None, const(1))                          : <'graphite function'> unsupported type <stream of DOUBLE> for argument '_sf_delimiter'
graphite('metric', None, data('x'))                         : <'graphite function'> unsupported type <stream of DOUBLE> for argument '_sf_delimiter'
graphite('metric', None, duration('1m'))                    : <'graphite function'> unsupported type 'duration' object for argument '_sf_delimiter'
graphite('metric', None, True)                              : <'graphite function'> unsupported type 'boolean' object for argument '_sf_delimiter'
graphite('metric', None, False)                             : <'graphite function'> unsupported type 'boolean' object for argument '_sf_delimiter'
graphite('metric', None, [])                                : <'graphite function'> unsupported type 'list' object for argument '_sf_delimiter'
graphite('metric', None, ['1'])                             : <'graphite function'> unsupported type 'list' object for argument '_sf_delimiter'
graphite('metric', None, [1])                               : <'graphite function'> unsupported type 'list' object for argument '_sf_delimiter'
graphite('metric', None, ())                                : <'graphite function'> unsupported type 'tuple' object for argument '_sf_delimiter'
graphite('metric', None, ('1',))                            : <'graphite function'> unsupported type 'tuple' object for argument '_sf_delimiter'
graphite('metric', None, ('1',1))                           : <'graphite function'> unsupported type 'tuple' object for argument '_sf_delimiter'
graphite('metric', None, {})                                : <'graphite function'> unsupported type 'dict' object for argument '_sf_delimiter'
graphite('metric', None, {'x': 'x'})                        : <'graphite function'> unsupported type 'dict' object for argument '_sf_delimiter'
graphite('metric', None, {'x': ['x']})                      : <'graphite function'> unsupported type 'dict' object for argument '_sf_delimiter'
graphite('metric', None, {'x': 1})                          : <'graphite function'> unsupported type 'dict' object for argument '_sf_delimiter'
graphite('metric', None, {'x': [1]})                        : <'graphite function'> unsupported type 'dict' object for argument '_sf_delimiter'
graphite('metric', None, lambda x: x)                       : <'graphite function'> unsupported type <lambda function> for argument '_sf_delimiter'

graphite('metric', None, '.', None)                         : Error executing SignalFlow program (ref: G8Oh5TZA0AA): [block : argument offset expects ['duration' object] value, got 'none' object, File "", line 1, in \ngraphite('metric', None, '.', None)\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', 1)                            : <'graphite function'> takes exactly 8 arguments (4 given)
graphite('metric', None, '.', 1.1)                          : Error executing SignalFlow program (ref: G8Oh5TZA0AI): [block : argument offset expects ['duration' object] value, got 'double' object, File "", line 1, in \ngraphite('metric', None, '.', 1.1)\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', '1')                          : Error executing SignalFlow program (ref: G8Oh5TZA0AM): [error parsing program graphite('metric', None, '.', '1'): invalid duration string 1]
graphite('metric', None, '.', const(1))                     : Error executing SignalFlow program (ref: G8Oh5TZA0AQ): [block : argument offset expects ['duration' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', const(1))\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', data('x'))                    : Error executing SignalFlow program (ref: G8Oh5TZA0AU): [block : argument offset expects ['duration' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', data('x'))\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', duration('1m'))               : <'graphite function'> takes exactly 8 arguments (4 given)
graphite('metric', None, '.', True)                         : Error executing SignalFlow program (ref: G8Oh6GhA0AA): [block : argument offset expects ['duration' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', True)\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', False)                        : Error executing SignalFlow program (ref: G8Oh7m6A4AA): [block : argument offset expects ['duration' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', False)\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', [])                           : Error executing SignalFlow program (ref: G8Oh7m_AwAA): [block : argument offset expects ['duration' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', [])\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', ['1'])                        : Error executing SignalFlow program (ref: G8Oh7xaA0AU): [block : argument offset expects ['duration' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', ['1'])\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', [1])                          : Error executing SignalFlow program (ref: G8Oh769A4AA): [block : argument offset expects ['duration' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', [1])\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', ())                           : Error executing SignalFlow program (ref: G8Oh_0vAwAA): [block : argument offset expects ['duration' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', ())\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', ('1',))                       : Error executing SignalFlow program (ref: G8Oh_0vAwAE): [block : argument offset expects ['duration' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', ('1',))\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', ('1',1))                      : Error executing SignalFlow program (ref: G8Oh_0vAwAI): [block : argument offset expects ['duration' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', ('1',1))\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', {})                           : Error executing SignalFlow program (ref: G8Oh_00A0AE): [block : argument offset expects ['duration' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', {})\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', {'x': 'x'})                   : Error executing SignalFlow program (ref: G8OiA8mA4AA): [block : argument offset expects ['duration' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', {'x': 'x'})\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', {'x': ['x']})                 : Error executing SignalFlow program (ref: G8OiBglAwAE): [block : argument offset expects ['duration' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', {'x': ['x']})\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', {'x': 1})                     : Error executing SignalFlow program (ref: G8OiCAfA0AA): [block : argument offset expects ['duration' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', {'x': 1})\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', {'x': [1]})                   : Error executing SignalFlow program (ref: G8OiDEaA4AA): [block : argument offset expects ['duration' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', {'x': [1]})\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]
graphite('metric', None, '.', lambda x: x)                  : Error executing SignalFlow program (ref: G8OiFIGAwBM): [block : argument offset expects ['duration' object] value, got <lambda function>, File "", line 1, in \ngraphite('metric', None, '.', lambda x: x)\nTypeError: <'graphite function'> takes exactly 8 arguments (4 given)]

graphite('metric', None, '.', 0, None)                      : <'graphite function'> takes exactly 8 arguments (5 given)
graphite('metric', None, '.', 0, 1)                         : Error executing SignalFlow program (ref: G8Oi3xYA0AQ): [block : argument rollup expects ['string' object, 'none' object] value, got 'long' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 1)\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, 1.1)                       : Error executing SignalFlow program (ref: G8Oi3xYA0AU): [block : argument rollup expects ['string' object, 'none' object] value, got 'double' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 1.1)\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, '1')                       : <'graphite function'> takes exactly 8 arguments (5 given)
graphite('metric', None, '.', 0, const(1))                  : Error executing SignalFlow program (ref: G8Oi3xYA0Ac): [block : argument rollup expects ['string' object, 'none' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, const(1))\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, data('x'))                 : Error executing SignalFlow program (ref: G8Oi3xYA0Ag): [block : argument rollup expects ['string' object, 'none' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, data('x'))\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, duration('1m'))            : Error executing SignalFlow program (ref: G8Oi3xYA0Ak): [block : argument rollup expects ['string' object, 'none' object] value, got 'duration' object, File "", line 1, in \ngraphite('metric', None, '.', 0, duration('1m'))\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, True)                      : Error executing SignalFlow program (ref: G8Oi3xYA0Ao): [block : argument rollup expects ['string' object, 'none' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, True)\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, False)                     : Error executing SignalFlow program (ref: G8Oi3xYA0As): [block : argument rollup expects ['string' object, 'none' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, False)\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, [])                        : Error executing SignalFlow program (ref: G8Oi3xYA0Aw): [block : argument rollup expects ['string' object, 'none' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, [])\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, ['1'])                     : Error executing SignalFlow program (ref: G8Oi3xYA0A0): [block : argument rollup expects ['string' object, 'none' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, ['1'])\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, [1])                       : Error executing SignalFlow program (ref: G8Oi3xYA0A4): [block : argument rollup expects ['string' object, 'none' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, [1])\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, ())                        : Error executing SignalFlow program (ref: G8Oi3xYA0A8): [block : argument rollup expects ['string' object, 'none' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, ())\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, ('1',))                    : Error executing SignalFlow program (ref: G8Oi3xYA0BA): [block : argument rollup expects ['string' object, 'none' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, ('1',))\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, ('1',1))                   : Error executing SignalFlow program (ref: G8Oi3xYA0BE): [block : argument rollup expects ['string' object, 'none' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, ('1',1))\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, {})                        : Error executing SignalFlow program (ref: G8Oi3xYA0BI): [block : argument rollup expects ['string' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, {})\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, {'x': 'x'})                : Error executing SignalFlow program (ref: G8Oi3xYA0BM): [block : argument rollup expects ['string' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, {'x': 'x'})\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, {'x': ['x']})              : Error executing SignalFlow program (ref: G8Oi3xYA0BQ): [block : argument rollup expects ['string' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, {'x': ['x']})\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, {'x': 1})                  : Error executing SignalFlow program (ref: G8Oi3xYA0BU): [block : argument rollup expects ['string' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, {'x': 1})\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, {'x': [1]})                : Error executing SignalFlow program (ref: G8Oi3xdA4AE): [block : argument rollup expects ['string' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, {'x': [1]})\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]
graphite('metric', None, '.', 0, lambda x: x)               : Error executing SignalFlow program (ref: G8Oi37DAwAA): [block : argument rollup expects ['string' object, 'none' object] value, got <lambda function>, File "", line 1, in \ngraphite('metric', None, '.', 0, lambda x: x)\nTypeError: <'graphite function'> takes exactly 8 arguments (5 given)]

graphite('metric', None, '.', 0, 'sum', None)               : Error executing SignalFlow program (ref: G8OiuVHA4AE): [block : argument extrapolation expects ['string' object] value, got 'none' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', None)\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', 1)                  : Error executing SignalFlow program (ref: G8OiusVAwAA): [block : argument extrapolation expects ['string' object] value, got 'long' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 1)\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', 1.1)                : Error executing SignalFlow program (ref: G8OivRRA0AA): [block : argument extrapolation expects ['string' object] value, got 'double' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 1.1)\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', '1')                : <'graphite function'> takes exactly 8 arguments (6 given)
graphite('metric', None, '.', 0, 'sum', const(1))           : Error executing SignalFlow program (ref: G8Oiv-ZAwAE): [block : argument extrapolation expects ['string' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', const(1))\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', data('x'))          : Error executing SignalFlow program (ref: G8OiwA3A0AA): [block : argument extrapolation expects ['string' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', data('x'))\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', duration('1m'))     : Error executing SignalFlow program (ref: G8Oiw5zA4AM): [block : argument extrapolation expects ['string' object] value, got 'duration' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', duration('1m'))\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', True)               : Error executing SignalFlow program (ref: G8Oiw50AwAY): [block : argument extrapolation expects ['string' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', True)\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', False)              : Error executing SignalFlow program (ref: G8OixQbA0AA): [block : argument extrapolation expects ['string' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', False)\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', [])                 : Error executing SignalFlow program (ref: G8Oix_3A4AA): [block : argument extrapolation expects ['string' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', [])\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', ['1'])              : Error executing SignalFlow program (ref: G8OiydmAwAA): [block : argument extrapolation expects ['string' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', ['1'])\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', [1])                : Error executing SignalFlow program (ref: G8OizeJA0AA): [block : argument extrapolation expects ['string' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', [1])\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', ())                 : Error executing SignalFlow program (ref: G8OizeNA4AA): [block : argument extrapolation expects ['string' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', ())\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', ('1',))             : Error executing SignalFlow program (ref: G8OizmAAwAA): [block : argument extrapolation expects ['string' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', ('1',))\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', ('1',1))            : Error executing SignalFlow program (ref: G8Oi0H8A4AA): [block : argument extrapolation expects ['string' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', ('1',1))\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', {})                 : Error executing SignalFlow program (ref: G8Oi1kuAwAA): [block : argument extrapolation expects ['string' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', {})\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', {'x': 'x'})         : Error executing SignalFlow program (ref: G8Oi1kuAwAE): [block : argument extrapolation expects ['string' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', {'x': 'x'})\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', {'x': ['x']})       : Error executing SignalFlow program (ref: G8Oi1kuAwAI): [block : argument extrapolation expects ['string' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', {'x': ['x']})\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', {'x': 1})           : Error executing SignalFlow program (ref: G8Oi1kuAwAM): [block : argument extrapolation expects ['string' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', {'x': 1})\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', {'x': [1]})         : Error executing SignalFlow program (ref: G8Oi1kyA0AA): [block : argument extrapolation expects ['string' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', {'x': [1]})\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]
graphite('metric', None, '.', 0, 'sum', lambda x: x)        : Error executing SignalFlow program (ref: G8Oi1lgA4AA): [block : argument extrapolation expects ['string' object] value, got <lambda function>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', lambda x: x)\nTypeError: <'graphite function'> takes exactly 8 arguments (6 given)]

graphite('metric', None, '.', 0, 'sum', 'zero', None)                                     : Error executing SignalFlow program (ref: G8Oi4YiA4Ag): [block : argument maxExtrapolations expects ['long' object] value, got 'none' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', None)\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 1)                                        : <'graphite function'> takes exactly 8 arguments (7 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 1.1)                                      : Error executing SignalFlow program (ref: G8Oi4YiA4Ao): [block : argument maxExtrapolations expects ['long' object] value, got 'double' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 1.1)\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', '1')                                      : Error executing SignalFlow program (ref: G8Oi4YiA4As): [block : argument maxExtrapolations expects ['long' object] value, got 'string' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', '1')\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', const(1))                                 : Error executing SignalFlow program (ref: G8Oi4YiA4Aw): [block : argument maxExtrapolations expects ['long' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', const(1))\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', data('x'))                                : Error executing SignalFlow program (ref: G8Oi5aAA0AE): [block : argument maxExtrapolations expects ['long' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', data('x'))\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', duration('1m'))                           : Error executing SignalFlow program (ref: G8Oi5aBA4Ak): [block : argument maxExtrapolations expects ['long' object] value, got 'duration' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', duration('1m'))\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', True)                                     : Error executing SignalFlow program (ref: G8Oi5isAwAA): [block : argument maxExtrapolations expects ['long' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', True)\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', False)                                    : Error executing SignalFlow program (ref: G8Oi61ZA0AI): [block : argument maxExtrapolations expects ['long' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', False)\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', [])                                       : Error executing SignalFlow program (ref: G8Oi61ZA0AM): [block : argument maxExtrapolations expects ['long' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', [])\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', ['1'])                                    : Error executing SignalFlow program (ref: G8Oi61ZA0AQ): [block : argument maxExtrapolations expects ['long' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', ['1'])\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', [1])                                      : Error executing SignalFlow program (ref: G8Oi7-kA0AA): [block : argument maxExtrapolations expects ['long' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', [1])\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', ())                                       : Error executing SignalFlow program (ref: G8Oi8fyA4AA): [block : argument maxExtrapolations expects ['long' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', ())\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', ('1',))                                   : Error executing SignalFlow program (ref: G8Oi8fyA4AE): [block : argument maxExtrapolations expects ['long' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', ('1',))\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', ('1',1))                                  : Error executing SignalFlow program (ref: G8Oi8fyA4AI): [block : argument maxExtrapolations expects ['long' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', ('1',1))\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', {})                                       : Error executing SignalFlow program (ref: G8Oi8fyA4AM): [block : argument maxExtrapolations expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', {})\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', {'x': 'x'})                               : Error executing SignalFlow program (ref: G8Oi8fyA4AQ): [block : argument maxExtrapolations expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', {'x': 'x'})\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', {'x': ['x']})                             : Error executing SignalFlow program (ref: G8Oi8fyA4AU): [block : argument maxExtrapolations expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', {'x': ['x']})\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', {'x': 1})                                 : Error executing SignalFlow program (ref: G8Oi-eRA0AA): [block : argument maxExtrapolations expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', {'x': 1})\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', {'x': [1]})                               : Error executing SignalFlow program (ref: G8Oi-eRA0AI): [block : argument maxExtrapolations expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', {'x': [1]})\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', lambda x: x)                              : Error executing SignalFlow program (ref: G8Oi-eRA0AM): [block : argument maxExtrapolations expects ['long' object] value, got <lambda function>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', lambda x: x)\nTypeError: <'graphite function'> takes exactly 8 arguments (7 given)]

graphite('metric', None, '.', 0, 'sum', 'zero', 0, None)                                  : <'graphite function'> takes exactly 8 arguments (8 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, 1)                                     : <'graphite function'> takes exactly 8 arguments (8 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, 1.1)                                   : Error executing SignalFlow program (ref: G8OjAk1AwAA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'double' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, 1.1)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1')                                   : Error executing SignalFlow program (ref: G8OjAveA0AE): [error parsing program graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1'): invalid duration string 1]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, const(1))                              : Error executing SignalFlow program (ref: G8OjA6mA4AA): [block : argument resolution expects ['duration' object, 'none' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, const(1))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, data('x'))                             : Error executing SignalFlow program (ref: G8OjA_fAwAA): [block : argument resolution expects ['duration' object, 'none' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, data('x'))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, duration('1m'))                        : <'graphite function'> takes exactly 8 arguments (8 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, True)                                  : Error executing SignalFlow program (ref: G8OjCGLA4AA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, True)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, False)                                 : Error executing SignalFlow program (ref: G8OjCYvAwAA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, False)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, [])                                    : Error executing SignalFlow program (ref: G8OjCkKA0AA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, [])\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, ['1'])                                 : Error executing SignalFlow program (ref: G8OjDBuAwAs): [block : argument resolution expects ['duration' object, 'none' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, ['1'])\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, [1])                                   : Error executing SignalFlow program (ref: G8OjEk9A0AA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, [1])\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, ())                                    : Error executing SignalFlow program (ref: G8OjEsVA4AE): [block : argument resolution expects ['duration' object, 'none' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, ())\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, ('1',))                                : Error executing SignalFlow program (ref: G8OjEw6AwAA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, ('1',))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, ('1',1))                               : Error executing SignalFlow program (ref: G8OjEz7A0AA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, ('1',1))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, {})                                    : Error executing SignalFlow program (ref: G8OjE_bA4AE): [block : argument resolution expects ['duration' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, {})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, {'x': 'x'})                            : Error executing SignalFlow program (ref: G8OjFo0AwAE): [block : argument resolution expects ['duration' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, {'x': 'x'})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, {'x': ['x']})                          : Error executing SignalFlow program (ref: G8OjGQ4A0AA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, {'x': ['x']})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, {'x': 1})                              : Error executing SignalFlow program (ref: G8OjGquAwAI): [block : argument resolution expects ['duration' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, {'x': 1})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, {'x': [1]})                            : Error executing SignalFlow program (ref: G8OjG1vA0AA): [block : argument resolution expects ['duration' object, 'none' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, {'x': [1]})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, lambda x: x)                           : Error executing SignalFlow program (ref: G8OjH7RA4AA): [block : argument resolution expects ['duration' object, 'none' object] value, got <lambda function>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, lambda x: x)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]

graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', None)                            : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', 1)                               : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', 1.1)                             : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', '1')                             : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', const(1))                        : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', data('x'))                       : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', duration('1m'))                  : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', True)                            : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', False)                           : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', [])                              : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', ['1'])                           : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', [1])                             : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', ())                              : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', ('1',))                          : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', ('1',1))                         : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', {})                              : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', {'x': 'x'})                      : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', {'x': ['x']})                    : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', {'x': 1})                        : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', {'x': [1]})                      : <'graphite function'> takes exactly 8 arguments (9 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', lambda x: x)                     : <'graphite function'> takes exactly 8 arguments (9 given)

graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=None)                        : Error executing SignalFlow program (ref: G8OjgwmA4AE): [block : argument foo expects ['long' object] value, got 'none' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=None)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=1)                           : <'graphite function'> takes exactly 8 arguments (8 given)
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=1.1)                         : Error executing SignalFlow program (ref: G8OjhHUA0AA): [block : argument foo expects ['long' object] value, got 'double' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=1.1)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo='1')                         : Error executing SignalFlow program (ref: G8OjhIgA4AA): [block : argument foo expects ['long' object] value, got 'string' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo='1')\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=const(1))                    : Error executing SignalFlow program (ref: G8OjhKKAwAA): [block : argument foo expects ['long' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=const(1))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=data('x'))                   : Error executing SignalFlow program (ref: G8OjhL2A0AA): [block : argument foo expects ['long' object] value, got <stream of DOUBLE>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=data('x'))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=duration('1m'))              : Error executing SignalFlow program (ref: G8OjhNhA4AA): [block : argument foo expects ['long' object] value, got 'duration' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=duration('1m'))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=True)                        : Error executing SignalFlow program (ref: G8OjhPLAwAA): [block : argument foo expects ['long' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=True)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=False)                       : Error executing SignalFlow program (ref: G8OjhQ2A0AA): [block : argument foo expects ['long' object] value, got 'boolean' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=False)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=[])                          : Error executing SignalFlow program (ref: G8OjhShA4AA): [block : argument foo expects ['long' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=[])\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=['1'])                       : Error executing SignalFlow program (ref: G8OjhUOAwAA): [block : argument foo expects ['long' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=['1'])\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=[1])                         : Error executing SignalFlow program (ref: G8OjhV5A0AE): [block : argument foo expects ['long' object] value, got 'list' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=[1])\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=())                          : Error executing SignalFlow program (ref: G8OjhXlA4AA): [block : argument foo expects ['long' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=())\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=('1',))                      : Error executing SignalFlow program (ref: G8OjhZQAwAQ): [block : argument foo expects ['long' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=('1',))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=('1',1))                     : Error executing SignalFlow program (ref: G8OjhcwA4AA): [block : argument foo expects ['long' object] value, got 'tuple' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=('1',1))\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={})                          : Error executing SignalFlow program (ref: G8OjhebAwAA): [block : argument foo expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={'x': 'x'})                  : Error executing SignalFlow program (ref: G8OjhgGA0AE): [block : argument foo expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={'x': 'x'})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={'x': ['x']})                : Error executing SignalFlow program (ref: G8OjhhwA4AE): [block : argument foo expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={'x': ['x']})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={'x': 1})                    : Error executing SignalFlow program (ref: G8OjhjbAwAA): [block : argument foo expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={'x': 1})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={'x': [1]})                  : Error executing SignalFlow program (ref: G8OjhlHA0AM): [block : argument foo expects ['long' object] value, got 'dict' object, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo={'x': [1]})\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]
graphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=lambda x: x)                 : Error executing SignalFlow program (ref: G8Ojhm0A4AA): [block : argument foo expects ['long' object] value, got <lambda function>, File "", line 1, in \ngraphite('metric', None, '.', 0, 'sum', 'zero', 0, '1m', foo=lambda x: x)\nTypeError: <'graphite function'> takes exactly 8 arguments (8 given)]

graphite('metric', foo=None)                                                              : <'graphite function'> unsupported type 'none' object for argument 'foo'
graphite('metric', foo=1)                                                                 :
graphite('metric', foo=1.1)                                                               : <'graphite function'> unsupported type 'double' object for argument 'foo'
graphite('metric', foo='1')                                                               : <'graphite function'> unsupported type 'string' object for argument 'foo'
graphite('metric', foo=const(1))                                                          : <'graphite function'> unsupported type <stream of DOUBLE> for argument 'foo'
graphite('metric', foo=data('x'))                                                         : <'graphite function'> unsupported type <stream of DOUBLE> for argument 'foo'
graphite('metric', foo=duration('1m'))                                                    : <'graphite function'> unsupported type 'duration' object for argument 'foo'
graphite('metric', foo=True)                                                              : <'graphite function'> unsupported type 'boolean' object for argument 'foo'
graphite('metric', foo=False)                                                             : <'graphite function'> unsupported type 'boolean' object for argument 'foo'
graphite('metric', foo=[])                                                                : <'graphite function'> unsupported type 'list' object for argument 'foo'
graphite('metric', foo=['1'])                                                             : <'graphite function'> unsupported type 'list' object for argument 'foo'
graphite('metric', foo=[1])                                                               : <'graphite function'> unsupported type 'list' object for argument 'foo'
graphite('metric', foo=())                                                                : <'graphite function'> unsupported type 'tuple' object for argument 'foo'
graphite('metric', foo=('1',))                                                            : <'graphite function'> unsupported type 'tuple' object for argument 'foo'
graphite('metric', foo=('1',1))                                                           : <'graphite function'> unsupported type 'tuple' object for argument 'foo'
graphite('metric', foo={})                                                                : <'graphite function'> unsupported type 'dict' object for argument 'foo'
graphite('metric', foo={'x': 'x'})                                                        : <'graphite function'> unsupported type 'dict' object for argument 'foo'
graphite('metric', foo={'x': ['x']})                                                      : <'graphite function'> unsupported type 'dict' object for argument 'foo'
graphite('metric', foo={'x': 1})                                                          : <'graphite function'> unsupported type 'dict' object for argument 'foo'
graphite('metric', foo={'x': [1]})                                                        : <'graphite function'> unsupported type 'dict' object for argument 'foo'
graphite('metric', foo=lambda x: x)                                                       : <'graphite function'> unsupported type <lambda function> for argument 'foo'

Notes:
- 8 parameters, and **kwargs
  - metric: string, required
  - filter: filter, None (default)
  - _sf_delimiter: magic, must be '.', can't be passed through a keyword
  - offset: duration (int, string), or missing
  - rollup: string or none
  - extrapolation: string or missing
  - maxExtrapolations: int or missing
  - resolution: duration (int, string), none
- _sf_delimiter is a pain, we can't model it accurately.  We'll handle it by ignoring it.

*/

type ffiGraphite struct {
	Metric *primitive.ObjectString           `ffi:"metric"`
	Filter ftypes.ThingOrNone[filter.Filter] `ffi:"filter"`
	Offset struct {
		Missing *interpreter.ObjectMissing
		/*Duration duration.Duration
		Int      *primitive.ObjectInt
		String   *primitive.ObjectString*/
	} `ffi:"offset"`
	Rollup            ftypes.ThingOrNone[*primitive.ObjectString] `ffi:"rollup"`
	Extrapolation     ftypes.ThingOrNone[*primitive.ObjectString] `ffi:"extrapolation"`
	MaxExtrapolations ftypes.ThingOrNone[*primitive.ObjectInt]    `ffi:"maxExtrapolations"`
	Resolution        struct {
		Duration *duration.Duration
		Int      *primitive.ObjectInt
		String   *primitive.ObjectString
		None     *primitive.ObjectNone
	} `ffi:"resolution"`
	KWArgs *primitive.ObjectDict `ffi:"kw,kwargs"`
}

func NewFFIGraphite() itypes.FlowCall {
	return ffi.NewFFI(ffiGraphite{
		Filter: ftypes.NewThingOrNoneNone[filter.Filter](),
		Offset: struct {
			/*Duration duration.Duration
			Int      *primitive.ObjectInt
			String   *primitive.ObjectString*/
			Missing *interpreter.ObjectMissing
		}{
			Missing: interpreter.NewObjectMissing(),
		},
		Rollup:            ftypes.NewThingOrNoneNone[*primitive.ObjectString](),
		Extrapolation:     ftypes.NewThingOrNoneNone[*primitive.ObjectString](),
		MaxExtrapolations: ftypes.NewThingOrNoneNone[*primitive.ObjectInt](),
		Resolution: struct {
			Duration *duration.Duration
			Int      *primitive.ObjectInt
			String   *primitive.ObjectString
			None     *primitive.ObjectNone
		}{
			None: primitive.NewObjectNone(),
		},
	})
}

func (f ffiGraphite) Call(i itypes.Interpreter) (itypes.Object, error) {
	if offset, err := f.resolveOffset(); err != nil {
		return nil, err
	} else if rollup, err := f.resolveRollup(); err != nil {
		return nil, err
	} else if extrapolation, err := f.resolveExtrapolation(); err != nil {
		return nil, err
	} else if maxExtrapolations, err := f.resolveMaxExtrapolations(); err != nil {
		return nil, err
	} else if resolution, err := f.resolveResolution(); err != nil {
		return nil, err
	} else if segments, err := f.resolveSegments(); err != nil {
		return nil, err
	} else {
		return NewStreamFuncGraphite(
			prototypeStreamDouble,
			f.Metric.Value,
			f.resolveFilter(),
			offset,
			rollup,
			extrapolation,
			maxExtrapolations,
			resolution,
			segments,
			0,
		), nil
	}
}

func (f ffiGraphite) resolveFilter() filter.Filter {
	if f.Filter.Thing != nil {
		return f.Filter.Thing
	} else {
		return nil
	}
}

func (f ffiGraphite) resolveOffset() (time.Duration, error) {
	if f.Offset.Missing == nil {
		return 0, errors.New("ffiGraphite.resolveOffset: param `offset` is expected to be missing")
	} else {
		return 0, nil
	}
}

func (f ffiGraphite) resolveRollup() (string, error) {
	if f.Rollup.None != nil {
		return "", nil
	} else if f.Rollup.Thing.Value == "average" {
		return "average", nil
	} else if f.Rollup.Thing.Value == "count" {
		return "count", nil
	} else if f.Rollup.Thing.Value == "latest" {
		return "latest", nil
	} else if f.Rollup.Thing.Value == "max" {
		return "max", nil
	} else if f.Rollup.Thing.Value == "min" {
		return "min", nil
	} else if f.Rollup.Thing.Value == "rate" {
		return "rate", nil
	} else if f.Rollup.Thing.Value == "sum" {
		return "sum", nil
	} else {
		return "", fmt.Errorf("ffiGraphite.resolveRollup: param `rollup` is %s, not [average, count, latest, max, min, rate, sum]", f.Rollup.Thing.Value)
	}
}

func (f ffiGraphite) resolveExtrapolation() (string, error) {
	if f.Extrapolation.None != nil {
		return "null", nil
	} else if f.Extrapolation.Thing.Value == "last_value" {
		return "last_value", nil
	} else if f.Extrapolation.Thing.Value == "null" {
		return "null", nil
	} else if f.Extrapolation.Thing.Value == "zero" {
		return "zero", nil
	} else {
		return "", fmt.Errorf("ffiGraphite.resolveExtrapolation: param `extrapolation` is %s, not [last_value, null, zero]", f.Extrapolation.Thing.Value)
	}
}

func (f ffiGraphite) resolveMaxExtrapolations() (int, error) {
	if f.MaxExtrapolations.None != nil {
		return -1, nil
	} else if f.MaxExtrapolations.Thing.Value < -1 {
		return 0, fmt.Errorf("ffiGraphite.resolveMaxExtrapolations: param `maxExtrapolations` must be above -2, got %d", f.MaxExtrapolations.Thing.Value)
	} else {
		return f.MaxExtrapolations.Thing.Value, nil
	}
}

func (f ffiGraphite) resolveResolution() (time.Duration, error) {
	if f.Resolution.None != nil {
		return 0, nil
	} else if f.Resolution.Int != nil {
		return time.Duration(f.Resolution.Int.Value) * time.Millisecond, nil
	} else if f.Resolution.Duration != nil {
		return f.Resolution.Duration.Duration, nil
	} else if resolution, err := duration.ParseDuration(f.Resolution.String.Value); err != nil {
		return 0, fmt.Errorf("ffiGraphite.resolveResolution: param `resolution` failed to parse: %w", err)
	} else {
		return resolution, nil
	}
}

func (f ffiGraphite) resolveSegments() (map[string]int, error) {
	m := map[string]int{}
	for idx, item := range f.KWArgs.Items {
		if strKey, ok := item.Key.(*primitive.ObjectString); !ok {
			return nil, fmt.Errorf("ffiGraphite.resolveSegments: key idx %d is %T not %T", idx, item.Key, strKey)
		} else if intValue, ok := item.Value.(*primitive.ObjectInt); !ok {
			return nil, fmt.Errorf("ffiGraphite.resolveSegments: key %s is %T not %T", strKey.Value, item.Value, intValue)
		} else {
			m[strKey.Value] = intValue.Value
		}
	}
	return m, nil
}

func (f ffiGraphite) Repr() string {
	return "ffiGraphite()"
}
