package interpreter

import (
	"fmt"
)

type ObjectString struct {
	Object

	Value string
}

func NewObjectString(s string) Object {
	return &ObjectString{
		Object: NewObject(map[string]Object{
			"__add__": methodStringAdd{Object: NewObject(nil)},
		}),
		Value: s,
	}
}

func (os *ObjectString) String(i *Interpreter) (string, error) {
	return os.Value, nil
}

var _ = FlowStringable(&ObjectString{})

type methodStringAdd struct {
	Object
}

func (msa methodStringAdd) Params(i *Interpreter) (*Params, error) {
	return BinaryParams, nil
}

func (msa methodStringAdd) Call(i *Interpreter) (Object, error) {
	if self, err := ArgAs[*ObjectString](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right := right.(type) {
		case *ObjectString:
			return NewObjectString(self.Value + right.Value), nil
		default:
			return nil, fmt.Errorf("methodStringAdd: unknown type %T", right)
		}
	}
}
