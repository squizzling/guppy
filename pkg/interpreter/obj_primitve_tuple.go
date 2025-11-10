package interpreter

// TODO: Proper interface
type ObjectTuple struct {
	Object

	Items []Object
}

func NewObjectTuple(items ...Object) *ObjectTuple {
	return &ObjectTuple{
		Object: NewObject(nil),
		Items:  items,
	}
}
