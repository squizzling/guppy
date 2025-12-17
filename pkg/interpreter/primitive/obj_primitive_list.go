package primitive

import (
	"fmt"
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
	if len(f.Self.Items) < f.Start.Value+1 || f.Start.Value < 0 {
		// TODO: Does flow support x[-1] for last item?
		return nil, fmt.Errorf("index %d out of range (0 - %d)", f.Start.Value, len(f.Self.Items)-1)
	} else {
		return f.Self.Items[f.Start.Value], nil
	}
}

type ffiObjectListSubscriptRange struct {
	Self  *ObjectList `ffi:"self"`
	Start struct {
		None *ObjectNone
		Int  *ObjectInt
	} `ffi:"start"`
	End struct {
		None *ObjectNone
		Int  *ObjectInt
	} `ffi:"end"`
}

func clamp(minVal int, val int, maxVal int) int {
	return max(minVal, min(maxVal, val))
}

func (f ffiObjectListSubscriptRange) Call(i itypes.Interpreter) (itypes.Object, error) {
	var start int
	if f.Start.Int == nil {
		start = 0
	} else if f.Start.Int.Value < 0 {
		start = len(f.Self.Items) + f.Start.Int.Value
	} else {
		start = f.Start.Int.Value
	}

	var end int
	if f.End.Int == nil {
		end = len(f.Self.Items)
	} else if f.End.Int.Value < 0 {
		end = len(f.Self.Items) + f.End.Int.Value
	} else {
		end = f.End.Int.Value
	}

	// No IndexError for range
	start = clamp(0, start, len(f.Self.Items))
	end = clamp(0, end, len(f.Self.Items))

	if end < start {
		return NewObjectList(), nil
	} else {
		return NewObjectList(f.Self.Items[start:end]...), nil
	}
}
