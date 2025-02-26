package filter

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodBinaryAnd struct {
	interpreter.Object
}

func (mba methodBinaryAnd) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "right"},
		},
	}, nil
}

func (mba methodBinaryAnd) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else if right, err := interpreter.ArgAs[Filter](i, "right"); err != nil {
		return nil, err
	} else {
		return NewAnd(self, right), nil
	}
}

type and struct {
	interpreter.Object

	left  Filter
	right Filter
}

func NewAnd(left Filter, right Filter) Filter {
	return &and{
		Object: newFilterObject(),
		left:   left,
		right:  right,
	}
}

func (a *and) RenderFilter() string {
	return fmt.Sprintf("(%s and %s)", a.left.RenderFilter(), a.right.RenderFilter())
}
