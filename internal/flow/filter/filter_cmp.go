package filter

import (
	"guppy/internal/interpreter"
)

type methodBinaryEqual struct {
	interpreter.Object

	invert bool
}

func (mbe methodBinaryEqual) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (mbe methodBinaryEqual) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if _, err := interpreter.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else if err := interpreter.ArgAsNone(i, "right"); err != nil {
		return nil, err
	} else {
		return interpreter.NewObjectBool(mbe.invert), nil
	}
}
