package primitive

import (
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/parser/ast"
)

type ObjectBool struct {
	itypes.Object

	Value bool
}

var prototypeObjectBool = itypes.NewObject(map[string]itypes.Object{
	"__binary_and__":       ffi.NewFFI(ffiObjectBoolBinaryOp{op: 0}),
	"__binary_or__":        ffi.NewFFI(ffiObjectBoolBinaryOp{op: 1}),
	"__unary_binary_not__": ffi.NewFFI(ffiObjectBoolUnaryBinaryNot{}),
	"__is__":               ffi.NewFFI(ffiObjectBoolIs{invert: false}),
	"__isnot__":            ffi.NewFFI(ffiObjectBoolIs{invert: true}),
})

func NewObjectBool(v bool) *ObjectBool {
	return &ObjectBool{
		Object: prototypeObjectBool,
		Value:  v,
	}
}

func (ob *ObjectBool) Repr() string {
	if ob.Value {
		return "True"
	} else {
		return "False"
	}
}

func (ob *ObjectBool) String(i itypes.Interpreter) (string, error) {
	return ob.Repr(), nil
}

func (ob *ObjectBool) VisitExpressionTernary(i itypes.Interpreter, left ast.Expression, cond itypes.Object, right ast.Expression) (any, error) {
	if ob.Value {
		return left.Accept(i)
	} else {
		return right.Accept(i)
	}
}

type ffiObjectBoolUnaryBinaryNot struct {
	Self *ObjectBool `ffi:"self"`
}

func (f ffiObjectBoolUnaryBinaryNot) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewObjectBool(!f.Self.Value), nil
}

var _ itypes.FlowTernary = (*ObjectBool)(nil)

type ffiObjectBoolIs struct {
	Self  *ObjectBool `ffi:"self"`
	Right struct {
		Bool   *ObjectBool
		Object itypes.Object
	} `ffi:"right"`

	invert bool
}

func (f ffiObjectBoolIs) Call(i itypes.Interpreter) (itypes.Object, error) {
	if f.Right.Bool != nil {
		return NewObjectBool((f.Self.Value == f.Right.Bool.Value) != f.invert), nil
	} else {
		return NewObjectBool(f.invert), nil
	}
}

type ffiObjectBoolBinaryOp struct {
	Self  *ObjectBool `ffi:"self"`
	Right *ObjectBool `ffi:"right"`

	op int
}

func (f ffiObjectBoolBinaryOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	switch f.op {
	case 0:
		return NewObjectBool(f.Self.Value && f.Right.Value), nil
	default:
		return NewObjectBool(f.Self.Value || f.Right.Value), nil
	}
}
