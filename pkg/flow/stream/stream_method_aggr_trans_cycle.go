package stream

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/squizzling/guppy/pkg/flow/duration"
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/ftypes"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiStreamAggregateTransformCycleMethod struct {
	Self Stream `ffi:"self"`
	By   struct {
		None   *primitive.ObjectNone
		String *primitive.ObjectString
		List   *primitive.ObjectList
	} `ffi:"by"`
	AllowMissing struct {
		None   *primitive.ObjectNone
		Bool   *primitive.ObjectBool
		String *primitive.ObjectString
	} `ffi:"allow_missing"`
	Over struct {
		None     *primitive.ObjectNone
		String   *primitive.ObjectString
		Duration *duration.Duration
	} `ffi:"over"`
	Cycle struct {
		None   *primitive.ObjectNone
		String *primitive.ObjectString
	} `ffi:"cycle"`
	CycleStart struct {
		None   *primitive.ObjectNone
		String *primitive.ObjectString
	} `ffi:"cycle_start"`
	TimeZone struct {
		None   *primitive.ObjectNone
		String *primitive.ObjectString
	} `ffi:"timezone"`
	PartialValues struct {
		None   *primitive.ObjectNone
		Bool   *primitive.ObjectBool
		String *primitive.ObjectString
	} `ffi:"partial_values"`
	ShiftCycles ftypes.ThingOrNone[*primitive.ObjectInt] `ffi:"shift_cycles"`

	fn string
}

func NewFFIStreamAggregateTransformCycleMethod(fn string) itypes.FlowCall {
	// TODO: Make this less messy.  For now we have expanded types instead of ThingOrXyz because we
	//       don't know all the types.  Once it's fully resolved, we can make better types.
	return ffi.NewFFI(ffiStreamAggregateTransformCycleMethod{
		By: struct {
			None   *primitive.ObjectNone
			String *primitive.ObjectString
			List   *primitive.ObjectList
		}{None: primitive.NewObjectNone()},
		AllowMissing: struct {
			None   *primitive.ObjectNone
			Bool   *primitive.ObjectBool
			String *primitive.ObjectString
		}{None: primitive.NewObjectNone()},
		Over: struct {
			None     *primitive.ObjectNone
			String   *primitive.ObjectString
			Duration *duration.Duration
		}{None: primitive.NewObjectNone()},
		Cycle: struct {
			None   *primitive.ObjectNone
			String *primitive.ObjectString
		}{None: primitive.NewObjectNone()},
		CycleStart: struct {
			None   *primitive.ObjectNone
			String *primitive.ObjectString
		}{None: primitive.NewObjectNone()},
		TimeZone: struct {
			None   *primitive.ObjectNone
			String *primitive.ObjectString
		}{None: primitive.NewObjectNone()},
		PartialValues: struct {
			None   *primitive.ObjectNone
			Bool   *primitive.ObjectBool
			String *primitive.ObjectString
		}{None: primitive.NewObjectNone()},

		ShiftCycles: ftypes.NewThingOrNoneNone[*primitive.ObjectInt](),
		fn:          fn,
	})
}

func (f ffiStreamAggregateTransformCycleMethod) Call(i itypes.Interpreter) (itypes.Object, error) {
	groupBy := f.By.None == nil || f.AllowMissing.None == nil
	groupOver := f.Over.None == nil
	groupCycle := f.Cycle.None == nil || f.CycleStart.None == nil || f.TimeZone.None == nil || f.PartialValues.None == nil || f.ShiftCycles.None == nil
	switch {
	case !groupBy && !groupOver && !groupCycle:
		return f.callBy()
	case !groupBy && !groupOver && groupCycle:
		return f.callCycle()
	case !groupBy && groupOver && !groupCycle:
		return f.callOver()
	case !groupBy && groupOver && groupCycle:
		return nil, errors.New("ffiStreamAggregateTransformCycleMethod: over and cycle parameter groups provided")
	case groupBy && !groupOver && !groupCycle:
		return f.callBy()
	case groupBy && !groupOver && groupCycle:
		return nil, errors.New("ffiStreamAggregateTransformCycleMethod: by and cycle parameter groups provided")
	case groupBy && groupOver && !groupCycle:
		return nil, errors.New("ffiStreamAggregateTransformCycleMethod: by and over parameter groups provided")
	case groupBy && groupOver && groupCycle:
		return nil, errors.New("ffiStreamAggregateTransformCycleMethod: by, over, and cycle parameter groups provided")
	}
	panic("unreachable")
}

