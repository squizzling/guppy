package primitive

import (
	"fmt"
	"strings"

	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
)

// TODO: Proper interface
type ObjectList struct {
	itypes.Object

	Items []itypes.Object
}

var prototypeObjectList = itypes.NewObject(map[string]itypes.Object{
	"__add__":       ffi.NewFFI(ffiObjectListAdd{}),
	"__subscript__": ffi.NewFFI(ffiObjectListSubscript{}),
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
	if len(f.Self.Items) < f.Start.Value+1 || f.Start.Value < 0 {
		// TODO: Does flow support x[-1] for last item?
		return nil, fmt.Errorf("index %d out of range (0 - %d)", f.Start.Value, len(f.Self.Items)-1)
	} else {
		return f.Self.Items[f.Start.Value], nil
	}
}
