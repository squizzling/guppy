package itypes

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/parser/ast"
)

type Interpreter interface {
	ast.VisitorStatement
	ast.VisitorExpression
	Execute(sp *ast.StatementProgram) error

	SetGlobal(name string, value Object) error
	GetGlobal(argName string) (Object, error)
	SetIntrinsic(name string, value Object) error
	Set(name string, value Object) error
	Get(name string) (Object, error)
	GetArg(name string) (Object, error)
	Debug(f string, args ...any)
	DoString(o Object) (string, error)
	DoParams(o Object) (*Params, error)
	DoCall(o Object) (Object, error)

	PushIntrinsicScope()
	PopScope()
}

type Object interface {
	// We include the root Object so the default behavior knows its object type, instead of just *flowObject
	Member(i Interpreter, obj Object, memberName string) (Object, error)
	Repr() string
}

type ParamDef struct {
	Name    string
	Default Object
}

// Params defines a set of parameters to a call.  By definition the following is the order:
//   - Params: these are a list of named parameters, they are filled with positional arguments, keyword arguments,
//     and defaults
//   - StarParam: this is a single parameter, and represents a `*param`.  If this is non-empty, then it will be passed an
//     empty tuple if nothing is present.
//   - KWParams: These are a list of named parameters, they are only filled with keyword arguments, and defaults.
//     Notably they will never be populated with positional arguments, even if a *starParam is provided.
//   - KWParam: This is a single parameter, and represents a `**param`.  If this is non-empty, then it will be passed an
//     empty dict if nothing is present.
//
// Examples:
//
// def func(params1, params2, *starParam, kwParams1, kwParams2, **kwParam): pass
//
// If invoked as either:
// func(p1, p2, p3, p4, kwParams1=p5, kwParams2=p6, kwParams3=p7, kwParams4=p8)
// func(p1, p2, *(p3, p4), kwParams1=p5, kwParams2=p6, kwParams3=p7, kwParams4=p8)
//
// These are both equivalent, and will fill:
// params1 = p1, params2 = p2, starParam = (p4, p5), kwParams1=p6, kwParams2=p7, kwParam={kwParam3: p7, kwParams4: p8}.
//
// If invoked as:
// func(p1, p2, *(p3, p4), p5, p6)
//
// This will be an error, because there is a `KWParams` after a `StarParam`.
//
// Notes:
//   - We currently do not parse flow parameters in to this structure, as we don't handle flow functions, only
//     FFI functions.
//   - This comment is formatted weird because gofmt is stupid.
type Params struct {
	Params    []ParamDef
	StarParam string
	KWParams  []ParamDef // KWParams can only be passed via keyword
	KWParam   string
}

func (pd *Params) Dump(i Interpreter) {
	i.Debug("Params:")
	for _, param := range pd.Params {
		if param.Default != nil {
			s, err := i.DoString(param.Default)
			if err != nil {
				s = err.Error()
			}
			i.Debug("- %s (%s)", param.Name, s)
		} else {
			i.Debug("- %s required", param.Name)
		}
	}
	i.Debug("StarParam: %s", pd.StarParam)
	i.Debug("KWParams:")
	for _, param := range pd.KWParams {
		if param.Default != nil {
			s, err := i.DoString(param.Default)
			if err != nil {
				s = err.Error()
			}
			i.Debug("- %s (%s)", param.Name, s)
		} else {
			i.Debug("- %s required", param.Name)
		}
	}
	i.Debug("KWParam: %s", pd.KWParam)

}

var BinaryParams = &Params{
	Params: []ParamDef{
		{Name: "self"},
		{Name: "right"},
	},
}

var UnaryParams = &Params{
	Params: []ParamDef{
		{Name: "self"},
	},
}

type FlowTernary interface {
	VisitExpressionTernary(i Interpreter, left ast.Expression, cond Object, right ast.Expression) (any, error)
}

func ArgAs[T any](i Interpreter, name string) (T, error) {
	var zero T
	if v, err := i.GetArg(name); err != nil {
		return zero, err
	} else if o, ok := v.(T); !ok {
		return zero, fmt.Errorf("arg %s is %T not %T", name, v, zero)
	} else {
		return o, nil
	}
}

type FlowCall interface {
	Object
	Params(i Interpreter) (*Params, error)
	Call(i Interpreter) (Object, error)
}

func Repr(o any) string {
	if repr, ok := o.(Object); ok {
		return repr.Repr()
	} else {
		return fmt.Sprintf("%#v", o)
	}
}
