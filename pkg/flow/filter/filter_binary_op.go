package filter

import (
	"fmt"
	"strings"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type and struct {
	itypes.Object

	filters []Filter
}

func NewAnd(left Filter, right Filter) Filter {
	var filters []Filter
	if leftAnd, ok := left.(*and); ok {
		filters = append(filters, leftAnd.filters...)
	} else {
		filters = append(filters, left)
	}
	if rightAnd, ok := right.(*and); ok {
		filters = append(filters, rightAnd.filters...)
	} else {
		filters = append(filters, right)
	}

	return &and{
		Object:  prototypeFilter,
		filters: filters,
	}
}

func (a *and) Repr() string {
	var s []string
	for _, f := range a.filters {
		s = append(s, f.Repr())
	}
	return fmt.Sprintf("(%s)", strings.Join(s, " and "))
}

type or struct {
	itypes.Object

	filters []Filter
}

func NewOr(left Filter, right Filter) Filter {
	var filters []Filter
	if leftOr, ok := left.(*or); ok {
		filters = append(filters, leftOr.filters...)
	} else {
		filters = append(filters, left)
	}
	if rightOr, ok := right.(*or); ok {
		filters = append(filters, rightOr.filters...)
	} else {
		filters = append(filters, right)
	}

	return &or{
		Object:  prototypeFilter,
		filters: filters,
	}
}

func (o *or) Repr() string {
	var s []string
	for _, f := range o.filters {
		s = append(s, f.Repr())
	}
	return fmt.Sprintf("(%s)", strings.Join(s, " or "))
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
