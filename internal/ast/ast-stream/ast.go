package aststream

import (
	"guppy/internal/ast"
)

const Package = "stream"

var Imports = []string{
	"time",
	"",
	"guppy/pkg/flow/filter",
	"guppy/pkg/interpreter/itypes",
}

var Interfaces = []ast.Interface{
	{"Stream", []ast.Node{
		// Top level
		{"FuncAbs", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Sources", "[]Stream", false},
		}},
		{"FuncAlerts", true, []ast.Field{
			{"Object", "itypes.Object", true},
		}},
		{"FuncCombine", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Mode", "string", false},
		}},
		{"FuncConstDouble", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Value", "float64", false},
			{"Key", "map[string]string", false},
		}},
		{"FuncConstInt", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Value", "int", false},
			{"Key", "map[string]string", false},
		}},
		{"FuncCount", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Sources", "[]Stream", false},
		}},
		{"FuncData", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"MetricName", "string", false},
			{"Filter", "filter.Filter", false},
			{"Rollup", "string", false},
			{"Extrapolation", "string", false},
			{"MaxExtrapolations", "int", false},
			{"TimeShift", "time.Duration", false},
		}},
		{"FuncDetect", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"On", "Stream", false},
			{"Off", "Stream", false},
			{"Mode", "string", false},
			{"Annotations", "itypes.Object", false},
			{"EventAnnotations", "itypes.Object", false},
			{"AutoResolveAfter", "*time.Duration", false},
		}},
		{"FuncEvents", true, []ast.Field{
			{"Object", "itypes.Object", true},
		}},
		{"FuncMax", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Sources", "[]Stream", false},
			{"Value", "itypes.Object", false},
		}},
		{"FuncMean", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Sources", "[]Stream", false},
			{"Constants", "[]itypes.Object", false},
		}},
		{"FuncMedian", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Sources", "[]Stream", false},
			{"Constants", "[]itypes.Object", false},
		}},
		{"FuncMin", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Sources", "[]Stream", false},
			{"Value", "itypes.Object", false},
		}},
		{"FuncSum", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Sources", "[]Stream", false},
			{"Constant", "float64", false},
		}},
		{"FuncThreshold", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Value", "float64", false},
		}},
		{"FuncUnion", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Sources", "[]Stream", false},
		}},
		{"FuncWhen", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Predicate", "Stream", false},
			{"Lasting", "*time.Duration", false},
			{"AtLeast", "float64", false},
		}},

		// Chained methods
		{"MethodAbove", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
		}},
		{"MethodAbs", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
		}},
		{"MethodAggregate", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Fn", "string", false},
			{"By", "[]string", false},
			{"AllowAllMissing", "bool", false},
			{"AllowMissing", "[]string", false},
		}},
		{"MethodBelow", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
		}},
		{"MethodFill", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Value", "itypes.Object", false},
			{"Duration", "int", false},
			{"MaxCount", "int", false},
		}},
		{"MethodGeneric", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Call", "string", false},
		}},
		{"MethodPercentile", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
		}},
		{"MethodPublish", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Label", "string", false},
			{"Enable", "bool", false},
		}},
		{"MethodScale", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Multiple", "float64", false},
		}},
		{"MethodTimeShift", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Offset", "time.Duration", false},
		}},
		{"MethodTop", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
		}},
		{"MethodTransform", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Fn", "string", false},
			{"Over", "time.Duration", false},
		}},
		{"MethodTransformCycle", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Source", "Stream", false},
			{"Fn", "string", false},
			{"Cycle", "string", false},
			{"CycleStart", "*string", false},
			{"Timezone", "*string", false},
			{"PartialValues", "bool", false},
			{"ShiftCycles", "int", false},
		}},

		// Operations
		{"BinaryOpDouble", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"ObjectStreamTernary", "*ObjectStreamTernary", true},
			{"Stream", "Stream", false},
			{"Op", "string", false},
			{"Other", "float64", false},
			{"Reverse", "bool", false},
		}},
		{"BinaryOpInt", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"ObjectStreamTernary", "*ObjectStreamTernary", true},
			{"Stream", "Stream", false},
			{"Op", "string", false},
			{"Other", "int", false},
			{"Reverse", "bool", false},
		}},
		{"BinaryOpStream", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"ObjectStreamTernary", "*ObjectStreamTernary", true},
			{"Left", "Stream", false},
			{"Op", "string", false},
			{"Right", "Stream", false},
		}},
		{"IsNone", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"ObjectStreamTernary", "*ObjectStreamTernary", true},
			{"Source", "Stream", false},
			{"Invert", "bool", false},
		}},
		{"Ternary", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Left", "Stream", false},
			{"Condition", "Stream", false},
			{"Right", "Stream", false},
		}},
		{"UnaryOpMinus", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Stream", "Stream", false},
		}},
	}, []string{
		"itypes.Object",
	}},
}
