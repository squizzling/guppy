package filter

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type methodBinaryAnd struct {
	itypes.Object
}

func (mba methodBinaryAnd) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return interpreter.BinaryParams, nil
}

func (mba methodBinaryAnd) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := interpreter.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else if right, err := interpreter.ArgAs[Filter](i, "right"); err != nil {
		return nil, err
	} else {
		return NewAnd(self, right), nil
	}
}

type and struct {
	itypes.Object

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
