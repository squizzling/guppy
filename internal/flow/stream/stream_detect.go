package stream

import (
	"fmt"
	"time"

	"guppy/internal/flow/duration"
	"guppy/internal/interpreter"
)

type FFIDetect struct {
	interpreter.Object
}

func (f FFIDetect) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "on"},
			{Name: "off", Default: interpreter.NewObjectNone()},
			{Name: "mode", Default: interpreter.NewObjectString("paired")},
			{Name: "annotations", Default: interpreter.NewObjectNone()},
			{Name: "event_annotations", Default: interpreter.NewObjectNone()},
			{Name: "auto_resolve_after", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (f FFIDetect) resolveOn(i *interpreter.Interpreter) (Stream, error) {
	// TODO: Make sure it's a stream of bool somehow
	if predicate, err := interpreter.ArgAs[Stream](i, "on"); err != nil {
		return nil, err
	} else {
		return predicate, err
	}
}

func (f FFIDetect) resolveOff(i *interpreter.Interpreter) (Stream, error) {
	// TODO: Make sure it's a stream of bool somehow
	if off, err := interpreter.ArgAs[Stream](i, "off"); err != nil {
		return nil, err
	} else {
		return off, err
	}
}

func (f FFIDetect) resolveMode(i *interpreter.Interpreter) (string, error) {
	if mode, err := interpreter.ArgAsString(i, "mode"); err != nil {
		return "", err
	} else if mode != "split" && mode != "paired" {
		return "", fmt.Errorf("detect() mode is %s not split or paired", mode)
	} else {
		return mode, err
	}
}

func (f FFIDetect) resolveAnnotations(i *interpreter.Interpreter) (interpreter.Object, error) {
	// TODO: Check type
	if annotations, err := i.Scope.Get("annotations"); err != nil {
		return nil, err
	} else {
		return annotations, err
	}
}

func (f FFIDetect) resolveEventAnnotations(i *interpreter.Interpreter) (interpreter.Object, error) {
	// TODO: Check type
	if eventAnnotations, err := i.Scope.Get("event_annotations"); err != nil {
		return nil, err
	} else {
		return eventAnnotations, err
	}
}

func (f FFIDetect) resolveAutoResolveAfter(i *interpreter.Interpreter) (*time.Duration, error) {
	if autoResolveAfter, err := i.Scope.GetArg("auto_resolve_after"); err != nil {
		return nil, err
	} else if _, isNone := autoResolveAfter.(*interpreter.ObjectNone); isNone {
		return nil, nil
	} else if dur, isDuration := autoResolveAfter.(*duration.Duration); isDuration {
		return &dur.Duration, nil
	} else if rawMilliseconds, isDuration := autoResolveAfter.(*interpreter.ObjectInt); isDuration {
		d := time.Millisecond * time.Duration(rawMilliseconds.Value)
		return &d, nil
	} else {
		return nil, fmt.Errorf("auto_resolve_after is %T not *interpreter.ObjectNone or *duration.Duration, or *interpreter.ObjectInt", autoResolveAfter)
	}
}

func (f FFIDetect) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
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
		return NewStreamDetect(newStreamAlertObject(), on, off, mode, annotations, eventAnnotations, autoResolveAfter), nil
	}
}

var _ = interpreter.FlowCall(FFIDetect{})
