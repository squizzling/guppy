package interpreter

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"

	"guppy/internal/parser/ast"
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
	// TODO: Technically this isn't necessary, but it protects globals from
	//       modification in case the caller doesn't push their own scope.
	i.pushScope()

	return i
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

func (i *Interpreter) pushScope() {
	i.Scope = &scope{
		isDefined:      make(map[string]bool),
		vars:           make(map[string]Object),
		deferredAssign: make(map[string]deferAssign),
		popChain:       i.Scope,
		lookupChain:    i.Scope,
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

func (i *Interpreter) resolveDeferred() error {
	for len(i.Scope.deferredAssign) > 0 {
		progress := false
		for key, da := range i.Scope.deferredAssign {
			maybeResolved, err := r(da.object.expr.Accept(i))
			if err != nil {
				return fmt.Errorf("deferred resolution failed for keys %s: %w", da.vars, err)
			}
			if _, ok := maybeResolved.(*ObjectDeferred); !ok {
				for idx, value := range maybeResolved.(*ObjectList).Items {
					i.Scope.vars[da.vars[idx]] = value
				}
				progress = true
				delete(i.Scope.deferredAssign, key)
			}
		}
		if !progress {
			break
		}
	}

	if len(i.Scope.deferredAssign) > 0 {
		return fmt.Errorf("%d remaining to resolve", len(i.Scope.deferredAssign))
	}

	for _, anon := range i.Scope.deferred {
		if o, err := anon.expr.Accept(i); err != nil {
			return fmt.Errorf("deferred anonymous resolution failed: %w", err)
		} else if od, ok := o.(*ObjectDeferred); ok {
			return fmt.Errorf("deferred anonymous resolution missing variables: %s", od.desired)
		}
	}
	i.Scope.deferredAssign = nil
	return nil
}

type deferAssign struct {
	vars   []string
	object *ObjectDeferred
}

type scope struct {
	isDefined      map[string]bool
	vars           map[string]Object
	deferredAssign map[string]deferAssign
	deferred       []*ObjectDeferred
	popChain       *scope // Used when popping
	lookupChain    *scope // Used for lookup
}

func (s *scope) Set(key string, value Object) error {
	// Set is not allowed to look up the chain to find something, it can only set in the current context, and only if
	// there's nothing there already.
	if s.isDefined[key] {
		return errors.New("scope contains multiple bindings of " + key)
	}
	s.isDefined[key] = true
	s.vars[key] = value
	return nil
}

func (s *scope) SetDefers(keys []string, d *ObjectDeferred) error {
	for _, key := range keys {
		if s.isDefined[key] {
			return errors.New("scope contains multiple bindings of " + key)
		}
		s.isDefined[key] = true
	}
	hashKey := strings.Join(keys, "|")
	s.deferredAssign[hashKey] = deferAssign{
		vars:   keys,
		object: d,
	}
	return nil
}

func (s *scope) Get(key string) (Object, error) {
	// Get is allowed to look up the chain to find something
	if val, ok := s.vars[key]; ok {
		return val, nil
	}

	// deferredAssign isn't efficient to search, so only do it if we know the key is present
	if s.isDefined[key] {
		for _, da := range s.deferredAssign {
			if slices.Contains(da.vars, key) {
				return da.object, nil
			}
		}
	}

	if s.lookupChain == nil {
		return NewObjectDeferred(ast.NewExpressionVariable(key), key), nil
	}
	return s.lookupChain.Get(key)
}

func (s *scope) DeferAnonymous(d *ObjectDeferred) {
	s.deferred = append(s.deferred, d)
}

func (i *Interpreter) Execute(sp *ast.StatementProgram) error {
	_, err := sp.Accept(i)
	return err
}
