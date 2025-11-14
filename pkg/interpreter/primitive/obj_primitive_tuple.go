package primitive

import (
	"guppy/pkg/interpreter/itypes"
)

// TODO: Proper interface
type ObjectTuple struct {
	itypes.Object

	Items []itypes.Object
}

func NewObjectTuple(items ...itypes.Object) *ObjectTuple {
	return &ObjectTuple{
		Object: itypes.NewObject(nil),
		Items:  items,
	}
}
