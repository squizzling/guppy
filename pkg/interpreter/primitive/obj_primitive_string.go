package primitive

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type ObjectString struct {
	itypes.Object

	Value string
}

var prototypeObjectString = itypes.NewObject(map[string]itypes.Object{
	"__add__": ffi.NewFFI(ffiObjectStringAdd{}),
	"__eq__":  ffi.NewFFI(ffiObjectStringRelOp{op: 0, invert: false}),
	"__ne__":  ffi.NewFFI(ffiObjectStringRelOp{op: 0, invert: true}),
})

func NewObjectString(s string) *ObjectString {
	return &ObjectString{
		Object: prototypeObjectString,
		Value:  s,
	}
}

func (os *ObjectString) String(i itypes.Interpreter) (string, error) {
	return os.Value, nil
}

func (os *ObjectString) Repr() string {
	return fmt.Sprintf("string(%s)", os.Value)
}

type ffiObjectStringAdd struct {
	Self  *ObjectString `ffi:"self"`
	Right *ObjectString `ffi:"right"`
}

func (f ffiObjectStringAdd) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewObjectString(f.Self.Value + f.Right.Value), nil
}

type ffiObjectStringRelOp struct {
	Self  *ObjectString `ffi:"self"`
	Right struct {
		String *ObjectString
		Object itypes.Object
	} `ffi:"right"`

	op     int // unused for now
	invert bool
}

func (f ffiObjectStringRelOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	if f.Right.String != nil {
		return NewObjectBool(f.Self.Value == f.Right.String.Value != f.invert), nil
	} else {
		return NewObjectBool(f.invert), nil
	}
}
