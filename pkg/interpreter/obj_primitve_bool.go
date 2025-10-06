package interpreter

type ObjectBool struct {
	Object

	Value bool
}

func NewObjectBool(v bool) Object {
	return &ObjectBool{
		Object: NewObject(nil),
		Value:  v,
	}
}
