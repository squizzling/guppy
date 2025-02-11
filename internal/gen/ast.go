package gen

import (
	"strings"
)

const Package = "ast"

var Interfaces = []Interface{
	{"Data", []ASTNode{
		{"Argument", true, []Field{
			{"Assign", "string"},
			{"Expr", "Expression"},
		}},
	}},
	{"Statement", []ASTNode{
		{"Program", true, []Field{
			{"Statements", "StatementList"},
		}},
		{"Expression", false, []Field{
			{"Assign", "[]string"},
			{"Expr", "ExpressionList"},
		}},
		{"Return", false, []Field{
			{"Expr", "Expression"},
		}},
		{"Import", false, []Field{
			// TODO
		}},
		{"Assert", false, []Field{
			{"Test", "Expression"},
			{"Raise", "Expression"},
		}},
		{"If", false, []Field{
			{"Condition", "[]Expression"},
			{"Statement", "[]StatementList"},
			{"StatementElse", "StatementList"},
		}},

		{"Function", false, []Field{
			{"Token", "string"},
			// {"Params", "[]tokenizer.Token"}, // TODO
			{"Body", "StatementList"},
		}},
		{"Decorated", false, []Field{
			// TODO
		}},

		{"List", true, []Field{
			{"Statements", "[]Statement"},
		}},
	}},
	{"Expression", []ASTNode{
		{"List", true, []Field{ // Make this a concrete return type
			{"Expressions", "[]Expression"},
			{"Tuple", "bool"},
		}},
		{"Binary", false, []Field{
			{"Left", "Expression"},
			{"Op", "tokenizer.Token"},
			{"Right", "Expression"},
		}},
		{"Grouping", false, []Field{
			{"Expr", "Expression"},
		}},
		{"Literal", false, []Field{
			{"Value", "any"},
		}},
		{"Logical", false, []Field{
			{"Left", "Expression"},
			{"Op", "tokenizer.Token"},
			{"Right", "Expression"},
		}},
		{"Ternary", false, []Field{
			{"Left", "Expression"},
			{"Cond", "Expression"},
			{"Right", "Expression"},
		}},
		{"Unary", false, []Field{
			{"Op", "tokenizer.TokenType"},
			{"Right", "Expression"},
		}},
		{"Variable", false, []Field{
			{"Identifier", "string"},
		}},
		{"Member", false, []Field{
			{"Expr", "Expression"},
			{"Identifier", "string"},
		}},
		{"Subscript", false, []Field{
			{"Expr", "Expression"},
			{"Identifier", "string"},
		}},
		{"Call", false, []Field{
			{"Expr", "Expression"},
			{"Args", "[]DataArgument"},
			{"StarArgs", "Expression"},
			{"KeywordArgs", "Expression"},
		}},
	}},
}

type Interface struct {
	Name  string
	Nodes []ASTNode
}

type ASTNode struct {
	Name        string
	NewConcrete bool
	Fields      []Field
}

type Field struct {
	Name string
	Type string
}

func IsInterface(name string) bool {
	for _, i := range Interfaces {
		if i.Name == name {
			return true
		}
	}
	return false
}

func IsInterfaceArray(name string) bool {
	if strings.HasPrefix(name, "[]") {
		return IsInterface(name[2:])
	}
	return false
}

func IsConcrete(name string) bool {
	for _, i := range Interfaces {
		for _, t := range i.Nodes {
			if i.Name+t.Name == name {
				return true
			}
		}
	}
	return false
}

func IsConcreteArray(name string) bool {
	if strings.HasPrefix(name, "[]") {
		return IsConcrete(name[2:])
	}
	return false
}
