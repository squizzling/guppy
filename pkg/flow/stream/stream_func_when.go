package stream

import (
	"fmt"
	"time"

	"github.com/squizzling/guppy/pkg/flow/duration"
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type FFIWhen struct {
	itypes.Object
}

func (f FFIWhen) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "predicate"},
			{Name: "lasting", Default: primitive.NewObjectNone()},
			{Name: "at_least", Default: primitive.NewObjectDouble(1.0)},
		},
	}, nil
}

func (f FFIWhen) resolvePredicate(i itypes.Interpreter) (Stream, error) {
	if predicate, err := itypes.ArgAs[Stream](i, "predicate"); err != nil {
		return nil, err
	} else {
		return predicate, err
	}
}

func (f FFIWhen) resolveLasting(i itypes.Interpreter) (*time.Duration, error) {
	if lasting, err := i.GetArg("lasting"); err != nil {
		return nil, err
	} else if _, isNone := lasting.(*primitive.ObjectNone); isNone {
		return nil, nil
	} else if dur, isDuration := lasting.(*duration.Duration); isDuration {
		return &dur.Duration, nil
	} else if objStr, ok := lasting.(*primitive.ObjectString); !ok {
		return nil, fmt.Errorf("lasting is %T not *interpreter.ObjectNone, *duration.Duration, or *interpreter.ObjectString", lasting)
	} else if dur, err := duration.ParseDuration(objStr.Value); err != nil {
		return nil, err
	} else {
		return &dur, nil
	}
}

func (f FFIWhen) resolveAtLeast(i itypes.Interpreter) (float64, error) {
	if atLeast, err := interpreter.ArgAsDouble(i, "at_least"); err != nil {
		return 0, err
	} else {
		// TODO: Validate range
		return atLeast, err
	}
}

func (f FFIWhen) Call(i itypes.Interpreter) (itypes.Object, error) {
	if predicate, err := f.resolvePredicate(i); err != nil {
		return nil, err
	} else if lasting, err := f.resolveLasting(i); err != nil {
		return nil, err
	} else if atLeast, err := f.resolveAtLeast(i); err != nil {
		return nil, err
	} else {
		return NewStreamFuncWhen(newStreamBoolObject(), predicate, lasting, atLeast), nil
	}
}

var _ = itypes.FlowCall(FFIWhen{})
