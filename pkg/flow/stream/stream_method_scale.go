package stream

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type methodScale struct {
	itypes.Object
}

func (ms methodScale) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "multiple"},
		},
	}, nil
}

func (ms methodScale) resolveMultiple(i itypes.Interpreter) (float64, error) {
	if multiple, err := i.GetArg("multiple"); err != nil {
		return 0, err
	} else {
		switch multiple := multiple.(type) {
		case *primitive.ObjectInt:
			return float64(multiple.Value), nil
		case *primitive.ObjectDouble:
			return multiple.Value, nil
		default:
			return 0, fmt.Errorf("duration is %T not *interpreter.ObjectInt", multiple)
		}
	}
}

func (ms methodScale) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if multiple, err := ms.resolveMultiple(i); err != nil {
		return nil, err
	} else {
		return NewStreamMethodScale(newStreamObject(), unpublish(self), multiple), nil
	}
}
