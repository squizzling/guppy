package stream

import (
	"guppy/pkg/interpreter"
)

type FFIThreshold struct {
	interpreter.Object

	value float64
}

func (f FFIThreshold) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "object"},
		},
	}, nil
}

func (f FFIThreshold) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if value, err := interpreter.ArgAsDouble(i, "object"); err != nil {
		return nil, err
	} else {
		return NewStreamThreshold(newStreamObject(), value), nil
	}
}

var _ = interpreter.FlowCall(FFIThreshold{})
