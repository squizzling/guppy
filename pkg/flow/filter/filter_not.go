package filter

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type ffiFilterUnaryBinaryNot struct {
	Self Filter `ffi:"self"`
}

func (f ffiFilterUnaryBinaryNot) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewFilterNot(prototypeFilter, f.Self), nil
}

func (fn *FilterNot) Repr() string {
	return repr(fn)
}
