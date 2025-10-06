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
		Object: NewObject(nil),
		Value:  v,
	}
}

func (ob *ObjectBool) VisitExpressionTernary(i *Interpreter, left ast.Expression, right ast.Expression) (any, error) {
	if ob.Value {
		return left.Accept(i)
	} else {
		return right.Accept(i)
	}
}
