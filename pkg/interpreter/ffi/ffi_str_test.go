package ffi

import (
	"fmt"
	"testing"

	"guppy/pkg/interpreter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type failingStr struct{}

func (f failingStr) Member(i *interpreter.Interpreter, obj interpreter.Object, memberName string) (interpreter.Object, error) {
	return nil, nil
}

func (f failingStr) String(i *interpreter.Interpreter) (string, error) {
	return "", fmt.Errorf("failingStr error")
}

func TestFFINewStr(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, NewFFIStr())
}

func TestFFIStr(t *testing.T) {
	t.Parallel()

	i := interpreter.NewInterpreter(false)

	s := FFIStr{interpreter.NewObjectBool(true)}
	o, err := s.Call(i)
	require.NoError(t, err)
	assert.Equal(t, "true", o.(*interpreter.ObjectString).Value)

	s = FFIStr{failingStr{}}
	_, err = s.Call(i)
	assert.ErrorContains(t, err, "failingStr error")
}
