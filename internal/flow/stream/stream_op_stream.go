package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type streamMathStream struct {
	interpreter.Object

	left  Stream
	op    string
	right Stream
}

func newStreamMathStream(left Stream, op string, right Stream) Stream {
	return &streamMathStream{
		Object: newStreamObject(),

		left:  unpublish(left),
		op:    op,
		right: unpublish(right),
	}
}

func (s streamMathStream) RenderStream() string {
	return fmt.Sprintf("(%s %s %s)", s.left.RenderStream(), s.op, s.right.RenderStream())
}
