package interpreter

import (
	"strings"

	"guppy/internal/parser/ast"
)

type ObjectDeferred struct {
	expr    ast.Expression
	desired []string
}

func NewObjectDeferred(expr ast.Expression, desired ...string) Object {
	return &ObjectDeferred{
		expr:    expr,
		desired: desired,
	}
}

func (o *ObjectDeferred) Repr() string {
	var sb strings.Builder
	sb.WriteString("deferred(")
	for idx, desired := range o.desired {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(desired)
	}
	sb.WriteString(")")
	return sb.String()
}

func (o *ObjectDeferred) Params(i *Interpreter) (*Params, error) {
	return &Params{
		StarParam: "s",
		//KWParam:   "k",
	}, nil
}

func (o *ObjectDeferred) Call(i *Interpreter) (Object, error) {
	return o, nil
}

func (o *ObjectDeferred) Member(i *Interpreter, obj Object, memberName string) (Object, error) {
	return o, nil
}
