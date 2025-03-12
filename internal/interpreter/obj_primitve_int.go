package interpreter

import (
	"fmt"
)

type ObjectInt struct {
	Object

	Value int
}

type methodIntOp struct {
	Object

	op      string
	reverse string
}

func NewObjectInt(i int) Object {
	return &ObjectInt{
		Object: NewObject(map[string]Object{
			"__add__":     methodIntOp{Object: NewObject(nil), op: "+", reverse: "__radd__"},
			"__mul__":     methodIntOp{Object: NewObject(nil), op: "*", reverse: "__rmul__"},
			"__sub__":     methodIntOp{Object: NewObject(nil), op: "-", reverse: "__rsub__"},
			"__truediv__": methodIntOp{Object: NewObject(nil), op: "/", reverse: "__rtruediv__"},
		}),
		Value: i,
	}
}

func (oi *ObjectInt) Repr() string {
	return fmt.Sprintf("int(%d)", oi.Value)
}

func (mio methodIntOp) Params(i *Interpreter) (*Params, error) {
	return BinaryParams, nil
}

func (mio methodIntOp) Call(i *Interpreter) (Object, error) {
	if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else if reverseOp, err := right.Member(i, right, mio.reverse); err == nil {
		// If it exists, we always use the reverse method, because it's more likely to be the intended behavior.
		// We explicitly don't expose reverse methods for primitives though.
		if reverseOpCall, ok := reverseOp.(FlowCall); ok {
			return reverseOpCall.Call(i)
		}
	}

	if self, err := ArgAs[*ObjectInt](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right := right.(type) {
		case *ObjectInt:
			switch mio.op {
			case "+":
				return NewObjectInt(self.Value + right.Value), nil
			case "-":
				return NewObjectInt(self.Value - right.Value), nil
			case "/":
				return NewObjectInt(self.Value / right.Value), nil
			case "*":
				return NewObjectInt(self.Value * right.Value), nil
			default:
				return nil, fmt.Errorf("methodIntOp: unknown op %s", mio.op)
			}
		default:
			return nil, fmt.Errorf("methodIntOp: unknown type %T", right)
		}
	}
}

var _ = FlowCall(methodIntOp{})
