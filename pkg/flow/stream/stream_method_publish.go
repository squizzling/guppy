package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type methodPublish struct {
	itypes.Object
}

func (mp methodPublish) Params(i itypes.Interpreter) (*itypes.Params, error) {
	// TODO: Convert to new FFI.  When converting, find out what other args are being passed, because I think there's a bunch.
	// This is 50 shades of wonky.  publish takes two parameters: label and additional_dimensions.  enable is handled through **args.
	/**
	const(1).publish(None)                                      :
	const(1).publish(1)                                         : <'streampublish function'> unsupported type 'long' object for argument 'label'
	const(1).publish(1.1)                                       : <'streampublish function'> unsupported type 'double' object for argument 'label'
	const(1).publish('1')                                       :
	const(1).publish(const(1))                                  : <'streampublish function'> unsupported type <stream of DOUBLE> for argument 'label'
	const(1).publish(data('x'))                                 : <'streampublish function'> unsupported type <stream of DOUBLE> for argument 'label'
	const(1).publish(duration('1m'))                            : <'streampublish function'> unsupported type 'duration' object for argument 'label'
	const(1).publish(True)                                      : <'streampublish function'> unsupported type 'boolean' object for argument 'label'
	const(1).publish(False)                                     : <'streampublish function'> unsupported type 'boolean' object for argument 'label'
	const(1).publish([])                                        : <'streampublish function'> unsupported type 'list' object for argument 'label'
	const(1).publish(['1'])                                     : <'streampublish function'> unsupported type 'list' object for argument 'label'
	const(1).publish([1])                                       : <'streampublish function'> unsupported type 'list' object for argument 'label'
	const(1).publish(())                                        : <'streampublish function'> unsupported type 'tuple' object for argument 'label'
	const(1).publish(('1',))                                    : <'streampublish function'> unsupported type 'tuple' object for argument 'label'
	const(1).publish(('1',1))                                   : <'streampublish function'> unsupported type 'tuple' object for argument 'label'
	const(1).publish({})                                        : <'streampublish function'> unsupported type 'dict' object for argument 'label'
	const(1).publish({'x': 'x'})                                : <'streampublish function'> unsupported type 'dict' object for argument 'label'
	const(1).publish({'x': ['x']})                              : <'streampublish function'> unsupported type 'dict' object for argument 'label'
	const(1).publish({'x': 1})                                  : <'streampublish function'> unsupported type 'dict' object for argument 'label'
	const(1).publish({'x': [1]})                                : <'streampublish function'> unsupported type 'dict' object for argument 'label'
	const(1).publish(lambda x: x)                               : <'streampublish function'> unsupported type <lambda function> for argument 'label'

	const(1).publish('x', None)                                 : <'streampublish function'> unsupported type 'none' object for argument 'additional_dimensions'
	const(1).publish('x', 1)                                    : <'streampublish function'> unsupported type 'long' object for argument 'additional_dimensions'
	const(1).publish('x', 1.1)                                  : <'streampublish function'> unsupported type 'double' object for argument 'additional_dimensions'
	const(1).publish('x', '1')                                  : <'streampublish function'> unsupported type 'string' object for argument 'additional_dimensions'
	const(1).publish('x', const(1))                             : <'streampublish function'> unsupported type <stream of DOUBLE> for argument 'additional_dimensions'
	const(1).publish('x', data('x'))                            : <'streampublish function'> unsupported type <stream of DOUBLE> for argument 'additional_dimensions'
	const(1).publish('x', duration('1m'))                       : <'streampublish function'> unsupported type 'duration' object for argument 'additional_dimensions'
	const(1).publish('x', True)                                 : <'streampublish function'> unsupported type 'boolean' object for argument 'additional_dimensions'
	const(1).publish('x', False)                                : <'streampublish function'> unsupported type 'boolean' object for argument 'additional_dimensions'
	const(1).publish('x', [])                                   : <'streampublish function'> unsupported type 'list' object for argument 'additional_dimensions'
	const(1).publish('x', ['1'])                                : <'streampublish function'> unsupported type 'list' object for argument 'additional_dimensions'
	const(1).publish('x', [1])                                  : <'streampublish function'> unsupported type 'list' object for argument 'additional_dimensions'
	const(1).publish('x', ())                                   : <'streampublish function'> unsupported type 'tuple' object for argument 'additional_dimensions'
	const(1).publish('x', ('1',))                               : <'streampublish function'> unsupported type 'tuple' object for argument 'additional_dimensions'
	const(1).publish('x', ('1',1))                              : <'streampublish function'> unsupported type 'tuple' object for argument 'additional_dimensions'
	const(1).publish('x', {})                                   : <'streampublish function'> argument 'additional_dimensions' got invalid value 'an empty dimension map'; expected a non-empty dimension map
	const(1).publish('x', {'x': 'x'})                           :
	const(1).publish('x', {'x': ['x']})                         : <'streampublish function'> unsupported type 'list' object for argument 'additional_dimensions'
	const(1).publish('x', {'x': 1})                             : <'streampublish function'> unsupported type 'long' object for argument 'additional_dimensions'
	const(1).publish('x', {'x': [1]})                           : <'streampublish function'> unsupported type 'list' object for argument 'additional_dimensions'
	const(1).publish('x', lambda x: x)                          : <'streampublish function'> unsupported type <lambda function> for argument 'additional_dimensions'

	const(1).publish('x', {'x': 'x'}, None)                     : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, 1)                        : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, 1.1)                      : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, '1')                      : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, const(1))                 : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, data('x'))                : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, duration('1m'))           : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, True)                     : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, False)                    : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, [])                       : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, ['1'])                    : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, [1])                      : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, ())                       : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, ('1',))                   : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, ('1',1))                  : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, {})                       : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, {'x': 'x'})               : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, {'x': ['x']})             : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, {'x': 1})                 : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, {'x': [1]})               : <'streampublish function'> takes exactly 2 arguments (3 given)
	const(1).publish('x', {'x': 'x'}, lambda x: x)              : <'streampublish function'> takes exactly 2 arguments (3 given)

	const(1).publish('x', {'x': 'x'}, derp=None)                : <'streampublish function'> unsupported type 'none' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=1)                   : <'streampublish function'> unsupported type 'long' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=1.1)                 : <'streampublish function'> unsupported type 'double' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp='1')                 :
	const(1).publish('x', {'x': 'x'}, derp=const(1))            : <'streampublish function'> unsupported type <stream of DOUBLE> for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=data('x'))           : <'streampublish function'> unsupported type <stream of DOUBLE> for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=duration('1m'))      : <'streampublish function'> unsupported type 'duration' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=True)                : <'streampublish function'> unsupported type 'boolean' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=False)               : <'streampublish function'> unsupported type 'boolean' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=[])                  : <'streampublish function'> unsupported type 'list' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=['1'])               : <'streampublish function'> unsupported type 'list' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=[1])                 : <'streampublish function'> unsupported type 'list' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=())                  : <'streampublish function'> unsupported type 'tuple' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=('1',))              : <'streampublish function'> unsupported type 'tuple' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=('1',1))             : <'streampublish function'> unsupported type 'tuple' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp={})                  : <'streampublish function'> unsupported type 'dict' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp={'x': 'x'})          : <'streampublish function'> unsupported type 'dict' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp={'x': ['x']})        : <'streampublish function'> unsupported type 'dict' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp={'x': 1})            : <'streampublish function'> unsupported type 'dict' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp={'x': [1]})          : <'streampublish function'> unsupported type 'dict' object for argument 'derp'
	const(1).publish('x', {'x': 'x'}, derp=lambda x: x)         : <'streampublish function'> unsupported type <lambda function> for argument 'derp'

	const(1).publish('x', {'x': 'x'}, enable=None)              : <'streampublish function'> unsupported type 'none' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=1)                 : <'streampublish function'> unsupported type 'long' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=1.1)               : <'streampublish function'> unsupported type 'double' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable='1')               : <'streampublish function'> unsupported type 'string' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=const(1))          : <'streampublish function'> unsupported type <stream of DOUBLE> for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=data('x'))         : <'streampublish function'> unsupported type <stream of DOUBLE> for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=duration('1m'))    : <'streampublish function'> unsupported type 'duration' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=True)              :
	const(1).publish('x', {'x': 'x'}, enable=False)             :
	const(1).publish('x', {'x': 'x'}, enable=[])                : <'streampublish function'> unsupported type 'list' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=['1'])             : <'streampublish function'> unsupported type 'list' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=[1])               : <'streampublish function'> unsupported type 'list' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=())                : <'streampublish function'> unsupported type 'tuple' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=('1',))            : <'streampublish function'> unsupported type 'tuple' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=('1',1))           : <'streampublish function'> unsupported type 'tuple' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable={})                : <'streampublish function'> unsupported type 'dict' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable={'x': 'x'})        : <'streampublish function'> unsupported type 'dict' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable={'x': ['x']})      : <'streampublish function'> unsupported type 'dict' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable={'x': 1})          : <'streampublish function'> unsupported type 'dict' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable={'x': [1]})        : <'streampublish function'> unsupported type 'dict' object for argument 'enable'
	const(1).publish('x', {'x': 'x'}, enable=lambda x: x)       : <'streampublish function'> unsupported type <lambda function> for argument 'enable'
	*/
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "label", Default: primitive.NewObjectString("")}, // TODO: Validate "" vs None
			{Name: "enable", Default: primitive.NewObjectBool(true)},
		},
		//KWParam: "additional_dimensions", // Maybe, I don't fully know this one.
		KWParam: "_",
	}, nil
}

func (mp methodPublish) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if label, err := interpreter.ArgAsString(i, "label"); err != nil {
		return nil, err
	} else if enable, err := itypes.ArgAs[*primitive.ObjectBool](i, "enable"); err != nil {
		return nil, err
	} else {
		// TODO: This whole thing is a hack to expose published data
		pub := NewStreamMethodPublish(newStreamObject(), unpublish(self), label, enable.Value)
		if rawPublished, err := i.GetGlobal("_published"); err != nil {
			return nil, err
		} else if published, ok := rawPublished.(interface{ Append(s *StreamMethodPublish) }); !ok {
			return nil, fmt.Errorf("invalid type")
		} else {
			published.Append(pub)
		}
		return pub, nil
	}
}

func (smp *StreamMethodPublish) Repr() string {
	// TODO: Better
	return ".publish()"
}
