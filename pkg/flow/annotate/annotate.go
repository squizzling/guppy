package annotate

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIAnnotate struct {
	itypes.Object
}

func (f FFIAnnotate) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "value"},
			{Name: "label"},
			{Name: "extra_props", Default: interpreter.NewObjectNone()},
			{Name: "publish", Default: interpreter.NewObjectMissing()},
		},
	}, nil
}

func (f FFIAnnotate) resolveValue(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: This is probably a Stream, try and assert it
	if value, err := i.GetArg("value"); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (f FFIAnnotate) resolveLabel(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: This is probably a string, try and assert it
	if label, err := i.GetArg("label"); err != nil {
		return nil, err
	} else {
		return label, nil
	}
}

func (f FFIAnnotate) resolveExtraProps(i itypes.Interpreter) (*interpreter.ObjectDict, error) {
	if extraProps, err := i.GetArg("extra_props"); err != nil {
		return nil, err
	} else if _, ok := extraProps.(*interpreter.ObjectMissing); ok {
		return interpreter.NewObjectDict(nil).(*interpreter.ObjectDict), nil
	} else if extraProps, ok := extraProps.(*interpreter.ObjectDict); ok {
		return extraProps, nil
	} else {
		return nil, fmt.Errorf("annotate(extraProps) must be missing or dict")
	}
}

func (f FFIAnnotate) resolvePublish(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: No idea what this is, maybe a string?
	if publish, err := i.GetArg("publish"); err != nil {
		return nil, err
	} else {
		return publish, nil
	}
}

func (f FFIAnnotate) Call(i itypes.Interpreter) (itypes.Object, error) {
	if value, err := f.resolveValue(i); err != nil {
		return nil, err
	} else if label, err := f.resolveLabel(i); err != nil {
		return nil, err
	} else if extraProps, err := f.resolveExtraProps(i); err != nil {
		return nil, err
	} else if publish, err := f.resolvePublish(i); err != nil {
		return nil, err
	} else {
		return NewAnnotated(value, label, extraProps, publish), nil
	}
}

var _ = interpreter.FlowCall(FFIAnnotate{})

type Annotated struct {
	itypes.Object

	Value      itypes.Object
	Label      itypes.Object
	ExtraProps *interpreter.ObjectDict
	Publish    itypes.Object
}

func NewAnnotated(value itypes.Object, label itypes.Object, extraProps *interpreter.ObjectDict, publish itypes.Object) *Annotated {
	return &Annotated{
		Object:     itypes.NewObject(nil),
		Value:      value,
		Label:      label,
		ExtraProps: extraProps,
		Publish:    publish,
	}
}
