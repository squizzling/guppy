package interpreter

import (
	"fmt"

	"guppy/pkg/parser/ast"
)

var TernaryParams = &Params{
	Params: []ParamDef{
		{Name: "self"},
		{Name: "left"},
		{Name: "right"},
	},
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

func (i *Interpreter) doAnd(left Object, right Object) (Object, error) {
	if and, err := left.Member(i, left, "__binary_and__"); err != nil {
		return nil, err
	} else {
		i.pushScope()
		defer i.popScope()

		if err := i.Scope.Set("self", left); err != nil {
			return nil, err
		}
		if err := i.Scope.Set("right", right); err != nil {
			return nil, err
		}
		return i.doCall(and)
	}
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

func (pd *Params) Dump(i *Interpreter) {
	i.debug("Params:")
	for _, param := range pd.Params {
		if param.Default != nil {
			s, err := i.DoString(param.Default)
			if err != nil {
				panic(err)
			}
			i.debug("- %s (%s)", param.Name, s)
		} else {
			i.debug("- %s required", param.Name)
		}
	}
	i.debug("StarParam: %s", pd.StarParam)
	i.debug("KWParams:")
	for _, param := range pd.KWParams {
		if param.Default != nil {
			s, err := i.DoString(param.Default)
			if err != nil {
				panic(err)
			}
			i.debug("- %s (%s)", param.Name, s)
		} else {
			i.debug("- %s required", param.Name)
		}
	}
	i.debug("KWParam: %s", pd.KWParam)

}

type FlowCall interface {
	Object
	Params(i *Interpreter) (*Params, error)
	Call(i *Interpreter) (Object, error)
}

func (i *Interpreter) doParams(fo Object) (*Params, error) {
	if fc, ok := fo.(FlowCall); ok {
		return fc.Params(i)
	} else {
		// TODO: Test this
		return nil, fmt.Errorf("%T is not callable", fo)
	}
}

func (i *Interpreter) doCall(fo Object) (Object, error) {
	if fc, ok := fo.(FlowCall); ok {
		return fc.Call(i)
	} else {
		// TODO: Test this
		return nil, fmt.Errorf("%T is not callable", fo)
	}
}

type FlowStringable interface {
	String(i *Interpreter) (string, error)
}

func (i *Interpreter) DoString(o Object) (string, error) {
	if s, ok := o.(FlowStringable); !ok {
		return "", fmt.Errorf("%T is not stringable", o)
	} else {
		return s.String(i)
	}
}

func ArgAsString(i *Interpreter, argName string) (string, error) {
	if objArg, err := i.Scope.GetArg(argName); err != nil {
		return "", err
	} else if strArg, err := i.DoString(objArg); err != nil {
		return "", err
	} else {
		return strArg, nil
	}
}

func ArgAsBool(i *Interpreter, argName string) (bool, error) {
	if objArg, err := i.Scope.GetArg(argName); err != nil {
		return false, err
	} else if boolArg, ok := objArg.(*ObjectBool); !ok {
		return false, fmt.Errorf("%T is not *interpreter.ObjectBool", objArg)
	} else {
		return boolArg.Value, nil
	}
}

func ArgAsNone(i *Interpreter, argName string) error {
	if objArg, err := i.Scope.GetArg(argName); err != nil {
		return err
	} else if _, ok := objArg.(*ObjectNone); !ok {
		return fmt.Errorf("%T is not *interpreter.ObjectNone", objArg)
	} else {
		return nil
	}
}

func ArgAsDouble(i *Interpreter, argName string) (float64, error) {
	if objArg, err := i.Scope.GetArg(argName); err != nil {
		return 0, err
	} else if doubleArg, ok := objArg.(*ObjectDouble); !ok {
		if intArg, ok := objArg.(*ObjectInt); !ok {
			return 0, fmt.Errorf("%T is not *interpreter.ObjectDouble or *interpreter.ObjectInt", objArg)
		} else {
			return float64(intArg.Value), nil
		}
	} else {
		return doubleArg.Value, nil
	}
}

func ArgAsLong(i *Interpreter, argName string) (int, error) {
	if objArg, err := i.Scope.GetArg(argName); err != nil {
		return 0, err
	} else if intArg, ok := objArg.(*ObjectInt); !ok {
		return 0, fmt.Errorf("%T is not *interpreter.ObjectInt", objArg)
	} else {
		return intArg.Value, nil
	}
}

func ArgAsOptLong(i *Interpreter, argName string) (*int, error) {
	if objArg, err := i.Scope.GetArg(argName); err != nil {
		return nil, err
	} else if intArg, ok := objArg.(*ObjectInt); ok {
		return &intArg.Value, nil
	} else if _, ok := objArg.(*ObjectNone); ok {
		return nil, nil
	} else {
		return nil, fmt.Errorf("arg %s is %T not *interpreter.ObjectInt or *interpreter.ObjectNone", argName, objArg)
	}
}

func ArgAs[T any](i *Interpreter, name string) (T, error) {
	var zero T
	if v, err := i.Scope.GetArg(name); err != nil {
		return zero, err
	} else if o, ok := v.(T); !ok {
		return zero, fmt.Errorf("arg %s is %T not %T", name, v, zero)
	} else {
		return o, nil
	}
}

func r(a any, err error) (Object, error) {
	if err != nil {
		return nil, err
	}
	return a.(Object), nil
}

func isTruthy(o Object) (bool, error) {
	switch o := o.(type) {
	case *ObjectBool:
		return o.Value, nil
	default:
		return false, fmt.Errorf("isTruthy condition is %T not *interpreter.ObjectBool", o)
	}
}

type FlowTernary interface {
	VisitExpressionTernary(i *Interpreter, left ast.Expression, cond Object, right ast.Expression) (any, error)
}
