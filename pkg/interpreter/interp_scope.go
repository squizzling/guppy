package interpreter

import (
	"github.com/squizzling/guppy/pkg/interpreter/scope"
)

func (i *interpreter) pushScope() {
	i.Scope = i.Scope.Child()
}

func (i *interpreter) PushIntrinsicScope() {
	i.Scope = i.Scope.Lookup(i.Intrinsics)
}

func (i *interpreter) withScope(fn func() error) error {
	i.pushScope()
	defer i.PopScope()
	return fn()
}

func (i *interpreter) pushClosure(s *scope.Scope) {
	i.Scope = i.Scope.Lookup(s)
}

func (i *interpreter) PopScope() {
	i.Scope = i.Scope.Parent()
}
