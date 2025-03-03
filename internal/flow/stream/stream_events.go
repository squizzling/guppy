package stream

import (
	"fmt"

	"guppy/internal/interpreter"
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
	return NewEvents(), nil
}

var _ = interpreter.FlowCall(FFIEvents{})

type events struct {
	interpreter.Object
}

func NewEvents() Stream {
	return &events{
		Object: newStreamObject(),
	}
}

func (e *events) RenderStream() string {
	return fmt.Sprintf("events()")
}
