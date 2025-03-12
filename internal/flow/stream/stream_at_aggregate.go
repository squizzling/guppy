package stream

import (
	"fmt"
	"strings"

	"guppy/internal/interpreter"
)

type streamAggregate struct {
	interpreter.Object

	source Stream
	fn     string
	by     []string
}

func newStreamAggregate(source Stream, fn string, by []string) Stream {
	return &streamAggregate{
		Object: newStreamObject(),

		source: unpublish(source),
		fn:     fn,
		by:     by,
	}
}

func (sa *streamAggregate) RenderStream() string {
	var bys []string
	if sa.by == nil {
		return fmt.Sprintf("%s.%s()", sa.source.RenderStream(), sa.fn)
	}
	for _, by := range sa.by {
		bys = append(bys, "'"+by+"'")
	}
	return fmt.Sprintf("%s.%s(by=[%s])", sa.source.RenderStream(), sa.fn, strings.Join(bys, ", "))
}
