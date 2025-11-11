package interpreter

import (
	"fmt"
	"strconv"

	"guppy/pkg/interpreter/itypes"
)

type ObjectDouble struct {
	itypes.Object

	Value float64
}

type methodDoubleOp struct {
	itypes.Object

	op      string
	reverse string
}

type methodDoubleNeg struct {
	itypes.Object
}

func NewObjectDouble(f float64) itypes.Object {
	return &ObjectDouble{
		Object: itypes.NewObject(map[string]itypes.Object{
			"__add__":         methodDoubleOp{Object: itypes.NewObject(nil), op: "+", reverse: "__radd__"},
			"__mul__":         methodDoubleOp{Object: itypes.NewObject(nil), op: "*", reverse: "__rmul__"},
			"__sub__":         methodDoubleOp{Object: itypes.NewObject(nil), op: "-", reverse: "__rsub__"},
			"__truediv__":     methodDoubleOp{Object: itypes.NewObject(nil), op: "/", reverse: "__rtruediv__"},
			"__unary_minus__": methodDoubleNeg{Object: itypes.NewObject(nil)},

			"__lt__": methodDoubleOp{Object: itypes.NewObject(nil), op: "<"},
			"__gt__": methodDoubleOp{Object: itypes.NewObject(nil), op: ">"},
			"__le__": methodDoubleOp{Object: itypes.NewObject(nil), op: "<="},
			"__ge__": methodDoubleOp{Object: itypes.NewObject(nil), op: ">="},
		}),
		Value: f,
	}
}

func (od *ObjectDouble) Repr() string {
	return fmt.Sprintf("double(%f)", od.Value)
}

func (od *ObjectDouble) String(i itypes.Interpreter) (string, error) {
	return strconv.FormatFloat(od.Value, 'f', 6, 64), nil
}

func (mdo methodDoubleOp) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (mdo methodDoubleOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else if reverseOp, err := right.Member(i, right, mdo.reverse); err == nil {
		// If it exists, we always use the reverse method, because it's more likely to be the intended behavior.
		// We explicitly don't expose reverse methods for primitives though.
		if reverseOpCall, ok := reverseOp.(FlowCall); ok {
			return reverseOpCall.Call(i)
		}
	}

	if self, err := ArgAs[*ObjectDouble](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else {
		var rightVal float64
		switch right := right.(type) {
		case *ObjectInt:
			rightVal = float64(right.Value)
		case *ObjectDouble:
			rightVal = right.Value
		default:
			return nil, fmt.Errorf("methodDoubleOp: unknown type %T", right)
		}

		switch mdo.op {
		case "+":
			return NewObjectDouble(self.Value + rightVal), nil
		case "-":
			return NewObjectDouble(self.Value - rightVal), nil
		case "/":
			return NewObjectDouble(self.Value / rightVal), nil
		case "*":
			return NewObjectDouble(self.Value * rightVal), nil
		case "<":
			return NewObjectBool(self.Value < rightVal), nil
		case ">":
			return NewObjectBool(self.Value > rightVal), nil
		case "<=":
			return NewObjectBool(self.Value <= rightVal), nil
		case ">=":
			return NewObjectBool(self.Value >= rightVal), nil
		default:
			return nil, fmt.Errorf("methodDoubleOp: unknown op %s", mdo.op)
		}
	}
}

var _ = FlowCall(methodDoubleOp{})

func (mdn methodDoubleNeg) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.UnaryParams, nil
}

func (mdn methodDoubleNeg) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := ArgAs[*ObjectDouble](i, "self"); err != nil {
		return nil, err
	} else {
		return NewObjectDouble(-self.Value), nil
	}
}

var _ = FlowCall(methodDoubleNeg{})
