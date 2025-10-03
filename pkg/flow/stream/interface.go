package stream

import (
	"guppy/pkg/interpreter"
)

// Stream[Alert]
func newStreamAlertObject() interpreter.Object {
	return interpreter.NewObject(map[string]interpreter.Object{
		"publish": methodPublish{interpreter.NewObject(nil)},
	})
}

// Stream[bool]
func newStreamBoolObject() interpreter.Object {
	return interpreter.NewObject(map[string]interpreter.Object{
		"publish":   methodPublish{interpreter.NewObject(nil)},
		"timeshift": methodTimeShift{interpreter.NewObject(nil)},

		// generic
		"dimensions": methodGeneric{interpreter.NewObject(nil), "dimensions"},
		"equals":     methodGeneric{interpreter.NewObject(nil), "equals"},
		"map":        methodGeneric{interpreter.NewObject(nil), "map"},
		"not_equals": methodGeneric{interpreter.NewObject(nil), "not_equals"},
		"promote":    methodGeneric{interpreter.NewObject(nil), "promote"},

		// Aggregations + transforms
		"count": methodStreamAggregateTransform{interpreter.NewObject(nil), "count"},

		"__binary_and__": methodStreamOpBool{interpreter.NewObject(nil), "and", false},
		"__binary_or__":  methodStreamOpBool{interpreter.NewObject(nil), "or", false},

		// Comparison operations
		"__eq__":  methodStreamOpBool{interpreter.NewObject(nil), "==", false},
		"__req__": methodStreamOpBool{interpreter.NewObject(nil), "==", true},
		"__ne__":  methodStreamOpBool{interpreter.NewObject(nil), "!=", false},
		"__rne__": methodStreamOpBool{interpreter.NewObject(nil), "!=", true},

		// is/is not
		"__is__":     methodStreamIs{interpreter.NewObject(nil), false, false},
		"__isnot__":  methodStreamIs{interpreter.NewObject(nil), true, false},
		"__ris__":    methodStreamIs{interpreter.NewObject(nil), false, true},
		"__risnot__": methodStreamIs{interpreter.NewObject(nil), true, true},

		// Ternary
		"__ternary__": methodStreamOpTernary{interpreter.NewObject(nil)},
	})
}

// Stream[Unspecified]
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
		"variance":             methodGeneric{interpreter.NewObject(nil), "variance"},

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
		"__ge__":  methodStreamOpBool{interpreter.NewObject(nil), ">=", false},
		"__rge__": methodStreamOpBool{interpreter.NewObject(nil), ">=", true},
		"__gt__":  methodStreamOpBool{interpreter.NewObject(nil), ">", false},
		"__rgt__": methodStreamOpBool{interpreter.NewObject(nil), ">", true},
		"__le__":  methodStreamOpBool{interpreter.NewObject(nil), "<=", false},
		"__rle__": methodStreamOpBool{interpreter.NewObject(nil), "<=", true},
		"__lt__":  methodStreamOpBool{interpreter.NewObject(nil), "<", false},
		"__rlt__": methodStreamOpBool{interpreter.NewObject(nil), "<", true},
		"__eq__":  methodStreamOpBool{interpreter.NewObject(nil), "==", false},
		"__req__": methodStreamOpBool{interpreter.NewObject(nil), "==", true},
		"__ne__":  methodStreamOpBool{interpreter.NewObject(nil), "!=", false},
		"__rne__": methodStreamOpBool{interpreter.NewObject(nil), "!=", true},

		// is/is not
		"__is__":     methodStreamIs{interpreter.NewObject(nil), false, false},
		"__isnot__":  methodStreamIs{interpreter.NewObject(nil), true, false},
		"__ris__":    methodStreamIs{interpreter.NewObject(nil), false, true},
		"__risnot__": methodStreamIs{interpreter.NewObject(nil), true, true},

		// Ternary
		"__ternary__": methodStreamOpTernary{interpreter.NewObject(nil)},
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
	return s // TODO: Understand if this should be done or not.
	if p, ok := s.(*StreamMethodPublish); ok {
		return unpublish(p.Source)
	}
	return s
}
