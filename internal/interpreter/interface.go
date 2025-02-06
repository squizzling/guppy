package interpreter

import (
	"fmt"

	"github.com/squizzling/types/pkg/result"
)

func (i *Interpreter) doAnd(left Object, right Object) result.Result[Object] {
	if resultAnd := left.Member(i, left, "__binary_and__"); !resultAnd.Ok() {
		return resultAnd
	} else {
		i.pushScope()
		defer i.popScope()

		i.Scope.DeclareSet("self", left)
		i.Scope.DeclareSet("right", right)
		return i.doCall(resultAnd.Value())
	}
}

type ArgData struct {
	Name    string
	Default Object
}

type FlowCall interface {
	Args(i *Interpreter) result.Result[[]ArgData]
	Call(i *Interpreter) result.Result[Object]
}

func (i *Interpreter) doArgs(fo Object) result.Result[[]ArgData] {
	if fc, ok := fo.(FlowCall); ok {
		return fc.Args(i)
	} else {
		// TODO: Test this
		return result.Err[[]ArgData](fmt.Errorf("%T is not callable", fo))
	}
}

func (i *Interpreter) doCall(fo Object) result.Result[Object] {
	if fc, ok := fo.(FlowCall); ok {
		return fc.Call(i)
	} else {
		// TODO: Test this
		return result.Err[Object](fmt.Errorf("%T is not callable", fo))
	}
}

type FlowStringable interface {
	String(i *Interpreter) result.Result[string]
}

func (i *Interpreter) doString(o Object) result.Result[string] {
	if s, ok := o.(FlowStringable); !ok {
		return result.Err[string](fmt.Errorf("%T is not stringable", &s))
	} else {
		return s.String(i)
	}
}

func ArgAsString(i *Interpreter, argName string) result.Result[string] {
	if resultObjArg := i.Scope.Get(argName); !resultObjArg.Ok() {
		return result.Err[string](resultObjArg.Err())
	} else if objArg := i.doString(resultObjArg.Value()); !objArg.Ok() {
		return result.Err[string](objArg.Err())
	} else {
		return result.Ok(objArg.Value())
	}
}

func ArgAs[T any](i *Interpreter, name string) result.Result[T] {
	if resultV := i.Scope.Get(name); !resultV.Ok() {
		return result.Err[T](resultV.Err())
	} else if v, ok := resultV.Value().(T); !ok {
		return result.Err[T](fmt.Errorf("arg %s is %T not %T", name, resultV.Value(), &v))
	} else {
		return result.Ok[T](v)
	}
}

func r(a any) result.Result[Object] {
	return a.(result.Result[Object])
}
