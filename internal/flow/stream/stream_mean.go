package stream

import (
	"guppy/internal/interpreter"
)

type methodMean struct {
	interpreter.Object
}

func (mm methodMean) Args(i *interpreter.Interpreter) ([]interpreter.ArgData, error) {
	return argsAggregate(i)
}

func (mm methodMean) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return callAggregate(i, NewMean)
}

type mean struct {
	interpreter.Object

	source Stream
	by     []string
}

func NewMean(source Stream, by []string) Stream {
	return &mean{
		Object: newStreamObject(),
		source: unpublish(source),
		by:     by,
	}
}

func (m *mean) RenderStream() string {
	return renderAggregate(m.source, "mean", m.by)
}
