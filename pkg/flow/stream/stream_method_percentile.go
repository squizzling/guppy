package stream

import (
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

// TODO: All of this.

type methodPercentile struct {
	itypes.Object
}

func (mp methodPercentile) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "pct", Default: primitive.NewObjectNone()},
			{Name: "allow_missing", Default: primitive.NewObjectNone()},
			{Name: "by", Default: primitive.NewObjectNone()},
			{Name: "over", Default: primitive.NewObjectNone()},
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
