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

func resolveMultiple(i *interpreter.Interpreter) (int, error) {
	if multiple, err := i.Scope.GetArg("multiple"); err != nil {
		return 0, err
	} else {
		switch multiple := multiple.(type) {
		case *interpreter.ObjectInt:
			return multiple.Value, nil
		default:
			return 0, fmt.Errorf("duration is %T not *interpreter.ObjectInt", multiple)
		}
	}
}

func (ms methodScale) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if multiple, err := resolveMultiple(i); err != nil {
		return nil, err
	} else {
		return NewScale(self, multiple), nil
	}
}

type scale struct {
	interpreter.Object

	source   Stream
	multiple int
}

func NewScale(source Stream, multiple int) Stream {
	return &scale{
		Object:   newStreamObject(),
		source:   unpublish(source),
		multiple: multiple,
	}
}

func (s *scale) RenderStream() string {
	return fmt.Sprintf("%s.scale(%d)", s.source.RenderStream(), s.multiple)
}
