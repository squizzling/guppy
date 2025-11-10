package interpreter

import (
	"strings"

	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/parser/ast"
)

type ObjectDeferred struct {
	expr    ast.Expression
	desired []string
}

func NewObjectDeferred(expr ast.Expression, desired ...string) itypes.Object {
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

func (o *ObjectDeferred) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "s",
		KWParam:   "k",
	}, nil
}

func (o *ObjectDeferred) Call(i itypes.Interpreter) (itypes.Object, error) {
	return o, nil
}

func (o *ObjectDeferred) Member(i itypes.Interpreter, obj itypes.Object, memberName string) (itypes.Object, error) {
	return o, nil
}
