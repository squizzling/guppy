package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type opMul struct {
	interpreter.Object
}

func (om opMul) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (om opMul) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return opCall(i, "mul", NewMul, NewMulScalar)
}

type Mul struct {
	interpreter.Object

	left  Stream
	right Stream
}

func NewMul(left Stream, right Stream) Stream {
	return &Mul{
		Object: newStreamObject(),
		left:   unpublish(left),
		right:  unpublish(right),
	}
}

func (m *Mul) RenderStream() string {
	return fmt.Sprintf("(%s * %s)", m.left.RenderStream(), m.right.RenderStream())
}

type MulScalar struct {
	interpreter.Object

	left  Stream
	right int
}

func NewMulScalar(left Stream, right int) Stream {
	return &MulScalar{
		Object: newStreamObject(),
		left:   unpublish(left),
		right:  right,
	}
}

func (ms *MulScalar) RenderStream() string {
	return fmt.Sprintf("(%s * %d)", ms.left.RenderStream(), ms.right)
}
