package interpreter

type ObjectNone struct {
	Object
}

func NewObjectNone() Object {
	return &ObjectNone{
		NewObject(map[string]Object{
			"__eq__":    methodNoneEqual{Object: NewObject(nil)},
			"__ne__":    methodNoneNotEqual{Object: NewObject(nil)},
			"__is__":    methodNoneIs{Object: NewObject(nil), invert: false, reverseInvert: "__ris__"},
			"__isnot__": methodNoneIs{Object: NewObject(nil), invert: true, reverseInvert: "__risnot__"},
		}),
	}
}

func (on *ObjectNone) Repr() string {
	return "None"
}

type methodNoneEqual struct {
	Object
}

func (mne methodNoneEqual) Params(i *Interpreter) (*Params, error) {
	return BinaryParams, nil
}

func (mne methodNoneEqual) Call(i *Interpreter) (Object, error) {
	if right, err := i.Scope.GetArg("right"); err != nil {
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
	Object
}

func (mnne methodNoneNotEqual) Params(i *Interpreter) (*Params, error) {
	return BinaryParams, nil
}

func (mne methodNoneNotEqual) Call(i *Interpreter) (Object, error) {
	if right, err := i.Scope.GetArg("right"); err != nil {
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
	Object

	invert        bool
	reverseInvert string
}

func (mni methodNoneIs) Params(i *Interpreter) (*Params, error) {
	return BinaryParams, nil
}

func (mni methodNoneIs) Call(i *Interpreter) (Object, error) {
	right, err := i.Scope.GetArg("right")
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
