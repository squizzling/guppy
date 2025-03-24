package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodStreamUnaryMinus struct {
	interpreter.Object
}

func (msum methodStreamUnaryMinus) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.UnaryParams, nil
}

func (msum methodStreamUnaryMinus) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamUnaryMinus(self), nil
	}
}

type streamUnaryMinus struct {
	interpreter.Object

	left Stream
}

func NewStreamUnaryMinus(left Stream) Stream {
	return &streamUnaryMinus{
		Object: newStreamObject(),

		left: unpublish(left),
	}
}

func (sum *streamUnaryMinus) RenderStream() string {
	return fmt.Sprintf("-(%s)", sum.left.RenderStream())
}
