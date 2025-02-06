package stream

import (
	"fmt"

	"github.com/squizzling/types/pkg/result"

	"guppy/internal/interpreter"
)

type methodPublish struct {
	interpreter.Object
}

func (mp methodPublish) Args(i *interpreter.Interpreter) result.Result[[]interpreter.ArgData] {
	return result.Ok([]interpreter.ArgData{
		{Name: "self"},
		{Name: "label"},
	})
}

func (mp methodPublish) Call(i *interpreter.Interpreter) result.Result[interpreter.Object] {
	if resultSelf := interpreter.ArgAs[Stream](i, "self"); !resultSelf.Ok() {
		return result.Err[interpreter.Object](resultSelf.Err())
	} else if resultLabel := interpreter.ArgAsString(i, "label"); !resultLabel.Ok() {
		return result.Err[interpreter.Object](resultLabel.Err())
	} else {
		return result.Ok[interpreter.Object](NewPublish(resultSelf.Value(), resultLabel.Value()))
	}
}

type publish struct {
	interpreter.Object

	source Stream
	label  string
}

func NewPublish(source Stream, label string) Stream {
	s := &publish{
		Object: newStreamObject(),
		source: source,
		label:  label,
	}
	fmt.Printf("%s\n", s.RenderStream())
	return s
}

func (p *publish) RenderStream() string {
	return fmt.Sprintf("%s.publish(label='%s')", p.source.RenderStream(), p.label)
}
