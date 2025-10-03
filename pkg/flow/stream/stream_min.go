package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFIMin struct {
	interpreter.Object
}

func (f FFIMin) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		StarParam: "values",
	}, nil
}

func (f FFIMin) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	var minConstant interpreter.Object
	var streamValues []Stream
	if values, err := interpreter.ArgAs[*interpreter.ObjectTuple](i, "values"); err != nil {
		return nil, err
	} else {
		for _, value := range values.Items {
			switch value := value.(type) {
			case *interpreter.ObjectInt:
				switch c := minConstant.(type) {
				case *interpreter.ObjectInt:
					if value.Value < c.Value {
						minConstant = value
					}
				case *interpreter.ObjectDouble:
					if float64(value.Value) < c.Value {
						minConstant = value
					}
				case nil:
					minConstant = value
				}
			case *interpreter.ObjectDouble:
				switch c := minConstant.(type) {
				case *interpreter.ObjectInt:
					if value.Value < float64(c.Value) {
						minConstant = value
					}
				case *interpreter.ObjectDouble:
					if value.Value < c.Value {
						minConstant = value
					}
				case nil:
					minConstant = value
				}
			case Stream:
				streamValues = append(streamValues, unpublish(value))
			default:
				return nil, fmt.Errorf("unexpected type: %T", value)
			}
		}
	}

	if len(streamValues) == 0 {
		if minConstant == nil {
			return nil, fmt.Errorf("invalid number of arguments to function min, expected at least 1")
		} else {
			return minConstant, nil
		}
	}

	// streamValues are already unpublished
	return NewStreamMin(newStreamObject(), streamValues, minConstant), nil
}

var _ = interpreter.FlowCall(FFIMin{})
