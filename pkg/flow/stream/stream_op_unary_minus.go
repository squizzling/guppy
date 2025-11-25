package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type methodStreamUnaryMinus struct {
	itypes.Object
}

func (msum methodStreamUnaryMinus) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.UnaryParams, nil
}

func (msum methodStreamUnaryMinus) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamUnaryOpMinus(prototypeStreamDouble, unpublish(self)), nil
	}
}
