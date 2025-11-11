package interpreter

import (
	"guppy/pkg/interpreter/itypes"
)

type ObjectNone struct {
	itypes.Object
}

func NewObjectNone() *ObjectNone {
	return &ObjectNone{
		itypes.NewObject(map[string]itypes.Object{
			"__eq__":    methodNoneEqual{Object: itypes.NewObject(nil)},
			"__ne__":    methodNoneNotEqual{Object: itypes.NewObject(nil)},
			"__is__":    methodNoneIs{Object: itypes.NewObject(nil), invert: false, reverseInvert: "__ris__"},
			"__isnot__": methodNoneIs{Object: itypes.NewObject(nil), invert: true, reverseInvert: "__risnot__"},
		}),
	}
}

func (on *ObjectNone) Repr() string {
	return "None"
}

func (on *ObjectNone) String(i itypes.Interpreter) (string, error) {
	return "None", nil
}

type methodNoneEqual struct {
	itypes.Object
}

func (mne methodNoneEqual) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return BinaryParams, nil
}

func (mne methodNoneEqual) Call(i itypes.Interpreter) (itypes.Object, error) {
	if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right.(type) {
		case *ObjectNone:
			return NewObjectBool(true), nil
		default:
			return NewObjectBool(false), nil
		}
	}
}

type methodNoneNotEqual struct {
	itypes.Object
}

func (mnne methodNoneNotEqual) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return BinaryParams, nil
}

func (mne methodNoneNotEqual) Call(i itypes.Interpreter) (itypes.Object, error) {
	if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right.(type) {
		case *ObjectNone:
			return NewObjectBool(false), nil
		default:
			return NewObjectBool(true), nil
		}
	}
}

type methodNoneIs struct {
	itypes.Object

	invert        bool
	reverseInvert string
}

func (mni methodNoneIs) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return BinaryParams, nil
}

func (mni methodNoneIs) Call(i itypes.Interpreter) (itypes.Object, error) {
	right, err := i.GetArg("right")
	if err != nil {
		return nil, err
	} else if reverseIs, err := right.Member(i, right, mni.reverseInvert); err == nil {
		// If it exists, we always use the reverse method, because it's more likely to be the intended behavior.
		if reverseIsCall, ok := reverseIs.(FlowCall); ok {
			return reverseIsCall.Call(i)
		}
	}

	switch right.(type) {
	case *ObjectNone:
		return NewObjectBool(!mni.invert), nil
	default:
		return NewObjectBool(mni.invert), nil
	}
}
