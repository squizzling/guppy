package filter

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiFilterRelOp struct {
	Self  Filter                `ffi:"self"`
	Right *primitive.ObjectNone `ffi:"right"`

	invert bool
}

func (f ffiFilterRelOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	// We can only compare to None, and we don't care if it's `==` or `is`,
	// so ultimately we're returning a fixed value.
	return primitive.NewObjectBool(f.invert), nil
}
