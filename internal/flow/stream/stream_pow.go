package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodPow struct {
	interpreter.Object
}

func (m methodPow) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodPow) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewPow(self), nil
	}
}

type Pow struct {
	interpreter.Object

	source Stream
}

func NewPow(source Stream) Stream {
	return &Pow{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Pow) RenderStream() string {
	return s.source.RenderStream() + ".pow()"
}
