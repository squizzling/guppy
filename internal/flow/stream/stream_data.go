package stream

import (
	"fmt"

	"guppy/internal/flow/filter"
	"guppy/internal/interpreter"
)

type FFIData struct {
	interpreter.Object
}

func (f FFIData) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "metric"},
			{Name: "filter", Default: interpreter.NewObjectNone()},
			{Name: "rollup", Default: interpreter.NewObjectNone()},
			{Name: "extrapolation", Default: interpreter.NewObjectString("null")},
			{Name: "maxExtrapolations", Default: interpreter.NewObjectInt(0)},
			{Name: "resolution", Default: interpreter.NewObjectNone()}, // TODO: Handle
		},
	}, nil
}

func resolveFilter(i *interpreter.Interpreter) (filter.Filter, error) {
	if fltr, err := i.Scope.GetArg("filter"); err != nil {
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
	if rollup, err := i.Scope.GetArg("rollup"); err != nil {
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

func resolveExtrapolation(i *interpreter.Interpreter) (string, error) {
	if extrapolation, err := interpreter.ArgAsString(i, "extrapolation"); err != nil {
		return "", err
	} else {
		return extrapolation, nil
	}
}

func resolveMaxExtrapolations(i *interpreter.Interpreter) (int, error) {
	if maxExtrapolations, err := interpreter.ArgAs[*interpreter.ObjectInt](i, "maxExtrapolations"); err != nil {
		return 0, err
	} else {
		return maxExtrapolations.Value, nil
	}
}

func (f FFIData) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if metricName, err := interpreter.ArgAsString(i, "metric"); err != nil {
		return nil, err
	} else if fltr, err := resolveFilter(i); err != nil {
		return nil, err
	} else if rollup, err := resolveRollup(i); err != nil {
		return nil, err
	} else if extrapolation, err := resolveExtrapolation(i); err != nil {
		return nil, err
	} else if maxExtrapolations, err := resolveMaxExtrapolations(i); err != nil {
		return nil, err
	} else {
		return NewData(metricName, fltr, rollup, extrapolation, maxExtrapolations), nil
	}
}

var _ = interpreter.FlowCall(FFIData{})

type data struct {
	interpreter.Object

	metricName        string
	filter            filter.Filter
	rollup            string
	extrapolation     string
	maxExtrapolations int
}

func NewData(metricName string, filter filter.Filter, rollup string, extrapolation string, maxExtrapolations int) Stream {
	return &data{
		Object:            newStreamObject(),
		metricName:        metricName,
		filter:            filter,
		rollup:            rollup,
		extrapolation:     extrapolation,
		maxExtrapolations: maxExtrapolations,
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

	return fmt.Sprintf("data('%s'%s%s, extrapolation='%s', maxExtrapolations=%d)", d.metricName, filterStr, rollupStr, d.extrapolation, d.maxExtrapolations)
}
