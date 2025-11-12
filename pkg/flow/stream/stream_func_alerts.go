package stream

import (
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type FFIAlerts struct {
	itypes.Object
}

func (f FFIAlerts) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "detector_id", Default: primitive.NewObjectNone()},
			{Name: "detector_name", Default: primitive.NewObjectNone()},
			{Name: "autodetect_id", Default: primitive.NewObjectNone()},
		},
		KWParams: []itypes.ParamDef{
			{Name: "filter", Default: primitive.NewObjectNone()},
		},
	}, nil
}

func (f FFIAlerts) Call(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: Process arguments
	// TODO: This doesn't have the usual stream methods
	return NewStreamFuncAlerts(newStreamObject()), nil
}

var _ = itypes.FlowCall(FFIAlerts{})
