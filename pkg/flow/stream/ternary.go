package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
	"github.com/squizzling/guppy/pkg/parser/ast"
)

type ObjectStreamTernary struct{}

func (ost *ObjectStreamTernary) resolveStream(i itypes.Interpreter, e ast.Expression) (Stream, error) {
	o, err := e.Accept(i)
	if err != nil {
		return nil, err
	}
	switch o := o.(type) {
	case Stream:
		return o, nil
	case *primitive.ObjectInt:
		return NewStreamFuncConstInt(prototypeStreamDouble, o.Value, nil), nil
	case *primitive.ObjectNone:
		return NewStreamConstNone(prototypeStreamObject), nil
	default:
		return nil, fmt.Errorf("ObjectStreamTernary.resolveStream got %T expecting Stream", o)
	}
}

func (ost *ObjectStreamTernary) VisitExpressionTernary(i itypes.Interpreter, left ast.Expression, cond itypes.Object, right ast.Expression) (any, error) {
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

var _ itypes.FlowTernary = (*ObjectStreamTernary)(nil)
