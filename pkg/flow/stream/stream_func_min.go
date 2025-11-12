package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIMin struct {
	itypes.Object
}

func (f FFIMin) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "values",
	}, nil
}

func (f FFIMin) Call(i itypes.Interpreter) (itypes.Object, error) {
	var minConstant itypes.Object
	var streamValues []Stream
	if values, err := itypes.ArgAs[*interpreter.ObjectTuple](i, "values"); err != nil {
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
	return NewStreamFuncMin(newStreamObject(), streamValues, minConstant), nil
}

var _ = interpreter.FlowCall(FFIMin{})
