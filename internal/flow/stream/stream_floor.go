package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodFloor struct {
	interpreter.Object
}

func (m methodFloor) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodFloor) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewFloor(self), nil
	}
}

type Floor struct {
	interpreter.Object

	source Stream
}

func NewFloor(source Stream) Stream {
	return &Floor{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Floor) RenderStream() string {
	return s.source.RenderStream() + ".floor()"
}
