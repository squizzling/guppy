package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type methodFill struct {
	itypes.Object
}

func (mf methodFill) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "value", Default: primitive.NewObjectNone()},
			{Name: "duration", Default: primitive.NewObjectNone()},
			{Name: "maxCount", Default: primitive.NewObjectNone()},
		},
	}, nil
}

func (mf methodFill) resolveDuration(i itypes.Interpreter) (int, error) {
	if by, err := i.GetArg("duration"); err != nil {
		return 0, err
	} else {
		switch by := by.(type) {
		case *primitive.ObjectNone:
			return 0, nil // explicitly nil
		default:
			return 0, fmt.Errorf("duration is %T not *interpreter.ObjectNone", by)
		}
	}
}

func (mf methodFill) resolveMaxCount(i itypes.Interpreter) (int, error) {
	if by, err := i.GetArg("maxCount"); err != nil {
		return 0, err
	} else {
		switch by := by.(type) {
		case *primitive.ObjectNone:
			return 0, nil // explicitly nil
		default:
			return 0, fmt.Errorf("maxCount is %T not *interpreter.ObjectNone", by)
		}
	}
}

func (mf methodFill) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if value, err := i.GetArg("value"); err != nil {
		return nil, err
	} else if duration, err := mf.resolveDuration(i); err != nil {
		return nil, err
	} else if maxCount, err := mf.resolveMaxCount(i); err != nil {
		return nil, err
	} else {
		return NewStreamMethodFill(newStreamObject(), unpublish(self), value, duration, maxCount), nil
	}
}
