package flow

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/squizzling/guppy/pkg/parser/ast"
	"github.com/squizzling/guppy/pkg/parser/parser"
	"github.com/squizzling/guppy/pkg/parser/tokenizer"
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

var dataArgumentListTests = []grammarTest[*ast.DataArgumentList]{
	{"parseActualArgs", parseActualArgs},
}

var dataListForTests = []grammarTest[*ast.DataListFor]{
	{"parseListFor", parseListFor},
}

var dataListIfTests = []grammarTest[*ast.DataListIf]{
	{"parseListIf", parseListIf},
}

var dataListIterTests = []grammarTest[*ast.DataListIter]{
	{"parseListIter", parseListIter},
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

var dataSubscriptTests = []grammarTest[*ast.DataSubscript]{
	{"parseSubscript", parseSubscript},
}

var dataExpressionTests = []grammarTest[ast.Expression]{
	{"parseDictExpr", parseDictExpr},
	{"parseLambdef", parseLambdef},
	{"parseLambdefNoCond", parseLambdefNoCond},
	{"parseTest", parseTest},
	{"parseTestListComp", parseTestListComp},
	{"parseTestListNoCond", parseTestListNoCond},
	{"parseTestNoCond", parseTestNoCond},
	{"parseTupleExpr", parseTupleExpr},
}

var dataStatementTests = []grammarTest[ast.Statement]{
	{"parseExpressionStatement", parseExpressionStatement},
	{"parseFunctionDefinition", parseFunctionDefinition},
	{"parseIfStatement", parseIfStatement},
	{"parseImportStatement", parseImportStatement},
	{"parseImportName", parseImportName},
	{"parseImportFrom", parseImportFrom},
	{"parseReturnStatement", parseReturnStatement},
	{"parseSuite", parseSuite},
}

func renderDataArgumentList(d *ast.DataArgumentList) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderDataListFor(d *ast.DataListFor) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderDataListIf(d *ast.DataListIf) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderDataListIter(d *ast.DataListIter) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderDataParameter(d *ast.DataParameter) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderDataParameterList(d *ast.DataParameterList) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderDataSubscript(d *ast.DataSubscript) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderExpression(d ast.Expression) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func renderStatement(d ast.Statement) string {
	return must1(d.Accept(ast.DebugWriter{})).(string)
}

func TestDataArgumentList(t *testing.T) {
	t.Parallel()

	for _, tst := range dataArgumentListTests {
		testFromFile(t, "testdata/dataargumentlist/"+tst.fileName+".txt", tst.parse, renderDataArgumentList)
	}
}

func TestDataListFor(t *testing.T) {
	t.Parallel()

	for _, tst := range dataListForTests {
		testFromFile(t, "testdata/datalistfor/"+tst.fileName+".txt", tst.parse, renderDataListFor)
	}
}

func TestDataListIf(t *testing.T) {
	t.Parallel()

	for _, tst := range dataListIfTests {
		testFromFile(t, "testdata/datalistif/"+tst.fileName+".txt", tst.parse, renderDataListIf)
	}
}

func TestDataListIter(t *testing.T) {
	t.Parallel()

	for _, tst := range dataListIterTests {
		testFromFile(t, "testdata/datalistiter/"+tst.fileName+".txt", tst.parse, renderDataListIter)
	}
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

func TestDataSubscript(t *testing.T) {
	t.Parallel()

	for _, tst := range dataSubscriptTests {
		testFromFile(t, "testdata/datasubscript/"+tst.fileName+".txt", tst.parse, renderDataSubscript)
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
