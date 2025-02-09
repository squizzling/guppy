package flow

import (
	"os"
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

func testExpressionFromFile(t *testing.T, fullFilename string, parse func(p *parser.Parser) (ast.Expression, *parser.ParseError)) {
	f := string(must1(os.ReadFile(fullFilename)))
	tests := strings.Split(f, "=====\n")
	for _, test := range tests {
		parts := strings.Split(test, "-----\n")
		input, output := parts[0], parts[1]
		t.Run(input, func(t *testing.T) {
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

func TestExpressions(t *testing.T) {
	t.Parallel()

	testExpressionFromFile(t, "testdata/expressions/parseTestListComp.txt", parseTestListComp)
	testExpressionFromFile(t, "testdata/expressions/parseTupleExpr.txt", parseTupleExpr)
}

func TestParseIdList(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		input           string
		expectedIds     []string
		expectedError   string
		remainingTokens int
	}{
		{input: "a", expectedIds: []string{"a"}, remainingTokens: 1},
		{input: "a,", expectedIds: []string{"a"}, remainingTokens: 1},
		{input: "a,,", expectedIds: []string{"a"}, remainingTokens: 2},
		{input: "a,b", expectedIds: []string{"a", "b"}, remainingTokens: 1},
		{input: "a,b,", expectedIds: []string{"a", "b"}, remainingTokens: 1},
		{input: "a,b,,", expectedIds: []string{"a", "b"}, remainingTokens: 2},
		{input: "a,b,c", expectedIds: []string{"a", "b", "c"}, remainingTokens: 1},
		{input: "a,b,c,", expectedIds: []string{"a", "b", "c"}, remainingTokens: 1},
		{input: "a,b,c,,", expectedIds: []string{"a", "b", "c"}, remainingTokens: 2},
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
