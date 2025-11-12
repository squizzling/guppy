package stream

import (
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type FFIEvents struct {
	itypes.Object
}

func (f FFIEvents) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "eventType", Default: primitive.NewObjectNone()},
			{Name: "filter", Default: primitive.NewObjectNone()},
		},
	}, nil
}

func (f FFIEvents) Call(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: These probably aren't the right attributes for an event stream
	return NewStreamFuncEvents(newStreamObject()), nil
}

var _ = itypes.FlowCall(FFIEvents{})
