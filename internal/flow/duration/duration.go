package duration

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"guppy/internal/interpreter"
)

type FFIDuration struct {
	interpreter.Object
}

func (f FFIDuration) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "duration"},
		},
	}, nil
}

func splitAndParseType(value string, typeSpecifier string, typeName string, full string, scale time.Duration) (time.Duration, string, error) {
	parts := strings.SplitN(value, typeSpecifier, 2)
	if len(parts) == 1 {
		return 0, parts[0], nil
	} else if val, err := strconv.Atoi(parts[0]); err != nil {
		return 0, "", fmt.Errorf("failed to decode %s in %s (%s): %w", typeName, full, parts[0], err)
	} else {
		return time.Duration(val) * scale, parts[1], nil
	}
}

func parseDurationType(value string, typeName string, full string, scale time.Duration) (time.Duration, error) {
	if v, err := strconv.Atoi(value); err != nil {
		return 0, fmt.Errorf("failed to decode %s in %s (%s): %w", typeName, full, value, err)
	} else {
		return time.Duration(v) * scale, nil
	}
}

var fmts = map[int]string{
	6: "week",
	5: "day",
	4: "hour",
	3: "minute",
	2: "second",
	1: "millisecond",
}

var scales = map[int]time.Duration{
	6: 7 * 24 * time.Hour,
	5: 24 * time.Hour,
	4: time.Hour,
	3: time.Minute,
	2: time.Second,
	1: time.Millisecond,
}

func ParseDuration(s string) (time.Duration, error) {
	// The rules I have discovered:
	// - Remove all whitespace (including \t and \n)
	// - Mixing terms is permitted
	// - Units must be in decreasing size (week 'w' > day 'd' > hour 'h' > minute 'm' > second 's' > millisecond 'ms')
	// - Integer only

	// Remove whitespace
	d := strings.NewReplacer(" ", "", "\n", "", "\t", "").Replace(s)

	haveDur := false
	parseDur := time.Duration(0)
	accumDur := time.Duration(0)
	maxUnit := 6

	for idx := 0; idx < len(d); idx++ {
		ch := d[idx]

		if ch >= '0' && ch <= '9' {
			parseDur = (parseDur * 10) + time.Duration(ch-'0')
			haveDur = true
			continue
		} else if !haveDur {
			return 0, fmt.Errorf("format specifier (%c) without value", ch)
		}

		formatSpecifier := 0
		if ch == 'w' {
			formatSpecifier = 6
		} else if ch == 'd' {
			formatSpecifier = 5
		} else if ch == 'h' {
			formatSpecifier = 4
		} else if ch == 'm' {
			if len(d)-1 > idx && d[idx+1] == 's' {
				formatSpecifier = 1
				idx++
			} else {
				formatSpecifier = 3
			}
		} else if ch == 's' {
			formatSpecifier = 2
		}

		if formatSpecifier == 0 {
			return 0, fmt.Errorf("format specifier (%c) not recognized", ch)
		}

		if formatSpecifier > maxUnit {
			return 0, fmt.Errorf("format specifier (%s) is higher than max unit (%s)", fmts[formatSpecifier], fmts[maxUnit])
		}

		accumDur += parseDur * scales[formatSpecifier]
		parseDur = 0
		maxUnit = formatSpecifier - 1
		haveDur = false
	}

	return accumDur, nil
}

func (f FFIDuration) resolveDuration(i *interpreter.Interpreter) (time.Duration, error) {
	if arg, err := i.Scope.GetArg("duration"); err != nil {
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

func (f FFIDuration) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if duration, err := f.resolveDuration(i); err != nil {
		return nil, err
	} else {
		return NewDuration(duration), nil
	}
}

var _ = interpreter.FlowCall(FFIDuration{})

type Duration struct {
	interpreter.Object

	Duration time.Duration
}

func NewDuration(d time.Duration) *Duration {
	return &Duration{
		Object: interpreter.NewObject(map[string]interpreter.Object{
			"__lt__": methodDurationOp{Object: interpreter.NewObject(nil), op: "<"},
			"__gt__": methodDurationOp{Object: interpreter.NewObject(nil), op: ">"},
			"__le__": methodDurationOp{Object: interpreter.NewObject(nil), op: "<="},
			"__ge__": methodDurationOp{Object: interpreter.NewObject(nil), op: ">="},
		}),
		Duration: d,
	}
}

func (d *Duration) String(i *interpreter.Interpreter) (string, error) {
	// TODO: Render in SFX units
	return d.Duration.String(), nil
}

type methodDurationOp struct {
	interpreter.Object

	op      string
	reverse string
}

func (mdu methodDurationOp) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return interpreter.BinaryParams, nil
}
func (mdu methodDurationOp) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else if reverseOp, err := right.Member(i, right, mdu.reverse); err == nil {
		// If it exists, we always use the reverse method, because it's more likely to be the intended behavior.
		// We explicitly don't expose reverse methods for primitives though.
		if reverseOpCall, ok := reverseOp.(interpreter.FlowCall); ok {
			return reverseOpCall.Call(i)
		}
	}

	if self, err := interpreter.ArgAs[*Duration](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		var rightVal time.Duration
		switch right := right.(type) {
		case *Duration:
			rightVal = right.Duration
		default:
			return nil, fmt.Errorf("methodDurationOp: unknown type %T", right)
		}

		switch mdu.op {
		case "+":
			return NewDuration(self.Duration + rightVal), nil
		case "-":
			return NewDuration(self.Duration - rightVal), nil
		case "/":
			return NewDuration(self.Duration / rightVal), nil
		case "*":
			return NewDuration(self.Duration * rightVal), nil
		case "<":
			return interpreter.NewObjectBool(self.Duration < rightVal), nil
		case ">":
			return interpreter.NewObjectBool(self.Duration > rightVal), nil
		case "<=":
			return interpreter.NewObjectBool(self.Duration <= rightVal), nil
		case ">=":
			return interpreter.NewObjectBool(self.Duration >= rightVal), nil
		default:
			return nil, fmt.Errorf("methodDurationOp: unknown op %s", mdu.op)
		}
	}
}
