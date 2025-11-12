package stream

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

// TODO: All of this.

type methodAbove struct {
	itypes.Object
}

func (ma methodAbove) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "limit", Default: interpreter.NewObjectNone()},
			{Name: "inclusive", Default: interpreter.NewObjectNone()},
			{Name: "clamp", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (ma methodAbove) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamMethodAbove(newStreamObject(), unpublish(self)), nil
	}
}
