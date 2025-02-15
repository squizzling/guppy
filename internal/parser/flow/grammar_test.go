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

func testDataParameterFromFile(t *testing.T, fullFilename string, parse func(p *parser.Parser) (*ast.DataParameter, *parser.ParseError)) {
	filename := path.Base(fullFilename)
	f := string(must1(os.ReadFile(fullFilename)))
	tests := strings.Split(f, "=====\n")
	for idx, test := range tests {
		parts := strings.Split(test, "-----\n")
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
				actualTree := must1(actual.Accept(ast.DebugWriter{}))
				assert.Equal(t, output, actualTree)
			}
		})
	}
}

func testDataParameterListFromFile(t *testing.T, fullFilename string, parse func(p *parser.Parser) (*ast.DataParameterList, *parser.ParseError)) {
	filename := path.Base(fullFilename)
	f := string(must1(os.ReadFile(fullFilename)))
	tests := strings.Split(f, "=====\n")
	for idx, test := range tests {
		parts := strings.Split(test, "-----\n")
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
				actualTree := must1(actual.Accept(ast.DebugWriter{}))
				assert.Equal(t, output, actualTree)
			}
		})
	}
}

func testExpressionFromFile(t *testing.T, fullFilename string, parse func(p *parser.Parser) (ast.Expression, *parser.ParseError)) {
	filename := path.Base(fullFilename)
	f := string(must1(os.ReadFile(fullFilename)))
	tests := strings.Split(f, "=====\n")
	for idx, test := range tests {
		parts := strings.Split(test, "-----\n")
		if len(parts) != 2 {
			t.Fatalf("malformed test in %s/%d", filename, idx)
		}
		input, output := parts[0], parts[1]
		t.Run(filename+"/"+input, func(t *testing.T) {
			t.Parallel()

			if actual, err := parse(parser.NewParser(tokenizer.NewTokenizer(input))); err != nil {
				assert.Equal(t, strings.TrimRight(output, "\n"), err.Error())
			} else {
				actualTree := must1(actual.Accept(ast.DebugWriter{}))
				assert.Equal(t, output, actualTree)
			}
		})
	}
}

func testStatementFromFile(t *testing.T, fullFilename string, parse func(p *parser.Parser) (ast.Statement, *parser.ParseError)) {
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
				actualTree := must1(actual.Accept(ast.DebugWriter{}))
				assert.Equal(t, output, actualTree)
			}
		})
	}
}

func TestExpressions(t *testing.T) {
	t.Parallel()

	testDataParameterFromFile(t, "testdata/dataparameter/parseParamType.txt", parseParamType)
	testDataParameterFromFile(t, "testdata/dataparameter/parseVarArgsKwsParam.txt", parseVarArgsKwsParam)
	testDataParameterFromFile(t, "testdata/dataparameter/parseVarArgsListParamDef.txt", parseVarArgsListParamDef)
	testDataParameterFromFile(t, "testdata/dataparameter/parseVarArgsListParamName.txt", parseVarArgsListParamName)
	testDataParameterFromFile(t, "testdata/dataparameter/parseVarArgsStarParam.txt", parseVarArgsStarParam)
	testDataParameterListFromFile(t, "testdata/dataparameterlist/parseVarArgsList.txt", parseVarArgsList)
	testDataParameterListFromFile(t, "testdata/dataparameterlist/parseParameters.txt", parseParameters)
	testExpressionFromFile(t, "testdata/expressions/parseDictExpr.txt", parseDictExpr)
	testExpressionFromFile(t, "testdata/expressions/parseTest.txt", parseTest)
	testExpressionFromFile(t, "testdata/expressions/parseTestListComp.txt", parseTestListComp)
	testExpressionFromFile(t, "testdata/expressions/parseTupleExpr.txt", parseTupleExpr)
	testStatementFromFile(t, "testdata/statements/parseExpressionStatement.txt", parseExpressionStatement)
	testStatementFromFile(t, "testdata/statements/parseFunctionDefinition.txt", parseFunctionDefinition)
	testStatementFromFile(t, "testdata/statements/parseSuite.txt", parseSuite)
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
