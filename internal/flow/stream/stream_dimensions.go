package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodDimensions struct {
	interpreter.Object
}

func (m methodDimensions) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodDimensions) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewDimensions(self), nil
	}
}

type Dimensions struct {
	interpreter.Object

	source Stream
}

func NewDimensions(source Stream) Stream {
	return &Dimensions{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Dimensions) RenderStream() string {
	return s.source.RenderStream() + ".dimensions()"
}
