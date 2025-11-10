package builtin

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
)

type FFIRange struct {
	Start *interpreter.ObjectInt                  `ffi:"start"`
	Stop  ffi.ThingOrNone[*interpreter.ObjectInt] `ffi:"stop"`
	Step  *interpreter.ObjectInt                  `ffi:"step"`
}

func NewFFIRange() interpreter.FlowCall {
	return ffi.NewFFI(FFIRange{
		Start: nil,
		Stop: ffi.ThingOrNone[*interpreter.ObjectInt]{
			None: interpreter.NewObjectNone(),
		},
		Step: interpreter.NewObjectInt(1),
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
			items = append(items, interpreter.NewObjectInt(i))
		}
	} else if step > 0 {
		for i := start; i < stop; i += step {
			items = append(items, interpreter.NewObjectInt(i))
		}
	}
	return interpreter.NewObjectList(items...), nil
}
