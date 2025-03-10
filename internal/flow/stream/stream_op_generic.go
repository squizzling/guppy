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
	reverse bool,
) (interpreter.Object, error) {
	if self, err := i.Scope.GetArg("self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		if reverse {
			self, right = right, self
		}

		if selfStream, ok := self.(Stream); !ok {
			return nil, fmt.Errorf("arg self is %T not interpreter.Stream", selfStream)
		} else {
			switch right := right.(type) {
			case Stream:
				return newStream(selfStream, right), nil
			case *interpreter.ObjectInt:
				return newScalarStream(selfStream, right.Value), nil
			default:
				return nil, fmt.Errorf("opCall[%s]: unknown type %T", op, right)
			}
		}
	}
}
