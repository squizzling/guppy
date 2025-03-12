package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodBetween struct {
	interpreter.Object
}

func (m methodBetween) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodBetween) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewBetween(self), nil
	}
}

type Between struct {
	interpreter.Object

	source Stream
}

func NewBetween(source Stream) Stream {
	return &Between{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Between) RenderStream() string {
	return s.source.RenderStream() + ".between()"
}
