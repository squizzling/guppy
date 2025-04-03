package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodStreamOp struct {
	interpreter.Object

	op      string
	reverse bool
}

func (mso methodStreamOp) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (mso methodStreamOp) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := i.Scope.GetArg("self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		if mso.reverse {
			self, right = right, self
		}

		if selfStream, ok := self.(Stream); !ok {
			return nil, fmt.Errorf("arg self is %T not interpreter.Stream", selfStream)
		} else {
			switch right := right.(type) {
			case Stream:
				return NewStreamMathOpStream(newStreamObject(), unpublish(selfStream), mso.op, unpublish(right)), nil
			case *interpreter.ObjectInt:
				return NewStreamMathOpInt(newStreamObject(), unpublish(selfStream), mso.op, right.Value, mso.reverse), nil
			case *interpreter.ObjectDouble:
				return NewStreamMathOpDouble(newStreamObject(), unpublish(selfStream), mso.op, right.Value, mso.reverse), nil
			default:
				return nil, fmt.Errorf("opCall[%s]: unknown type %T", mso.op, right)
			}
		}
	}
}
