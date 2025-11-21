package filter

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
)

type not struct {
	itypes.Object

	right Filter
}

func NewNot(right Filter) Filter {
	return &not{
		Object: prototypeFilter,
		right:  right,
	}
}

func (n *not) Repr() string {
	return fmt.Sprintf("(not %s)", n.right.Repr())
}

type ffiFilterUnaryBinaryNot struct {
	Self Filter `ffi:"self"`
}

func (f ffiFilterUnaryBinaryNot) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewNot(f.Self), nil
}
