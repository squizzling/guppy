package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodAbs struct {
	interpreter.Object
}

func (m methodAbs) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodAbs) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewAbs(self), nil
	}
}

type Abs struct {
	interpreter.Object

	source Stream
}

func NewAbs(source Stream) Stream {
	return &Abs{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Abs) RenderStream() string {
	return s.source.RenderStream() + ".abs()"
}
