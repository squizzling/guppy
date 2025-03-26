package interpreter

import (
	"guppy/internal/parser/ast"
)

type ObjectFunction struct {
	Object

	name   string
	params *Params
	scope  *scope
	body   ast.Statement
}

func NewObjectFunction(name string, params *Params, scope *scope, body ast.Statement) Object {
	// TODO: Don't use scope, as it's not exported.  The visibility needs revisiting generally.
	return &ObjectFunction{
		Object: NewObject(nil),
		name:   name,
		params: params,
		scope:  scope,
		body:   body,
	}
}

func (of *ObjectFunction) Params(i *Interpreter) (*Params, error) {
	return of.params, nil
}

func (of *ObjectFunction) Call(i *Interpreter) (Object, error) {
	// This is where we would evaluate deferred statements, except SFX doesn't.
	o, err := of.body.Accept(i)
	if err != nil {
		panic(err)
	}
	if o == nil {
		return NewObjectNone(), nil
	} else {
		return o.(Object), nil
	}
}
