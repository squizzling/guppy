package stream

import (
	"fmt"
	"time"

	"github.com/squizzling/guppy/pkg/flow/duration"
	"github.com/squizzling/guppy/pkg/flow/filter"
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/ftypes"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiData struct {
	Metric            *primitive.ObjectString                     `ffi:"metric"`
	Filter            ftypes.ThingOrNone[filter.Filter]           `ffi:"filter"`
	Rollup            ftypes.ThingOrNone[*primitive.ObjectString] `ffi:"rollup"`
	Extrapolation     *primitive.ObjectString                     `ffi:"extrapolation"`
	MaxExtrapolations *primitive.ObjectInt                        `ffi:"maxExtrapolations"`
	Resolution        struct {
		None     *primitive.ObjectNone
		Duration *duration.Duration
		Int      *primitive.ObjectInt
		String   *primitive.ObjectString
	} `ffi:"resolution"`
}

func NewFFIData() itypes.FlowCall {
	return ffi.NewFFI(ffiData{
		Filter:            ftypes.NewThingOrNoneNone[filter.Filter](),
		Rollup:            ftypes.NewThingOrNoneNone[*primitive.ObjectString](),
		Extrapolation:     primitive.NewObjectString("null"),
		MaxExtrapolations: primitive.NewObjectInt(-1),
		Resolution: struct {
			None     *primitive.ObjectNone
			Duration *duration.Duration
			Int      *primitive.ObjectInt
			String   *primitive.ObjectString
		}{
			None: primitive.NewObjectNone(),
		},
	})
}

func (f ffiData) Call(i itypes.Interpreter) (itypes.Object, error) {
	if filter, err := f.resolveFilter(); err != nil {
		return nil, err
	} else if rollup, err := f.resolveRollup(); err != nil {
		return nil, err
	} else if extrapolation, err := f.resolveExtrapolation(); err != nil {
		return nil, err
	} else if maxExtrapolations, err := f.resolveMaxExtrapolations(); err != nil {
		return nil, err
	} else if resolution, err := f.resolveResolution(); err != nil {
		return nil, err
	} else {
		return NewStreamFuncData(
			prototypeStreamDouble,
			f.Metric.Value,
			filter,
			rollup,
			extrapolation,
			maxExtrapolations,
			resolution,
			0,
		), nil
	}
}

func (f ffiData) resolveFilter() (filter.Filter, error) {
	return f.Filter.Thing, nil
}

func (f ffiData) resolveRollup() (string, error) {
	if f.Rollup.None != nil {
		return "", nil
	} else {
		return f.Rollup.Thing.Value, nil
	}
}

func (f ffiData) resolveExtrapolation() (string, error) {
	switch f.Extrapolation.Value {
	case "null", "last", "last_value", "zero":
		return f.Extrapolation.Value, nil
	default:
		return "", fmt.Errorf("ffiData.resolveExtrapolation: param `extrapolation` is %s, expected [null, last, last_value, zero]", f.Extrapolation.Value)
	}
}

func (f ffiData) resolveMaxExtrapolations() (int, error) {
	return f.MaxExtrapolations.Value, nil
}

func (f ffiData) resolveResolution() (*time.Duration, error) {
	if f.Resolution.None != nil {
		return nil, nil
	} else if f.Resolution.Int != nil {
		d := time.Duration(f.Resolution.Int.Value) * time.Millisecond
		return &d, nil
	} else if f.Resolution.Duration != nil {
		return &f.Resolution.Duration.Duration, nil
	} else if d, err := duration.ParseDuration(f.Resolution.String.Value); err != nil {
		return nil, fmt.Errorf("ffiData.resolveResolution: param `resolution` failed to parse: %w", err)
	} else {
		return &d, nil
	}
}

func (sfd StreamFuncData) Repr() string {
	// TODO: Better
	return "data()"
}
