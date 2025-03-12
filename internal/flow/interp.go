package flow

import (
	"guppy/internal/flow/debug"
	"guppy/internal/flow/filter"
	"guppy/internal/flow/stream"
	"guppy/internal/interpreter"
)

func NewInterpreter(enableTrace bool) *interpreter.Interpreter {
	i := interpreter.NewInterpreter(enableTrace)

	_ = i.Globals.Set("data", &stream.FFIData{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("events", &stream.FFIEvents{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("filter", &filter.FFIFilter{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("_print", &debug.FFIPrint{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("threshold", &stream.FFIThreshold{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("union", &stream.FFIUnion{Object: interpreter.NewObject(nil)})

	_ = i.Scope.Set("Args", interpreter.NewObjectDict(nil))

	return i
}
