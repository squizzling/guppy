package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type streamMathScalar struct {
	interpreter.Object

	left  Stream
	op    string
	right int

	reverse bool
}

func newStreamMathScalar(left Stream, op string, right int, reverse bool) Stream {
	return &streamMathScalar{
		Object: newStreamObject(),

		left:    unpublish(left),
		op:      op,
		right:   right,
		reverse: reverse,
	}
}

func (s *streamMathScalar) RenderStream() string {
	if s.reverse {
		return fmt.Sprintf("(%d %s %s)", s.right, s.op, s.left.RenderStream())
	} else {
		return fmt.Sprintf("(%s %s %d)", s.left.RenderStream(), s.op, s.right)
	}
}
