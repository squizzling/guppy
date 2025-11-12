package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIMax struct {
	itypes.Object
}

func (f FFIMax) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "values",
	}, nil
}

func (f FFIMax) Call(i itypes.Interpreter) (itypes.Object, error) {
	var maxConstant itypes.Object
	var streamValues []Stream
	if values, err := itypes.ArgAs[*interpreter.ObjectTuple](i, "values"); err != nil {
		return nil, err
	} else {
		for _, value := range values.Items {
			switch value := value.(type) {
			case *interpreter.ObjectInt:
				switch c := maxConstant.(type) {
				case *interpreter.ObjectInt:
					if value.Value > c.Value {
						maxConstant = value
					}
				case *interpreter.ObjectDouble:
					if float64(value.Value) > c.Value {
						maxConstant = value
					}
				case nil:
					maxConstant = value
				}
			case *interpreter.ObjectDouble:
				switch c := maxConstant.(type) {
				case *interpreter.ObjectInt:
					if value.Value > float64(c.Value) {
						maxConstant = value
					}
				case *interpreter.ObjectDouble:
					if value.Value > c.Value {
						maxConstant = value
					}
				case nil:
					maxConstant = value
				}
			case Stream:
				streamValues = append(streamValues, unpublish(value))
			default:
				return nil, fmt.Errorf("unexpected type: %T", value)
			}
		}
	}

	if len(streamValues) == 0 {
		if maxConstant == nil {
			return nil, fmt.Errorf("invalid number of arguments to function max, expected at least 1")
		} else {
			return maxConstant, nil
		}
	}

	// streamValues are already unpublished
	return NewStreamFuncMax(newStreamObject(), streamValues, maxConstant), nil
}

var _ = itypes.FlowCall(FFIMax{})
