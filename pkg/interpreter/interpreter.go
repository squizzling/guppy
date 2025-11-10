package interpreter

import (
	"bytes"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	"guppy/pkg/parser/ast"
)

type Interpreter struct {
	Globals *scope
	Scope   *scope

	debugDepth  int
	enableTrace bool
}

func NewInterpreter(enableTrace bool) *Interpreter {
	i := &Interpreter{
		enableTrace: enableTrace,
	}
	i.pushScope()
	i.Globals = i.Scope
	return i
}

type MethodStr interface {
	MethodStr() (string, error)
}

type Reprer interface {
	Repr() string
}

func Repr(o any) string {
	if repr, ok := o.(Reprer); ok {
		return repr.Repr()
	} else {
		return fmt.Sprintf("%#v", o)
	}
}

func (i *Interpreter) Debug(f string, args ...any) {
	if !i.enableTrace {
		return
	}
	fmt.Printf("%sdebug %s\n", strings.Repeat(" ", i.debugDepth), fmt.Sprintf(f, args...))
}

func (i *Interpreter) trace(a ...any) func(returnValue *any, err *error) {
	if !i.enableTrace {
		return func(returnValue *any, err *error) {}
	}

	rpc := []uintptr{0}
	s := runtime.Callers(2, rpc)
	n := "unknown"
	if s > 0 {
		frame, _ := runtime.CallersFrames(rpc).Next()
		n = frame.Function
		ns := strings.Split(n, ".")
		n = ns[len(ns)-1]
	}

	ss := ""
	if len(a) > 0 {
		ss = fmt.Sprintf(a[0].(string), a[1:]...)
		if ss != "" {
			ss = " - " + ss
		}
	}
	fmt.Printf("%senter %s%s\n", strings.Repeat(" ", i.debugDepth), n, ss)
	i.debugDepth++
	return func(returnValue *any, err *error) {
		i.debugDepth--
		if !bytes.Contains(debug.Stack(), []byte("panic")) {
			if *returnValue != nil {
				fmt.Printf("%sleave %s%s -> %s\n", strings.Repeat(" ", i.debugDepth), n, ss, Repr(*returnValue))
			} else if *err != nil {
				fmt.Printf("%sleave %s%s -> error(%s)\n", strings.Repeat(" ", i.debugDepth), n, ss, *err)
			} else {
				fmt.Printf("%sleave %s%s -> (nil, nil)\n", strings.Repeat(" ", i.debugDepth), n, ss)
			}
		}
	}
}

func (i *Interpreter) Execute(sp *ast.StatementProgram) error {
	_, err := sp.Accept(i)
	return err
}
