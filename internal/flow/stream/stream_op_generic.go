package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

func opCall(
	i *interpreter.Interpreter,
	op string,
	newStream func(self Stream, right Stream) Stream,
	newScalarStream func(self Stream, right int) Stream,
) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		_ = self
		switch right := right.(type) {
		case Stream:
			return newStream(self, right), nil
		case *interpreter.ObjectInt:
			return newScalarStream(self, right.Value), nil
		default:
			return nil, fmt.Errorf("opCall[%s]: unknown type %T", op, right)
		}
	}
}
