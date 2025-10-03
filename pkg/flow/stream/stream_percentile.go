package stream

import (
	"guppy/pkg/interpreter"
)

// TODO: All of this.

type methodPercentile struct {
	interpreter.Object
}

func (mp methodPercentile) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "pct", Default: interpreter.NewObjectNone()},
			{Name: "allow_missing", Default: interpreter.NewObjectNone()},
			{Name: "by", Default: interpreter.NewObjectNone()},
			{Name: "over", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (mp methodPercentile) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamMethodPercentile(newStreamObject(), unpublish(self)), nil
	}
}
