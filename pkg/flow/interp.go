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
	"guppy/pkg/interpreter/primitive"
)

func NewInterpreter(enableTrace bool) itypes.Interpreter {
	i := interpreter.NewInterpreter(enableTrace)

	// New style
	_ = i.SetGlobal("len", builtin.NewFFILen())
	_ = i.SetGlobal("range", builtin.NewFFIRange())
	_ = i.SetGlobal("repr", builtin.NewFFIRepr())
	_ = i.SetGlobal("str", builtin.NewFFIStr())

	// Old style
	_ = i.SetGlobal("abs", &stream.FFIAbs{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("alerts", &stream.FFIAlerts{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("annotate", &annotate.FFIAnnotate{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("combine", &stream.FFICombine{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("const", &stream.FFIConst{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("count", &stream.FFICount{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("data", &stream.FFIData{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("detect", &stream.FFIDetect{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("events", &stream.FFIEvents{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("filter", &filter.FFIFilter{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("max", &stream.FFIMax{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("mean", &stream.FFIMean{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("median", &stream.FFIMedian{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("min", &stream.FFIMin{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("partition_filter", &filter.FFIPartitionFilter{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("sum", &stream.FFISum{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("_print", &debug.FFIPrint{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("_published", NewPublished())
	_ = i.SetGlobal("threshold", &stream.FFIThreshold{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("duration", &duration.FFIDuration{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("union", &stream.FFIUnion{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("when", &stream.FFIWhen{Object: itypes.NewObject(nil)})

	_ = i.Set("Args", primitive.NewObjectDict(nil))

	return i
}
