package stream

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type methodGeneric struct {
	itypes.Object

	Function string
}

func (mg methodGeneric) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
		},
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (mg methodGeneric) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamMethodGeneric(newStreamObject(), unpublish(self), mg.Function), nil
	}
}
