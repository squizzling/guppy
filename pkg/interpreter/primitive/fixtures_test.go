package primitive_test

import (
	"os"
	"path"
	"slices"
	"strings"
	"testing"

	"guppy/pkg/flow/stream"
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/builtin"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/parser/flow"
	"guppy/pkg/parser/parser"
	"guppy/pkg/parser/tokenizer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func must1[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func testFromFile(
	t *testing.T,
	fullFilename string,
) {
	filename := path.Base(fullFilename)
	fileContents := string(must1(os.ReadFile(fullFilename)))

	fileContents = strings.Join(slices.DeleteFunc(slices.Collect(strings.Lines(fileContents)),
		func(s string) bool {
			return strings.HasPrefix(s, "#")
		},
	), "")
	tests := strings.Split(fileContents, "=====\n")
	for idx, test := range tests {
		parts := strings.Split(test, "\n-----\n")
		if len(parts) != 2 {
			t.Fatalf("malformed test in %s/%d", filename, idx)
		}
		input, output := parts[0], parts[1]
		output = strings.TrimRight(output, "\n")
		t.Run(filename+"/"+input, func(t *testing.T) {
			t.Parallel()

			p, parseErr := flow.ParseProgram(parser.NewParser(tokenizer.NewTokenizer(input)))
			if parseErr != nil {
				panic("parse failed: " + parseErr.Error())
			}
			i := interpreter.NewInterpreter(false)
			_ = i.SetGlobal("repr", builtin.NewFFIRepr())
			_ = i.SetGlobal("str", builtin.NewFFIStr())
			_ = i.SetGlobal("const", &stream.FFIConst{Object: itypes.NewObject(nil)}) // So we can use __ris__

			// p.Accept() will push/pop a scope.  p.Statements.Accept() will use the current scope,
			// which we need to retrieve `o`
			_, err := p.Statements.Accept(i)

			if strings.HasPrefix(output, "*") {
				require.Error(t, err)
				assert.EqualValues(t, output[1:], err.Error())
			} else {
				require.NoError(t, err)
				oVal, err := i.Get("o")
				require.NoError(t, err)
				s, err := i.DoString(oVal)
				// TODO: Make everything repr(), so we can test via that.
				// TODO: Add a repr() builtin function so we can test that.
				if strings.HasPrefix(output, "?") {
					require.Error(t, err)
					assert.EqualValues(t, output[1:], err.Error())
				} else {
					require.NoError(t, err)
					assert.Equal(t, output, s)
				}
			}
		})
	}
}
