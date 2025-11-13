package stream

import (
	"fmt"
	"time"

	"guppy/pkg/flow/duration"
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type methodStreamAggregateTransform struct {
	itypes.Object

	name string
}

func (msat methodStreamAggregateTransform) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "by", Default: primitive.NewObjectNone()},             // by
			{Name: "allow_missing", Default: primitive.NewObjectNone()},  // by
			{Name: "over", Default: primitive.NewObjectNone()},           // over
			{Name: "cycle", Default: primitive.NewObjectNone()},          // cycle
			{Name: "cycle_start", Default: primitive.NewObjectNone()},    // cycle
			{Name: "timezone", Default: primitive.NewObjectNone()},       // cycle
			{Name: "partial_values", Default: primitive.NewObjectNone()}, // cycle
			{Name: "shift_cycles", Default: primitive.NewObjectNone()},   // cycle
		},
	}, nil
}

func (msat methodStreamAggregateTransform) resolveBy(i itypes.Interpreter) ([]string, error) {
	if by, err := i.GetArg("by"); err != nil {
		return nil, err
	} else {
		switch by := by.(type) {
		case *primitive.ObjectNone:
			return nil, nil // explicitly nil
		case *interpreter.ObjectString:
			return []string{by.Value}, nil
		case *interpreter.ObjectList:
			actualBy := make([]string, 0, len(by.Items)) // explicitly not nil
			for idx, item := range by.Items {
				if s, ok := item.(*interpreter.ObjectString); ok {
					actualBy = append(actualBy, s.Value)
				} else {
					return nil, fmt.Errorf("methodStreamAggregateTransform(by) element %d is %T not *interpreter.ObjectString", idx, item)
				}
			}
			return actualBy, nil
		default:
			return nil, fmt.Errorf("methodStreamAggregateTransform(by) is %T not *interpreter.ObjectNone, *interpreter.ObjectString, or *interpreter.ObjectList", by)
		}
	}
}

func (msat methodStreamAggregateTransform) resolveAllowMissing(i itypes.Interpreter) (bool, []string, error) {
	if allowMissing, err := i.GetArg("allow_missing"); err != nil {
		return false, nil, err
	} else {
		switch allowMissing := allowMissing.(type) {
		case *primitive.ObjectNone:
			return false, nil, nil
		case *primitive.ObjectBool:
			if allowMissing.Value {
				return true, nil, nil
			} else {
				return false, nil, nil
			}
		case *interpreter.ObjectString:
			return false, []string{allowMissing.Value}, nil
		case *interpreter.ObjectList:
			actualAllowMissing := make([]string, 0, len(allowMissing.Items))
			for idx, item := range allowMissing.Items {
				if s, ok := item.(*interpreter.ObjectString); ok {
					actualAllowMissing = append(actualAllowMissing, s.Value)
				} else {
					return false, nil, fmt.Errorf("methodStreamAggregateTransform(allowMissing) element %d is %T not *interpreter.ObjectString", idx, item)
				}
			}
			return false, actualAllowMissing, nil

		default:
			return false, nil, fmt.Errorf("methodStreamAggregateTransform(allowMissing) is %T not *interpreter.ObjectNone, *interpreter.ObjectList, *interpreter.ObjectString, or *interpreter.ObjectBool", allowMissing)
		}
	}
}

func (msat methodStreamAggregateTransform) resolveOver(i itypes.Interpreter) (*time.Duration, error) {
	if over, err := i.GetArg("over"); err != nil {
		return nil, err
	} else {
		switch over := over.(type) {
		case *primitive.ObjectNone:
			return nil, nil
		case *interpreter.ObjectString:
			if d, err := duration.ParseDuration(over.Value); err != nil {
				return nil, err
			} else {
				return &d, nil
			}
		case *duration.Duration:
			return &over.Duration, nil
		default:
			return nil, fmt.Errorf("methodStreamAggregateTransform(over) is %T not *flow.Duration, *interpreter.ObjectNone, or *interpreter.ObjectString", over)
		}
	}
}

