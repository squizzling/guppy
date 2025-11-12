package interpreter

import (
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
	"guppy/pkg/parser/ast"
)

type ObjectFunction struct {
	itypes.Object

	name   string
	params *itypes.Params
	scope  *scope
	body   ast.Statement
}

func NewObjectFunction(name string, params *itypes.Params, scope *scope, body ast.Statement) itypes.Object {
	// TODO: Don't use scope, as it's not exported.  The visibility needs revisiting generally.
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
