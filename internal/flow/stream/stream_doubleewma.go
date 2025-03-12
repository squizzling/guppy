package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodDoubleEWMA struct {
	interpreter.Object
}

func (m methodDoubleEWMA) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodDoubleEWMA) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewDoubleEWMA(self), nil
	}
}

type DoubleEWMA struct {
	interpreter.Object

	source Stream
}

func NewDoubleEWMA(source Stream) Stream {
	return &DoubleEWMA{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *DoubleEWMA) RenderStream() string {
	return s.source.RenderStream() + ".doubleEWMA()"
}
