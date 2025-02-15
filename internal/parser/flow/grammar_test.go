package flow

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"guppy/internal/parser/ast"
	"guppy/internal/parser/parser"
	"guppy/internal/parser/tokenizer"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func must1[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func testFromFile[T any](
	t *testing.T,
	fullFilename string,
	parse func(p *parser.Parser) (T, *parser.ParseError),
	render func(t T) string,
) {
	// TODO: Pull out common code with rebuildForFile, so we can have a consistent view between both worlds

	filename := path.Base(fullFilename)
	f := string(must1(os.ReadFile(fullFilename)))
	tests := strings.Split(f, "=====\n")
	for idx, test := range tests {
		parts := strings.Split(test, "\n-----\n")
		if len(parts) != 2 {
			t.Fatalf("malformed test in %s/%d", filename, idx)
		}
		input, output := parts[0], parts[1]
		t.Run(filename+"/"+input, func(t *testing.T) {
			t.Parallel()

			if actual, err := parse(parser.NewParser(tokenizer.NewTokenizer(input))); err != nil {
				if !assert.Equal(t, strings.TrimRight(output, "\n"), err.Error()) {
					ss := err.Stack()
					for _, s := range ss {
						_, _ = fmt.Fprintf(os.Stderr, "%s %s\n", s.Location, s.Message)
					}
				}
			} else {
				actualTree := render(actual)
				assert.Equal(t, output, actualTree)
			}
		})
	}
}

type grammarTest[T any] struct {
	fileName string
	parse    func(p *parser.Parser) (T, *parser.ParseError)
}

var dataParameterTests = []grammarTest[*ast.DataParameter]{
	{"parseParamType", parseParamType},
	{"parseVarArgsKwsParam", parseVarArgsKwsParam},
	{"parseVarArgsListParamDef", parseVarArgsListParamDef},
	{"parseVarArgsListParamName", parseVarArgsListParamName},
	{"parseVarArgsStarParam", parseVarArgsStarParam},
}

var dataParameterListTests = []grammarTest[*ast.DataParameterList]{
	{"parseVarArgsList", parseVarArgsList},
	{"parseParameters", parseParameters},
}

var dataExpressionTests = []grammarTest[ast.Expression]{
	{"parseDictExpr", parseDictExpr},
	{"parseTest", parseTest},
	{"parseTestListComp", parseTestListComp},
	{"parseTupleExpr", parseTupleExpr},
}

var dataStatementTests = []grammarTest[ast.Statement]{
	{"parseExpressionStatement", parseExpressionStatement},
	{"parseFunctionDefinition", parseFunctionDefinition},
	{"parseSuite", parseSuite},
}

func renderDataParameter(d *ast.DataParameter) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderDataParameterList(d *ast.DataParameterList) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderExpression(d ast.Expression) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderStatement(d ast.Statement) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func TestDataParameter(t *testing.T) {
	t.Parallel()

	for _, tst := range dataParameterTests {
		testFromFile(t, "testdata/dataparameter/"+tst.fileName+".txt", tst.parse, renderDataParameter)
	}
}

func TestDataParameterList(t *testing.T) {
	t.Parallel()

	for _, tst := range dataParameterListTests {
		testFromFile(t, "testdata/dataparameterlist/"+tst.fileName+".txt", tst.parse, renderDataParameterList)
	}
}

func TestExpression(t *testing.T) {
	t.Parallel()

	for _, tst := range dataExpressionTests {
		testFromFile(t, "testdata/expressions/"+tst.fileName+".txt", tst.parse, renderExpression)
	}
}

func TestStatement(t *testing.T) {
	t.Parallel()

	for _, tst := range dataStatementTests {
		testFromFile(t, "testdata/statements/"+tst.fileName+".txt", tst.parse, renderStatement)
	}
}

func TestParseIdList(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		input           string
		expectedIds     []string
		expectedError   string
		remainingTokens int
	}{
		{input: "a", expectedIds: []string{"a"}, remainingTokens: 2},
		{input: "a,", expectedIds: []string{"a"}, remainingTokens: 2},
		{input: "a,,", expectedIds: []string{"a"}, remainingTokens: 3},
		{input: "a,b", expectedIds: []string{"a", "b"}, remainingTokens: 2},
		{input: "a,b,", expectedIds: []string{"a", "b"}, remainingTokens: 2},
		{input: "a,b,,", expectedIds: []string{"a", "b"}, remainingTokens: 3},
		{input: "a,b,c", expectedIds: []string{"a", "b", "c"}, remainingTokens: 2},
		{input: "a,b,c,", expectedIds: []string{"a", "b", "c"}, remainingTokens: 2},
		{input: "a,b,c,,", expectedIds: []string{"a", "b", "c"}, remainingTokens: 3},
		{input: ",a", expectedError: "expecting ID, found COMMA", remainingTokens: 3},
	} {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			p := parser.NewParser(tokenizer.NewTokenizer(tc.input))
			ids, err := parseIdList(p)
			if tc.expectedError != "" {
				require.NotNil(t, err)
				assert.Equal(t, tc.expectedError, err.Message)
			} else {
				require.Nil(t, err)
			}
			assert.Equal(t, tc.remainingTokens, p.RemainingTokens())
			assert.Equal(t, tc.expectedIds, ids)
		})
	}
}
