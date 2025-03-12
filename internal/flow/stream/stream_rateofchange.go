package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodRateofChange struct {
	interpreter.Object
}

func (m methodRateofChange) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodRateofChange) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewRateofChange(self), nil
	}
}

type RateofChange struct {
	interpreter.Object

	source Stream
}

func NewRateofChange(source Stream) Stream {
	return &RateofChange{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *RateofChange) RenderStream() string {
	return s.source.RenderStream() + ".rateofChange()"
}
