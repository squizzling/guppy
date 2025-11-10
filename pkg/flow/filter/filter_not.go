package filter

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type methodBinaryNot struct {
	itypes.Object
}

func (mba methodBinaryNot) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (mba methodBinaryNot) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := interpreter.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else {
		return NewNot(self), nil
	}
}

type not struct {
	itypes.Object

	right Filter
}

func NewNot(right Filter) Filter {
	return &not{
		Object: newFilterObject(),
		right:  right,
	}
}

func (n *not) RenderFilter() string {
	return fmt.Sprintf("(not %s)", n.right.RenderFilter())
}
