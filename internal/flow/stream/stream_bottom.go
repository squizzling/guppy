package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodBottom struct {
	interpreter.Object
}

func (m methodBottom) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodBottom) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewBottom(self), nil
	}
}

type Bottom struct {
	interpreter.Object

	source Stream
}

func NewBottom(source Stream) Stream {
	return &Bottom{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Bottom) RenderStream() string {
	return s.source.RenderStream() + ".bottom()"
}
