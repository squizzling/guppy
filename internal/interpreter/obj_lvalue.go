package interpreter

import (
	"github.com/squizzling/types/pkg/result"
)

type ObjectLValue struct {
	left  Object
	right Object
}

func NewLValue(left Object, right Object) Object {
	return &ObjectLValue{
		left:  left,
		right: right,
	}
}

func (lv *ObjectLValue) Args(i *Interpreter) result.Result[[]ArgData] {
	return i.doArgs(lv.right)
}

func (lv *ObjectLValue) Call(i *Interpreter) result.Result[Object] {
	return i.doCall(lv.right)
}

func (lv *ObjectLValue) Member(i *Interpreter, obj Object, memberName string) result.Result[Object] {
	return lv.right.Member(i, lv.right, memberName)
}
