package scope

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"guppy/pkg/interpreter/deferred"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/parser/ast"
)

type deferAssign struct {
	vars   []string
	object *deferred.ObjectDeferred
}

type Scope struct {
	isDefined      map[string]bool
	vars           map[string]itypes.Object
	deferredAssign map[string]deferAssign
	deferred       []*deferred.ObjectDeferred
	popChain       *Scope // Used when popping
	lookupChain    *Scope // Used for lookup
}

func (s *Scope) Child() *Scope {
	// Implementation note: s may be nil if it's the root.  This is allowed.
	return &Scope{
		isDefined:      make(map[string]bool),
		vars:           make(map[string]itypes.Object),
		deferredAssign: make(map[string]deferAssign),
		popChain:       s,
		lookupChain:    s,
	}
}

func (s *Scope) Closure(lookup *Scope) *Scope {
	return &Scope{
		isDefined:      make(map[string]bool),
		vars:           make(map[string]itypes.Object),
		deferredAssign: make(map[string]deferAssign),
		popChain:       s,
		lookupChain:    lookup,
	}
}

func (s *Scope) Parent() *Scope {
	return s.popChain
}

func (s *Scope) ResolveDeferred(i itypes.Interpreter) error {
	var pending []string
	for len(s.deferredAssign) > 0 {
		progress := false
		for key, da := range s.deferredAssign {
			maybeResolved, err := da.object.Expr.Accept(i)
			if err != nil {
				return fmt.Errorf("deferred resolution failed for keys %s: %w", da.vars, err)
			}
			if od, ok := maybeResolved.(*deferred.ObjectDeferred); ok {
				pending = append(pending, od.Desired...)
			} else if obj, ok := maybeResolved.(itypes.Object); ok {
				s.vars[da.vars[0]] = obj
				progress = true
				delete(s.deferredAssign, key)
			} else {
				return fmt.Errorf("expecting ObjectDeferred or Object, found %T", maybeResolved)
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
		if o, err := anon.Expr.Accept(i); err != nil {
			return fmt.Errorf("deferred anonymous resolution failed: %w", err)
		} else if od, ok := o.(*deferred.ObjectDeferred); ok {
			return fmt.Errorf("deferred anonymous resolution missing variables: %s", od.Desired)
		}
	}
	s.deferredAssign = nil
	return nil
}

func (s *Scope) Set(key string, value itypes.Object) error {
	// Set is not allowed to look up the chain to find something, it can only set in the current context, and only if
	// there's nothing there already.
	if s.isDefined[key] {
		return errors.New("scope contains multiple bindings of " + key)
	}
	s.isDefined[key] = true
	s.vars[key] = value
	return nil
}

func (s *Scope) SetDefers(keys []string, d *deferred.ObjectDeferred) error {
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
func (s *Scope) GetArg(key string) (itypes.Object, error) {
	// TODO: Refactor.
	if val, ok := s.vars[key]; ok {
		return val, nil
	}

	if s.lookupChain == nil {
		return nil, fmt.Errorf("argument %s does not exist", key)
	}
	return s.lookupChain.GetArg(key)
}

func (s *Scope) Get(key string) (itypes.Object, error) {
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
		return deferred.NewObjectDeferred(ast.NewExpressionVariable(key), key), nil
	}
	return s.lookupChain.Get(key)
}

func (s *Scope) DeferAnonymous(d *deferred.ObjectDeferred) {
	s.deferred = append(s.deferred, d)
}
