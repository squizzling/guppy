package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

// Stream[Alert]
func newStreamAlertObject() itypes.Object {
	return itypes.NewObject(map[string]itypes.Object{
		"publish": methodPublish{itypes.NewObject(nil)},
	})
}

// Stream[bool]
func newStreamBoolObject() itypes.Object {
	return itypes.NewObject(map[string]itypes.Object{
		"publish":   methodPublish{itypes.NewObject(nil)},
		"timeshift": methodTimeShift{itypes.NewObject(nil)},

		// generic
		"dimensions": methodGeneric{itypes.NewObject(nil), "dimensions"},
		"equals":     methodGeneric{itypes.NewObject(nil), "equals"},
		"map":        methodGeneric{itypes.NewObject(nil), "map"},
		"not_equals": methodGeneric{itypes.NewObject(nil), "not_equals"},
		"promote":    methodGeneric{itypes.NewObject(nil), "promote"},

		// Aggregations + transforms
		"count": methodStreamAggregateTransform{itypes.NewObject(nil), "count"},

		"__binary_and__": methodStreamOpBool{itypes.NewObject(nil), "and", false},
		"__binary_or__":  methodStreamOpBool{itypes.NewObject(nil), "or", false},

		// Comparison operations
		"__eq__":  methodStreamOpBool{itypes.NewObject(nil), "==", false},
		"__req__": methodStreamOpBool{itypes.NewObject(nil), "==", true},
		"__ne__":  methodStreamOpBool{itypes.NewObject(nil), "!=", false},
		"__rne__": methodStreamOpBool{itypes.NewObject(nil), "!=", true},

		// is/is not
		"__is__":     methodStreamIs{itypes.NewObject(nil), false, false},
		"__isnot__":  methodStreamIs{itypes.NewObject(nil), true, false},
		"__ris__":    methodStreamIs{itypes.NewObject(nil), false, true},
		"__risnot__": methodStreamIs{itypes.NewObject(nil), true, true},
	})
}

