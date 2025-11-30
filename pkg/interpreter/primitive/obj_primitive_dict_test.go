package primitive_test

import (
	"testing"

	"github.com/squizzling/guppy/pkg/interpreter/primitive"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDict(t *testing.T) {
	testFromFile(t, "testdata/dict.txt")
}

func TestObjectDictAsMapStringString(t *testing.T) {
	d := primitive.NewObjectDict([]primitive.DictItem{
		{
			Key:   primitive.NewObjectString("foo"),
			Value: primitive.NewObjectString("bar"),
		},
	})

	m, err := d.AsMapStringString()
	require.NoError(t, err)
	assert.EqualValues(t, map[string]string{"foo": "bar"}, m)
}

func TestObjectDictAsMapStringStringInvalidKey(t *testing.T) {
	d := primitive.NewObjectDict([]primitive.DictItem{
		{
			Key:   primitive.NewObjectInt(5),
			Value: primitive.NewObjectString("bar"),
		},
	})

	m, err := d.AsMapStringString()
	require.Error(t, err)
	assert.Equal(t, "dict idx 0 (int(5)) is *primitive.ObjectInt not *interpreter.ObjectString", err.Error())
	assert.Nil(t, m)
}

func TestObjectDictAsMapStringStringInvalidValue(t *testing.T) {
	d := primitive.NewObjectDict([]primitive.DictItem{
		{
			Key:   primitive.NewObjectString("foo"),
			Value: primitive.NewObjectInt(5),
		},
	})

	m, err := d.AsMapStringString()
	require.Error(t, err)
	assert.Equal(t, "dict idx 0 (foo) is *primitive.ObjectInt not *interpreter.ObjectString", err.Error())
	assert.Nil(t, m)
}
