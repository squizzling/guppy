package duration

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIDuration struct {
	itypes.Object
}

func (f FFIDuration) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "duration"},
		},
	}, nil
}

var unitScale = map[string]time.Duration{
	"w":   7 * 24 * time.Hour,
	"d":   24 * time.Hour,
	"h":   time.Hour,
	"m":   time.Minute,
	"min": time.Minute,
	"s":   time.Second,
	"ms":  time.Millisecond,
}

var unitRank = map[string]int{
	"w":   6,
	"d":   5,
	"h":   4,
	"m":   3,
	"min": 3,
	"s":   2,
	"ms":  1,
}

var rankUnit = map[int]string{
	6: "w",
	5: "d",
	4: "h",
	3: "m",
	2: "s",
	1: "ms",
}

func ParseDuration(s string) (time.Duration, error) {
	// The rules I have discovered:
	// - Remove all whitespace (including \t and \n)
	// - Mixing terms is permitted
	// - Units must be in decreasing size (week 'w' > day 'd' > hour 'h' > minute 'm' > second 's' > millisecond 'ms')
	// - Integer only

	// Remove whitespace
	d := strings.NewReplacer(" ", "", "\n", "", "\t", "").Replace(s)

	var values []time.Duration
	var units []string

	value := time.Duration(0)
	var unit strings.Builder
	readValue := true

	if len(d) == 0 {
		return 0, errors.New("empty duration")
	} else if d[0] < '0' || d[0] > '9' { // Make sure it starts with a digit, because it simplifies everything the main loop
		return 0, fmt.Errorf("duration without value: %s", d)
	}

	for _, ch := range d {
		if readValue {
			if ch >= '0' && ch <= '9' {
				value = (value * 10) + time.Duration(ch-'0')
			} else {
				values = append(values, value) // Store the value
				value = 0                      // Reset the value for next time
				readValue = false              // Switch to reading units
				unit.WriteRune(ch)             // Store the first unit
			}
		} else {
			if ch >= '0' && ch <= '9' {
				units = append(units, unit.String()) // Store the unit
				unit.Reset()                         // Reset the unit for next time
				readValue = true                     // Switch to reading values
				value = time.Duration(ch - '0')      // Store the first value digit
			} else {
				unit.WriteRune(ch)
			}
		}
	}

	if readValue {
		// If we're reading a value, then we don't have a unit
		return 0, fmt.Errorf("duration without unit: %s", d)
	}
	units = append(units, unit.String())

	maxRank := 6
	accumDur := time.Duration(0)
	for i, value := range values {
		unit := units[i]

		rank, ok := unitRank[unit]
		if !ok {
			return 0, fmt.Errorf("unit %s not recognized in %s", unit, d)
		}
		if rank > maxRank {
			return 0, fmt.Errorf("unit %s is higher than max unit (%s)", unit, rankUnit[maxRank])
		}

		accumDur += value * unitScale[unit]
		maxRank = rank - 1
	}

	return accumDur, nil
}

func (f FFIDuration) resolveDuration(i itypes.Interpreter) (time.Duration, error) {
	if arg, err := i.GetArg("duration"); err != nil {
		return 0, err
	} else {
		switch arg := arg.(type) {
		case *interpreter.ObjectInt: // milliseconds
			return time.Duration(arg.Value) * time.Millisecond, nil
		case *interpreter.ObjectString:
			return ParseDuration(arg.Value)
		case *Duration:
			return arg.Duration, nil
		default:
			return 0, fmt.Errorf("arg duration is %T not *interpreter.ObjectInt, *interpreter.ObjectString, or *flow.Duration", arg)
		}
	}
}

func (f FFIDuration) Call(i itypes.Interpreter) (itypes.Object, error) {
	if duration, err := f.resolveDuration(i); err != nil {
		return nil, err
	} else {
		return NewDuration(duration), nil
	}
}

var _ = interpreter.FlowCall(FFIDuration{})

type Duration struct {
	itypes.Object

	Duration time.Duration
}

func NewDuration(d time.Duration) *Duration {
	return &Duration{
		Object: itypes.NewObject(map[string]itypes.Object{
			"__rmul__": methodDurationOp{Object: itypes.NewObject(nil), op: "*", reverse: true},

			"__lt__": methodDurationOp{Object: itypes.NewObject(nil), op: "<"},
			"__gt__": methodDurationOp{Object: itypes.NewObject(nil), op: ">"},
			"__le__": methodDurationOp{Object: itypes.NewObject(nil), op: "<="},
			"__ge__": methodDurationOp{Object: itypes.NewObject(nil), op: ">="},
		}),
		Duration: d,
	}
}

func (d *Duration) String(i itypes.Interpreter) (string, error) {
	// TODO: Render in SFX units
	return d.Duration.String(), nil
}

type methodDurationOp struct {
	itypes.Object

	op      string
	reverse bool
}

func (mdu methodDurationOp) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return interpreter.BinaryParams, nil
}
func (mdu methodDurationOp) Call(i itypes.Interpreter) (itypes.Object, error) {
	// TODO: I don't love this, but I'm being lazy right now.
	// Proper fix is to not use ArgAs[*Duration]
	selfArg := "self"
	rightArg := "right"
	if mdu.reverse {
		selfArg, rightArg = rightArg, selfArg
	}

	if self, err := interpreter.ArgAs[*Duration](i, selfArg); err != nil {
		return nil, err
	} else if right, err := i.GetArg(rightArg); err != nil {
		return nil, err
	} else {
		var rightVal time.Duration
		switch right := right.(type) {
		case *Duration: // TODO: Look in to what it means for duration('2d') * duration('3d') in SFX
			rightVal = right.Duration
		case *interpreter.ObjectInt:
			rightVal = time.Duration(right.Value)
		default:
			return nil, fmt.Errorf("methodDurationOp: unknown type %T", right)
		}

		switch mdu.op {
		case "+":
			return NewDuration(self.Duration + rightVal), nil
		case "-":
			if mdu.reverse {
				return NewDuration(rightVal - self.Duration), nil
			} else {
				return NewDuration(self.Duration - rightVal), nil
			}
		case "/":
			if mdu.reverse {
				return NewDuration(rightVal / self.Duration), nil
			} else {
				return NewDuration(self.Duration / rightVal), nil
			}
		case "*":
			return NewDuration(self.Duration * rightVal), nil
		case "<":
			if mdu.reverse {
				panic("handle this")
			}
			return interpreter.NewObjectBool(self.Duration < rightVal), nil
		case ">":
			if mdu.reverse {
				panic("handle this")
			}
			return interpreter.NewObjectBool(self.Duration > rightVal), nil
		case "<=":
			if mdu.reverse {
				panic("handle this")
			}
			return interpreter.NewObjectBool(self.Duration <= rightVal), nil
		case ">=":
			if mdu.reverse {
				panic("handle this")
			}
			return interpreter.NewObjectBool(self.Duration >= rightVal), nil
		default:
			return nil, fmt.Errorf("methodDurationOp: unknown op %s", mdu.op)
		}
	}
}
