package builtin

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/ftypes"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type FFIRange struct {
	Start *primitive.ObjectInt                     `ffi:"start"`
	Stop  ftypes.ThingOrNone[*primitive.ObjectInt] `ffi:"stop"`
	Step  *primitive.ObjectInt                     `ffi:"step"`
}

func NewFFIRange() itypes.FlowCall {
	return ffi.NewFFI(FFIRange{
		Start: nil,
		Stop: ftypes.ThingOrNone[*primitive.ObjectInt]{
			None: primitive.NewObjectNone(),
		},
		Step: primitive.NewObjectInt(1),
	})
}

func (f FFIRange) Call(i itypes.Interpreter) (itypes.Object, error) {
	var start int
	var stop int
	if f.Stop.Thing != nil {
		start = f.Start.Value
		stop = f.Stop.Thing.Value
	} else if f.Stop.None != nil {
		start = 0
		stop = f.Start.Value
	} else {
		return nil, fmt.Errorf("FFIRange.Call: FFIRange.Stop is not set")
	}
	step := f.Step.Value
	return f.newRange(start, stop, step)
}

func (f FFIRange) newRange(start int, stop int, step int) (itypes.Object, error) {
	if step == 0 {
		return nil, fmt.Errorf("invalid step in range(%d, %d, %d)", start, stop, step)
	}

	var items []itypes.Object
	if stop < start && step < 0 {
		for i := start; i > stop; i += step {
			items = append(items, primitive.NewObjectInt(i))
		}
	} else if step > 0 {
		for i := start; i < stop; i += step {
			items = append(items, primitive.NewObjectInt(i))
		}
	}
	return interpreter.NewObjectList(items...), nil
}
