package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type methodStreamOpTernary struct {
	interpreter.Object
}

func (msot methodStreamOpTernary) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.TernaryParams, nil
}

func (msot methodStreamOpTernary) resolveStream(i *interpreter.Interpreter, argName string) (Stream, error) {
	if in, err := i.Scope.GetArg(argName); err != nil {
		return nil, err
	} else {
		switch in := in.(type) {
		case Stream:
			return in, nil
		case *interpreter.ObjectInt:
			return NewStreamFuncConstInt(newStreamObject(), in.Value, nil), nil
		case *interpreter.ObjectDouble:
			return NewStreamFuncConstDouble(newStreamObject(), in.Value, nil), nil
		default:
			return nil, fmt.Errorf("%s is %T not Stream, *interpreter.ObjectInt, or *interpreter.ObjectDouble", argName, in)
		}
	}
}

func (msot methodStreamOpTernary) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if left, err := msot.resolveStream(i, "left"); err != nil {
		return nil, err
	} else if right, err := msot.resolveStream(i, "right"); err != nil {
		return nil, err
	} else {
		return NewStreamTernary(newStreamObject(), self, left, right), nil
	}
}
