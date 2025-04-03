package stream

import (
	"guppy/internal/interpreter"
)

func newStreamObject() interpreter.Object {
	return interpreter.NewObject(map[string]interpreter.Object{
		// misc
		"above":      methodAbove{interpreter.NewObject(nil)},
		"abs":        methodAbs{interpreter.NewObject(nil)},
		"below":      methodBelow{interpreter.NewObject(nil)},
		"percentile": methodPercentile{interpreter.NewObject(nil)},
		"publish":    methodPublish{interpreter.NewObject(nil)},
		"timeshift":  methodTimeShift{interpreter.NewObject(nil)},
		"top":        methodTop{interpreter.NewObject(nil)},

		// generic
		"between":              methodGeneric{interpreter.NewObject(nil), "between"},
		"bottom":               methodGeneric{interpreter.NewObject(nil), "bottom"},
		"ceil":                 methodGeneric{interpreter.NewObject(nil), "ceil"},
		"delta":                methodGeneric{interpreter.NewObject(nil), "delta"},
		"dimensions":           methodGeneric{interpreter.NewObject(nil), "dimensions"},
		"double_ewma":          methodGeneric{interpreter.NewObject(nil), "double_ewma"},
		"equals":               methodGeneric{interpreter.NewObject(nil), "equals"},
		"ewma":                 methodGeneric{interpreter.NewObject(nil), "ewma"},
		"fill":                 methodGeneric{interpreter.NewObject(nil), "fill"},
		"floor":                methodGeneric{interpreter.NewObject(nil), "floor"},
		"histogram_percentile": methodGeneric{interpreter.NewObject(nil), "histogram_percentile"},
		"integrate":            methodGeneric{interpreter.NewObject(nil), "integrate"},
		"log10":                methodGeneric{interpreter.NewObject(nil), "log10"},
		"log":                  methodGeneric{interpreter.NewObject(nil), "log"},
		"map":                  methodGeneric{interpreter.NewObject(nil), "map"},
		"mean_plus_stddev":     methodGeneric{interpreter.NewObject(nil), "mean_plus_stddev"},
		"not_equals":           methodGeneric{interpreter.NewObject(nil), "not_equals"},
		"pow":                  methodGeneric{interpreter.NewObject(nil), "pow"},
		"promote":              methodGeneric{interpreter.NewObject(nil), "promote"},
		"rateofchange":         methodGeneric{interpreter.NewObject(nil), "rateofchange"},
		"scale":                methodGeneric{interpreter.NewObject(nil), "scale"},
		"sqrt":                 methodGeneric{interpreter.NewObject(nil), "sqrt"},
		"stddev":               methodGeneric{interpreter.NewObject(nil), "stddev"},

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

		// Comparison operations
		"__ge__":  methodStreamOp{interpreter.NewObject(nil), ">=", false},
		"__rge__": methodStreamOp{interpreter.NewObject(nil), ">=", true},
		"__gt__":  methodStreamOp{interpreter.NewObject(nil), ">", false},
		"__rgt__": methodStreamOp{interpreter.NewObject(nil), ">", true},
		"__le__":  methodStreamOp{interpreter.NewObject(nil), "<=", false},
		"__rle__": methodStreamOp{interpreter.NewObject(nil), "<=", true},
		"__lt__":  methodStreamOp{interpreter.NewObject(nil), "<", false},
		"__rlt__": methodStreamOp{interpreter.NewObject(nil), "<", true},
		"__eq__":  methodStreamOp{interpreter.NewObject(nil), "==", false},
		"__req__": methodStreamOp{interpreter.NewObject(nil), "==", true},
		"__ne__":  methodStreamOp{interpreter.NewObject(nil), "!=", false},
		"__rne__": methodStreamOp{interpreter.NewObject(nil), "!=", true},
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
	if p, ok := s.(*StreamPublish); ok {
		return unpublish(p.Source)
	}
	return s
}
