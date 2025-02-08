package interpreter

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

func (lv *ObjectLValue) Args(i *Interpreter) ([]ArgData, error) {
	return i.doArgs(lv.right)
}

func (lv *ObjectLValue) Call(i *Interpreter) (Object, error) {
	return i.doCall(lv.right)
}

func (lv *ObjectLValue) Member(i *Interpreter, obj Object, memberName string) (Object, error) {
	return lv.right.Member(i, lv.right, memberName)
}
