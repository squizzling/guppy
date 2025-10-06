package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/parser/ast"
)

type methodStreamIs struct {
	interpreter.Object

	invert  bool
	reverse bool
}

func (msi methodStreamIs) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (msi methodStreamIs) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := i.Scope.GetArg("self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		if msi.reverse {
			self, right = right, self
		}

		if selfStream, ok := self.(Stream); !ok {
			return nil, fmt.Errorf("arg self is %T not Stream", self)
		} else {
			return NewStreamIsNone(newStreamObject(), selfStream, msi.invert), nil
		}
	}
}

func (sin *StreamIsNone) resolveStream(o any) (Stream, error) {
	switch o := o.(type) {
	case Stream:
		return o, nil
	case *interpreter.ObjectInt:
		return NewStreamFuncConstInt(newStreamObject(), o.Value, nil), nil
	default:
		return nil, fmt.Errorf("StreamIsNone.resolveStream got %T expecting Stream", o)
	}
}

func (sin *StreamIsNone) VisitExpressionTernary(i *interpreter.Interpreter, left ast.Expression, right ast.Expression) (any, error) {
	if left, err := left.Accept(i); err != nil {
		return nil, err
	} else if right, err := right.Accept(i); err != nil {
		return nil, err
	} else if leftStream, err := sin.resolveStream(left); err != nil {
		return nil, err
	} else if rightStream, err := sin.resolveStream(right); err != nil {
		return nil, err
	} else {
		return NewStreamTernary(newStreamObject(), sin, leftStream, rightStream), nil
	}
}
