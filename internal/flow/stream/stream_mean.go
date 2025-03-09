package stream

import (
	"guppy/internal/interpreter"
)

type methodMean struct {
	interpreter.Object
}

func (mm methodMean) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return argsAggregateOver(i)
}

func (mm methodMean) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return callAggregateOver(i, NewMeanAggregate, NewMeanTransform)
}

type meanAggregate struct {
	interpreter.Object

	source Stream
	by     []string
}

func NewMeanAggregate(source Stream, by []string) Stream {
	return &meanAggregate{
		Object: newStreamObject(),
		source: unpublish(source),
		by:     by,
	}
}

func (ma *meanAggregate) RenderStream() string {
	return renderAggregate(ma.source, "mean", ma.by)
}

type meanTransform struct {
	interpreter.Object

	source Stream
	over   string
}

func NewMeanTransform(source Stream, over string) Stream {
	return &meanTransform{
		Object: newStreamObject(),
		source: unpublish(source),
		over:   over,
	}
}

func (mt *meanTransform) RenderStream() string {
	return renderTransform(mt.source, "mean", mt.over)
}
