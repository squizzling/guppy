package stream

import (
	"fmt"
	"strings"

	"guppy/internal/interpreter"
)

func argsAggregate(i *interpreter.Interpreter) ([]interpreter.ParamData, error) {
	return []interpreter.ParamData{
		{Name: "self"},
		{Name: "by", Default: interpreter.NewObjectNone()},
	}, nil
}

func resolveBy(i *interpreter.Interpreter) ([]string, error) {
	if by, err := i.Scope.Get("by"); err != nil {
		return nil, err
	} else {
		switch by := by.(type) {
		case *interpreter.ObjectNone:
			return nil, nil // explicitly nil
		case *interpreter.ObjectString:
			s, _ := by.String(i)
			return []string{s}, nil
		case *interpreter.ObjectList:
			actualBy := make([]string, 0, len(by.Items)) // explicitly not nil
			for idx, item := range by.Items {
				if s, ok := item.(*interpreter.ObjectString); ok {
					s, _ := s.String(i)
					actualBy = append(actualBy, s)
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

func callAggregate(i *interpreter.Interpreter, new func(source Stream, by []string) Stream) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if by, err := resolveBy(i); err != nil {
		return nil, err
	} else {
		return new(self, by), nil
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
