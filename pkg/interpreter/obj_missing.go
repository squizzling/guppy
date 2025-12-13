package interpreter

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type ObjectMissing struct {
	itypes.Object
}

func NewObjectMissing() *ObjectMissing {
	return &ObjectMissing{
		Object: itypes.NewObject(map[string]itypes.Object{}),
	}
}

func (om *ObjectMissing) Repr() string {
	return "Missing"
}
