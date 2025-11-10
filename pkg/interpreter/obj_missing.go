package interpreter

import (
	"guppy/pkg/interpreter/itypes"
)

type ObjectMissing struct {
	itypes.Object
}

func NewObjectMissing() itypes.Object {
	return &ObjectMissing{
		NewObject(map[string]itypes.Object{}),
	}
}

func (om *ObjectMissing) Repr() string {
	return "Missing"
}
