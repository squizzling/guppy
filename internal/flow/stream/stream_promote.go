package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodPromote struct {
	interpreter.Object
}

func (m methodPromote) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodPromote) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewPromote(self), nil
	}
}

type Promote struct {
	interpreter.Object

	source Stream
}

func NewPromote(source Stream) Stream {
	return &Promote{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Promote) RenderStream() string {
	return s.source.RenderStream() + ".promote()"
}
