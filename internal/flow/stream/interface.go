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
		"fill":    methodFill{interpreter.NewObject(nil)},
		"max":     methodMax{interpreter.NewObject(nil)},
		"mean":    methodMax{interpreter.NewObject(nil)},
		"publish": methodPublish{interpreter.NewObject(nil)},
		"sum":     methodSum{interpreter.NewObject(nil)},
	})
}

var _ = interpreter.FlowCall(methodFill{})
var _ = interpreter.FlowCall(methodMax{})
var _ = interpreter.FlowCall(methodMean{})
var _ = interpreter.FlowCall(methodPublish{})
var _ = interpreter.FlowCall(methodSum{})

// unpublish will remove any publish called on a Stream. This is because a publish
// is not actually useful from a dataflow perspective.
func unpublish(s Stream) Stream {
	if p, ok := s.(*publish); ok {
		return unpublish(p.source)
	}
	return s
}
