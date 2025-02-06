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
		"__binary_and__": methodBinaryAnd{},
	})
}

var _ = interpreter.FlowCall(methodBinaryAnd{})
