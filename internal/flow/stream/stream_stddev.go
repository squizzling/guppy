package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodStdDev struct {
	interpreter.Object
}

func (m methodStdDev) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodStdDev) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStdDev(self), nil
	}
}

type StdDev struct {
	interpreter.Object

	source Stream
}

func NewStdDev(source Stream) Stream {
	return &StdDev{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *StdDev) RenderStream() string {
	return s.source.RenderStream() + ".stdDev()"
}
