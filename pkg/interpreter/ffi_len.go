package interpreter

import (
	"fmt"
)

type FFILen struct {
	Object
}

func (f FFILen) Repr() string {
	return "len"
}

func (f FFILen) Params(i *Interpreter) (*Params, error) {
	return &Params{
		Params: []ParamDef{
			{
				Name: "value",
			},
		},
	}, nil
}

func (f FFILen) Call(i *Interpreter) (Object, error) {
	if value, err := i.Scope.GetArg("value"); err != nil {
		return nil, err
	} else {
		switch value := value.(type) {
		case *ObjectList:
			return NewObjectInt(len(value.Items)), nil
		default:
			return nil, fmt.Errorf("len() value is %T, not *interpreter.ObjectList", value)
		}
	}
}
