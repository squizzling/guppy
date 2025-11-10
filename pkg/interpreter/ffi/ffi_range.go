package ffi

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFIRange struct {
	Start *interpreter.ObjectInt                          `ffi:"start"`
	Stop  interpreter.ThingOrNone[*interpreter.ObjectInt] `ffi:"stop"`
	Step  *interpreter.ObjectInt                          `ffi:"step"`
}

func NewFFIRange() interpreter.FlowCall {
	return interpreter.NewFFI(FFIRange{
		Start: nil,
		Stop: interpreter.ThingOrNone[*interpreter.ObjectInt]{
			None: interpreter.NewObjectNone(),
		},
		Step: interpreter.NewObjectInt(1),
	})
}

func (f FFIRange) Call() (interpreter.Object, error) {
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

func (f FFIRange) newRange(start int, stop int, step int) (interpreter.Object, error) {
	if step == 0 {
		return nil, fmt.Errorf("invalid step in range(%d, %d, %d)", start, stop, step)
	}

	var items []interpreter.Object
	if stop < start && step < 0 {
		for i := start; i > stop; i += step {
			items = append(items, interpreter.NewObjectInt(i))
		}
	} else if step > 0 {
		for i := start; i < stop; i += step {
			items = append(items, interpreter.NewObjectInt(i))
		}
	}
	return interpreter.NewObjectList(items...), nil
}
