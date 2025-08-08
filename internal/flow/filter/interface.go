package filter

import (
	"guppy/internal/interpreter"
)

type Filter interface {
	interpreter.Object

	RenderFilter() string
}

func newFilterObject() interpreter.Object {
	return interpreter.NewObject(map[string]interpreter.Object{
		"__binary_and__":       methodBinaryAnd{},
		"__binary_or__":        methodBinaryOr{},
		"__unary_binary_not__": methodBinaryNot{},
		"__eq__":               methodBinaryEqual{invert: false},
		"__ne__":               methodBinaryEqual{invert: true},
		"__is__":               methodBinaryIs{invert: false},
		"__isnot__":            methodBinaryIs{invert: true},
	})
}

var _ = interpreter.FlowCall(methodBinaryAnd{})
var _ = interpreter.FlowCall(methodBinaryOr{})
var _ = interpreter.FlowCall(methodBinaryNot{})
var _ = interpreter.FlowCall(methodBinaryEqual{})
