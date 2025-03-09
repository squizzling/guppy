package stream

import (
	"guppy/internal/interpreter"
)

type methodMax struct {
	interpreter.Object
}

func (mm methodMax) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return argsAggregateOver(i)
}

func (mm methodMax) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return callAggregateOver(i, NewMaxAggregate, NewMaxTransform)
}

type maxAggregate struct {
	interpreter.Object

	source Stream
	by     []string
}

func NewMaxAggregate(source Stream, by []string) Stream {
	return &maxAggregate{
		Object: newStreamObject(),
		source: unpublish(source),
		by:     by,
	}
}

func (ma *maxAggregate) RenderStream() string {
	return renderAggregate(ma.source, "max", ma.by)
}

type maxTransform struct {
	interpreter.Object

	source Stream
	over   string
}

func NewMaxTransform(source Stream, over string) Stream {
	return &maxTransform{
		Object: newStreamObject(),
		source: unpublish(source),
		over:   over,
	}
}

func (mt *maxTransform) RenderStream() string {
	return renderTransform(mt.source, "max", mt.over)
}
