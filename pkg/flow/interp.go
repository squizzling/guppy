package flow

import (
	"guppy/pkg/flow/annotate"
	"guppy/pkg/flow/debug"
	"guppy/pkg/flow/duration"
	"guppy/pkg/flow/filter"
	"guppy/pkg/flow/stream"
	"guppy/pkg/interpreter"
)

func NewInterpreter(enableTrace bool) *interpreter.Interpreter {
	i := interpreter.NewInterpreter(enableTrace)

	_ = i.Globals.Set("alerts", &stream.FFIAlerts{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("annotate", &annotate.FFIAnnotate{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("combine", &stream.FFICombine{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("const", &stream.FFIConst{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("count", &stream.FFICount{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("data", &stream.FFIData{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("detect", &stream.FFIDetect{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("events", &stream.FFIEvents{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("filter", &filter.FFIFilter{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("max", &stream.FFIMax{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("mean", &stream.FFIMean{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("median", &stream.FFIMedian{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("min", &stream.FFIMin{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("str", &interpreter.FFIStr{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("_print", &debug.FFIPrint{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("_published", NewPublished())
	_ = i.Globals.Set("threshold", &stream.FFIThreshold{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("duration", &duration.FFIDuration{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("union", &stream.FFIUnion{Object: interpreter.NewObject(nil)})
	_ = i.Globals.Set("when", &stream.FFIWhen{Object: interpreter.NewObject(nil)})

	_ = i.Scope.Set("Args", interpreter.NewObjectDict(nil))

	return i
}
