package deferred

import (
	"fmt"
	"strings"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/parser/ast"
)

type ObjectDeferred struct {
	Expr    ast.Expression
	Desired []string
}

func NewObjectDeferred(expr ast.Expression, desired ...string) itypes.Object {
	return &ObjectDeferred{
		Expr:    expr,
		Desired: desired,
	}
}

func (o *ObjectDeferred) Repr() string {
	return fmt.Sprintf("deferred(%s)", strings.Join(o.Desired, ", "))
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
