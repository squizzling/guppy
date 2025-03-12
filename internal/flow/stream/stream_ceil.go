package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodCeil struct {
	interpreter.Object
}

func (m methodCeil) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodCeil) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewCeil(self), nil
	}
}

type Ceil struct {
	interpreter.Object

	source Stream
}

func NewCeil(source Stream) Stream {
	return &Ceil{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Ceil) RenderStream() string {
	return s.source.RenderStream() + ".ceil()"
}
