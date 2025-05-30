package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodTop struct {
	interpreter.Object
}

func (mt methodTop) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "count", Default: interpreter.NewObjectNone()},
			{Name: "by", Default: interpreter.NewObjectNone()},
			{Name: "allow_missing", Default: interpreter.NewObjectNone()},
			{Name: "percentage", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (mt methodTop) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewTop(self), nil
	}
}

type Top struct {
	interpreter.Object

	source Stream
}

func NewTop(source Stream) Stream {
	return &Top{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (t *Top) RenderStream() string {
	return t.source.RenderStream() + ".top()"
}
