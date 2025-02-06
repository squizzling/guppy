package filter

import (
	"fmt"

	"github.com/squizzling/types/pkg/result"

	"guppy/internal/interpreter"
)

type methodBinaryAnd struct {
	interpreter.Object
}

func (mba methodBinaryAnd) Args(i *interpreter.Interpreter) result.Result[[]interpreter.ArgData] {
	return result.Ok([]interpreter.ArgData{
		{Name: "self"},
		{Name: "right"},
	})
}

func (mba methodBinaryAnd) Call(i *interpreter.Interpreter) result.Result[interpreter.Object] {
	if resultSelf := interpreter.ArgAs[Filter](i, "self"); !resultSelf.Ok() {
		return result.Err[interpreter.Object](resultSelf.Err())
	} else if resultRight := interpreter.ArgAs[Filter](i, "right"); !resultRight.Ok() {
		return result.Err[interpreter.Object](resultRight.Err())
	} else {
		return result.Ok[interpreter.Object](NewAnd(resultSelf.Value(), resultRight.Value()))
	}
}

type and struct {
	interpreter.Object

	left  Filter
	right Filter
}

func NewAnd(left Filter, right Filter) Filter {
	return &and{
		Object: newFilterObject(),
		left:   left,
		right:  right,
	}
}

func (a *and) RenderFilter() string {
	return fmt.Sprintf("(%s and %s)", a.left.RenderFilter(), a.right.RenderFilter())
}
