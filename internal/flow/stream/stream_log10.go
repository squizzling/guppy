package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodLog10 struct {
	interpreter.Object
}

func (m methodLog10) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodLog10) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewLog10(self), nil
	}
}

type Log10 struct {
	interpreter.Object

	source Stream
}

func NewLog10(source Stream) Stream {
	return &Log10{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Log10) RenderStream() string {
	return s.source.RenderStream() + ".log10()"
}
