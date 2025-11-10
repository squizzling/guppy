package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/parser/ast"
)

type ObjectStreamTernary struct{}

func (ost *ObjectStreamTernary) resolveStream(i *interpreter.Interpreter, e ast.Expression) (Stream, error) {
	o, err := e.Accept(i)
	if err != nil {
		return nil, err
	}
	switch o := o.(type) {
	case Stream:
		return o, nil
	case *interpreter.ObjectInt:
		return NewStreamFuncConstInt(newStreamObject(), o.Value, nil), nil
	default:
		return nil, fmt.Errorf("StreamIsNone.resolveStream got %T expecting Stream", o)
	}
}

func (ost *ObjectStreamTernary) VisitExpressionTernary(i *interpreter.Interpreter, left ast.Expression, cond itypes.Object, right ast.Expression) (any, error) {
	if leftStream, err := ost.resolveStream(i, left); err != nil {
		return nil, err
	} else if condStream, ok := cond.(Stream); !ok {
		return nil, fmt.Errorf("ObjectStreamTernary.VisitExpressionTernary cond is %T not Stream", cond)
	} else if rightStream, err := ost.resolveStream(i, right); err != nil {
		return nil, err
	} else {
		return NewStreamTernary(newStreamObject(), leftStream, condStream, rightStream), nil
	}
}

var _ interpreter.FlowTernary = (*ObjectStreamTernary)(nil)
