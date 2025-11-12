package interpreter

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
)

func (i *interpreter) GetArg(argName string) (itypes.Object, error) {
	return i.Scope.GetArg(argName)
}

func (i *interpreter) Get(argName string) (itypes.Object, error) {
	return i.Scope.Get(argName)
}

func (i *interpreter) Set(name string, value itypes.Object) error {
	return i.Scope.Set(name, value)
}

func (i *interpreter) GetGlobal(argName string) (itypes.Object, error) {
	return i.Globals.GetArg(argName)
}

func (i *interpreter) SetGlobal(name string, value itypes.Object) error {
	return i.Globals.Set(name, value)
}

func (i *interpreter) doAnd(left itypes.Object, right itypes.Object) (itypes.Object, error) {
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
		return i.DoCall(and)
	}
}

type FlowCall interface {
	itypes.Object
	Params(i itypes.Interpreter) (*itypes.Params, error)
	Call(i itypes.Interpreter) (itypes.Object, error)
}

func (i *interpreter) DoParams(fo itypes.Object) (*itypes.Params, error) {
	if fc, ok := fo.(FlowCall); ok {
		return fc.Params(i)
	} else {
		// TODO: Test this
		return nil, fmt.Errorf("%T is not callable", fo)
	}
}

func (i *interpreter) DoCall(fo itypes.Object) (itypes.Object, error) {
	if fc, ok := fo.(FlowCall); ok {
		return fc.Call(i)
	} else {
		// TODO: Test this
		return nil, fmt.Errorf("%T is not callable", fo)
	}
}

type FlowStringable interface {
	itypes.Object
	String(i itypes.Interpreter) (string, error)
}

func (i *interpreter) DoString(o itypes.Object) (string, error) {
	if s, ok := o.(FlowStringable); !ok {
		return "", fmt.Errorf("%T is not stringable", o)
	} else {
		return s.String(i)
	}
}

func ArgAsString(i itypes.Interpreter, argName string) (string, error) {
	if objArg, err := i.GetArg(argName); err != nil {
		return "", err
	} else if strArg, err := i.DoString(objArg); err != nil {
		return "", err
	} else {
		return strArg, nil
	}
}

func ArgAsDouble(i itypes.Interpreter, argName string) (float64, error) {
	if objArg, err := i.GetArg(argName); err != nil {
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

func r(a any, err error) (itypes.Object, error) {
	if err != nil {
		return nil, err
	}
	return a.(itypes.Object), nil
}

func isTruthy(o itypes.Object) (bool, error) {
	switch o := o.(type) {
	case *ObjectBool:
		return o.Value, nil
	default:
		return false, fmt.Errorf("isTruthy condition is %T not *interpreter.ObjectBool", o)
	}
}
