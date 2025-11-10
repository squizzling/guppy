package interpreter

import (
	"guppy/pkg/interpreter/itypes"
)

type ObjectLValue struct {
	left  itypes.Object
	right itypes.Object
}

func NewLValue(left itypes.Object, right itypes.Object) itypes.Object {
	return &ObjectLValue{
		left:  left,
		right: right,
	}
}

func (lv *ObjectLValue) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return i.DoParams(lv.right)
}

func (lv *ObjectLValue) Call(i itypes.Interpreter) (itypes.Object, error) {
	return i.DoCall(lv.right)
}

func (lv *ObjectLValue) Member(i itypes.Interpreter, obj itypes.Object, memberName string) (itypes.Object, error) {
	return lv.right.Member(i, lv.right, memberName)
}
