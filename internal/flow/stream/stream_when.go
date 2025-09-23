package stream

import (
	"fmt"
	"time"

	"guppy/internal/flow/duration"
	"guppy/internal/interpreter"
)

type FFIWhen struct {
	interpreter.Object
}

func (f FFIWhen) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "predicate"},
			{Name: "lasting", Default: interpreter.NewObjectNone()},
			{Name: "at_least", Default: interpreter.NewObjectDouble(1.0)},
		},
	}, nil
}

func (f FFIWhen) resolvePredicate(i *interpreter.Interpreter) (Stream, error) {
	if predicate, err := interpreter.ArgAs[Stream](i, "predicate"); err != nil {
		return nil, err
	} else {
		return predicate, err
	}
}

func (f FFIWhen) resolveLasting(i *interpreter.Interpreter) (*time.Duration, error) {
	if lasting, err := i.Scope.GetArg("lasting"); err != nil {
		return nil, err
	} else if _, isNone := lasting.(*interpreter.ObjectNone); isNone {
		return nil, nil
	} else if dur, isDuration := lasting.(*duration.Duration); isDuration {
		return &dur.Duration, nil
	} else {
		return nil, fmt.Errorf("lasting is %T not *interpreter.ObjectNone or *duration.Duration", lasting)
	}
}

func (f FFIWhen) resolveAtLeast(i *interpreter.Interpreter) (float64, error) {
	if atLeast, err := interpreter.ArgAsDouble(i, "at_least"); err != nil {
		return 0, err
	} else {
		// TODO: Validate range
		return atLeast, err
	}
}

func (f FFIWhen) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if predicate, err := f.resolvePredicate(i); err != nil {
		return nil, err
	} else if lasting, err := f.resolveLasting(i); err != nil {
		return nil, err
	} else if atLeast, err := f.resolveAtLeast(i); err != nil {
		return nil, err
	} else {
		return NewStreamWhen(newStreamBoolObject(), predicate, lasting, atLeast), nil
	}
}

var _ = interpreter.FlowCall(FFIWhen{})
