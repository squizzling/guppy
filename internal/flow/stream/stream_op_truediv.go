package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type opTrueDiv struct {
	interpreter.Object
}

func (otd opTrueDiv) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}

func (otd opTrueDiv) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	return opCall(i, "trueDiv", NewTrueDiv, NewTrueDivScalar)
}

type TrueDiv struct {
	interpreter.Object

	left  Stream
	right Stream
}

func NewTrueDiv(left Stream, right Stream) Stream {
	return &TrueDiv{
		Object: newStreamObject(),
		left:   unpublish(left),
		right:  unpublish(right),
	}
}

func (td *TrueDiv) RenderStream() string {
	return fmt.Sprintf("(%s / %s)", td.left.RenderStream(), td.right.RenderStream())
}

type TrueDivScalar struct {
	interpreter.Object

	left  Stream
	right int
}

func NewTrueDivScalar(left Stream, right int) Stream {
	return &TrueDivScalar{
		Object: newStreamObject(),
		left:   unpublish(left),
		right:  right,
	}
}

func (tds *TrueDivScalar) RenderStream() string {
	return fmt.Sprintf("(%s / %d)", tds.left.RenderStream(), tds.right)
}
