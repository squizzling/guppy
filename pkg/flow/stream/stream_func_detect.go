package stream

import (
	"fmt"
	"time"

	"github.com/squizzling/guppy/pkg/flow/duration"
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type FFIDetect struct {
	itypes.Object
}

func (f FFIDetect) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "on"},
			{Name: "off", Default: primitive.NewObjectNone()},
			{Name: "mode", Default: primitive.NewObjectString("paired")},
			{Name: "annotations", Default: primitive.NewObjectNone()},
			{Name: "event_annotations", Default: primitive.NewObjectNone()},
			{Name: "auto_resolve_after", Default: primitive.NewObjectNone()},
		},
	}, nil
}

func (f FFIDetect) resolveOn(i itypes.Interpreter) (Stream, error) {
	// TODO: Make sure it's a stream of bool somehow
	if on, err := itypes.ArgAs[Stream](i, "on"); err != nil {
		return nil, err
	} else {
		return on, err
	}
}

func (f FFIDetect) resolveOff(i itypes.Interpreter) (Stream, error) {
	// TODO: Make sure it's a stream of bool somehow
	if off, err := i.GetArg("off"); err != nil {
		return nil, err
	} else if _, isNone := off.(*primitive.ObjectNone); isNone {
		return nil, nil
	} else if off == nil {
		return nil, nil
	} else if offStream, isStream := off.(Stream); isStream {
		return offStream, nil
	} else {
		return nil, fmt.Errorf("detect(off) is %T not *interpreter.ObjectNone, stream.Stream, or nil", off)
	}
}

func (f FFIDetect) resolveMode(i itypes.Interpreter) (string, error) {
	if mode, err := interpreter.ArgAsString(i, "mode"); err != nil {
		return "", err
	} else if mode != "split" && mode != "paired" {
		return "", fmt.Errorf("detect() mode is %s not split or paired", mode)
	} else {
		return mode, err
	}
}

func (f FFIDetect) resolveAnnotations(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: Check type
	if annotations, err := i.GetArg("annotations"); err != nil {
		return nil, err
	} else {
		return annotations, err
	}
}

func (f FFIDetect) resolveEventAnnotations(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: Check type
	if eventAnnotations, err := i.GetArg("event_annotations"); err != nil {
		return nil, err
	} else {
		return eventAnnotations, err
	}
}

func (f FFIDetect) resolveAutoResolveAfter(i itypes.Interpreter) (*time.Duration, error) {
	if autoResolveAfter, err := i.GetArg("auto_resolve_after"); err != nil {
		return nil, err
	} else if _, isNone := autoResolveAfter.(*primitive.ObjectNone); isNone {
		return nil, nil
	} else if dur, isDuration := autoResolveAfter.(*duration.Duration); isDuration {
		return &dur.Duration, nil
	} else if rawMilliseconds, isDuration := autoResolveAfter.(*primitive.ObjectInt); isDuration {
		d := time.Millisecond * time.Duration(rawMilliseconds.Value)
		return &d, nil
	} else {
		return nil, fmt.Errorf("auto_resolve_after is %T not *interpreter.ObjectNone or *duration.Duration, or *interpreter.ObjectInt", autoResolveAfter)
	}
}

func (f FFIDetect) Call(i itypes.Interpreter) (itypes.Object, error) {
	if on, err := f.resolveOn(i); err != nil {
		return nil, err
	} else if off, err := f.resolveOff(i); err != nil {
		return nil, err
	} else if mode, err := f.resolveMode(i); err != nil {
		return nil, err
	} else if annotations, err := f.resolveAnnotations(i); err != nil {
		return nil, err
	} else if eventAnnotations, err := f.resolveEventAnnotations(i); err != nil {
		return nil, err
	} else if autoResolveAfter, err := f.resolveAutoResolveAfter(i); err != nil {
		return nil, err
	} else {
		return NewStreamFuncDetect(prototypeStreamAlert, on, off, mode, annotations, eventAnnotations, autoResolveAfter), nil
	}
}

var _ = itypes.FlowCall(FFIDetect{})
