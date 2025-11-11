package interpreter

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/parser/ast"
)

type ObjectBool struct {
	itypes.Object

	Value bool
}

func NewObjectBool(v bool) *ObjectBool {
	return &ObjectBool{
		Object: NewObject(map[string]itypes.Object{
			"__unary_binary_not__": methodBoolUnaryBinaryNot{Object: NewObject(nil)},
		}),
		Value: v,
	}
}

func (ob *ObjectBool) String(i itypes.Interpreter) (string, error) {
	return fmt.Sprintf("%t", ob.Value), nil
}

func (ob *ObjectBool) VisitExpressionTernary(i itypes.Interpreter, left ast.Expression, cond itypes.Object, right ast.Expression) (any, error) {
	if ob.Value {
		return left.Accept(i)
	} else {
		return right.Accept(i)
	}
}

var _ FlowTernary = (*ObjectBool)(nil)

type methodBoolUnaryBinaryNot struct {
	itypes.Object
}

func (mbubn methodBoolUnaryBinaryNot) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return UnaryParams, nil
}

func (mbubn methodBoolUnaryBinaryNot) Call(i itypes.Interpreter) (itypes.Object, error) {
	if selfValue, err := ArgAsBool(i, "self"); err != nil {
		return nil, err
	} else {
		return NewObjectBool(!selfValue), nil
	}
}
