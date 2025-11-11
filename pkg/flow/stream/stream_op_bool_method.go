package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type methodStreamOpBool struct {
	itypes.Object

	op      string
	reverse bool
}

func (msob methodStreamOpBool) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (msob methodStreamOpBool) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := i.GetArg("self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
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
