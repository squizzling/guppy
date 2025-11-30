package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type FFIThreshold struct {
	itypes.Object

	value float64
}

func (f FFIThreshold) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "object"},
		},
	}, nil
}

func (f FFIThreshold) Call(i itypes.Interpreter) (itypes.Object, error) {
	if value, err := interpreter.ArgAsDouble(i, "object"); err != nil {
		return nil, err
	} else {
		return NewStreamFuncThreshold(newStreamObject(), value), nil
	}
}

var _ = itypes.FlowCall(FFIThreshold{})
