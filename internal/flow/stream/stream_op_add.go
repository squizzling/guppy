package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type opAdd struct {
	interpreter.Object
}

func (oa opAdd) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (oa opAdd) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return opCall(i, "add", NewAdd, NewAddScalar)
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

	left  Stream
	right int
}

func NewAddScalar(left Stream, right int) Stream {
	return &AddScalar{
		Object: newStreamObject(),
		left:   unpublish(left),
		right:  right,
	}
}

func (as *AddScalar) RenderStream() string {
	return fmt.Sprintf("(%s + %d)", as.left.RenderStream(), as.right)
}
