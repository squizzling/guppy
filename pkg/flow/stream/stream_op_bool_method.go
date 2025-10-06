package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type methodStreamOpBool struct {
	interpreter.Object

	op      string
	reverse bool
}

func (msob methodStreamOpBool) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (msob methodStreamOpBool) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := i.Scope.GetArg("self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		if msob.reverse {
			self, right = right, self
		}

		if selfStream, ok := self.(Stream); !ok {
			return nil, fmt.Errorf("arg self is %T not interpreter.Stream", selfStream)
		} else {
			switch right := right.(type) {
			case Stream:
				return NewStreamBinaryOpStream(newStreamBoolObject(), &ObjectStreamTernary{}, unpublish(selfStream), msob.op, unpublish(right)), nil
			case *interpreter.ObjectInt:
				return NewStreamBinaryOpInt(newStreamBoolObject(), &ObjectStreamTernary{}, unpublish(selfStream), msob.op, right.Value, msob.reverse), nil
			case *interpreter.ObjectDouble:
				return NewStreamBinaryOpDouble(newStreamBoolObject(), &ObjectStreamTernary{}, unpublish(selfStream), msob.op, right.Value, msob.reverse), nil
			default:
				return nil, fmt.Errorf("opBinaryCall[%s]: unknown type %T", msob.op, right)
			}
		}
	}
}
