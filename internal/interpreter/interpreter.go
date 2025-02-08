package interpreter

import (
	"bytes"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	"guppy/internal/parser/ast"
)

type Interpreter struct {
	Globals *scope
	Scope   *scope

	debugDepth int
}

func NewInterpreter() *Interpreter {
	i := &Interpreter{}
	i.pushScope()
	i.Globals = i.Scope

	return i
}

func (i *Interpreter) trace(a ...any) func() {
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
	//fmt.Printf("%senter %s%s\n", strings.Repeat(" ", i.debugDepth), n, ss)
	i.debugDepth++
	return func() {
		i.debugDepth--
		if !bytes.Contains(debug.Stack(), []byte("panic")) {
			//fmt.Printf("%sleave %s%s\n", strings.Repeat(" ", i.debugDepth), n, ss)
		}
	}
}

func (i *Interpreter) pushScope() {
	i.Scope = &scope{
		vars:        make(map[string]Object),
		popChain:    i.Scope,
		lookupChain: i.Scope,
	}
}

func (i *Interpreter) pushNewScope(s *scope) {
	i.Scope = &scope{
		vars:        make(map[string]Object),
		popChain:    i.Scope,
		lookupChain: s,
	}
}

func (i *Interpreter) popScope() {
	i.Scope = i.Scope.popChain
}

type scope struct {
	vars        map[string]Object
	popChain    *scope // Used when popping
	lookupChain *scope // Used for lookup
}

func (s *scope) Declare(key string) {
	if _, ok := s.vars[key]; !ok {
		s.vars[key] = nil
	}
}

func (s *scope) Set(key string, value Object) bool {
	// TODO: This whole model might be wrong for Signalflow.
	if _, ok := s.vars[key]; ok {
		s.vars[key] = value
		return true
	}
	if s.lookupChain == nil {
		return false
	}
	return s.lookupChain.Set(key, value)
}

func (s *scope) DeclareSet(key string, value Object) {
	s.Declare(key)
	s.Set(key, value)
}

func (s *scope) Get(key string) (Object, error) {
	if val, ok := s.vars[key]; ok {
		return val, nil
	}
	if s.lookupChain == nil {
		return nil, fmt.Errorf("unknown variable %s", key)
	}
	return s.lookupChain.Get(key)
}

func (i *Interpreter) Execute(sp ast.StatementProgram) {
	_, err := sp.Accept(i)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
