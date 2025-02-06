package interpreter

// TODO: Proper interface
type ObjectList struct {
	Object

	Items []Object
}

func NewObjectList(items ...Object) Object {
	return &ObjectList{
		Items: items,
	}
}
