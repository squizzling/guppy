package aststream

import (
	"guppy/internal/ast"
)

const Package = "stream"

var Imports = []string{
	"time",
	"",
	"guppy/pkg/flow/filter",
	"guppy/pkg/interpreter",
}

var Interfaces = []ast.Interface{
	{"Stream", []ast.Node{
		{"Above", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
		}},
		{"Abs", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
		}},
		{"Aggregate", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Fn", "string", false},
			{"By", "[]string", false},
			{"AllowAllMissing", "bool", false},
			{"AllowMissing", "[]string", false},
		}},
		{"Alerts", true, []ast.Field{
			{"Object", "interpreter.Object", true},
		}},
		{"Below", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
		}},
		{"BinaryOpDouble", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Stream", "Stream", false},
			{"Op", "string", false},
			{"Other", "float64", false},
			{"Reverse", "bool", false},
		}},
		{"BinaryOpInt", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Stream", "Stream", false},
			{"Op", "string", false},
			{"Other", "int", false},
			{"Reverse", "bool", false},
		}},
		{"BinaryOpStream", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Left", "Stream", false},
			{"Op", "string", false},
			{"Right", "Stream", false},
		}},
		{"Combine", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Mode", "string", false},
		}},
		{"ConstDouble", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Value", "float64", false},
			{"Key", "map[string]string", false},
		}},
		{"ConstInt", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Value", "int", false},
			{"Key", "map[string]string", false},
		}},
		{"Count", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Sources", "[]Stream", false},
		}},
		{"Data", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"MetricName", "string", false},
			{"Filter", "filter.Filter", false},
			{"Rollup", "string", false},
			{"Extrapolation", "string", false},
			{"MaxExtrapolations", "int", false},
			{"TimeShift", "time.Duration", false},
		}},
		{"Detect", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"On", "Stream", false},
			{"Off", "Stream", false},
			{"Mode", "string", false},
			{"Annotations", "interpreter.Object", false},
			{"EventAnnotations", "interpreter.Object", false},
			{"AutoResolveAfter", "*time.Duration", false},
		}},
		{"Events", true, []ast.Field{
			{"Object", "interpreter.Object", true},
		}},
		{"Fill", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Value", "interpreter.Object", false},
			{"Duration", "int", false},
			{"MaxCount", "int", false},
		}},
		{"Generic", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Call", "string", false},
		}},
		{"IsNone", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Invert", "bool", false},
		}},
		{"Max", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Sources", "[]Stream", false},
			{"Value", "interpreter.Object", false},
		}},
		{"Mean", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Sources", "[]Stream", false},
			{"Constants", "[]interpreter.Object", false},
		}},
		{"Median", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Sources", "[]Stream", false},
			{"Constants", "[]interpreter.Object", false},
		}},
		{"Min", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Sources", "[]Stream", false},
			{"Value", "interpreter.Object", false},
		}},
		{"Percentile", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
		}},
		{"Publish", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Label", "string", false},
			{"Enable", "bool", false},
		}},
		{"Scale", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Multiple", "float64", false},
		}},
		{"Ternary", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Condition", "Stream", false},
			{"Left", "Stream", false},
			{"Right", "Stream", false},
		}},
		{"Threshold", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Value", "float64", false},
		}},
		{"TimeShift", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Offset", "time.Duration", false},
		}},
		{"Top", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
		}},
		{"Transform", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Fn", "string", false},
			{"Over", "time.Duration", false},
		}},
		{"TransformCycle", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Fn", "string", false},
			{"Cycle", "string", false},
			{"CycleStart", "*string", false},
			{"Timezone", "*string", false},
			{"PartialValues", "bool", false},
			{"ShiftCycles", "int", false},
		}},
		{"UnaryOpMinus", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Stream", "Stream", false},
		}},
		{"Union", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Sources", "[]Stream", false},
		}},
		{"When", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Predicate", "Stream", false},
			{"Lasting", "*time.Duration", false},
			{"AtLeast", "float64", false},
		}},
	}, []string{
		"interpreter.Object",
	}},
}
