package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodTimeShift struct {
	interpreter.Object
}

func (mts methodTimeShift) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "offset"},
		},
	}, nil
}

func resolveOffset(i *interpreter.Interpreter) (string, error) {
	if offset, err := interpreter.ArgAsString(i, "offset"); err != nil {
		return "", err
	} else {
		return offset, nil
	}
}

func (mts methodTimeShift) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if offset, err := resolveOffset(i); err != nil {
		return nil, err
	} else {
		return NewTimeShift(self, offset), nil
	}
}

type TimeShift struct {
	interpreter.Object

	source Stream
	offset string
}

func NewTimeShift(source Stream, offset string) Stream {
	return &TimeShift{
		Object: newStreamObject(),
		source: unpublish(source),
		offset: offset,
	}
}

func (s *TimeShift) RenderStream() string {
	return fmt.Sprintf("%s.timeshift('%s')", s.source.RenderStream(), s.offset)
}
