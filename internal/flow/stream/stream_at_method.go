package stream

import (
	"fmt"

	"guppy/internal/interpreter"
)

type methodStreamAggregateTransform struct {
	interpreter.Object

	name string
}

func (msat methodStreamAggregateTransform) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "by", Default: interpreter.NewObjectNone()},
			{Name: "over", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

func resolveBy(i *interpreter.Interpreter) ([]string, error) {
	if by, err := i.Scope.GetArg("by"); err != nil {
		return nil, err
	} else {
		switch by := by.(type) {
		case *interpreter.ObjectNone:
			return nil, nil // explicitly nil
		case *interpreter.ObjectString:
			return []string{by.Value}, nil
		case *interpreter.ObjectList:
			actualBy := make([]string, 0, len(by.Items)) // explicitly not nil
			for idx, item := range by.Items {
				if s, ok := item.(*interpreter.ObjectString); ok {
					actualBy = append(actualBy, s.Value)
				} else {
					return nil, fmt.Errorf("by element %d is %T not *interpreter.ObjectString", idx, item)
				}
			}
			return actualBy, nil
		default:
			return nil, fmt.Errorf("by is %T not *interpreter.ObjectNone, *interpreter.ObjectString, or *interpreter.ObjectList", by)
		}
	}
}

func resolveOver(i *interpreter.Interpreter) (*string, error) {
	if over, err := i.Scope.GetArg("over"); err != nil {
		return nil, err
	} else {
		switch over := over.(type) {
		case *interpreter.ObjectNone:
			return nil, nil // explicitly nil
		case *interpreter.ObjectString:
			s := over.Value
			return &s, nil
		default:
			return nil, fmt.Errorf("over is %T not *interpreter.ObjectNone, or *interpreter.ObjectString", over)
		}
	}
}

func (msat methodStreamAggregateTransform) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if by, err := resolveBy(i); err != nil {
		return nil, err
	} else if over, err := resolveOver(i); err != nil {
		return nil, err
	} else if over != nil && len(by) > 0 {
		return nil, fmt.Errorf("only one argument of [by, over] may be specified")
	} else if over != nil {
		return newStreamTransform(self, msat.name, *over), nil
	} else {
		return newStreamAggregate(self, msat.name, by), nil
	}
}