func (f ffiStreamAggregateTransformCycleMethod) resolveBy() ([]string, error) {
	// TODO: Can it take a tuple of strings?
	if f.By.String != nil {
		return []string{f.By.String.Value}, nil
	} else if f.By.List != nil {
		var items []string
		for idx, item := range f.By.List.Items {
			if byItem, ok := item.(*primitive.ObjectString); !ok {
				return nil, fmt.Errorf("ffiStreamAggregateTransformCycleMethod.resolveBy: index %d of by contains a %T not a *primitive.ObjectString", idx, item)
			} else {
				items = append(items, byItem.Value)
			}
		}
		return items, nil
	} else {
		return nil, nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) resolveAllowMissing() (bool, []string, error) {
	// TODO: It can take a list of strings, which returns false, []list, nil
	// TODO: Can it take a tuple?
	// TODO: If allowmissing is 'True', is that allow_missing=True, or allow_missing=['true']?
	if f.AllowMissing.None != nil {
		return false, nil, nil
	} else if f.AllowMissing.Bool != nil {
		return f.AllowMissing.Bool.Value, nil, nil
	} else {
		return false, []string{f.AllowMissing.String.Value}, nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) callBy() (itypes.Object, error) {
	if by, err := f.resolveBy(); err != nil {
		return nil, err
	} else if allowMissing, allowMissingValues, err := f.resolveAllowMissing(); err != nil {
		return nil, err
	} else {
		return NewStreamMethodAggregate(
			prototypeStreamDouble,
			f.Self,
			f.fn,
			by,
			allowMissing,
			allowMissingValues,
		), nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) resolveCycle() (string, error) {
	switch f.Cycle.String.Value {
	case "hour":
		return "hour", nil
	case "day":
		return "day", nil
	case "week":
		return "week", nil
	case "month":
		return "month", nil
	case "quarter":
		return "quarter", nil
	default:
		return "", errors.New("ffiStreamAggregateTransformCycleMethod.resolveCycle: param `cycle` is not [hour, day, week, month, quarter]")
	}
}

func (f ffiStreamAggregateTransformCycleMethod) resolveCycleStart() (*string, error) {
	if f.CycleStart.String != nil {
		return &f.CycleStart.String.Value, nil
	} else {
		return nil, nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) resolveTimezone() (*string, error) {
	if f.TimeZone.String != nil {
		return &f.TimeZone.String.Value, nil
	} else {
		return nil, nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) resolvePartialValues() (bool, error) {
	if f.PartialValues.Bool != nil {
		return f.PartialValues.Bool.Value, nil
	} else if f.PartialValues.String != nil {
		return strings.EqualFold(f.PartialValues.String.Value, "true"), nil
	} else {
		return false, nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) resolveShiftCycles() (int, error) {
	if f.ShiftCycles.Thing != nil {
		return f.ShiftCycles.Thing.Value, nil
	} else {
		return 0, nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) callCycle() (itypes.Object, error) {
	if cycle, err := f.resolveCycle(); err != nil {
		return nil, err
	} else if cycleStart, err := f.resolveCycleStart(); err != nil {
		return nil, err
	} else if timeZone, err := f.resolveTimezone(); err != nil {
		return nil, err
	} else if partialValues, err := f.resolvePartialValues(); err != nil {
		return nil, err
	} else if shiftCycles, err := f.resolveShiftCycles(); err != nil {
		return nil, err
	} else {
		return NewStreamMethodTransformCycle(
			prototypeStreamDouble,
			f.Self,
			f.fn,
			cycle,
			cycleStart,
			timeZone,
			partialValues,
			shiftCycles,
		), nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) resolveOver() (time.Duration, error) {
	if f.Over.Duration != nil {
		return f.Over.Duration.Duration, nil
	} else if d, err := duration.ParseDuration(f.Over.String.Value); err != nil {
		return 0, err
	} else {
		return d, nil
	}
}

func (f ffiStreamAggregateTransformCycleMethod) callOver() (itypes.Object, error) {
	if over, err := f.resolveOver(); err != nil {
		return nil, err
	} else {
		return NewStreamMethodTransform(
			prototypeStreamDouble,
			f.Self,
			f.fn,
			over,
		), nil
	}
}
