package primitive

import (
	"fmt"
	"strconv"

	"guppy/pkg/interpreter/itypes"
)

type ObjectInt struct {
	itypes.Object

	Value int
}

type methodIntOp struct {
	itypes.Object

	op      string
	reverse string
}

type methodIntNeg struct {
	itypes.Object
}

func NewObjectInt(i int) *ObjectInt {
	return &ObjectInt{
		Object: itypes.NewObject(map[string]itypes.Object{
			"__add__":         methodIntOp{Object: itypes.NewObject(nil), op: "+", reverse: "__radd__"},
			"__mul__":         methodIntOp{Object: itypes.NewObject(nil), op: "*", reverse: "__rmul__"},
			"__sub__":         methodIntOp{Object: itypes.NewObject(nil), op: "-", reverse: "__rsub__"},
			"__truediv__":     methodIntOp{Object: itypes.NewObject(nil), op: "/", reverse: "__rtruediv__"},
			"__unary_minus__": methodIntNeg{Object: itypes.NewObject(nil)},

			"__lt__": methodIntOp{Object: itypes.NewObject(nil), op: "<"},
			"__gt__": methodIntOp{Object: itypes.NewObject(nil), op: ">"},
			"__le__": methodIntOp{Object: itypes.NewObject(nil), op: "<="},
			"__ge__": methodIntOp{Object: itypes.NewObject(nil), op: ">="},
			"__eq__": methodIntOp{Object: itypes.NewObject(nil), op: "=="},
			"__ne__": methodIntOp{Object: itypes.NewObject(nil), op: "!="},
		}),
		Value: i,
	}
}

func (oi *ObjectInt) Repr() string {
	return fmt.Sprintf("int(%d)", oi.Value)
}

func (oi *ObjectInt) String(i itypes.Interpreter) (string, error) {
	return strconv.Itoa(oi.Value), nil
}

func (mio methodIntOp) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (mio methodIntOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else if reverseOp, err := right.Member(i, right, mio.reverse); err == nil {
		// If it exists, we always use the reverse method, because it's more likely to be the intended behavior.
		// We explicitly don't expose reverse methods for primitives though.
		if reverseOpCall, ok := reverseOp.(itypes.FlowCall); ok {
			return reverseOpCall.Call(i)
		}
	}

	if self, err := itypes.ArgAs[*ObjectInt](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else {
		var rightVal int
		switch right := right.(type) {
		case *ObjectInt:
			rightVal = right.Value
		case *ObjectDouble:
			rightVal = int(right.Value)
		default:
			return nil, fmt.Errorf("methodIntOp: unknown type %T op %s", right, mio.op)
		}

		switch mio.op {
		case "+":
			return NewObjectInt(self.Value + rightVal), nil
		case "-":
			return NewObjectInt(self.Value - rightVal), nil
		case "/":
			return NewObjectInt(self.Value / rightVal), nil
		case "*":
			return NewObjectInt(self.Value * rightVal), nil
		case "<":
			return NewObjectBool(self.Value < rightVal), nil
		case ">":
			return NewObjectBool(self.Value > rightVal), nil
		case "<=":
			return NewObjectBool(self.Value <= rightVal), nil
		case ">=":
			return NewObjectBool(self.Value >= rightVal), nil
		case "==":
			return NewObjectBool(self.Value == rightVal), nil
		case "!=":
			return NewObjectBool(self.Value != rightVal), nil
		default:
			return nil, fmt.Errorf("methodIntOp: unknown op %s", mio.op)
		}
	}
}

var _ = itypes.FlowCall(methodIntOp{})

func (min methodIntNeg) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.UnaryParams, nil
}

func (min methodIntNeg) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[*ObjectInt](i, "self"); err != nil {
		return nil, err
	} else {
		return NewObjectInt(-self.Value), nil
	}
}

var _ = itypes.FlowCall(methodIntNeg{})
