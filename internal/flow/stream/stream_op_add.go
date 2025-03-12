package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type opAdd struct {
	interpreter.Object

	reverse bool
}

func (oa opAdd) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (oa opAdd) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return opCall(i, "add", NewAdd, NewAddScalar, oa.reverse)
}

type add struct {
	interpreter.Object

	left  Stream
	right Stream
}

func NewAdd(left Stream, right Stream) Stream {
	return &add{
		Object: newStreamObject(),
		left:   unpublish(left),
		right:  unpublish(right),
	}
}

func (a *add) RenderStream() string {
	return fmt.Sprintf("(%s + %s)", a.left.RenderStream(), a.right.RenderStream())
}

type AddScalar struct {
	interpreter.Object

	left    Stream
	right   int
	reverse bool
}

func NewAddScalar(left Stream, right int, reverse bool) Stream {
	return &AddScalar{
		Object:  newStreamObject(),
		left:    unpublish(left),
		right:   right,
		reverse: reverse,
	}
}

func (as *AddScalar) RenderStream() string {
	if as.reverse {
		return fmt.Sprintf("(%d + %s)", as.right, as.left.RenderStream())
	} else {
		return fmt.Sprintf("(%s + %d)", as.left.RenderStream(), as.right)
	}
}
