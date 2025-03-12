package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodDelta struct {
	interpreter.Object
}

func (m methodDelta) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodDelta) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewDelta(self), nil
	}
}

type Delta struct {
	interpreter.Object

	source Stream
}

func NewDelta(source Stream) Stream {
	return &Delta{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Delta) RenderStream() string {
	return s.source.RenderStream() + ".delta()"
}
