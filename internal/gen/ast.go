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
		{"ArgumentList", true, []Field{
			{"Args", "[]Expression"},
			{"NamedArgs", "[]*DataArgument"},
			{"StarArg", "Expression"},
			{"KeywordArg", "Expression"},
		}},
		{"ImportAs", true, []Field{
			{"Name", "[]string"},
			{"As", "string"},
		}},
		{"ListIter", true, []Field{
			{"For", "*DataListFor"},
			{"If", "*DataListIf"},
		}},
		{"ListFor", true, []Field{
			{"Idents", "[]string"},
			{"Expr", "Expression"},
			{"Iter", "*DataListIter"},
		}},
		{"ListIf", true, []Field{
			{"Expr", "Expression"},
			{"Iter", "*DataListIter"},
		}},
		{"Parameter", true, []Field{
			{"Name", "string"},
			{"Type", "string"},
			{"Default", "Expression"},
			{"StarArg", "bool"},
			{"KeywordArg", "bool"},
		}},
		{"ParameterList", true, []Field{
			{"List", "[]*DataParameter"},
		}},
		{"Subscript", true, []Field{
			{"Start", "Expression"},
			{"End", "Expression"},
			{"Range", "bool"},
		}},
	}},
	{"Statement", []ASTNode{
		{"Program", true, []Field{
			{"Statements", "*StatementList"},
		}},
		{"Expression", false, []Field{
			{"Assign", "[]string"},
			{"Expr", "Expression"},
		}},
		{"Return", false, []Field{
			{"Expr", "Expression"},
		}},
		{"ImportFrom", false, []Field{
			{"From", "[]string"},
			{"Imports", "[]*DataImportAs"},
		}},
		{"ImportFromStar", false, []Field{
			{"From", "[]string"},
		}},
		{"ImportNames", false, []Field{
			{"Imports", "[]*DataImportAs"},
		}},
		{"Assert", false, []Field{
			{"Test", "Expression"},
			{"Raise", "Expression"},
		}},
		{"If", false, []Field{
			{"Condition", "[]Expression"},
			{"Statement", "[]Statement"},
			{"StatementElse", "Statement"},
		}},

		{"Function", false, []Field{
			{"Token", "string"},
			{"Params", "*DataParameterList"},
			{"Body", "Statement"},
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
		{"ListMaker", false, []Field{
			{"Expr", "Expression"},
			{"For", "*DataListFor"},
		}},
		{"Binary", false, []Field{
			{"Left", "Expression"},
			{"Op", "tokenizer.Token"},
			{"Right", "Expression"},
		}},
		{"Dict", false, []Field{
			{"Keys", "[]Expression"},
			{"Values", "[]Expression"},
		}},
		{"Grouping", false, []Field{
			{"Expr", "Expression"},
		}},
		{"Lambda", false, []Field{
			{"Identifier", "string"},
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
			{"Start", "Expression"},
			{"End", "Expression"},
			{"Range", "bool"},
		}},
		{"Call", false, []Field{
			{"Expr", "Expression"},
			{"Args", "[]Expression"},
			{"NamedArgs", "[]*DataArgument"},
			{"StarArg", "Expression"},
			{"KeywordArg", "Expression"},
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
	name = strings.TrimLeft(name, "*")
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
