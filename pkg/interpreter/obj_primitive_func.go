package interpreter

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
	"github.com/squizzling/guppy/pkg/interpreter/scope"
	"github.com/squizzling/guppy/pkg/parser/ast"
)

type ObjectFunction struct {
	itypes.Object

	name   string
	params *itypes.Params
	scope  *scope.Scope
	body   ast.Statement
}

func NewObjectFunction(name string, params *itypes.Params, scope *scope.Scope, body ast.Statement) itypes.Object {
	return &ObjectFunction{
		Object: itypes.NewObject(nil),
		name:   name,
		params: params,
		scope:  scope,
		body:   body,
	}
}

func (of *ObjectFunction) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return of.params, nil
}

func (of *ObjectFunction) Call(i itypes.Interpreter) (itypes.Object, error) {
	// This is where we would evaluate deferred statements, except SFX doesn't.
	o, err := of.body.Accept(i)
	if err != nil {
		//panic(err)
		return nil, err
	}
	if o == nil {
		return primitive.NewObjectNone(), nil
	} else {
		return o.(itypes.Object), nil
	}
}

func (of *ObjectFunction) Repr() string {
	// TODO: More information
	return fmt.Sprintf("func(%s)", of.name)
}
