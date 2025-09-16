package interpreter

import (
	"fmt"
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
	} else if valueStr, ok := value.(FlowStringable); !ok {
		return nil, fmt.Errorf("[FFIStr] %T is not FlowStringable", value)
	} else if s, err := valueStr.String(i); err != nil {
		return nil, fmt.Errorf("[FFIStr] %T String() failed: %w", value, err)
	} else {
		return NewObjectString(s), nil
	}
}
