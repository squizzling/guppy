package flow

import (
	"github.com/squizzling/guppy/pkg/flow/annotate"
	"github.com/squizzling/guppy/pkg/flow/debug"
	"github.com/squizzling/guppy/pkg/flow/duration"
	"github.com/squizzling/guppy/pkg/flow/filter"
	"github.com/squizzling/guppy/pkg/flow/stream"
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/builtin"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

func NewInterpreter(enableTrace bool) itypes.Interpreter {
	i := interpreter.NewInterpreter(enableTrace)

	// New style
	_ = i.SetGlobal("len", builtin.NewFFILen())
	_ = i.SetGlobal("range", builtin.NewFFIRange())
	_ = i.SetGlobal("repr", builtin.NewFFIRepr())
	_ = i.SetGlobal("str", builtin.NewFFIStr())

	_ = i.SetGlobal("filter", filter.NewFFIFilter())
	_ = i.SetGlobal("partition_filter", filter.NewFFIPartitionFilter())

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
	_ = i.SetGlobal("max", &stream.FFIMax{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("mean", &stream.FFIMean{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("median", &stream.FFIMedian{Object: itypes.NewObject(nil)})
	_ = i.SetGlobal("min", &stream.FFIMin{Object: itypes.NewObject(nil)})
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