// Stream[Unspecified]
func newStreamObject() itypes.Object {
	return itypes.NewObject(map[string]itypes.Object{
		// misc
		"above":      methodAbove{itypes.NewObject(nil)},
		"abs":        methodAbs{itypes.NewObject(nil)},
		"below":      methodBelow{itypes.NewObject(nil)},
		"percentile": methodPercentile{itypes.NewObject(nil)},
		"publish":    methodPublish{itypes.NewObject(nil)},
		"timeshift":  methodTimeShift{itypes.NewObject(nil)},
		"top":        methodTop{itypes.NewObject(nil)},

		// generic
		"between":              methodGeneric{itypes.NewObject(nil), "between"},
		"bottom":               methodGeneric{itypes.NewObject(nil), "bottom"},
		"ceil":                 methodGeneric{itypes.NewObject(nil), "ceil"},
		"delta":                methodGeneric{itypes.NewObject(nil), "delta"},
		"dimensions":           methodGeneric{itypes.NewObject(nil), "dimensions"},
		"double_ewma":          methodGeneric{itypes.NewObject(nil), "double_ewma"},
		"equals":               methodGeneric{itypes.NewObject(nil), "equals"},
		"ewma":                 methodGeneric{itypes.NewObject(nil), "ewma"},
		"fill":                 methodGeneric{itypes.NewObject(nil), "fill"},
		"floor":                methodGeneric{itypes.NewObject(nil), "floor"},
		"histogram_percentile": methodGeneric{itypes.NewObject(nil), "histogram_percentile"},
		"integrate":            methodGeneric{itypes.NewObject(nil), "integrate"},
		"log10":                methodGeneric{itypes.NewObject(nil), "log10"},
		"log":                  methodGeneric{itypes.NewObject(nil), "log"},
		"map":                  methodGeneric{itypes.NewObject(nil), "map"},
		"mean_plus_stddev":     methodGeneric{itypes.NewObject(nil), "mean_plus_stddev"},
		"not_equals":           methodGeneric{itypes.NewObject(nil), "not_equals"},
		"pow":                  methodGeneric{itypes.NewObject(nil), "pow"},
		"promote":              methodGeneric{itypes.NewObject(nil), "promote"},
		"rateofchange":         methodGeneric{itypes.NewObject(nil), "rateofchange"},
		"scale":                methodGeneric{itypes.NewObject(nil), "scale"},
		"sqrt":                 methodGeneric{itypes.NewObject(nil), "sqrt"},
		"stddev":               methodGeneric{itypes.NewObject(nil), "stddev"},
		"variance":             methodGeneric{itypes.NewObject(nil), "variance"},

		// Aggregations + transforms
		"count":  methodStreamAggregateTransform{itypes.NewObject(nil), "count"},
		"max":    methodStreamAggregateTransform{itypes.NewObject(nil), "max"},
		"median": methodStreamAggregateTransform{itypes.NewObject(nil), "median"},
		"min":    methodStreamAggregateTransform{itypes.NewObject(nil), "min"},
		"mean":   methodStreamAggregateTransform{itypes.NewObject(nil), "mean"},
		"sum":    methodStreamAggregateTransform{itypes.NewObject(nil), "sum"},

		// Math operations
		"__add__":      methodStreamOp{itypes.NewObject(nil), "+", false},
		"__radd__":     methodStreamOp{itypes.NewObject(nil), "+", true},
		"__mul__":      methodStreamOp{itypes.NewObject(nil), "*", false},
		"__rmul__":     methodStreamOp{itypes.NewObject(nil), "*", true},
		"__sub__":      methodStreamOp{itypes.NewObject(nil), "-", false},
		"__rsub__":     methodStreamOp{itypes.NewObject(nil), "-", true},
		"__truediv__":  methodStreamOp{itypes.NewObject(nil), "/", false},
		"__rtruediv__": methodStreamOp{itypes.NewObject(nil), "/", true},

		"__unary_minus__": methodStreamUnaryMinus{itypes.NewObject(nil)},

		// Comparison operations
		"__ge__":  methodStreamOpBool{itypes.NewObject(nil), ">=", false},
		"__rge__": methodStreamOpBool{itypes.NewObject(nil), ">=", true},
		"__gt__":  methodStreamOpBool{itypes.NewObject(nil), ">", false},
		"__rgt__": methodStreamOpBool{itypes.NewObject(nil), ">", true},
		"__le__":  methodStreamOpBool{itypes.NewObject(nil), "<=", false},
		"__rle__": methodStreamOpBool{itypes.NewObject(nil), "<=", true},
		"__lt__":  methodStreamOpBool{itypes.NewObject(nil), "<", false},
		"__rlt__": methodStreamOpBool{itypes.NewObject(nil), "<", true},
		"__eq__":  methodStreamOpBool{itypes.NewObject(nil), "==", false},
		"__req__": methodStreamOpBool{itypes.NewObject(nil), "==", true},
		"__ne__":  methodStreamOpBool{itypes.NewObject(nil), "!=", false},
		"__rne__": methodStreamOpBool{itypes.NewObject(nil), "!=", true},

		// is/is not
		"__is__":     methodStreamIs{itypes.NewObject(nil), false, false},
		"__isnot__":  methodStreamIs{itypes.NewObject(nil), true, false},
		"__ris__":    methodStreamIs{itypes.NewObject(nil), false, true},
		"__risnot__": methodStreamIs{itypes.NewObject(nil), true, true},
	})
}

var _ = itypes.FlowCall(methodFill{})
var _ = itypes.FlowCall(methodPublish{})
var _ = itypes.FlowCall(methodScale{})
var _ = itypes.FlowCall(methodStreamAggregateTransform{})
var _ = itypes.FlowCall(methodStreamOp{})

// unpublish will remove any publish called on a Stream. This is because a publish
// is not actually useful from a dataflow perspective.
func unpublish(s Stream) Stream {
	return s // TODO: Understand if this should be done or not.
	if p, ok := s.(*StreamMethodPublish); ok {
		return unpublish(p.Source)
	}
	return s
}
