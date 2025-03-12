package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type streamTransform struct {
	interpreter.Object

	source Stream
	fn     string
	over   string
}

func newStreamTransform(source Stream, fn string, over string) Stream {
	return &streamTransform{
		Object: newStreamObject(),

		source: unpublish(source),
		fn:     fn,
		over:   over,
	}
}

func (st *streamTransform) RenderStream() string {
	return fmt.Sprintf("%s.%s(over='%s')", st.source.RenderStream(), st.fn, st.over)
}
