package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type streamMathScalar[T any] struct {
	interpreter.Object

	left  Stream
	op    string
	right T

	reverse bool
}

func newStreamMathScalar[T any](left Stream, op string, right T, reverse bool) Stream {
	return &streamMathScalar[T]{
		Object: newStreamObject(),

		left:    unpublish(left),
		op:      op,
		right:   right,
		reverse: reverse,
	}
}

func (s *streamMathScalar[T]) RenderStream() string {
	if s.reverse {
		return fmt.Sprintf("(%v %s %s)", s.right, s.op, s.left.RenderStream())
	} else {
		return fmt.Sprintf("(%s %s %v)", s.left.RenderStream(), s.op, s.right)
	}
}
