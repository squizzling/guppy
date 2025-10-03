package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFISum struct {
	interpreter.Object
}

func (f FFISum) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		StarParam: "values",
	}, nil
}

func (f FFISum) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	haveConstant := false
	var sumConstant float64
	var streamValues []Stream
	if values, err := interpreter.ArgAs[*interpreter.ObjectTuple](i, "values"); err != nil {
		return nil, err
	} else {
		for _, value := range values.Items {
			switch value := value.(type) {
			case *interpreter.ObjectInt:
				sumConstant += float64(value.Value)
				haveConstant = true
			case *interpreter.ObjectDouble:
				sumConstant += value.Value
				haveConstant = true
			case Stream:
				streamValues = append(streamValues, unpublish(value))
			default:
				return nil, fmt.Errorf("unexpected type: %T", value)
			}
		}
	}

	if len(streamValues) == 0 && !haveConstant {
		return nil, fmt.Errorf("invalid number of arguments to function min, expected at least 1")
	} else if len(streamValues) == 0 {
		return interpreter.NewObjectDouble(sumConstant), nil
	} else {
		return NewStreamFuncSum(newStreamObject(), streamValues, sumConstant), nil
	}
}

var _ = interpreter.FlowCall(FFISum{})
