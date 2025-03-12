package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodAbove struct {
	interpreter.Object
}

func (ma methodAbove) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "limit", Default: interpreter.NewObjectNone()},
			{Name: "inclusive", Default: interpreter.NewObjectNone()},
			{Name: "clamp", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (ma methodAbove) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewAbove(self), nil
	}
}

type Above struct {
	interpreter.Object

	source Stream
}

func NewAbove(source Stream) Stream {
	return &Above{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (a *Above) RenderStream() string {
	return a.source.RenderStream() + ".above()"
}
