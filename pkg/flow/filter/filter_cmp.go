package filter

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type methodBinaryEqual struct {
	itypes.Object

	invert bool
}

func (mbe methodBinaryEqual) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (mbe methodBinaryEqual) Call(i itypes.Interpreter) (itypes.Object, error) {
	if _, err := itypes.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else if _, err := itypes.ArgAs[*interpreter.ObjectNone](i, "right"); err != nil {
		return nil, err
	} else {
		return primitive.NewObjectBool(mbe.invert), nil
	}
}

type methodBinaryIs struct {
	itypes.Object

	invert bool
}

func (mbi methodBinaryIs) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (mbi methodBinaryIs) Call(i itypes.Interpreter) (itypes.Object, error) {
	if _, err := itypes.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else if _, err := itypes.ArgAs[*interpreter.ObjectNone](i, "right"); err != nil {
		return nil, err
	} else {
		return primitive.NewObjectBool(mbi.invert), nil
	}
}
