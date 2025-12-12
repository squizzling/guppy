package astfilter

import (
	"github.com/squizzling/guppy/internal/ast"
)

const Package = "filter"

var Imports = []string{
	"github.com/squizzling/guppy/pkg/interpreter/itypes",
}

var Interfaces = []ast.Interface{
	{"Filter", []ast.Node{
		{"Not", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Right", "Filter", false},
		}},
		{"KeyValue", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Key", "string", false},
			{"Values", "[]string", false},
			{"MatchMissing", "bool", false},
		}},
		{"Partition", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Index", "int", false},
			{"Total", "int", false},
		}},
		{"And", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Filters", "[]Filter", false},
		}},
		{"Or", true, []ast.Field{
			{"Object", "itypes.Object", true},
			{"Filters", "[]Filter", false},
		}},
	}, []string{
		"itypes.Object",
	}},
}
