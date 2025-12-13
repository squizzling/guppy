package primitive

import (
	"strings"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
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

func (ot *ObjectTuple) Repr() string {
	var sb strings.Builder
	sb.WriteString("tuple(")
	for idx, item := range ot.Items {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(itypes.Repr(item))
	}
	sb.WriteString(")")
	return sb.String()
}
