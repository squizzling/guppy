package stream

import (
	"guppy/internal/interpreter"
)

type Stream interface {
	interpreter.Object

	RenderStream() string
}

func newStreamObject() interpreter.Object {
	return interpreter.NewObject(map[string]interpreter.Object{
		"fill":         methodFill{interpreter.NewObject(nil)},
		"max":          methodMax{interpreter.NewObject(nil)},
		"mean":         methodMean{interpreter.NewObject(nil)},
		"publish":      methodPublish{interpreter.NewObject(nil)},
		"sum":          methodSum{interpreter.NewObject(nil)},
		"__add__":      opAdd{interpreter.NewObject(nil), false},
		"__radd__":     opAdd{interpreter.NewObject(nil), true},
		"__mul__":      opMul{interpreter.NewObject(nil), false},
		"__rmul__":     opMul{interpreter.NewObject(nil), true},
		"__sub__":      opSub{interpreter.NewObject(nil), false},
		"__rsub__":     opSub{interpreter.NewObject(nil), true},
		"__truediv__":  opTrueDiv{interpreter.NewObject(nil), false},
		"__rtruediv__": opTrueDiv{interpreter.NewObject(nil), true},
	})
}

var _ = interpreter.FlowCall(methodFill{})
var _ = interpreter.FlowCall(methodMax{})
var _ = interpreter.FlowCall(methodMean{})
var _ = interpreter.FlowCall(methodPublish{})
var _ = interpreter.FlowCall(methodSum{})
var _ = interpreter.FlowCall(opAdd{})
var _ = interpreter.FlowCall(opMul{})
var _ = interpreter.FlowCall(opSub{})
var _ = interpreter.FlowCall(opTrueDiv{})

// unpublish will remove any publish called on a Stream. This is because a publish
// is not actually useful from a dataflow perspective.
func unpublish(s Stream) Stream {
	if p, ok := s.(*publish); ok {
		return unpublish(p.source)
	}
	return s
}
