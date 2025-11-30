//go:build rebuild

package flow

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/squizzling/guppy/pkg/parser/parser"
	"github.com/squizzling/guppy/pkg/parser/tokenizer"
)

func rebuildForFile[T any](
	fullFileName string,
	parse func(p *parser.Parser) (T, *parser.ParseError),
	render func(t T) string,
) {
	// TODO: Pull out common code with testFromFile, so we can have a consistent view between both worlds
	f := string(must1(os.ReadFile(fullFileName)))

	var b bytes.Buffer

	tests := strings.Split(f, "=====\n")
	for idx, test := range tests {
		parts := strings.Split(test, "\n-----\n")
		input, output := parts[0], parts[1]
		if idx > 0 {
			_, _ = fmt.Fprintf(&b, "=====\n")
		}

		expr, err := parse(parser.NewParser(tokenizer.NewTokenizer(input)))
		actualOutput := ""
		if err != nil {
			actualOutput = err.Error()
		} else {
			actualOutput = strings.TrimRight(render(expr), "\n")
		}

		if strings.Trim(actualOutput, "\n") != strings.Trim(output, "\n") {
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
	for _, tst := range dataArgumentListTests {
		rebuildForFile("testdata/dataargumentlist/"+tst.fileName+".txt", tst.parse, renderDataArgumentList)
	}
	for _, tst := range dataListForTests {
		rebuildForFile("testdata/datalistfor/"+tst.fileName+".txt", tst.parse, renderDataListFor)
	}
	for _, tst := range dataListIfTests {
		rebuildForFile("testdata/datalistif/"+tst.fileName+".txt", tst.parse, renderDataListIf)
	}
	for _, tst := range dataListIterTests {
		rebuildForFile("testdata/datalistiter/"+tst.fileName+".txt", tst.parse, renderDataListIter)
	}
	for _, tst := range dataParameterTests {
		rebuildForFile("testdata/dataparameter/"+tst.fileName+".txt", tst.parse, renderDataParameter)
	}
	for _, tst := range dataParameterListTests {
		rebuildForFile("testdata/dataparameterlist/"+tst.fileName+".txt", tst.parse, renderDataParameterList)
	}
	for _, tst := range dataSubscriptTests {
		rebuildForFile("testdata/datasubscript/"+tst.fileName+".txt", tst.parse, renderDataSubscript)
	}
	for _, tst := range dataExpressionTests {
		rebuildForFile("testdata/expressions/"+tst.fileName+".txt", tst.parse, renderExpression)
	}
	for _, tst := range dataStatementTests {
		rebuildForFile("testdata/statements/"+tst.fileName+".txt", tst.parse, renderStatement)
	}
}
