package interpreter

import (
	"github.com/squizzling/types/pkg/result"
)

type ObjectString struct {
	Object

	s string
}

func NewObjectString(s string) Object {
	return &ObjectString{
		s: s,
	}
}

func (os *ObjectString) String(i *Interpreter) result.Result[string] {
	return result.Ok(os.s)
}

var _ = FlowStringable(&ObjectString{})
