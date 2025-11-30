package primitive

import (
	"fmt"
	"strconv"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type ObjectInt struct {
	itypes.Object

	Value int
}

var prototypeObjectInt = itypes.NewObject(map[string]itypes.Object{
	"__add__":     ffi.NewFFI(ffiObjectIntMathOp{op: 0, reverseMethod: "__radd__"}),
	"__sub__":     ffi.NewFFI(ffiObjectIntMathOp{op: 1, reverseMethod: "__rsub__"}),
	"__mul__":     ffi.NewFFI(ffiObjectIntMathOp{op: 2, reverseMethod: "__rmul__"}),
	"__truediv__": ffi.NewFFI(ffiObjectIntMathOp{op: 3, reverseMethod: "__rtruediv__"}),

	"__unary_minus__": ffi.NewFFI(ffiObjectIntMathNeg{}),

	"__lt__": ffi.NewFFI(ffiObjectIntRelOp{op: 0, invert: false}),
	"__gt__": ffi.NewFFI(ffiObjectIntRelOp{op: 1, invert: false}),
	"__eq__": ffi.NewFFI(ffiObjectIntRelOp{op: 2, invert: false}),

	"__ge__": ffi.NewFFI(ffiObjectIntRelOp{op: 0, invert: true}),
	"__le__": ffi.NewFFI(ffiObjectIntRelOp{op: 1, invert: true}),
	"__ne__": ffi.NewFFI(ffiObjectIntRelOp{op: 2, invert: true}),
})

func NewObjectInt(i int) *ObjectInt {
	return &ObjectInt{
		Object: prototypeObjectInt,
		Value:  i,
	}
}

func (oi *ObjectInt) Repr() string {
	return fmt.Sprintf("int(%d)", oi.Value)
}

func (oi *ObjectInt) String(i itypes.Interpreter) (string, error) {
	return strconv.Itoa(oi.Value), nil
}

type ffiObjectIntRelOp struct {
	Self  *ObjectInt `ffi:"self"`
	Right struct {
		Int    *ObjectInt
		Double *ObjectDouble
	} `ffi:"right"`

	op     int
	invert bool
}

func (f ffiObjectIntRelOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	var right int
	if f.Right.Int != nil {
		right = f.Right.Int.Value
	} else {
		right = int(f.Right.Double.Value)
	}
	switch f.op {
	case 0:
		return NewObjectBool(f.Self.Value < right != f.invert), nil
	case 1:
		return NewObjectBool(f.Self.Value > right != f.invert), nil
	default:
		return NewObjectBool(f.Self.Value == right != f.invert), nil
	}
}

type ffiObjectIntMathOp struct {
	Self  *ObjectInt `ffi:"self"`
	Right struct {
		Int    *ObjectInt
		Double *ObjectDouble
		Object itypes.Object
	} `ffi:"right"`

	op            int
	reverseMethod string
}

func (f ffiObjectIntMathOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	var right int

	switch {
	case f.Right.Int != nil:
		right = f.Right.Int.Value
	case f.Right.Double != nil:
		right = int(f.Right.Double.Value)
	default:
		if reverseOp, err := f.Right.Object.Member(i, f.Right.Object, f.reverseMethod); err == nil {
			// If it exists, we always use the reverse method, because it's more likely to be the intended behavior.
			// We explicitly don't expose reverse methods for primitives though.
			if reverseOpCall, ok := reverseOp.(itypes.FlowCall); ok {
				return reverseOpCall.Call(i)
			}
		}
		return nil, fmt.Errorf("param `right` for ffiObjectIntMathOp.Right is %T not *primitive.ObjectInt, *primitive.ObjectDouble, or an itypes.Object with %s", f.Right.Object, f.reverseMethod)
	}

	switch f.op {
	case 0:
		return NewObjectInt(f.Self.Value + right), nil
	case 1:
		return NewObjectInt(f.Self.Value - right), nil
	case 2:
		return NewObjectInt(f.Self.Value * right), nil
	default:
		return NewObjectInt(f.Self.Value / right), nil
	}
}

type ffiObjectIntMathNeg struct {
	Self *ObjectInt `ffi:"self"`
}

func (f ffiObjectIntMathNeg) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewObjectInt(-f.Self.Value), nil
}
