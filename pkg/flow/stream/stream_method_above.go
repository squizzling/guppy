package stream

import (
	"guppy/pkg/interpreter"
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
		return NewStreamMethodAbove(newStreamObject(), unpublish(self)), nil
	}
}
