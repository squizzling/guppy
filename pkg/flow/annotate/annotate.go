package annotate

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFIAnnotate struct {
	interpreter.Object
}

func (f FFIAnnotate) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "value"},
			{Name: "label"},
			{Name: "extra_props", Default: interpreter.NewObjectNone()},
			{Name: "publish", Default: interpreter.NewObjectMissing()},
		},
	}, nil
}

func (f FFIAnnotate) resolveValue(i *interpreter.Interpreter) (interpreter.Object, error) {
	// TODO: This is probably a Stream, try and assert it
	if value, err := i.Scope.GetArg("value"); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (f FFIAnnotate) resolveLabel(i *interpreter.Interpreter) (interpreter.Object, error) {
	// TODO: This is probably a string, try and assert it
	if label, err := i.Scope.GetArg("label"); err != nil {
		return nil, err
	} else {
		return label, nil
	}
}

func (f FFIAnnotate) resolveExtraProps(i *interpreter.Interpreter) (*interpreter.ObjectDict, error) {
	if extraProps, err := i.Scope.GetArg("extra_props"); err != nil {
		return nil, err
	} else if _, ok := extraProps.(*interpreter.ObjectMissing); ok {
		return interpreter.NewObjectDict(nil).(*interpreter.ObjectDict), nil
	} else if extraProps, ok := extraProps.(*interpreter.ObjectDict); ok {
		return extraProps, nil
	} else {
		return nil, fmt.Errorf("annotate(extraProps) must be missing or dict")
	}
}

func (f FFIAnnotate) resolvePublish(i *interpreter.Interpreter) (interpreter.Object, error) {
	// TODO: No idea what this is, maybe a string?
	if publish, err := i.Scope.GetArg("publish"); err != nil {
		return nil, err
	} else {
		return publish, nil
	}
}

func (f FFIAnnotate) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
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
	interpreter.Object

	Value      interpreter.Object
	Label      interpreter.Object
	ExtraProps *interpreter.ObjectDict
	Publish    interpreter.Object
}

func NewAnnotated(value interpreter.Object, label interpreter.Object, extraProps *interpreter.ObjectDict, publish interpreter.Object) *Annotated {
	return &Annotated{
		Object:     interpreter.NewObject(nil),
		Value:      value,
		Label:      label,
		ExtraProps: extraProps,
		Publish:    publish,
	}
}
