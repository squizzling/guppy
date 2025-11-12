package stream

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

// TODO: All of this.

type methodPercentile struct {
	itypes.Object
}

func (mp methodPercentile) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "pct", Default: interpreter.NewObjectNone()},
			{Name: "allow_missing", Default: interpreter.NewObjectNone()},
			{Name: "by", Default: interpreter.NewObjectNone()},
			{Name: "over", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (mp methodPercentile) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamMethodPercentile(newStreamObject(), unpublish(self)), nil
	}
}
