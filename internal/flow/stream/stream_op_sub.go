package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type opSub struct {
	interpreter.Object

	reverse bool
}

func (os opSub) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (os opSub) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return opCall(i, "sub", NewSub, NewSubScalar, os.reverse)
}

type sub struct {
	interpreter.Object

	left  Stream
	right Stream
}

func NewSub(left Stream, right Stream) Stream {
	return &sub{
		Object: newStreamObject(),
		left:   unpublish(left),
		right:  unpublish(right),
	}
}

func (s *sub) RenderStream() string {
	return fmt.Sprintf("(%s - %s)", s.left.RenderStream(), s.right.RenderStream())
}

type SubScalar struct {
	interpreter.Object

	left    Stream
	right   int
	reverse bool
}

func NewSubScalar(left Stream, right int, reverse bool) Stream {
	return &SubScalar{
		Object:  newStreamObject(),
		left:    unpublish(left),
		right:   right,
		reverse: reverse,
	}
}

func (tds *SubScalar) RenderStream() string {
	if tds.reverse {
		return fmt.Sprintf("(%d - %s)", tds.right, tds.left.RenderStream())
	} else {
		return fmt.Sprintf("(%s - %d)", tds.left.RenderStream(), tds.right)
	}
}
