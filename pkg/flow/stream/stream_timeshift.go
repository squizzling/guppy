package stream

import (
	"time"

	"guppy/pkg/flow/duration"
	"guppy/pkg/interpreter"
)

type methodTimeShift struct {
	interpreter.Object
}

func (mts methodTimeShift) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "offset"},
		},
	}, nil
}

func resolveOffset(i *interpreter.Interpreter) (string, error) {
	if offset, err := interpreter.ArgAsString(i, "offset"); err != nil {
		return "", err
	} else {
		return offset, nil
	}
}

func (mts methodTimeShift) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
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
