package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodIntegrate struct {
	interpreter.Object
}

func (m methodIntegrate) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodIntegrate) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewIntegrate(self), nil
	}
}

type Integrate struct {
	interpreter.Object

	source Stream
}

func NewIntegrate(source Stream) Stream {
	return &Integrate{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Integrate) RenderStream() string {
	return s.source.RenderStream() + ".integrate()"
}
