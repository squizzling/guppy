package interpreter

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"guppy/internal/parser/ast"
)

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

func (i *Interpreter) pushScope() {
	i.Scope = &scope{
		isDefined:      make(map[string]bool),
		vars:           make(map[string]Object),
		deferredAssign: make(map[string]deferAssign),
		popChain:       i.Scope,
		lookupChain:    i.Scope,
	}
}

func (i *Interpreter) popScope() {
	i.Scope = i.Scope.popChain
}

func (s *scope) resolveDeferred(i *Interpreter) error {
	var pending []string
	for len(s.deferredAssign) > 0 {
		progress := false
		for key, da := range s.deferredAssign {
			maybeResolved, err := r(da.object.expr.Accept(i))
			if err != nil {
				return fmt.Errorf("deferred resolution failed for keys %s: %w", da.vars, err)
			}
			if od, ok := maybeResolved.(*ObjectDeferred); !ok {
				for idx, value := range maybeResolved.(*ObjectList).Items {
					s.vars[da.vars[idx]] = value
				}
				progress = true
				delete(s.deferredAssign, key)
			} else {
				pending = append(pending, od.desired...)
			}
		}
		if !progress {
			break
		}
	}

	if len(s.deferredAssign) > 0 {
		slices.Sort(pending)
		return fmt.Errorf("deferred assignment missing variables: %s", slices.Compact(pending))
	}

	for _, anon := range s.deferred {
		if o, err := anon.expr.Accept(i); err != nil {
			return fmt.Errorf("deferred anonymous resolution failed: %w", err)
		} else if od, ok := o.(*ObjectDeferred); ok {
			return fmt.Errorf("deferred anonymous resolution missing variables: %s", od.desired)
		}
	}
	s.deferredAssign = nil
	return nil
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

// GetArg is a hack to look up variables without deferred resolution.  It fails if it can't find
// something.  It's meant for arguments, which we know much exist.  I don't love the interface
// and it would be good to refactor at some point.
func (s *scope) GetArg(key string) (Object, error) {
	// TODO: Refactor.
	if val, ok := s.vars[key]; ok {
		return val, nil
	}

	if s.lookupChain == nil {
		return nil, fmt.Errorf("argument %s does not exist", key)
	}
	return s.lookupChain.GetArg(key)
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
