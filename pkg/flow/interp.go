package flow

import (
	"guppy/pkg/flow/annotate"
	"guppy/pkg/flow/debug"
	"guppy/pkg/flow/duration"
	"guppy/pkg/flow/filter"
	"guppy/pkg/flow/stream"
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/builtin"
	"guppy/pkg/interpreter/itypes"
)

func NewInterpreter(enableTrace bool) itypes.Interpreter {
	i := interpreter.NewInterpreter(enableTrace)

	// New style
	_ = i.SetGlobal("len", builtin.NewFFILen())
	_ = i.SetGlobal("range", builtin.NewFFIRange())
	_ = i.SetGlobal("str", builtin.NewFFIStr())

	// Old style
	_ = i.SetGlobal("abs", &stream.FFIAbs{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("alerts", &stream.FFIAlerts{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("annotate", &annotate.FFIAnnotate{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("combine", &stream.FFICombine{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("const", &stream.FFIConst{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("count", &stream.FFICount{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("data", &stream.FFIData{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("detect", &stream.FFIDetect{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("events", &stream.FFIEvents{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("filter", &filter.FFIFilter{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("max", &stream.FFIMax{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("mean", &stream.FFIMean{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("median", &stream.FFIMedian{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("min", &stream.FFIMin{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("partition_filter", &filter.FFIPartitionFilter{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("sum", &stream.FFISum{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("_print", &debug.FFIPrint{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("_published", NewPublished())
	_ = i.SetGlobal("threshold", &stream.FFIThreshold{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("duration", &duration.FFIDuration{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("union", &stream.FFIUnion{Object: interpreter.NewObject(nil)})
	_ = i.SetGlobal("when", &stream.FFIWhen{Object: interpreter.NewObject(nil)})

	_ = i.Set("Args", interpreter.NewObjectDict(nil))

	return i
}
