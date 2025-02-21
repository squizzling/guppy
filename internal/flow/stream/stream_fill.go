package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodFill struct {
	interpreter.Object
}

func (mf methodFill) Args(i *interpreter.Interpreter) ([]interpreter.ArgData, error) {
	return []interpreter.ArgData{
		{Name: "self"},
		{Name: "value", Default: interpreter.NewObjectNone()},
		{Name: "duration", Default: interpreter.NewObjectNone()},
		{Name: "maxCount", Default: interpreter.NewObjectNone()},
	}, nil
}

func resolveDuration(i *interpreter.Interpreter) (int, error) {
	if by, err := i.Scope.Get("duration"); err != nil {
		return 0, err
	} else {
		switch by := by.(type) {
		case *interpreter.ObjectNone:
			return 0, nil // explicitly nil
		default:
			return 0, fmt.Errorf("duration is %T not *interpreter.ObjectNone", by)
		}
	}
}

func resolveMaxCount(i *interpreter.Interpreter) (int, error) {
	if by, err := i.Scope.Get("maxCount"); err != nil {
		return 0, err
	} else {
		switch by := by.(type) {
		case *interpreter.ObjectNone:
			return 0, nil // explicitly nil
		default:
			return 0, fmt.Errorf("maxCount is %T not *interpreter.ObjectNone", by)
		}
	}
}

func (mf methodFill) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if value, err := i.Scope.Get("value"); err != nil {
		return nil, err
	} else if duration, err := resolveDuration(i); err != nil {
		return nil, err
	} else if maxCount, err := resolveMaxCount(i); err != nil {
		return nil, err
	} else {
		return NewFill(self, value, duration, maxCount), nil
	}
}

type fill struct {
	interpreter.Object

	source   Stream
	value    interpreter.Object
	duration int
	maxCount int
}

func NewFill(source Stream, value interpreter.Object, duration int, maxCount int) Stream {
	return &fill{
		Object:   newStreamObject(),
		source:   unpublish(source),
		value:    value,
		duration: duration,
		maxCount: maxCount,
	}
}

func (f *fill) RenderStream() string {
	s := ""
	if f.value != nil {
		s += fmt.Sprintf("value=%#v", f.value)
	}
	if f.duration > 0 {
		s += fmt.Sprintf("duration=%d", f.duration)
	}
	if f.maxCount > 0 {
		s += fmt.Sprintf("maxCount=%d", f.maxCount)
	}
	return f.source.RenderStream() + ".fill(" + s + ")"
}
