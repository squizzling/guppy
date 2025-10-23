package interpreter

import (
	"fmt"
)

type FFIRange struct {
	Object
}

func (f FFIRange) Repr() string {
	return "range"
}

func (f FFIRange) Params(i *Interpreter) (*Params, error) {
	return &Params{
		Params: []ParamDef{
			{Name: "start"},
			{Name: "stop", Default: NewObjectNone()},
			{Name: "step", Default: NewObjectInt(1)},
		},
	}, nil
}

func (f FFIRange) Call(i *Interpreter) (Object, error) {
	if start, err := ArgAsLong(i, "start"); err != nil {
		return nil, err
	} else if pStop, err := ArgAsOptLong(i, "stop"); err != nil {
		return nil, err
	} else if pStop == nil {
		// generate range from 0 to start with step 1
		return f.newRange(0, start, 1)
	} else if pStep, err := ArgAsOptLong(i, "step"); err != nil {
		return nil, err
	} else if pStep == nil {
		// generate range from start to *pStop with step 1
		return f.newRange(start, *pStop, 1)
	} else {
		// generate range from start to *pStop with step *pStep
		return f.newRange(start, *pStop, *pStep)
	}
}

func (f FFIRange) newRange(start int, stop int, step int) (Object, error) {
	if step == 0 {
		return nil, fmt.Errorf("invalid step in range(%d, %d, %d)", start, stop, step)
	}

	var items []Object
	if stop < start && step < 0 {
		for i := start; i > stop; i += step {
			items = append(items, NewObjectInt(i))
		}
	} else if step > 0 {
		for i := start; i < stop; i += step {
			items = append(items, NewObjectInt(i))
		}
	}
	return NewObjectList(items...), nil
}
