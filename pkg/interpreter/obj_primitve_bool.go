package interpreter

type ObjectBool struct {
	Object

	Value bool
}

func NewObjectBool(v bool) Object {
	return &ObjectBool{
		Object: NewObject(map[string]Object{
			"__ternary__": methodBoolTernary{Object: NewObject(nil)},
		}),
		Value: v,
	}
}

type methodBoolTernary struct {
	Object
}

func (mbt methodBoolTernary) Params(i *Interpreter) (*Params, error) {
	return TernaryParams, nil
}

func (mbt methodBoolTernary) Call(i *Interpreter) (Object, error) {
	if value, err := ArgAsBool(i, "self"); err != nil {
		return nil, err
	} else if value {
		return i.Scope.GetArg("left")
	} else {
		return i.Scope.GetArg("right")
	}
}
