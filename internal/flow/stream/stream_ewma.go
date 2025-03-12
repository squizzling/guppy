package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodEWMA struct {
	interpreter.Object
}

func (m methodEWMA) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodEWMA) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewEWMA(self), nil
	}
}

type EWMA struct {
	interpreter.Object

	source Stream
}

func NewEWMA(source Stream) Stream {
	return &EWMA{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *EWMA) RenderStream() string {
	return s.source.RenderStream() + ".eWMA()"
}
