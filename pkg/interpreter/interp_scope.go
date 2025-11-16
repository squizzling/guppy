package interpreter

import (
	"guppy/pkg/interpreter/scope"
)

func (i *interpreter) pushScope() {
	i.Scope = i.Scope.Child()
}

func (i *interpreter) withScope(fn func() error) error {
	i.pushScope()
	defer i.popScope()
	return fn()
}

func (i *interpreter) pushClosure(s *scope.Scope) {
	i.Scope = i.Scope.Closure(s)
}

func (i *interpreter) popScope() {
	i.Scope = i.Scope.Parent()
}
