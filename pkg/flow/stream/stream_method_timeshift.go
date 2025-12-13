package stream

import (
	"fmt"
	"time"

	"github.com/squizzling/guppy/pkg/flow/duration"
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type methodTimeShift struct {
	itypes.Object
}

func (mts methodTimeShift) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "offset"},
		},
	}, nil
}

func resolveOffset(i itypes.Interpreter) (string, error) {
	if offset, err := interpreter.ArgAsString(i, "offset"); err != nil {
		return "", err
	} else {
		return offset, nil
	}
}

func (mts methodTimeShift) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if offset, err := resolveOffset(i); err != nil {
		return nil, err
	} else if dur, err := duration.ParseDuration(offset); err != nil {
		return nil, err
	} else {
		return NewStreamMethodTimeShift(newStreamObject(), self, dur), nil
		//return self.CloneTimeShift(dur), nil
	}
}

func (mts methodTimeShift) Repr() string {
	return "methodTimeShift()"
}

func cloneTimeshift(s Stream, amount time.Duration) Stream {
	if s == nil {
		panic("I just want to see if this happens")
		return s
	}
	switch s := s.(type) {
	// TODO: Handle other generating commands
	case *StreamFuncData:
		newStream := s.CloneTimeShift(amount).(*StreamFuncData)
		newStream.TimeShift += amount
		return newStream
	case *StreamMethodPublish: // Remove Publish from the time-shifted graph
		return s.Source.CloneTimeShift(amount)
	default:
		return s.CloneTimeShift(amount)
	}
}

func (smts *StreamMethodTimeShift) Repr() string {
	// TODO: Better
	return fmt.Sprintf("timeshift()")
}
