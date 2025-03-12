package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodEquals struct {
	interpreter.Object
}

func (m methodEquals) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodEquals) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewEquals(self), nil
	}
}

type Equals struct {
	interpreter.Object

	source Stream
}

func NewEquals(source Stream) Stream {
	return &Equals{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Equals) RenderStream() string {
	return s.source.RenderStream() + ".equals()"
}
