package stream

import (
	"guppy/pkg/interpreter"
)

type FFIEvents struct {
	interpreter.Object
}

func (f FFIEvents) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "eventType", Default: interpreter.NewObjectNone()},
			{Name: "filter", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (f FFIEvents) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	// TODO: These probably aren't the right attributes for an event stream
	return NewStreamEvents(newStreamObject()), nil
}

var _ = interpreter.FlowCall(FFIEvents{})
