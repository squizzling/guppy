package filter

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type methodBinaryOr struct {
	interpreter.Object
}

func (mbo methodBinaryOr) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (mbo methodBinaryOr) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else if right, err := interpreter.ArgAs[Filter](i, "right"); err != nil {
		return nil, err
	} else {
		return NewAnd(self, right), nil
	}
}

type or struct {
	interpreter.Object

	left  Filter
	right Filter
}

func NewOr(left Filter, right Filter) Filter {
	return &or{
		Object: newFilterObject(),
		left:   left,
		right:  right,
	}
}

func (o *or) RenderFilter() string {
	return fmt.Sprintf("(%s or %s)", o.left.RenderFilter(), o.right.RenderFilter())
}
