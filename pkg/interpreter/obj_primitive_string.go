package interpreter

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
)

type ObjectString struct {
	itypes.Object

	Value string
}

func NewObjectString(s string) *ObjectString {
	return &ObjectString{
		Object: NewObject(map[string]itypes.Object{
			"__add__": methodStringAdd{Object: NewObject(nil)},
			"__eq__":  methodStringEqual{Object: NewObject(nil)},
			"__ne__":  methodStringNotEqual{Object: NewObject(nil)},
		}),
		Value: s,
	}
}

func (os *ObjectString) String(i itypes.Interpreter) (string, error) {
	return os.Value, nil
}

var _ = FlowStringable(&ObjectString{})

type methodStringAdd struct {
	itypes.Object
}

func (msa methodStringAdd) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return BinaryParams, nil
}

func (msa methodStringAdd) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := ArgAs[*ObjectString](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
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

type methodStringEqual struct {
	itypes.Object
}

func (mse methodStringEqual) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return BinaryParams, nil
}

func (mse methodStringEqual) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := ArgAs[*ObjectString](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right := right.(type) {
		case *ObjectString:
			return NewObjectBool(self.Value == right.Value), nil
		default:
			return nil, fmt.Errorf("methodStringAdd: unknown type %T", right)
		}
	}
}

type methodStringNotEqual struct {
	itypes.Object
}

func (msne methodStringNotEqual) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return BinaryParams, nil
}

func (msne methodStringNotEqual) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := ArgAs[*ObjectString](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right := right.(type) {
		case *ObjectString:
			return NewObjectBool(self.Value != right.Value), nil
		default:
			return nil, fmt.Errorf("methodStringAdd: unknown type %T", right)
		}
	}
}
