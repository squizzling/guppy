package stream

import (
	"fmt"

	"github.com/squizzling/types/pkg/result"

	"guppy/internal/flow/filter"
	"guppy/internal/interpreter"
)

type FFIData struct {
	interpreter.Object
}

func (f FFIData) Args(i *interpreter.Interpreter) result.Result[[]interpreter.ArgData] {
	return result.Ok([]interpreter.ArgData{
		{Name: "metric"},
		{Name: "filter", Default: interpreter.NewObjectNone()},
		//{Name: "rollup", Default: interpreter.NewObjectNone()},
		//{Name: "extrapolation", Default: interpreter.NewObjectString("null")},
		//{Name: "maxExtrapolations", Default: interpreter.NewObjectNone()}, // TODO: Check what this has for a default
		//{Name: "resolution", Default: interpreter.NewObjectNone()},
	})
}

func (f FFIData) Call(i *interpreter.Interpreter) result.Result[interpreter.Object] {
	if resultMetricName := interpreter.ArgAsString(i, "metric"); !resultMetricName.Ok() {
		return result.Err[interpreter.Object](resultMetricName.Err())
	} else if resultObjFilter := i.Scope.Get("filter"); !resultObjFilter.Ok() {
		return result.Err[interpreter.Object](resultObjFilter.Err())
	} else {
		var actualFilter filter.Filter
		switch fltr := resultObjFilter.Value().(type) {
		case *interpreter.ObjectNone:
			actualFilter = nil
		case filter.Filter:
			actualFilter = fltr
		default:
			return result.Err[interpreter.Object](fmt.Errorf("filter is %T not *interpreter.ObjectNone, or filter.Filter", fltr))
		}
		return result.Ok[interpreter.Object](NewData(resultMetricName.Value(), actualFilter))
	}
}

var _ = interpreter.FlowCall(FFIData{})

type data struct {
	interpreter.Object

	metricName string
	filter     filter.Filter
}

func NewData(metricName string, filter filter.Filter) Stream {
	return &data{
		Object:     newStreamObject(),
		metricName: metricName,
		filter:     filter,
	}
}

func (d *data) RenderStream() string {
	filterStr := ""
	if d.filter != nil {
		filterStr = fmt.Sprintf(", filter=%s", d.filter.RenderFilter())
	}
	return fmt.Sprintf("data('%s'%s)", d.metricName, filterStr)
}
