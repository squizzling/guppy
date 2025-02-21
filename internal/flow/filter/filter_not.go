package filter

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodBinaryNot struct {
	interpreter.Object
}

func (mba methodBinaryNot) Args(i *interpreter.Interpreter) ([]interpreter.ArgData, error) {
	return []interpreter.ArgData{
		{Name: "self"},
	}, nil
}

func (mba methodBinaryNot) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Filter](i, "self"); err != nil {
		return nil, err
	} else {
		return NewNot(self), nil
	}
}

type not struct {
	interpreter.Object

	right Filter
}

func NewNot(right Filter) Filter {
	return &not{
		Object: newFilterObject(),
		right:  right,
	}
}

func (n *not) RenderFilter() string {
	return fmt.Sprintf("(not %s)", n.right.RenderFilter())
}
