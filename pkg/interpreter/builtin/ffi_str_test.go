package builtin

import (
	"fmt"
	"testing"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type failingStr struct{}

func (f failingStr) Member(i itypes.Interpreter, obj itypes.Object, memberName string) (itypes.Object, error) {
	return nil, nil
}

func (f failingStr) String(i itypes.Interpreter) (string, error) {
	return "", fmt.Errorf("failingStr error")
}

func (f failingStr) Repr() string {
	return "failingStr()"
}

func TestFFINewStr(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, NewFFIStr())
}

func TestFFIStr(t *testing.T) {
	t.Parallel()

	i := interpreter.NewInterpreter(false)

	s := FFIStr{primitive.NewObjectBool(true)}
	o, err := s.Call(i)
	require.NoError(t, err)
	assert.Equal(t, "True", o.(*primitive.ObjectString).Value)

	s = FFIStr{failingStr{}}
	_, err = s.Call(i)
	assert.ErrorContains(t, err, "failingStr error")
}
