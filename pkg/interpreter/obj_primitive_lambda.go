package interpreter

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/parser/ast"
)

type ObjectLambda struct {
	itypes.Object

	Identifier string
	Expression ast.Expression
}

func NewObjectLambda(identifier string, expr ast.Expression) itypes.Object {
	return &ObjectLambda{
		Object:     itypes.NewObject(nil),
		Identifier: identifier,
		Expression: expr,
	}
}

func (ol *ObjectLambda) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{Params: []itypes.ParamDef{{Name: ol.Identifier}}}, nil
}

func (ol *ObjectLambda) Call(i itypes.Interpreter) (itypes.Object, error) {
	if o, err := ol.Expression.Accept(i); err != nil {
		return nil, err
	} else {
		return o.(itypes.Object), nil
	}
}

func (ol *ObjectLambda) Repr() string {
	// TODO: More
	return fmt.Sprintf("lambda(%s)", ol.Identifier)
}
