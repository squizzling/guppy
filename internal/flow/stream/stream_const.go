package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type FFIConst struct {
	interpreter.Object
}

func (f FFIConst) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "value"},
		},
	}, nil
}

func (f FFIConst) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if value, err := i.Scope.Get("value"); err != nil {
		return nil, err
	} else {
		switch value := value.(type) {
		case *interpreter.ObjectInt:
			return NewConstStream(value.Value), nil
		case *interpreter.ObjectDouble:
			return NewConstStream(value.Value), nil
		default:
			return nil, fmt.Errorf("value is %T not *interpreter.ObjectInt, or *interpreter.ObjectDouble", value)
		}
	}
}

var _ = interpreter.FlowCall(FFIConst{})

type constStream[T any] struct {
	interpreter.Object

	value T
}

func NewConstStream[T any](i T) Stream {
	return &constStream[T]{
		Object: newStreamObject(),
		value:  i,
	}
}

func (cs *constStream[T]) RenderStream() string {
	return fmt.Sprintf("const(%v)", cs.value)
}
