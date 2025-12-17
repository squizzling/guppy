package primitive

import (
	"strings"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

// TODO: Proper interface
type ObjectTuple struct {
	itypes.Object

	Items []itypes.Object
}

var prototypeObjectTuple = itypes.NewObject(map[string]itypes.Object{
	"__add__":             ffi.NewFFI(ffiObjectTupleAdd{}),
	"__subscript__":       ffi.NewFFI(ffiObjectTupleSubscript{}),
	"__subscript_range__": ffi.NewFFI(ffiObjectTupleSubscriptRange{}),
})

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

type ffiObjectTupleAdd struct {
	Self  *ObjectTuple `ffi:"self"`
	Right *ObjectTuple `ffi:"right"`
}

func (f ffiObjectTupleAdd) Call(i itypes.Interpreter) (itypes.Object, error) {
	items := make([]itypes.Object, 0, len(f.Self.Items)+len(f.Right.Items))
	items = append(items, f.Self.Items...)
	items = append(items, f.Right.Items...)
	return NewObjectTuple(items...), nil
}

type ffiObjectTupleSubscript struct {
	Self  *ObjectTuple `ffi:"self"`
	Start *ObjectInt   `ffi:"start"`
}

func (f ffiObjectTupleSubscript) Call(i itypes.Interpreter) (itypes.Object, error) {
	return subscript(f.Self.Items, f.Start.Value)
}

type ffiObjectTupleSubscriptRange struct {
	Self  *ObjectTuple `ffi:"self"`
	Start struct {
		None *ObjectNone
		Int  *ObjectInt
	} `ffi:"start"`
	End struct {
		None *ObjectNone
		Int  *ObjectInt
	} `ffi:"end"`
}

func (f ffiObjectTupleSubscriptRange) Call(i itypes.Interpreter) (itypes.Object, error) {
	var start *int
	if f.Start.Int != nil {
		start = &f.Start.Int.Value
	}

	var end *int
	if f.End.Int != nil {
		end = &f.End.Int.Value
	}
	return NewObjectTuple(subscriptRange(f.Self.Items, start, end)...), nil
}
