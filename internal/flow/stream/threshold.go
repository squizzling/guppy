package stream

import (
	"fmt"

	"guppy/internal/interpreter"
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
		return NewThreshold(value), nil
	}
}

var _ = interpreter.FlowCall(FFIThreshold{})

// threshold is not a stream, but then it becomes a stream, case 3697507
type threshold struct {
	interpreter.Object

	value float64
}

func NewThreshold(value float64) Stream {
	return &threshold{
		Object: newStreamObject(),
		value:  value,
	}
}

func (t *threshold) RenderStream() string {
	return fmt.Sprintf("threshold(%f)", t.value)
}
