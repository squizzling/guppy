package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodPublish struct {
	interpreter.Object
}

func (mp methodPublish) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "label"},
		},
	}, nil
}

func (mp methodPublish) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if label, err := interpreter.ArgAsString(i, "label"); err != nil {
		return nil, err
	} else {
		return NewPublish(self, label), nil
	}
}

type publish struct {
	interpreter.Object

	source Stream
	label  string
}

func NewPublish(source Stream, label string) Stream {
	s := &publish{
		Object: newStreamObject(),
		source: source,
		label:  label,
	}
	fmt.Printf("%s\n", s.RenderStream())
	return s
}

func (p *publish) RenderStream() string {
	return fmt.Sprintf("%s.publish(label='%s')", p.source.RenderStream(), p.label)
}
