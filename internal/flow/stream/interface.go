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
		// misc
		"fill":      methodFill{interpreter.NewObject(nil)},
		"publish":   methodPublish{interpreter.NewObject(nil)},
		"scale":     methodScale{interpreter.NewObject(nil)},
		"timeshift": methodTimeShift{interpreter.NewObject(nil)},

		// Aggregations + transforms
		"count": methodStreamAggregateTransform{interpreter.NewObject(nil), "count"},
		"max":   methodStreamAggregateTransform{interpreter.NewObject(nil), "max"},
		"min":   methodStreamAggregateTransform{interpreter.NewObject(nil), "min"},
		"mean":  methodStreamAggregateTransform{interpreter.NewObject(nil), "mean"},
		"sum":   methodStreamAggregateTransform{interpreter.NewObject(nil), "sum"},

		// Math operations
		"__add__":      methodStreamOp{interpreter.NewObject(nil), "+", false},
		"__radd__":     methodStreamOp{interpreter.NewObject(nil), "+", true},
		"__mul__":      methodStreamOp{interpreter.NewObject(nil), "*", false},
		"__rmul__":     methodStreamOp{interpreter.NewObject(nil), "*", true},
		"__sub__":      methodStreamOp{interpreter.NewObject(nil), "-", false},
		"__rsub__":     methodStreamOp{interpreter.NewObject(nil), "-", true},
		"__truediv__":  methodStreamOp{interpreter.NewObject(nil), "/", false},
		"__rtruediv__": methodStreamOp{interpreter.NewObject(nil), "/", true},
	})
}

var _ = interpreter.FlowCall(methodFill{})
var _ = interpreter.FlowCall(methodPublish{})
var _ = interpreter.FlowCall(methodScale{})
var _ = interpreter.FlowCall(methodStreamAggregateTransform{})
var _ = interpreter.FlowCall(methodStreamOp{})

// unpublish will remove any publish called on a Stream. This is because a publish
// is not actually useful from a dataflow perspective.
func unpublish(s Stream) Stream {
	if p, ok := s.(*publish); ok {
		return unpublish(p.source)
	}
	return s
}
