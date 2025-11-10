package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIMean struct {
	itypes.Object
}

func (f FFIMean) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "values",
	}, nil
}

func (f FFIMean) Call(i itypes.Interpreter) (itypes.Object, error) {
	var meanConstants []itypes.Object
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
