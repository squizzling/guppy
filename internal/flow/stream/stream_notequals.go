package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodNotEquals struct {
	interpreter.Object
}

func (m methodNotEquals) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (m methodNotEquals) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewNotEquals(self), nil
	}
}

type NotEquals struct {
	interpreter.Object

	source Stream
}

func NewNotEquals(source Stream) Stream {
	return &NotEquals{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *NotEquals) RenderStream() string {
	return s.source.RenderStream() + ".notEquals()"
}
