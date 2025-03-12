package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodMap struct {
	interpreter.Object
}

func (m methodMap) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodMap) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewMap(self), nil
	}
}

type Map struct {
	interpreter.Object

	source Stream
}

func NewMap(source Stream) Stream {
	return &Map{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Map) RenderStream() string {
	return s.source.RenderStream() + ".map()"
}
