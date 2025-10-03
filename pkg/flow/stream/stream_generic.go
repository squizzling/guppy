package stream

import (
	"guppy/pkg/interpreter"
)

type methodGeneric struct {
	interpreter.Object

	Function string
}

func (mg methodGeneric) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (mg methodGeneric) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamGeneric(newStreamObject(), unpublish(self), mg.Function), nil
	}
}
