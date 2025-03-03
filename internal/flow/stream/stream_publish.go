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
			{Name: "enable", Default: interpreter.NewObjectBool(true)},
		},
		//KWParam: "additional_dimensions", // Maybe, I don't fully know this one.
	}, nil
}

func (mp methodPublish) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if label, err := interpreter.ArgAsString(i, "label"); err != nil {
		return nil, err
	} else if enable, err := interpreter.ArgAsBool(i, "enable"); err != nil {
		return nil, err
	} else {
		return NewPublish(self, label, enable), nil
	}
}

type publish struct {
	interpreter.Object

	source Stream
	label  string
	enable bool
}

func NewPublish(source Stream, label string, enable bool) Stream {
	s := &publish{
		Object: newStreamObject(),
		source: source,
		label:  label,
		enable: enable,
	}
	fmt.Printf("%s\n", s.RenderStream())
	return s
}

func (p *publish) RenderStream() string {
	return fmt.Sprintf("%s.publish(label='%s')", p.source.RenderStream(), p.label)
}
