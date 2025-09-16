package interpreter

import (
	"fmt"
	"strconv"
)

type FFIStr struct {
	Object
}

func (f FFIStr) Repr() string {
	return "str"
}

func (f FFIStr) Params(i *Interpreter) (*Params, error) {
	return &Params{
		Params: []ParamDef{
			{
				Name:    "value",
				Default: NewObjectMissing(),
			},
		},
	}, nil
}

func (f FFIStr) Call(i *Interpreter) (Object, error) {
	if value, err := i.Scope.GetArg("value"); err != nil {
		return nil, err
	} else {
		switch value := value.(type) {
		case *ObjectInt:
			return NewObjectString(strconv.Itoa(value.Value)), nil
		case *ObjectDouble:
			return NewObjectString(strconv.FormatFloat(value.Value, 'f', 6, 64)), nil
		case *ObjectString:
			return value, nil
		case *ObjectNone:
			return NewObjectString("None"), nil
		default:
			return nil, fmt.Errorf("[FFIStr] %T is not *interpreter.ObjectInt, *interpreter.ObjectDouble, or *interpreter.ObjectString", value)
		}
	}
}
