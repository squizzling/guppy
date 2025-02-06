package stream

import (
	"fmt"
	"strings"

	"github.com/squizzling/types/pkg/result"

	"guppy/internal/interpreter"
)

func argsAggregate(i *interpreter.Interpreter) result.Result[[]interpreter.ArgData] {
	return result.Ok([]interpreter.ArgData{
		{Name: "self"},
		{Name: "by", Default: interpreter.NewObjectNone()},
	})
}

func callAggregate(i *interpreter.Interpreter, new func(source Stream, by []string) Stream) result.Result[interpreter.Object] {
	if resultSelf := interpreter.ArgAs[Stream](i, "self"); !resultSelf.Ok() {
		return result.Err[interpreter.Object](resultSelf.Err())
	} else if resultBy := i.Scope.Get("by"); !resultBy.Ok() {
		return resultBy
	} else {
		var actualBy []string
		switch by := resultBy.Value().(type) {
		case *interpreter.ObjectNone:
			actualBy = nil // explicitly nil
		case *interpreter.ObjectString:
			actualBy = []string{by.String(i).Value()}
		case *interpreter.ObjectList:
			actualBy = make([]string, 0, len(by.Items)) // explicitly not nil
			for idx, item := range by.Items {
				if s, ok := item.(*interpreter.ObjectString); ok {
					actualBy = append(actualBy, s.String(i).Value())
				} else {
					return result.Err[interpreter.Object](fmt.Errorf("by element %d is %T not *interpreter.ObjectString", idx, item))
				}
			}
		default:
			return result.Err[interpreter.Object](fmt.Errorf("by is %T not *interpreter.ObjectNone, *interpreter.ObjectString, or *interpreter.ObjectList", resultBy.Value()))
		}
		return result.Ok[interpreter.Object](new(resultSelf.Value(), actualBy))
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
