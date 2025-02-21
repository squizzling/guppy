package stream

import (
	"fmt"

	"guppy/internal/flow/filter"
	"guppy/internal/interpreter"
)

type FFIData struct {
	interpreter.Object
}

func (f FFIData) Params(i *interpreter.Interpreter) ([]interpreter.ParamData, error) {
	return []interpreter.ParamData{
		{Name: "metric"},
		{Name: "filter", Default: interpreter.NewObjectNone()},
		{Name: "rollup", Default: interpreter.NewObjectNone()},
		{Name: "extrapolation", Default: interpreter.NewObjectString("null")},
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

func resolveRollup(i *interpreter.Interpreter) (string, error) {
	if rollup, err := i.Scope.Get("rollup"); err != nil {
		return "", err
	} else {
		switch rollup := rollup.(type) {
		case *interpreter.ObjectNone:
			return "", nil
		case interpreter.FlowStringable:
			return rollup.String(i)
		default:
			return "", fmt.Errorf("rollup is %T not *interpreter.ObjectNone, or interpreter.FlowStringable", rollup)
		}
	}
}

func (f FFIData) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if metricName, err := interpreter.ArgAsString(i, "metric"); err != nil {
		return nil, err
	} else if fltr, err := resolveFilter(i); err != nil {
		return nil, err
	} else if rollup, err := resolveRollup(i); err != nil {
		return nil, err
	} else {
		return NewData(metricName, fltr, rollup), nil
	}
}

var _ = interpreter.FlowCall(FFIData{})

type data struct {
	interpreter.Object

	metricName string
	filter     filter.Filter
	rollup     string
}

func NewData(metricName string, filter filter.Filter, rollup string) Stream {
	return &data{
		Object:     newStreamObject(),
		metricName: metricName,
		filter:     filter,
		rollup:     rollup,
	}
}

func (d *data) RenderStream() string {
	filterStr := ""
	if d.filter != nil {
		filterStr = fmt.Sprintf(", filter=%s", d.filter.RenderFilter())
	}
	rollupStr := ""
	if d.rollup != "" {
		rollupStr = fmt.Sprintf(", rollup='%s'", d.rollup)
	}
	return fmt.Sprintf("data('%s'%s%s)", d.metricName, filterStr, rollupStr)
}
