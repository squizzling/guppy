package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodBelow struct {
	interpreter.Object
}

func (mb methodBelow) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "limit", Default: interpreter.NewObjectNone()},
			{Name: "inclusive", Default: interpreter.NewObjectNone()},
			{Name: "clamp", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (mb methodBelow) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewBelow(self), nil
	}
}

type Below struct {
	interpreter.Object

	source Stream
}

func NewBelow(source Stream) Stream {
	return &Below{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (b *Below) RenderStream() string {
	return b.source.RenderStream() + ".below()"
}
