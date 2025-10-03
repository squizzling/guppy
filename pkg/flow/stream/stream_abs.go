package stream

import (
	"guppy/pkg/interpreter"
)

// TODO: All of this.

type methodAbs struct {
	interpreter.Object
}

func (m methodAbs) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodAbs) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamAbs(newStreamObject(), unpublish(self)), nil
	}
}
