package stream

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

// TODO: All of this.

type methodTop struct {
	itypes.Object
}

func (mt methodTop) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "count", Default: interpreter.NewObjectNone()},
			{Name: "by", Default: interpreter.NewObjectNone()},
			{Name: "allow_missing", Default: interpreter.NewObjectNone()},
			{Name: "percentage", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (mt methodTop) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamMethodTop(newStreamObject(), unpublish(self)), nil
	}
}
