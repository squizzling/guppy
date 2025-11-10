package stream

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIEvents struct {
	itypes.Object
}

func (f FFIEvents) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "eventType", Default: interpreter.NewObjectNone()},
			{Name: "filter", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (f FFIEvents) Call(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: These probably aren't the right attributes for an event stream
	return NewStreamFuncEvents(newStreamObject()), nil
}

var _ = interpreter.FlowCall(FFIEvents{})
