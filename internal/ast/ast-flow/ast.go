package astflow

import (
	"guppy/internal/ast"
)

const Package = "ast"

var Interfaces = ast.Interfaces{
	{"Data", []ast.Node{
		{"Argument", true, []ast.Field{
			{"Assign", "string"},
			{"Expr", "Expression"},
		}},
		{"ArgumentList", true, []ast.Field{
			{"Args", "[]Expression"},
			{"NamedArgs", "[]*DataArgument"},
			{"StarArg", "Expression"},
			{"KeywordArg", "Expression"},
		}},
		{"ImportAs", true, []ast.Field{
			{"Name", "[]string"},
			{"As", "string"},
		}},
		{"ListIter", true, []ast.Field{
			{"For", "*DataListFor"},
			{"If", "*DataListIf"},
		}},
		{"ListFor", true, []ast.Field{
			{"Idents", "[]string"},
			{"Expr", "Expression"},
			{"Iter", "*DataListIter"},
		}},
		{"ListIf", true, []ast.Field{
			{"Expr", "Expression"},
			{"Iter", "*DataListIter"},
		}},
		{"Parameter", true, []ast.Field{
			{"Name", "string"},
			{"Type", "string"},
			{"Default", "Expression"},
			{"StarArg", "bool"},
			{"KeywordArg", "bool"},
		}},
		{"ParameterList", true, []ast.Field{
			{"List", "[]*DataParameter"},
		}},
		{"Subscript", true, []ast.Field{
			{"Start", "Expression"},
			{"End", "Expression"},
			{"Range", "bool"},
		}},
	}},
	{"Statement", []ast.Node{
		{"Program", true, []ast.Field{
			{"Statements", "*StatementList"},
		}},
		{"Expression", false, []ast.Field{
			{"Assign", "[]string"},
			{"Expr", "Expression"},
		}},
		{"Return", false, []ast.Field{
			{"Expr", "Expression"},
		}},
		{"ImportFrom", false, []ast.Field{
			{"From", "[]string"},
			{"Imports", "[]*DataImportAs"},
		}},
		{"ImportFromStar", false, []ast.Field{
			{"From", "[]string"},
		}},
		{"ImportNames", false, []ast.Field{
			{"Imports", "[]*DataImportAs"},
		}},
		{"Assert", false, []ast.Field{
			{"Test", "Expression"},
			{"Raise", "Expression"},
		}},
		{"If", false, []ast.Field{
			{"Condition", "[]Expression"},
			{"Statement", "[]Statement"},
			{"StatementElse", "Statement"},
		}},

		{"Function", false, []ast.Field{
			{"Token", "string"},
			{"Params", "*DataParameterList"},
			{"Body", "Statement"},
		}},
		{"Decorated", false, []ast.Field{
			// TODO
		}},

		{"List", true, []ast.Field{
			{"Statements", "[]Statement"},
		}},
	}},
	{"Expression", []ast.Node{
		{"List", true, []ast.Field{ // Make this a concrete return type
			{"Expressions", "[]Expression"},
			{"Tuple", "bool"},
		}},
		{"ListMaker", false, []ast.Field{
			{"Expr", "Expression"},
			{"For", "*DataListFor"},
		}},
		{"Binary", false, []ast.Field{
			{"Left", "Expression"},
			{"Op", "tokenizer.Token"},
			{"Right", "Expression"},
		}},
		{"Dict", false, []ast.Field{
			{"Keys", "[]Expression"},
			{"Values", "[]Expression"},
		}},
		{"Grouping", false, []ast.Field{
			{"Expr", "Expression"},
		}},
		{"Lambda", false, []ast.Field{
			{"Identifier", "string"},
			{"Expr", "Expression"},
		}},
		{"Literal", false, []ast.Field{
			{"Value", "any"},
		}},
		{"Logical", false, []ast.Field{
			{"Left", "Expression"},
			{"Op", "tokenizer.Token"},
			{"Right", "Expression"},
		}},
		{"Ternary", false, []ast.Field{
			{"Left", "Expression"},
			{"Cond", "Expression"},
			{"Right", "Expression"},
		}},
		{"Unary", false, []ast.Field{
			{"Op", "tokenizer.TokenType"},
			{"Right", "Expression"},
		}},
		{"Variable", false, []ast.Field{
			{"Identifier", "string"},
		}},
		{"Member", false, []ast.Field{
			{"Expr", "Expression"},
			{"Identifier", "string"},
		}},
		{"Subscript", false, []ast.Field{
			{"Expr", "Expression"},
			{"Start", "Expression"},
			{"End", "Expression"},
			{"Range", "bool"},
		}},
		{"Call", false, []ast.Field{
			{"Expr", "Expression"},
			{"Args", "[]Expression"},
			{"NamedArgs", "[]*DataArgument"},
			{"StarArg", "Expression"},
			{"KeywordArg", "Expression"},
		}},
	}},
}
