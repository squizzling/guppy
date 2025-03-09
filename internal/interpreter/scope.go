package interpreter

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"guppy/internal/parser/ast"
)

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
