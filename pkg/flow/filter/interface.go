package filter

import (
	"guppy/pkg/interpreter/itypes"
)

type Filter interface {
	itypes.Object

	RenderFilter() string
}

func newFilterObject() itypes.Object {
	return itypes.NewObject(map[string]itypes.Object{
		"__binary_and__":       methodBinaryAnd{},
		"__binary_or__":        methodBinaryOr{},
		"__unary_binary_not__": methodBinaryNot{},
		"__eq__":               methodBinaryEqual{invert: false},
		"__ne__":               methodBinaryEqual{invert: true},
		"__is__":               methodBinaryIs{invert: false},
		"__isnot__":            methodBinaryIs{invert: true},
	})
}

var _ = itypes.FlowCall(methodBinaryAnd{})
var _ = itypes.FlowCall(methodBinaryOr{})
var _ = itypes.FlowCall(methodBinaryNot{})
var _ = itypes.FlowCall(methodBinaryEqual{})
