package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodMeanPlusStdDev struct {
	interpreter.Object
}

func (m methodMeanPlusStdDev) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodMeanPlusStdDev) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewMeanPlusStdDev(self), nil
	}
}

type MeanPlusStdDev struct {
	interpreter.Object

	source Stream
}

func NewMeanPlusStdDev(source Stream) Stream {
	return &MeanPlusStdDev{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *MeanPlusStdDev) RenderStream() string {
	return s.source.RenderStream() + ".meanPlusStdDev()"
}
