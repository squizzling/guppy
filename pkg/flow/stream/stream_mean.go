package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFIMean struct {
	interpreter.Object
}

func (f FFIMean) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		StarParam: "values",
	}, nil
}

func (f FFIMean) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	var meanConstants []interpreter.Object
	var streamValues []Stream
	if values, err := interpreter.ArgAs[*interpreter.ObjectTuple](i, "values"); err != nil {
		return nil, err
	} else {
		for _, value := range values.Items {
			switch value := value.(type) {
			case *interpreter.ObjectInt:
				meanConstants = append(meanConstants, value)
			case *interpreter.ObjectDouble:
				meanConstants = append(meanConstants, value)
			case Stream:
				streamValues = append(streamValues, unpublish(value))
			default:
				return nil, fmt.Errorf("unexpected type: %T", value)
			}
		}
	}

	if len(streamValues) == 0 {
		if meanConstants == nil {
			return nil, fmt.Errorf("invalid number of arguments to function mean, expected at least 1")
		} else {
			return nil, fmt.Errorf("mean on constants is not implemented")
		}
	}

	// streamValues are already unpublished
	return NewStreamFuncMean(newStreamObject(), streamValues, meanConstants), nil
}

var _ = interpreter.FlowCall(FFIMean{})
