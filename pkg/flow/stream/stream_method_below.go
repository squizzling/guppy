package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

// TODO: All of this.

type methodBelow struct {
	itypes.Object
}

func (mb methodBelow) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "limit", Default: primitive.NewObjectNone()},
			{Name: "inclusive", Default: primitive.NewObjectNone()},
			{Name: "clamp", Default: primitive.NewObjectNone()},
		},
	}, nil
}

func (mb methodBelow) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamMethodBelow(newStreamObject(), unpublish(self)), nil
	}
}
