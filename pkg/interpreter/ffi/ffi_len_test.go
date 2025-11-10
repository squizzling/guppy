package ffi

import (
	"testing"

	"guppy/pkg/interpreter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFFINewLen(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, NewFFILen())
}

func TestFFILen(t *testing.T) {
	t.Parallel()

	l := FFILen{}
	l.Value.List = interpreter.NewObjectList(interpreter.NewObjectInt(0))
	o, err := l.Call()
	require.NoError(t, err)
	assert.Equal(t, 1, o.(*interpreter.ObjectInt).Value)

	l = FFILen{}
	l.Value.Tuple = interpreter.NewObjectTuple(interpreter.NewObjectInt(0))
	o, err = l.Call()
	require.NoError(t, err)
	assert.Equal(t, 1, o.(*interpreter.ObjectInt).Value)

	l = FFILen{}
	l.Value.String = interpreter.NewObjectString("test")
	o, err = l.Call()
	require.NoError(t, err)
	assert.Equal(t, 4, o.(*interpreter.ObjectInt).Value)

	l = FFILen{}
	o, err = l.Call()
	assert.ErrorContains(t, err, "FFILen.Value is not set")
}
