package stream

import (
	"guppy/internal/interpreter"
)

type methodSum struct {
	interpreter.Object
}

func (ms methodSum) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return argsAggregateOver(i)
}

func (ms methodSum) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return callAggregateOver(i, NewSumAggregate, NewSumTransform)
}

type sumAggregate struct {
	interpreter.Object

	source Stream
	by     []string
}

func NewSumAggregate(source Stream, by []string) Stream {
	return &sumAggregate{
		Object: newStreamObject(),
		source: unpublish(source),
		by:     by,
	}
}

func (sa *sumAggregate) RenderStream() string {
	return renderAggregate(sa.source, "sum", sa.by)
}

type sumTransform struct {
	interpreter.Object

	source Stream
	over   string
}

func NewSumTransform(source Stream, over string) Stream {
	return &sumTransform{
		Object: newStreamObject(),
		source: unpublish(source),
		over:   over,
	}
}

func (st *sumTransform) RenderStream() string {
	return renderTransform(st.source, "sum", st.over)
}
