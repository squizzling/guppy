package stream

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type methodStreamOp struct {
	itypes.Object

	op      string
	reverse bool
}

func (mso methodStreamOp) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (mso methodStreamOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := i.GetArg("self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
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
				return NewStreamBinaryOpStream(newStreamObject(), &ObjectStreamTernary{}, unpublish(selfStream), mso.op, unpublish(right)), nil
			case *primitive.ObjectInt:
				return NewStreamBinaryOpInt(newStreamObject(), &ObjectStreamTernary{}, unpublish(selfStream), mso.op, right.Value, mso.reverse), nil
			case *primitive.ObjectDouble:
				return NewStreamBinaryOpDouble(newStreamObject(), &ObjectStreamTernary{}, unpublish(selfStream), mso.op, right.Value, mso.reverse), nil
			default:
				return nil, fmt.Errorf("opCall[%s]: unknown type %T", mso.op, right)
			}
		}
	}
}
