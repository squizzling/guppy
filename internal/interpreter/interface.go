package interpreter

import (
	"fmt"
)

func (i *Interpreter) doAnd(left Object, right Object) (Object, error) {
	if and, err := left.Member(i, left, "__binary_and__"); err != nil {
		return nil, err
	} else {
		i.pushScope()
		defer i.popScope()

		i.Scope.DeclareSet("self", left)
		i.Scope.DeclareSet("right", right)
		return i.doCall(and)
	}
}

type ParamData struct {
	Name    string
	Default Object
}

type FlowCall interface {
	Params(i *Interpreter) ([]ParamData, error)
	Call(i *Interpreter) (Object, error)
}

func (i *Interpreter) doParams(fo Object) ([]ParamData, error) {
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

func (i *Interpreter) doString(o Object) (string, error) {
	if s, ok := o.(FlowStringable); !ok {
		return "", fmt.Errorf("%T is not stringable", &s)
	} else {
		return s.String(i)
	}
}

type FlowIntable interface {
	Int(i *Interpreter) (int, error)
}

func (i *Interpreter) doInt(o Object) (int, error) {
	if s, ok := o.(FlowIntable); !ok {
		return 0, fmt.Errorf("%T is not intable", &s)
	} else {
		return s.Int(i)
	}
}

func ArgAsString(i *Interpreter, argName string) (string, error) {
	if objArg, err := i.Scope.Get(argName); err != nil {
		return "", err
	} else if strArg, err := i.doString(objArg); err != nil {
		return "", err
	} else {
		return strArg, nil
	}
}

func ArgAs[T any](i *Interpreter, name string) (T, error) {
	var zero T
	if v, err := i.Scope.Get(name); err != nil {
		return zero, err
	} else if o, ok := v.(T); !ok {
		return zero, fmt.Errorf("arg %s is %T not %T", name, o, &v)
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
