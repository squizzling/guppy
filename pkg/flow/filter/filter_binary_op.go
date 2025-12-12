package filter

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

func newFilterAnd(left Filter, right Filter) Filter {
	var filters []Filter
	if leftAnd, ok := left.(*FilterAnd); ok {
		filters = append(filters, leftAnd.Filters...)
	} else {
		filters = append(filters, left)
	}
	if rightAnd, ok := right.(*FilterAnd); ok {
		filters = append(filters, rightAnd.Filters...)
	} else {
		filters = append(filters, right)
	}

	return NewFilterAnd(prototypeFilter, filters)
}

func (fa *FilterAnd) Repr() string {
	return repr(fa)
}

func newFilterOr(left Filter, right Filter) Filter {
	var filters []Filter
	if leftOr, ok := left.(*FilterOr); ok {
		filters = append(filters, leftOr.Filters...)
	} else {
		filters = append(filters, left)
	}

	if rightOr, ok := right.(*FilterOr); ok {
		filters = append(filters, rightOr.Filters...)
	} else {
		filters = append(filters, right)
	}

	return NewFilterOr(prototypeFilter, filters)
}

func (fo *FilterOr) Repr() string {
	return repr(fo)
}

type ffiFilterBinaryOp struct {
	Self  Filter `ffi:"self"`
	Right Filter `ffi:"right"`

	op int
}

func (f ffiFilterBinaryOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	switch f.op {
	case 0:
		return newFilterAnd(f.Self, f.Right), nil
	default:
		return newFilterOr(f.Self, f.Right), nil
	}
}
