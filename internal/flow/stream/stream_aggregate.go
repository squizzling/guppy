package stream

import (
	"fmt"
	"strings"

	"guppy/internal/interpreter"
)

// argsAggregate returns an argument list that supports `by` but not `over`
func argsAggregate(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "by", Default: interpreter.NewObjectNone()},
		},
	}, nil
}

// argsAggregate returns an argument list that supports `by` and `over`
func argsAggregateOver(i *interpreter.Interpreter) (*interpreter.Params, error) {
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

func callAggregate(i *interpreter.Interpreter, new func(source Stream, by []string) Stream) (interpreter.Object, error) {
	// TODO: Not sure if this will ever be used, but it might, histogram_percentile maybe?
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if by, err := resolveBy(i); err != nil {
		return nil, err
	} else {
		return new(self, by), nil
	}
}

func callAggregateOver(
	i *interpreter.Interpreter,
	newAggregate func(source Stream, by []string) Stream,
	newTransform func(source Stream, over string) Stream,
) (
	interpreter.Object,
	error,
) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if by, err := resolveBy(i); err != nil {
		return nil, err
	} else if over, err := resolveOver(i); err != nil {
		return nil, err
	} else if over != nil && len(by) > 0 {
		return nil, fmt.Errorf("only one argument of [by, over] may be specified")
	} else if over != nil {
		return newTransform(self, *over), nil
	} else {
		return newAggregate(self, by), nil
	}
}

func renderAggregate(source Stream, aggr string, by []string) string {
	var bys []string
	if by == nil {
		return fmt.Sprintf("%s.%s()", source.RenderStream(), aggr)
	}
	for _, by := range by {
		bys = append(bys, "'"+by+"'")
	}
	return fmt.Sprintf("%s.%s(by=[%s])", source.RenderStream(), aggr, strings.Join(bys, ", "))
}

func renderTransform(source Stream, aggr string, over string) string {
	return fmt.Sprintf("%s.%s(over='%s')", source.RenderStream(), aggr, over)
}
