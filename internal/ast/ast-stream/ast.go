package aststream

import (
	"guppy/internal/ast"
)

const Package = "stream"

var Imports = []string{
	"guppy/internal/flow/filter",
	"guppy/internal/interpreter",
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
		}},
		{"Alerts", true, []ast.Field{
			{"Object", "interpreter.Object", true},
		}},
		{"Below", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
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
		{"Data", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"MetricName", "string", false},
			{"Filter", "filter.Filter", false},
			{"Rollup", "string", false},
			{"Extrapolation", "string", false},
			{"MaxExtrapolations", "int", false},
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
		{"Max", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Sources", "[]Stream", false},
			{"Value", "interpreter.Object", false},
		}},
		{"MathOpDouble", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Stream", "Stream", false},
			{"Op", "string", false},
			{"Other", "float64", false},
			{"Reverse", "bool", false},
		}},
		{"MathOpInt", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Stream", "Stream", false},
			{"Op", "string", false},
			{"Other", "int", false},
			{"Reverse", "bool", false},
		}},
		{"MathOpStream", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Left", "Stream", false},
			{"Op", "string", false},
			{"Right", "Stream", false},
		}},
		{"MathOpUnaryMinus", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Stream", "Stream", false},
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
		{"Threshold", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Value", "float64", false},
		}},
		{"TimeShift", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Offset", "string", false},
		}},
		{"Top", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
		}},
		{"Transform", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Source", "Stream", false},
			{"Fn", "string", false},
			{"Over", "string", false},
		}},
		{"Union", true, []ast.Field{
			{"Object", "interpreter.Object", true},
			{"Sources", "[]Stream", false},
		}},
	}, []string{
		"interpreter.Object",
	}},
}
