package stream

import (
	"fmt"

	"guppy/internal/flow/filter"
	"guppy/internal/interpreter"
)

type FFIData struct {
	interpreter.Object
}

func (f FFIData) Args(i *interpreter.Interpreter) ([]interpreter.ArgData, error) {
	return []interpreter.ArgData{
		{Name: "metric"},
		{Name: "filter", Default: interpreter.NewObjectNone()},
		//{Name: "rollup", Default: interpreter.NewObjectNone()},
		//{Name: "extrapolation", Default: interpreter.NewObjectString("null")},
		//{Name: "maxExtrapolations", Default: interpreter.NewObjectNone()}, // TODO: Check what this has for a default
		//{Name: "resolution", Default: interpreter.NewObjectNone()},
	}, nil
}

func resolveFilter(i *interpreter.Interpreter) (filter.Filter, error) {
	if fltr, err := i.Scope.Get("filter"); err != nil {
		return nil, err
	} else {
		switch fltr := fltr.(type) {
		case *interpreter.ObjectNone:
			return nil, nil
		case filter.Filter:
			return fltr, nil
		default:
			return nil, fmt.Errorf("filter is %T not *interpreter.ObjectNone, or filter.Filter", fltr)
		}
	}
}

func (f FFIData) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if metricName, err := interpreter.ArgAsString(i, "metric"); err != nil {
		return nil, err
	} else if fltr, err := resolveFilter(i); err != nil {
		return nil, err
	} else {
		return NewData(metricName, fltr), nil
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
