package interpreter

type ObjectNone struct {
	Object
}

func NewObjectNone() Object {
	return &ObjectNone{
		NewObject(map[string]Object{
			"__eq__":    methodNoneEqual{Object: NewObject(nil)},
			"__ne__":    methodNoneNotEqual{Object: NewObject(nil)},
			"__is__":    methodNoneIs{Object: NewObject(nil), invert: false},
			"__isnot__": methodNoneIs{Object: NewObject(nil), invert: true},
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

	invert bool
}

func (mni methodNoneIs) Params(i *Interpreter) (*Params, error) {
	return BinaryParams, nil
}

func (mni methodNoneIs) Call(i *Interpreter) (Object, error) {
	if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right.(type) {
		case *ObjectNone:
			return NewObjectBool(!mni.invert), nil
		default:
			return NewObjectBool(mni.invert), nil
		}
	}
}
