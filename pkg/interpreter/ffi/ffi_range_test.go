package ffi

import (
	"testing"

	"guppy/pkg/interpreter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFFIRangeError(t *testing.T) {
	t.Parallel()

	i := interpreter.NewInterpreter(false)
	require.NoError(t, i.Scope.Set("start", interpreter.NewObjectInt(0)))
	require.NoError(t, i.Scope.Set("stop", interpreter.NewObjectInt(0)))
	require.NoError(t, i.Scope.Set("step", interpreter.NewObjectInt(0)))
	_, err := NewFFIRange().Call(i)
	assert.ErrorContains(t, err, "invalid step in range")
}

func TestFFIRangeErrorInvalidDate(t *testing.T) {
	t.Parallel()

	i := interpreter.NewInterpreter(false)

	f := FFIRange{
		Start: interpreter.NewObjectInt(0),
		Stop:  interpreter.ThingOrNone[*interpreter.ObjectInt]{},
		Step:  interpreter.NewObjectInt(0),
	}
	_, err := f.Call(i)
	assert.ErrorContains(t, err, "FFIRange.Stop is not set")
}

func TestFFIRange(t *testing.T) {
	t.Parallel()

	for _, ts := range []struct {
		input    FFIRange
		expected []int
	}{
		/**
		for stop in [-10, 0, 10, None]:
		 for start in [-10, 0, 10]:
		  for step in [-3, -1, 1, 3]:
		   #print(f'start={start} step={step} stop={stop}')
		   if stop is not None:
			print(f'{{input: FFIRange{{Start: NewObjectInt({start}), Stop: {f'NewThingOrNoneThing(NewObjectInt({stop}))' if stop is not None else 'NewThingOrNoneNone[*ObjectInt]()'}, Step: NewObjectInt({step})}}, expected: []int{str(list(range(start, stop, step))).replace("[", "{").replace("]", "}")}}},')
		   elif step == 1:
			print(f'{{input: FFIRange{{Start: NewObjectInt({start}), Stop: NewThingOrNoneNone[*ObjectInt](), Step: NewObjectInt(1)}}, expected: []int{str(list(range(start))).replace("[", "{").replace("]", "}")}}},')
		*/
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(-3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(-1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(-3)}, expected: []int{0, -3, -6, -9}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(-1)}, expected: []int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(-3)}, expected: []int{10, 7, 4, 1, -2, -5, -8}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(-1)}, expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(-10)), Step: interpreter.NewObjectInt(3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(-3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(-1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(1)}, expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(3)}, expected: []int{-10, -7, -4, -1}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(-3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(-1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(-3)}, expected: []int{10, 7, 4, 1}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(-1)}, expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(0)), Step: interpreter.NewObjectInt(3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(-3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(-1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(1)}, expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(3)}, expected: []int{-10, -7, -4, -1, 2, 5, 8}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(-3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(-1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(1)}, expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(3)}, expected: []int{0, 3, 6, 9}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(-3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(-1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneThing(interpreter.NewObjectInt(10)), Step: interpreter.NewObjectInt(3)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(-10), Stop: interpreter.NewThingOrNoneNone[*interpreter.ObjectInt](), Step: interpreter.NewObjectInt(1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(0), Stop: interpreter.NewThingOrNoneNone[*interpreter.ObjectInt](), Step: interpreter.NewObjectInt(1)}, expected: []int{}},
		{input: FFIRange{Start: interpreter.NewObjectInt(10), Stop: interpreter.NewThingOrNoneNone[*interpreter.ObjectInt](), Step: interpreter.NewObjectInt(1)}, expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	} {
		t.Run("", func(t *testing.T) {
			t.Parallel()

			i := interpreter.NewInterpreter(false)

			rng, err := ts.input.Call(i)
			require.NoError(t, err)
			is := []int{}
			for _, o := range rng.(*interpreter.ObjectList).Items {
				is = append(is, o.(*interpreter.ObjectInt).Value)
			}
			assert.Equal(t, ts.expected, is)
		})
	}
}
