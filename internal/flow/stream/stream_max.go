package stream

import (
	"guppy/internal/interpreter"
)

type methodMax struct {
	interpreter.Object
}

func (mm methodMax) Args(i *interpreter.Interpreter) ([]interpreter.ArgData, error) {
	return argsAggregate(i)
}

func (mm methodMax) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return callAggregate(i, NewMax)
}

type max struct {
	interpreter.Object

	source Stream
	by     []string
}

func NewMax(source Stream, by []string) Stream {
	return &max{
		Object: newStreamObject(),
		source: unpublish(source),
		by:     by,
	}
}

func (m *max) RenderStream() string {
	return renderAggregate(m.source, "max", m.by)
}
