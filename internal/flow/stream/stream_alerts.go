package stream

import (
	"guppy/internal/interpreter"
)

type FFIAlerts struct {
	interpreter.Object
}

func (f FFIAlerts) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "detector_id", Default: interpreter.NewObjectNone()},
			{Name: "detector_name", Default: interpreter.NewObjectNone()},
			{Name: "autodetect_id", Default: interpreter.NewObjectNone()},
		},
		KWParams: []interpreter.ParamDef{
			{Name: "filter", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (f FFIAlerts) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	// TODO: Process arguments
	// TODO: This doesn't have the usual stream methods
	return NewStreamAlerts(newStreamObject()), nil
}

var _ = interpreter.FlowCall(FFIAlerts{})
