package stream

import (
	"guppy/pkg/interpreter"
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
		return NewStreamUnaryOpMinus(newStreamObject(), unpublish(self)), nil
	}
}
