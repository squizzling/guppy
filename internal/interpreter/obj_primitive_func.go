package interpreter

import (
	"guppy/internal/parser/ast"
)

type ObjectFunction struct {
	Object

	name   string
	params *Params
	body   ast.Statement
}

func NewObjectFunction(name string, params *Params, body ast.Statement) Object {
	return &ObjectFunction{
		Object: NewObject(nil),
		name:   name,
		params: params,
		body:   body,
	}
}

func (of *ObjectFunction) Params(i *Interpreter) (*Params, error) {
	return of.params, nil
}

func (of *ObjectFunction) Call(i *Interpreter) (Object, error) {
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
