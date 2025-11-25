package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
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
	if values, err := itypes.ArgAs[*primitive.ObjectTuple](i, "values"); err != nil {
		return nil, err
	} else {
		for _, value := range values.Items {
			switch value := value.(type) {
			case *primitive.ObjectInt:
				switch c := minConstant.(type) {
				case *primitive.ObjectInt:
					if value.Value < c.Value {
						minConstant = value
					}
				case *primitive.ObjectDouble:
					if float64(value.Value) < c.Value {
						minConstant = value
					}
				case nil:
					minConstant = value
				}
			case *primitive.ObjectDouble:
				switch c := minConstant.(type) {
				case *primitive.ObjectInt:
					if value.Value < float64(c.Value) {
						minConstant = value
					}
				case *primitive.ObjectDouble:
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
	return NewStreamFuncMin(prototypeStreamDouble, streamValues, minConstant), nil
}

var _ = itypes.FlowCall(FFIMin{})
