package builtin

import (
	"testing"

	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type reprThing struct {
	itypes.Object
}

func (f reprThing) Repr() string {
	return "repr"
}

func TestFFINewRepr(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, NewFFIRepr())
}

func TestFFIRepr(t *testing.T) {
	t.Parallel()

	i := interpreter.NewInterpreter(false)

	s := FFIRepr{reprThing{}}
	o, err := s.Call(i)
	require.NoError(t, err)
	assert.Equal(t, "repr", o.(*primitive.ObjectString).Value)
}
