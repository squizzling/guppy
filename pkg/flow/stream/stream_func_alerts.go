package stream

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIAlerts struct {
	itypes.Object
}

func (f FFIAlerts) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "detector_id", Default: interpreter.NewObjectNone()},
			{Name: "detector_name", Default: interpreter.NewObjectNone()},
			{Name: "autodetect_id", Default: interpreter.NewObjectNone()},
		},
		KWParams: []itypes.ParamDef{
			{Name: "filter", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func (f FFIAlerts) Call(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: Process arguments
	// TODO: This doesn't have the usual stream methods
	return NewStreamFuncAlerts(newStreamObject()), nil
}

var _ = itypes.FlowCall(FFIAlerts{})
