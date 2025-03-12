package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodHistogramPercentile struct {
	interpreter.Object
}

func (m methodHistogramPercentile) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodHistogramPercentile) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewHistogramPercentile(self), nil
	}
}

type HistogramPercentile struct {
	interpreter.Object

	source Stream
}

func NewHistogramPercentile(source Stream) Stream {
	return &HistogramPercentile{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *HistogramPercentile) RenderStream() string {
	return s.source.RenderStream() + ".histogramPercentile()"
}
