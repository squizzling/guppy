package interpreter

import (
	"fmt"
)

type ObjectInt struct {
	Object

	Value int
}

func NewObjectInt(i int) Object {
	return &ObjectInt{
		Object: NewObject(map[string]Object{
			"__add__": methodIntAdd{NewObject(nil)},
		}),
		Value: i,
	}
}

type methodIntAdd struct {
	Object
}

func (mia methodIntAdd) Params(i *Interpreter) (*Params, error) {
	return BinaryParams, nil
}

func (mia methodIntAdd) Call(i *Interpreter) (Object, error) {
	if self, err := ArgAs[*ObjectInt](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.Get("right"); err != nil {
		return nil, err
	} else {
		switch right := right.(type) {
		case *ObjectInt:
			return NewObjectInt(self.Value + right.Value), nil
		default:
			return nil, fmt.Errorf("methodIntAdd: unknown type %T", right)
		}
	}
}

var _ = FlowCall(methodIntAdd{})
