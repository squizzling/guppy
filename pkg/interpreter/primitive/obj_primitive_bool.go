package primitive

import (
	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/parser/ast"
)

type ObjectBool struct {
	itypes.Object

	Value bool
}

var prototypeObjectBool = itypes.NewObject(map[string]itypes.Object{
	"__unary_binary_not__": ffi.NewFFI(ffiObjectBoolUnaryBinaryNot{}),
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
