package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFICombine struct {
	itypes.Object
}

func (f FFICombine) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "expression"},
			{Name: "mode", Default: interpreter.NewObjectMissing()},
		},
	}, nil
}

func (f FFICombine) resolveExpression(i itypes.Interpreter) (Stream, error) {
	if expression, err := itypes.ArgAs[Stream](i, "expression"); err != nil {
		return nil, err
	} else {
		return expression, nil
	}
}

func (f FFICombine) resolveMode(i itypes.Interpreter) error {
	if mode, err := i.GetArg("mode"); err != nil {
		return err
	} else if _, isMissing := mode.(*interpreter.ObjectMissing); !isMissing {
		return fmt.Errorf("got combine mode %s", mode)
	} else {
		return nil
	}
}

func (f FFICombine) Call(i itypes.Interpreter) (itypes.Object, error) {
	if expression, err := f.resolveExpression(i); err != nil {
		return nil, err
	} else if err = f.resolveMode(i); err != nil {
		return nil, err
	} else {
		return NewStreamFuncCombine(newStreamObject(), expression, ""), nil
	}
}

var _ = itypes.FlowCall(FFICombine{})
