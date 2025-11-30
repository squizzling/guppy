package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type FFIMedian struct {
	itypes.Object
}

func (f FFIMedian) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "values",
	}, nil
}

func (f FFIMedian) Call(i itypes.Interpreter) (itypes.Object, error) {
	var medianConstants []itypes.Object
	var streamValues []Stream
	if values, err := itypes.ArgAs[*primitive.ObjectTuple](i, "values"); err != nil {
		return nil, err
	} else {
		for _, value := range values.Items {
			switch value := value.(type) {
			case *primitive.ObjectInt:
				medianConstants = append(medianConstants, value)
			case *primitive.ObjectDouble:
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
	return NewStreamFuncMedian(newStreamObject(), streamValues, medianConstants), nil
}

var _ = itypes.FlowCall(FFIMedian{})
