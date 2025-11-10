package interpreter

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"testing"

	"guppy/pkg/parser/flow"
	"guppy/pkg/parser/parser"
	"guppy/pkg/parser/tokenizer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestFFI struct {
	Single *ObjectInt              `ffi:"single"`
	OneOf  ThingOrNone[*ObjectInt] `ffi:"oneof"`
}

func (t TestFFI) Call(i *Interpreter) (Object, error) {
	if t.OneOf.Thing != nil {
		return NewObjectInt(t.Single.Value + t.OneOf.Thing.Value), nil
	} else if t.OneOf.None != nil {
		return NewObjectInt(t.Single.Value), nil
	} else {
		panic("neither OneOf nor Thing are set")
	}
}

type TestFFISingleKW struct {
	Single *ObjectInt              `ffi:"single,kw"`
	OneOf  ThingOrNone[*ObjectInt] `ffi:"oneof"`
}

func (t TestFFISingleKW) Call(i *Interpreter) (Object, error) {
	if t.OneOf.Thing != nil {
		return NewObjectInt(t.Single.Value + t.OneOf.Thing.Value), nil
	} else if t.OneOf.None != nil {
		return NewObjectInt(t.Single.Value), nil
	} else {
		panic("neither OneOf nor Thing are set")
	}
}

type TestFFIOneOfKW struct {
	Single *ObjectInt              `ffi:"single"`
	OneOf  ThingOrNone[*ObjectInt] `ffi:"oneof,kw"`
}

func (t TestFFIOneOfKW) Call(i *Interpreter) (Object, error) {
	if t.OneOf.Thing != nil {
		return NewObjectInt(t.Single.Value + t.OneOf.Thing.Value), nil
	} else if t.OneOf.None != nil {
		return NewObjectInt(t.Single.Value), nil
	} else {
		panic("neither OneOf nor Thing are set")
	}
}

type TestFFISingleKWOneOfKW struct {
	Single *ObjectInt              `ffi:"single,kw"`
	OneOf  ThingOrNone[*ObjectInt] `ffi:"oneof,kw"`
}

func (t TestFFISingleKWOneOfKW) Call(i *Interpreter) (Object, error) {
	if t.OneOf.Thing != nil {
		return NewObjectInt(t.Single.Value + t.OneOf.Thing.Value), nil
	} else if t.OneOf.None != nil {
		return NewObjectInt(t.Single.Value), nil
	} else {
		panic("neither OneOf nor Thing are set")
	}
}

func TestNewtFFIDefaults(t *testing.T) {
	for _, fn := range []string{
		"ffi", "ffioneofkw", "ffisinglekw", "ffisinglekwoneofkw",
	} {
		t.Run(fn, func(t *testing.T) {
			for _, ts := range []struct {
				name          string
				defaultSingle *ObjectInt
				defaultOneOf  ThingOrNone[*ObjectInt]
			}{
				{
					"single-default-oneof-default",
					NewObjectInt(1),
					ThingOrNone[*ObjectInt]{nil, NewObjectInt(2)},
				}, {
					"single-default-oneof-missing",
					NewObjectInt(1),
					ThingOrNone[*ObjectInt]{nil, nil},
				}, {
					"single-default-oneof-none",
					NewObjectInt(1),
					ThingOrNone[*ObjectInt]{NewObjectNone(), nil},
				}, {
					"single-missing-oneof-default",
					nil,
					ThingOrNone[*ObjectInt]{nil, NewObjectInt(2)},
				}, {
					"single-missing-oneof-missing",
					nil,
					ThingOrNone[*ObjectInt]{nil, nil},
				}, {
					"single-missing-oneof-none",
					nil,
					ThingOrNone[*ObjectInt]{NewObjectNone(), nil},
				},
			} {
				t.Run(ts.name, func(t *testing.T) {
					ffi := NewFFI(TestFFI{Single: ts.defaultSingle, OneOf: ts.defaultOneOf})
					ffiOneOfKW := NewFFI(TestFFIOneOfKW{Single: ts.defaultSingle, OneOf: ts.defaultOneOf})
					ffiSingleKW := NewFFI(TestFFISingleKW{Single: ts.defaultSingle, OneOf: ts.defaultOneOf})
					ffiSingleKWOneOfKW := NewFFI(TestFFISingleKWOneOfKW{Single: ts.defaultSingle, OneOf: ts.defaultOneOf})
					testFromFile(
						t,
						fmt.Sprintf("testdata/%s/%s.txt", fn, ts.name),
						map[string]FlowCall{
							"ffi":                ffi,
							"ffioneofkw":         ffiOneOfKW,
							"ffisinglekw":        ffiSingleKW,
							"ffisinglekwoneofkw": ffiSingleKWOneOfKW,
						},
					)
				})
			}
		})
	}
}

func must1[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func testFromFile(
	t *testing.T,
	fullFilename string,
	calls map[string]FlowCall,
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
			i := NewInterpreter(false)
			for k, v := range calls {
				require.NoError(t, i.Scope.Set(k, v))
			}
			// p.Accept() will push/pop a scope.  p.Statements.Accept() will use the current scope,
			// which we need to retrieve `o`
			_, err := p.Statements.Accept(i)
			if strings.HasPrefix(output, "*") {
				require.Error(t, err)
				assert.EqualValues(t, output[1:], err.Error())
			} else {
				require.NoError(t, err)
				oVal, err := i.Scope.Get("o")
				require.NoError(t, err)
				assert.Equal(t, output, strconv.Itoa(oVal.(*ObjectInt).Value))
			}
		})
	}
}

func TestFFIError(t *testing.T) {
	for _, ts := range []struct {
		name          string
		single        Object
		oneOf         Object
		expectedError string
	}{
		{"single-wrong-type", NewObjectNone(), NewObjectNone(), "param `single` for TestFFI.Single is *interpreter.ObjectNone not *interpreter.ObjectInt"},
		{"single-missing", nil, NewObjectNone(), "param `single` for TestFFI.Single is missing, expecting *interpreter.ObjectInt"},
		{"oneOf-wrong-type", NewObjectInt(1), NewObjectString(""), "param `oneof` for TestFFI.OneOf is *interpreter.ObjectString not *interpreter.ObjectNone, or *interpreter.ObjectInt"},
		{"oneOf-missing", NewObjectInt(1), nil, "param `oneof` for TestFFI.OneOf is missing, expecting *interpreter.ObjectNone, or *interpreter.ObjectInt"},
	} {
		t.Run(ts.name, func(t *testing.T) {
			i := NewInterpreter(false)
			f := NewFFI(TestFFI{})

			if ts.single != nil {
				require.NoError(t, i.Scope.Set("single", ts.single))
			}
			if ts.oneOf != nil {
				require.NoError(t, i.Scope.Set("oneof", ts.oneOf))
			}
			_, err := f.Call(i)
			require.Error(t, err)
			assert.EqualValues(t, ts.expectedError, err.Error())
		})
	}
}
