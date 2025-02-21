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
		"__unary_binary_not__": methodBinaryNot{},
	})
}

var _ = interpreter.FlowCall(methodBinaryAnd{})
var _ = interpreter.FlowCall(methodBinaryNot{})
