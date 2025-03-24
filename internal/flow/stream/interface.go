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
		"above":                methodAbove{interpreter.NewObject(nil)},
		"abs":                  methodAbs{interpreter.NewObject(nil)},
		"below":                methodBelow{interpreter.NewObject(nil)},
		"between":              methodBetween{interpreter.NewObject(nil)},
		"bottom":               methodBottom{interpreter.NewObject(nil)},
		"ceil":                 methodCeil{interpreter.NewObject(nil)},
		"delta":                methodDelta{interpreter.NewObject(nil)},
		"dimensions":           methodDimensions{interpreter.NewObject(nil)},
		"double_ewma":          methodDoubleEWMA{interpreter.NewObject(nil)},
		"equals":               methodEquals{interpreter.NewObject(nil)},
		"ewma":                 methodEWMA{interpreter.NewObject(nil)},
		"fill":                 methodFill{interpreter.NewObject(nil)},
		"floor":                methodFloor{interpreter.NewObject(nil)},
		"histogram_percentile": methodHistogramPercentile{interpreter.NewObject(nil)},
		"integrate":            methodIntegrate{interpreter.NewObject(nil)},
		"log10":                methodLog10{interpreter.NewObject(nil)},
		"log":                  methodLog{interpreter.NewObject(nil)},
		"map":                  methodMap{interpreter.NewObject(nil)},
		"mean_plus_stddev":     methodMeanPlusStdDev{interpreter.NewObject(nil)},
		"not_equals":           methodNotEquals{interpreter.NewObject(nil)},
		"percentile":           methodPercentile{interpreter.NewObject(nil)},
		"pow":                  methodPow{interpreter.NewObject(nil)},
		"promote":              methodPromote{interpreter.NewObject(nil)},
		"publish":              methodPublish{interpreter.NewObject(nil)},
		"rateofchange":         methodRateofChange{interpreter.NewObject(nil)},
		"scale":                methodScale{interpreter.NewObject(nil)},
		"sqrt":                 methodSqrt{interpreter.NewObject(nil)},
		"stddev":               methodStdDev{interpreter.NewObject(nil)},
		"timeshift":            methodTimeShift{interpreter.NewObject(nil)},
		"top":                  methodTop{interpreter.NewObject(nil)},

		// Aggregations + transforms
		"count":  methodStreamAggregateTransform{interpreter.NewObject(nil), "count"},
		"max":    methodStreamAggregateTransform{interpreter.NewObject(nil), "max"},
		"median": methodStreamAggregateTransform{interpreter.NewObject(nil), "median"},
		"min":    methodStreamAggregateTransform{interpreter.NewObject(nil), "min"},
		"mean":   methodStreamAggregateTransform{interpreter.NewObject(nil), "mean"},
		"sum":    methodStreamAggregateTransform{interpreter.NewObject(nil), "sum"},

		// Math operations
		"__add__":      methodStreamOp{interpreter.NewObject(nil), "+", false},
		"__radd__":     methodStreamOp{interpreter.NewObject(nil), "+", true},
		"__mul__":      methodStreamOp{interpreter.NewObject(nil), "*", false},
		"__rmul__":     methodStreamOp{interpreter.NewObject(nil), "*", true},
		"__sub__":      methodStreamOp{interpreter.NewObject(nil), "-", false},
		"__rsub__":     methodStreamOp{interpreter.NewObject(nil), "-", true},
		"__truediv__":  methodStreamOp{interpreter.NewObject(nil), "/", false},
		"__rtruediv__": methodStreamOp{interpreter.NewObject(nil), "/", true},

		"__unary_minus__": methodStreamUnaryMinus{interpreter.NewObject(nil)},
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
