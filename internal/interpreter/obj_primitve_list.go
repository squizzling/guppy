package interpreter

// TODO: Proper interface
type ObjectList struct {
	Object

	Items []Object
}

func NewObjectList(items ...Object) Object {
	return &ObjectList{
		Object: NewObject(nil),
		Items:  items,
	}
}
