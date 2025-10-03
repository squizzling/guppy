package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type methodFill struct {
	interpreter.Object
}

func (mf methodFill) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "value", Default: interpreter.NewObjectNone()},
			{Name: "duration", Default: interpreter.NewObjectNone()},
			{Name: "maxCount", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (mf methodFill) resolveDuration(i *interpreter.Interpreter) (int, error) {
	if by, err := i.Scope.GetArg("duration"); err != nil {
		return 0, err
	} else {
		switch by := by.(type) {
		case *interpreter.ObjectNone:
			return 0, nil // explicitly nil
		default:
			return 0, fmt.Errorf("duration is %T not *interpreter.ObjectNone", by)
		}
	}
}

func (mf methodFill) resolveMaxCount(i *interpreter.Interpreter) (int, error) {
	if by, err := i.Scope.GetArg("maxCount"); err != nil {
		return 0, err
	} else {
		switch by := by.(type) {
		case *interpreter.ObjectNone:
			return 0, nil // explicitly nil
		default:
			return 0, fmt.Errorf("maxCount is %T not *interpreter.ObjectNone", by)
		}
	}
}

func (mf methodFill) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if value, err := i.Scope.Get("value"); err != nil {
		return nil, err
	} else if duration, err := mf.resolveDuration(i); err != nil {
		return nil, err
	} else if maxCount, err := mf.resolveMaxCount(i); err != nil {
		return nil, err
	} else {
		return NewStreamFill(newStreamObject(), unpublish(self), value, duration, maxCount), nil
	}
}
