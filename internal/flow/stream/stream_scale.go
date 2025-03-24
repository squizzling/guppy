package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodScale struct {
	interpreter.Object
}

func (ms methodScale) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "multiple"},
		},
	}, nil
}

func (ms methodScale) resolveMultiple(i *interpreter.Interpreter) (float64, error) {
	if multiple, err := i.Scope.GetArg("multiple"); err != nil {
		return 0, err
	} else {
		switch multiple := multiple.(type) {
		case *interpreter.ObjectInt:
			return float64(multiple.Value), nil
		case *interpreter.ObjectDouble:
			return multiple.Value, nil
		default:
			return 0, fmt.Errorf("duration is %T not *interpreter.ObjectInt", multiple)
		}
	}
}

func (ms methodScale) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if multiple, err := ms.resolveMultiple(i); err != nil {
		return nil, err
	} else {
		return NewScale(self, multiple), nil
	}
}

type scale struct {
	interpreter.Object

	source   Stream
	multiple float64
}

func NewScale(source Stream, multiple float64) Stream {
	return &scale{
		Object:   newStreamObject(),
		source:   unpublish(source),
		multiple: multiple,
	}
}

func (s *scale) RenderStream() string {
	return fmt.Sprintf("%s.scale(%f)", s.source.RenderStream(), s.multiple)
}
