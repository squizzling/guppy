package astflow

import (
	"guppy/internal/ast"
)

const Package = "ast"

var Interfaces = ast.Interfaces{
	{"Data", []ast.Node{
		{"Argument", true, []ast.Field{
			{"Assign", "string", false},
			{"Expr", "Expression", false},
		}},
		{"ArgumentList", true, []ast.Field{
			{"Args", "[]Expression", false},
			{"NamedArgs", "[]*DataArgument", false},
			{"StarArg", "Expression", false},
			{"KeywordArg", "Expression", false},
		}},
		{"ImportAs", true, []ast.Field{
			{"Name", "[]string", false},
			{"As", "string", false},
		}},
		{"ListIter", true, []ast.Field{
			{"For", "*DataListFor", false},
			{"If", "*DataListIf", false},
		}},
		{"ListFor", true, []ast.Field{
			{"Idents", "[]string", false},
			{"Expr", "Expression", false},
			{"Iter", "*DataListIter", false},
		}},
		{"ListIf", true, []ast.Field{
			{"Expr", "Expression", false},
			{"Iter", "*DataListIter", false},
		}},
		{"Parameter", true, []ast.Field{
			{"Name", "string", false},
			{"Type", "string", false},
			{"Default", "Expression", false},
			{"StarArg", "bool", false},
			{"KeywordArg", "bool", false},
		}},
		{"ParameterList", true, []ast.Field{
			{"List", "[]*DataParameter", false},
		}},
		{"Subscript", true, []ast.Field{
			{"Start", "Expression", false},
			{"End", "Expression", false},
			{"Range", "bool", false},
		}},
	}, nil},
	{"Statement", []ast.Node{
		{"Program", true, []ast.Field{
			{"Statements", "*StatementList", false},
		}},
		{"Expression", false, []ast.Field{
			{"Assign", "[]string", false},
			{"Expr", "Expression", false},
		}},
		{"Return", false, []ast.Field{
			{"Expr", "Expression", false},
		}},
		{"ImportFrom", false, []ast.Field{
			{"From", "[]string", false},
			{"Imports", "[]*DataImportAs", false},
		}},
		{"ImportFromStar", false, []ast.Field{
			{"From", "[]string", false},
		}},
		{"ImportNames", false, []ast.Field{
			{"Imports", "[]*DataImportAs", false},
		}},
		{"Assert", false, []ast.Field{
			{"Test", "Expression", false},
			{"Raise", "Expression", false},
		}},
		{"If", false, []ast.Field{
			{"Condition", "[]Expression", false},
			{"Statement", "[]Statement", false},
			{"StatementElse", "Statement", false},
		}},

		{"Function", false, []ast.Field{
			{"Token", "string", false},
			{"Params", "*DataParameterList", false},
			{"Body", "Statement", false},
		}},
		{"Decorated", false, []ast.Field{
			// TODO
		}},

		{"List", true, []ast.Field{
			{"Statements", "[]Statement", false},
		}},
	}, nil},
	{"Expression", []ast.Node{
		{"List", false, []ast.Field{
			{"Expressions", "[]Expression", false},
		}},
		{"ListMaker", false, []ast.Field{
			{"Expr", "Expression", false},
			{"For", "*DataListFor", false},
		}},
		{"Binary", false, []ast.Field{
			{"Left", "Expression", false},
			{"Op", "tokenizer.Token", false},
			{"Right", "Expression", false},
		}},
		{"Dict", false, []ast.Field{
			{"Keys", "[]Expression", false},
			{"Values", "[]Expression", false},
		}},
		{"Grouping", false, []ast.Field{
			{"Expr", "Expression", false},
		}},
		{"Lambda", false, []ast.Field{
			{"Identifier", "string", false},
			{"Expr", "Expression", false},
		}},
		{"Literal", false, []ast.Field{
			{"Value", "any", false},
		}},
		{"Logical", false, []ast.Field{
			{"Left", "Expression", false},
			{"Op", "tokenizer.Token", false},
			{"Right", "Expression", false},
		}},
		{"Ternary", false, []ast.Field{
			{"Left", "Expression", false},
			{"Cond", "Expression", false},
			{"Right", "Expression", false},
		}},
		{"Tuple", false, []ast.Field{
			{"Expressions", "[]Expression", false},
		}},
		{"Unary", false, []ast.Field{
			{"Op", "tokenizer.TokenType", false},
			{"Right", "Expression", false},
		}},
		{"Variable", false, []ast.Field{
			{"Identifier", "string", false},
		}},
		{"Member", false, []ast.Field{
			{"Expr", "Expression", false},
			{"Identifier", "string", false},
		}},
		{"Subscript", false, []ast.Field{
			{"Expr", "Expression", false},
			{"Start", "Expression", false},
			{"End", "Expression", false},
			{"Range", "bool", false},
		}},
		{"Call", false, []ast.Field{
			{"Expr", "Expression", false},
			{"Args", "[]Expression", false},
			{"NamedArgs", "[]*DataArgument", false},
			{"StarArg", "Expression", false},
			{"KeywordArg", "Expression", false},
		}},
	}, nil},
}
