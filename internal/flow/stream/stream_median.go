package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type FFIMedian struct {
	interpreter.Object
}

func (f FFIMedian) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		StarParam: "values",
	}, nil
}

func (f FFIMedian) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	var medianConstants []interpreter.Object
	var streamValues []Stream
	if values, err := interpreter.ArgAs[*interpreter.ObjectTuple](i, "values"); err != nil {
		return nil, err
	} else {
		for _, value := range values.Items {
			switch value := value.(type) {
			case *interpreter.ObjectInt:
				medianConstants = append(medianConstants, value)
			case *interpreter.ObjectDouble:
				medianConstants = append(medianConstants, value)
			case Stream:
				streamValues = append(streamValues, unpublish(value))
			default:
				return nil, fmt.Errorf("unexpected type: %T", value)
			}
		}
	}

	if len(streamValues) == 0 {
		if medianConstants == nil {
			return nil, fmt.Errorf("invalid number of arguments to function median, expected at least 1")
		} else {
			return nil, fmt.Errorf("median on constants is not implemented")
		}
	}

	// streamValues are already unpublished
	return NewStreamMedian(newStreamObject(), streamValues, medianConstants), nil
}

var _ = interpreter.FlowCall(FFIMedian{})
