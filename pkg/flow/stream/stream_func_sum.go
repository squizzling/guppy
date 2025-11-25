package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type FFISum struct {
	itypes.Object
}

func (f FFISum) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "values",
	}, nil
}

func (f FFISum) Call(i itypes.Interpreter) (itypes.Object, error) {
	haveConstant := false
	var sumConstant float64
	var streamValues []Stream
	if values, err := itypes.ArgAs[*primitive.ObjectTuple](i, "values"); err != nil {
		return nil, err
	} else {
		for _, value := range values.Items {
			switch value := value.(type) {
			case *primitive.ObjectInt:
				sumConstant += float64(value.Value)
				haveConstant = true
			case *primitive.ObjectDouble:
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
		return primitive.NewObjectDouble(sumConstant), nil
	} else {
		return NewStreamFuncSum(prototypeStreamDouble, streamValues, sumConstant), nil
	}
}

var _ = itypes.FlowCall(FFISum{})
