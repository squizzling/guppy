package filter

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
)

type and struct {
	itypes.Object

	left  Filter
	right Filter
}

func NewAnd(left Filter, right Filter) Filter {
	return &and{
		Object: prototypeFilter,
		left:   left,
		right:  right,
	}
}

func (a *and) Repr() string {
	return fmt.Sprintf("(%s and %s)", a.left.Repr(), a.right.Repr())
}

type or struct {
	itypes.Object

	left  Filter
	right Filter
}

func NewOr(left Filter, right Filter) Filter {
	return &or{
		Object: prototypeFilter,
		left:   left,
		right:  right,
	}
}

func (o *or) Repr() string {
	return fmt.Sprintf("(%s or %s)", o.left.Repr(), o.right.Repr())
}

type ffiFilterBinaryOp struct {
	Self  Filter `ffi:"self"`
	Right Filter `ffi:"right"`

	op int
}

func (f ffiFilterBinaryOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	switch f.op {
	case 0:
		return NewAnd(f.Self, f.Right), nil
	default:
		return NewOr(f.Self, f.Right), nil
	}
}

/*type methodBinaryAnd struct {
	itypes.Object
}

func (mba methodBinaryAnd) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (mba methodBinaryAnd) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else if right, err := itypes.ArgAs[Filter](i, "right"); err != nil {
		return nil, err
	} else {
		return NewAnd(self, right), nil
	}
}*/
