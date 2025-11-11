package itypes

type ObjectLValue struct {
	Left  Object
	Right Object
}

func NewLValue(left Object, right Object) Object {
	return &ObjectLValue{
		Left:  left,
		Right: right,
	}
}

func (lv *ObjectLValue) Params(i Interpreter) (*Params, error) {
	return i.DoParams(lv.Right)
}

func (lv *ObjectLValue) Call(i Interpreter) (Object, error) {
	return i.DoCall(lv.Right)
}

func (lv *ObjectLValue) Member(i Interpreter, obj Object, memberName string) (Object, error) {
	return lv.Right.Member(i, lv.Right, memberName)
}
