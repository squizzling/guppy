package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type methodStreamIs struct {
	itypes.Object

	invert  bool
	reverse bool
}

func (msi methodStreamIs) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (msi methodStreamIs) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := i.GetArg("self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else {
		if msi.reverse {
			self, right = right, self
		}

		if selfStream, ok := self.(Stream); !ok {
			return nil, fmt.Errorf("arg self is %T not Stream", self)
		} else {
			return NewStreamIsNone(newStreamObject(), &ObjectStreamTernary{}, selfStream, msi.invert), nil
		}
	}
}

func (sin *StreamIsNone) resolveStream(o any) (Stream, error) {
	switch o := o.(type) {
	case Stream:
		return o, nil
	case *primitive.ObjectInt:
		return NewStreamFuncConstInt(newStreamObject(), o.Value, nil), nil
	default:
		return nil, fmt.Errorf("StreamIsNone.resolveStream got %T expecting Stream", o)
	}
}
