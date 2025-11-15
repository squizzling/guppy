package primitive

import (
	"guppy/pkg/interpreter/itypes"
)

// TODO: Proper interface
type ObjectTuple struct {
	itypes.Object

	Items []itypes.Object
}

var prototypeObjectTuple = itypes.NewObject(nil)

func NewObjectTuple(items ...itypes.Object) *ObjectTuple {
	return &ObjectTuple{
		Object: prototypeObjectTuple,
		Items:  items,
	}
}
