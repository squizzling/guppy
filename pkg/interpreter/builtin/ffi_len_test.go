package builtin

import (
	"testing"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/primitive"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFFINewLen(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, NewFFILen())
}

func TestFFILen(t *testing.T) {
	t.Parallel()

	i := interpreter.NewInterpreter(false)

	l := FFILen{}
	l.Value.List = primitive.NewObjectList(primitive.NewObjectInt(0))
	o, err := l.Call(i)
	require.NoError(t, err)
	assert.Equal(t, 1, o.(*primitive.ObjectInt).Value)

	l = FFILen{}
	l.Value.Tuple = primitive.NewObjectTuple(primitive.NewObjectInt(0))
	o, err = l.Call(i)
	require.NoError(t, err)
	assert.Equal(t, 1, o.(*primitive.ObjectInt).Value)

	l = FFILen{}
	l.Value.String = primitive.NewObjectString("test")
	o, err = l.Call(i)
	require.NoError(t, err)
	assert.Equal(t, 4, o.(*primitive.ObjectInt).Value)

	l = FFILen{}
	o, err = l.Call(i)
	assert.ErrorContains(t, err, "FFILen.Value is not set")
}
