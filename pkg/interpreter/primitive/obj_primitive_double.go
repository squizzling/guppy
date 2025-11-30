package primitive

import (
	"fmt"
	"strconv"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type ObjectDouble struct {
	itypes.Object

	Value float64
}

var prototypeObjectDouble = itypes.NewObject(map[string]itypes.Object{
	"__add__":     ffi.NewFFI(ffiObjectDoubleMathOp{op: 0, reverseMethod: "__radd__"}),
	"__sub__":     ffi.NewFFI(ffiObjectDoubleMathOp{op: 1, reverseMethod: "__rsub__"}),
	"__mul__":     ffi.NewFFI(ffiObjectDoubleMathOp{op: 2, reverseMethod: "__rmul__"}),
	"__truediv__": ffi.NewFFI(ffiObjectDoubleMathOp{op: 3, reverseMethod: "__rtruediv__"}),

	"__unary_minus__": ffi.NewFFI(ffiObjectDoubleMathNeg{}),

	"__lt__": ffi.NewFFI(ffiObjectDoubleRelOp{op: 0, invert: false}),
	"__gt__": ffi.NewFFI(ffiObjectDoubleRelOp{op: 1, invert: false}),
	"__eq__": ffi.NewFFI(ffiObjectDoubleRelOp{op: 2, invert: false}),

	"__ge__": ffi.NewFFI(ffiObjectDoubleRelOp{op: 0, invert: true}),
	"__le__": ffi.NewFFI(ffiObjectDoubleRelOp{op: 1, invert: true}),
	"__ne__": ffi.NewFFI(ffiObjectDoubleRelOp{op: 2, invert: true}),
})

func NewObjectDouble(i float64) *ObjectDouble {
	return &ObjectDouble{
		Object: prototypeObjectDouble,
		Value:  i,
	}
}

func (od *ObjectDouble) Repr() string {
	return fmt.Sprintf("double(%f)", od.Value)
}

func (od *ObjectDouble) String(i itypes.Interpreter) (string, error) {
	return strconv.FormatFloat(od.Value, 'f', 6, 64), nil
}

type ffiObjectDoubleRelOp struct {
	Self  *ObjectDouble `ffi:"self"`
	Right struct {
		Double *ObjectDouble
		Int    *ObjectInt
	} `ffi:"right"`

	op     int
	invert bool
}

func (f ffiObjectDoubleRelOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	var right float64
	if f.Right.Double != nil {
		right = f.Right.Double.Value
	} else {
		right = float64(f.Right.Int.Value)
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

type ffiObjectDoubleMathOp struct {
	Self  *ObjectDouble `ffi:"self"`
	Right struct {
		Double *ObjectDouble
		Int    *ObjectInt
		Object itypes.Object
	} `ffi:"right"`

	op            int
	reverseMethod string
}

func (f ffiObjectDoubleMathOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	var right float64

	switch {
	case f.Right.Double != nil:
		right = f.Right.Double.Value
	case f.Right.Int != nil:
		right = float64(f.Right.Int.Value)
	default:
		if reverseOp, err := f.Right.Object.Member(i, f.Right.Object, f.reverseMethod); err == nil {
			// If it exists, we always use the reverse method, because it's more likely to be the intended behavior.
			// We explicitly don't expose reverse methods for primitives though.
			if reverseOpCall, ok := reverseOp.(itypes.FlowCall); ok {
				return reverseOpCall.Call(i)
			}
		}
		return nil, fmt.Errorf("param `right` for ffiObjectDoubleMathOp.Right is %T not *primitive.ObjectDouble, *primitive.ObjectInt, or an itypes.Object with %s", f.Right.Object, f.reverseMethod)
	}

	switch f.op {
	case 0:
		return NewObjectDouble(f.Self.Value + right), nil
	case 1:
		return NewObjectDouble(f.Self.Value - right), nil
	case 2:
		return NewObjectDouble(f.Self.Value * right), nil
	default:
		return NewObjectDouble(f.Self.Value / right), nil
	}
}

type ffiObjectDoubleMathNeg struct {
	Self *ObjectDouble `ffi:"self"`
}

func (f ffiObjectDoubleMathNeg) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewObjectDouble(-f.Self.Value), nil
}
