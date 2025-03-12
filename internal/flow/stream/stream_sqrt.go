package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodSqrt struct {
	interpreter.Object
}

func (m methodSqrt) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodSqrt) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewSqrt(self), nil
	}
}

type Sqrt struct {
	interpreter.Object

	source Stream
}

func NewSqrt(source Stream) Stream {
	return &Sqrt{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Sqrt) RenderStream() string {
	return s.source.RenderStream() + ".sqrt()"
}
