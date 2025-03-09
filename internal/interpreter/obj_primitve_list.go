package interpreter

import (
	"strings"
)

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

func (ol *ObjectList) Repr() string {
	var sb strings.Builder
	sb.WriteString("list(")
	for idx, item := range ol.Items {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(Repr(item))
	}
	sb.WriteString(")")
	return sb.String()
}
