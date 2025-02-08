package stream

import (
	"guppy/internal/interpreter"
)

type methodSum struct {
	interpreter.Object
}

func (ms methodSum) Args(i *interpreter.Interpreter) ([]interpreter.ArgData, error) {
	return argsAggregate(i)
}

func (ms methodSum) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return callAggregate(i, NewSum)
}

type sum struct {
	interpreter.Object

	source Stream
	by     []string
}

func NewSum(source Stream, by []string) Stream {
	return &sum{
		Object: newStreamObject(),
		source: unpublish(source),
		by:     by,
	}
}

func (s *sum) RenderStream() string {
	return renderAggregate(s.source, "sum", s.by)
}
