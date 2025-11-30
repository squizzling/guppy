package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/flow/filter"
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type FFIData struct {
	itypes.Object
}

func (f FFIData) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "metric"},
			{Name: "filter", Default: primitive.NewObjectNone()},
			{Name: "rollup", Default: primitive.NewObjectNone()},
			{Name: "extrapolation", Default: primitive.NewObjectString("null")},
			{Name: "maxExtrapolations", Default: primitive.NewObjectInt(-1)},
			{Name: "resolution", Default: primitive.NewObjectNone()}, // TODO: Handle
		},
	}, nil
}

func resolveFilter(i itypes.Interpreter) (filter.Filter, error) {
	if fltr, err := i.GetArg("filter"); err != nil {
		return nil, err
	} else {
		switch fltr := fltr.(type) {
		case *primitive.ObjectNone:
			return nil, nil
		case filter.Filter:
			return fltr, nil
		default:
			return nil, fmt.Errorf("filter is %T not *interpreter.ObjectNone, or filter.Filter", fltr)
		}
	}
}

func resolveRollup(i itypes.Interpreter) (string, error) {
	if rollup, err := i.GetArg("rollup"); err != nil {
		return "", err
	} else {
		switch rollup := rollup.(type) {
		case *primitive.ObjectNone:
			return "", nil
		case interpreter.FlowStringable:
			return rollup.String(i)
		default:
			return "", fmt.Errorf("rollup is %T not *interpreter.ObjectNone, or interpreter.FlowStringable", rollup)
		}
	}
}

func resolveExtrapolation(i itypes.Interpreter) (string, error) {
	if extrapolation, err := interpreter.ArgAsString(i, "extrapolation"); err != nil {
		return "", err
	} else {
		return extrapolation, nil
	}
}

func resolveMaxExtrapolations(i itypes.Interpreter) (int, error) {
	if maxExtrapolations, err := itypes.ArgAs[*primitive.ObjectInt](i, "maxExtrapolations"); err != nil {
		return 0, err
	} else {
		return maxExtrapolations.Value, nil
	}
}

func (f FFIData) Call(i itypes.Interpreter) (itypes.Object, error) {
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
		return NewStreamFuncData(newStreamObject(), metricName, fltr, rollup, extrapolation, maxExtrapolations, 0), nil
	}
}

var _ = itypes.FlowCall(FFIData{})
