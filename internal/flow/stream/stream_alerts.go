package stream

import (
	"fmt"

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
	}, nil
}

func (f FFIAlerts) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	// TODO: Process arguments
	return NewAlerts(), nil
}

var _ = interpreter.FlowCall(FFIAlerts{})

type Alerts struct {
	interpreter.Object
}

func NewAlerts() Stream {
	return &Alerts{
		Object: newStreamObject(), // TODO: This doesn't have the usual methods
	}
}

func (a *Alerts) RenderStream() string {
	// TODO: Render
	return fmt.Sprintf("alerts()")
}
