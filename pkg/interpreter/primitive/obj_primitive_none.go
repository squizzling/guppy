package primitive

import (
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type ObjectNone struct {
	itypes.Object
}

var prototypeObjectNone = itypes.NewObject(map[string]itypes.Object{
	"__eq__":    ffi.NewFFI(ffiObjectNoneEqual{invert: false}),
	"__ne__":    ffi.NewFFI(ffiObjectNoneEqual{invert: true}),
	"__is__":    ffi.NewFFI(ffiObjectNoneIs{invert: false, reverseMethod: "__ris__"}),
	"__isnot__": ffi.NewFFI(ffiObjectNoneIs{invert: true, reverseMethod: "__risnot__"}),
})

func NewObjectNone() *ObjectNone {
	return &ObjectNone{
		Object: prototypeObjectNone,
	}
}

func (on *ObjectNone) Repr() string {
	return "None"
}

func (on *ObjectNone) String(i itypes.Interpreter) (string, error) {
	return "None", nil
}

type ffiObjectNoneEqual struct {
	Self   *ObjectNone   `ffi:"self"`
	Right  itypes.Object `ffi:"right"`
	invert bool
}

func (f ffiObjectNoneEqual) Call(i itypes.Interpreter) (itypes.Object, error) {
	_, ok := f.Right.(*ObjectNone)
	return NewObjectBool(ok != f.invert), nil
}

type ffiObjectNoneIs struct {
	Self          *ObjectNone   `ffi:"self"`
	Right         itypes.Object `ffi:"right"`
	invert        bool
	reverseMethod string
}

func (f ffiObjectNoneIs) Call(i itypes.Interpreter) (itypes.Object, error) {
	if reverseIs, err := f.Right.Member(i, f.Right, f.reverseMethod); err == nil {
		// If it exists, we always use the reverse method, because it's more likely to be the intended behavior.
		if reverseIsCall, ok := reverseIs.(itypes.FlowCall); ok {
			return reverseIsCall.Call(i)
		}
	}

	_, ok := f.Right.(*ObjectNone)
	return NewObjectBool(ok != f.invert), nil
}
