package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFICombine struct {
	interpreter.Object
}

func (f FFICombine) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "expression"},
			{Name: "mode", Default: interpreter.NewObjectMissing()},
		},
	}, nil
}

func (f FFICombine) resolveExpression(i *interpreter.Interpreter) (Stream, error) {
	if expression, err := interpreter.ArgAs[Stream](i, "expression"); err != nil {
		return nil, err
	} else {
		return expression, nil
	}
}

func (f FFICombine) resolveMode(i *interpreter.Interpreter) error {
	if mode, err := i.Scope.Get("mode"); err != nil {
		return err
	} else if _, isMissing := mode.(*interpreter.ObjectMissing); !isMissing {
		return fmt.Errorf("got combine mode %s", mode)
	} else {
		return nil
	}
}

func (f FFICombine) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if expression, err := f.resolveExpression(i); err != nil {
		return nil, err
	} else if err = f.resolveMode(i); err != nil {
		return nil, err
	} else {
		return NewStreamCombine(newStreamObject(), expression, ""), nil
	}
}

var _ = interpreter.FlowCall(FFICombine{})
