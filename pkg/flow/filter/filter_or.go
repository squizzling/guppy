package filter

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
)

type methodBinaryOr struct {
	itypes.Object
}

func (mbo methodBinaryOr) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (mbo methodBinaryOr) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else if right, err := itypes.ArgAs[Filter](i, "right"); err != nil {
		return nil, err
	} else {
		return NewAnd(self, right), nil
	}
}

type or struct {
	itypes.Object

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
