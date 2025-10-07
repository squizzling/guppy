package interpreter

import (
	"guppy/pkg/parser/ast"
)

type ObjectBool struct {
	Object

	Value bool
}

func NewObjectBool(v bool) Object {
	return &ObjectBool{
		Object: NewObject(map[string]Object{
			"__unary_binary_not__": methodBoolUnaryBinaryNot{Object: NewObject(nil)},
		}),
		Value: v,
	}
}

func (ob *ObjectBool) VisitExpressionTernary(i *Interpreter, left ast.Expression, cond Object, right ast.Expression) (any, error) {
	if ob.Value {
		return left.Accept(i)
	} else {
		return right.Accept(i)
	}
}

var _ FlowTernary = (*ObjectBool)(nil)

type methodBoolUnaryBinaryNot struct {
	Object
}

func (mbubn methodBoolUnaryBinaryNot) Params(i *Interpreter) (*Params, error) {
	return UnaryParams, nil
}

func (mbubn methodBoolUnaryBinaryNot) Call(i *Interpreter) (Object, error) {
	if selfValue, err := ArgAsBool(i, "self"); err != nil {
		return nil, err
	} else {
		return NewObjectBool(!selfValue), nil
	}
}
