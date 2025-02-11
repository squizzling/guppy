//go:build rebuild

package flow

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"guppy/internal/parser/ast"
	"guppy/internal/parser/parser"
	"guppy/internal/parser/tokenizer"
)

func rebuildExpressionForFile(
	fullFileName string,
	parse func(p *parser.Parser) (ast.Expression, *parser.ParseError),
) {
	f := string(must1(os.ReadFile(fullFileName)))

	var b bytes.Buffer

	tests := strings.Split(f, "=====\n")
	for idx, test := range tests {
		parts := strings.Split(test, "-----\n")
		input, output := strings.TrimRight(parts[0], "\n"), strings.TrimRight(parts[1], "\n")
		if idx > 0 {
			_, _ = fmt.Fprintf(&b, "=====\n")
		}

		expr, err := parse(parser.NewParser(tokenizer.NewTokenizer(input)))
		actualOutput := ""
		if err != nil {
			actualOutput = err.Error()
		} else {
			actualOutput = strings.TrimRight(must1(expr.Accept(ast.DebugWriter{})).(string), "\n")
		}

		if actualOutput != output {
			fmt.Printf("updated: %s\n", input)
		} else {
			fmt.Printf("unchanged: %s\n", input)
		}

		_, _ = fmt.Fprintf(&b, "%s\n", input)
		_, _ = fmt.Fprintf(&b, "-----\n")
		_, _ = fmt.Fprintf(&b, "%s\n", actualOutput)
	}
	must(os.WriteFile(fullFileName, b.Bytes(), 0o644))
}

func rebuildStatementForFile(
	fullFileName string,
	parse func(p *parser.Parser) (ast.Statement, *parser.ParseError),
) {
	f := string(must1(os.ReadFile(fullFileName)))

	var b bytes.Buffer

	tests := strings.Split(f, "=====\n")
	for idx, test := range tests {
		parts := strings.Split(test, "-----\n")
		input, output := strings.TrimRight(parts[0], "\n"), strings.TrimRight(parts[1], "\n")
		if idx > 0 {
			_, _ = fmt.Fprintf(&b, "=====\n")
		}

		expr, err := parse(parser.NewParser(tokenizer.NewTokenizer(input)))
		actualOutput := ""
		if err != nil {
			actualOutput = err.Error()
		} else {
			actualOutput = strings.TrimRight(must1(expr.Accept(ast.DebugWriter{})).(string), "\n")
		}

		if actualOutput != output {
			fmt.Printf("updated: %s\n", input)
		} else {
			fmt.Printf("unchanged: %s\n", input)
		}

		_, _ = fmt.Fprintf(&b, "%s\n", input)
		_, _ = fmt.Fprintf(&b, "-----\n")
		_, _ = fmt.Fprintf(&b, "%s\n", actualOutput)
	}
	must(os.WriteFile(fullFileName, b.Bytes(), 0o644))
}

func TestRebuild(t *testing.T) {
	rebuildExpressionForFile("testdata/expressions/parseTest.txt", parseTest)
	rebuildExpressionForFile("testdata/expressions/parseTestListComp.txt", parseTestListComp)
	rebuildExpressionForFile("testdata/expressions/parseTupleExpr.txt", parseTupleExpr)
	rebuildStatementForFile("testdata/statements/parseExpressionStatement.txt", parseExpressionStatement)
}