func (msat methodStreamAggregateTransform) resolveCycle(i itypes.Interpreter) (*string, error) {
	if cycle, err := i.GetArg("cycle"); err != nil {
		return nil, err
	} else {
		switch over := cycle.(type) {
		case *primitive.ObjectNone:
			return nil, nil
		case *interpreter.ObjectString:
			if over.Value == "hour" || over.Value == "week" || over.Value == "month" || over.Value == "day" || over.Value == "quarter" {
				return &over.Value, nil
			} else {
				return nil, fmt.Errorf("methodStreamAggregateTransform(cycle) is %s not [hour, week, month, day, quarter]", over.Value)
			}
		default:
			return nil, fmt.Errorf("methodStreamAggregateTransform(cycle) is %T not *interpreter.ObjectNone, or *interpreter.ObjectString", over)
		}
	}
}

func (msat methodStreamAggregateTransform) resolveCycleStart(i itypes.Interpreter) (*string, error) {
	if cycleStart, err := i.GetArg("cycle_start"); err != nil {
		return nil, err
	} else {
		switch cycleStart := cycleStart.(type) {
		case *primitive.ObjectNone:
			return nil, nil
		case *interpreter.ObjectString:
			return &cycleStart.Value, nil
		default:
			return nil, fmt.Errorf("methodStreamAggregateTransform(cycleStart) is %T not *interpreter.ObjectNone, or *interpreter.ObjectString", cycleStart)
		}
	}
}

func (msat methodStreamAggregateTransform) resolveTimezone(i itypes.Interpreter) (*string, error) {
	if timezone, err := i.GetArg("timezone"); err != nil {
		return nil, err
	} else {
		switch timezone := timezone.(type) {
		case *primitive.ObjectNone:
			return nil, nil
		case *interpreter.ObjectString:
			return &timezone.Value, nil
		default:
			return nil, fmt.Errorf("methodStreamAggregateTransform(timezone) is %T not *interpreter.ObjectNone, or *interpreter.ObjectString", timezone)
		}
	}
}

func (msat methodStreamAggregateTransform) resolvePartialValues(i itypes.Interpreter) (bool, error) {
	if partialValues, err := i.GetArg("partial_values"); err != nil {
		return false, err
	} else {
		switch partialValues := partialValues.(type) {
		case *primitive.ObjectNone:
			return false, nil
		case *primitive.ObjectBool:
			return partialValues.Value, nil
		default:
			return false, fmt.Errorf("methodStreamAggregateTransform(partialValues) is %T not *interpreter.ObjectNone, or *interpreter.ObjectBool", partialValues)
		}
	}
}

func (msat methodStreamAggregateTransform) resolveShiftCycles(i itypes.Interpreter) (int, error) {
	if shiftCycles, err := i.GetArg("shift_cycles"); err != nil {
		return 0, err
	} else {
		switch shiftCycles := shiftCycles.(type) {
		case *primitive.ObjectNone:
			return 0, nil
		case *primitive.ObjectInt:
			return shiftCycles.Value, nil
		default:
			return 0, fmt.Errorf("methodStreamAggregateTransform(shiftCycles) is %T not *interpreter.ObjectNone, or *interpreter.ObjectBool", shiftCycles)
		}
	}
}

func (msat methodStreamAggregateTransform) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if by, err := msat.resolveBy(i); err != nil {
		return nil, err
	} else if allowAllMissing, allowMissing, err := msat.resolveAllowMissing(i); err != nil {
		return nil, err
	} else if over, err := msat.resolveOver(i); err != nil {
		return nil, err
	} else if cycle, err := msat.resolveCycle(i); err != nil {
		return nil, err
	} else if cycleStart, err := msat.resolveCycleStart(i); err != nil {
		return nil, err
	} else if timezone, err := msat.resolveTimezone(i); err != nil {
		return nil, err
	} else if partialValues, err := msat.resolvePartialValues(i); err != nil {
		return nil, err
	} else if shiftCycles, err := msat.resolveShiftCycles(i); err != nil {
		return nil, err
	} else if (over != nil && len(by) > 0) || (over != nil && cycle != nil) || (len(by) > 0 && cycle != nil) {
		return nil, fmt.Errorf("only one argument of [by, over, cycle] may be specified")
	} else if over != nil {
		return NewStreamMethodTransform(newStreamObject(), unpublish(self), msat.name, *over), nil
	} else if cycle != nil {
		return NewStreamMethodTransformCycle(newStreamObject(), unpublish(self), msat.name, *cycle, cycleStart, timezone, partialValues, shiftCycles), nil
	} else {
		return NewStreamMethodAggregate(newStreamObject(), unpublish(self), msat.name, by, allowAllMissing, allowMissing), nil
	}
}
