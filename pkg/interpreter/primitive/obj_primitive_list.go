package primitive

import (
	"strings"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

// TODO: Proper interface
type ObjectList struct {
	itypes.Object

	Items []itypes.Object
}

var prototypeObjectList = itypes.NewObject(map[string]itypes.Object{
	"__add__":             ffi.NewFFI(ffiObjectListAdd{}),
	"__subscript__":       ffi.NewFFI(ffiObjectListSubscript{}),
	"__subscript_range__": ffi.NewFFI(ffiObjectListSubscriptRange{}),
})

func NewObjectList(items ...itypes.Object) *ObjectList {
	return &ObjectList{
		Object: prototypeObjectList,
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
		sb.WriteString(itypes.Repr(item))
	}
	sb.WriteString(")")
	return sb.String()
}

type ffiObjectListAdd struct {
	Self  *ObjectList `ffi:"self"`
	Right *ObjectList `ffi:"right"`
}

func (f ffiObjectListAdd) Call(i itypes.Interpreter) (itypes.Object, error) {
	items := make([]itypes.Object, 0, len(f.Self.Items)+len(f.Right.Items))
	items = append(items, f.Self.Items...)
	items = append(items, f.Right.Items...)
	return NewObjectList(items...), nil
}

type ffiObjectListSubscript struct {
	Self  *ObjectList `ffi:"self"`
	Start *ObjectInt  `ffi:"start"`
}

func (f ffiObjectListSubscript) Call(i itypes.Interpreter) (itypes.Object, error) {
	return subscript(f.Self.Items, f.Start.Value)
}

type ffiObjectListSubscriptRange struct {
	Self  *ObjectList    `ffi:"self"`
	Start paramSubscript `ffi:"start"`
	End   paramSubscript `ffi:"end"`
}

func (f ffiObjectListSubscriptRange) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewObjectList(subscriptRange(f.Self.Items, f.Start, f.End)...), nil
}
