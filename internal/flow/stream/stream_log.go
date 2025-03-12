package stream

import (
	"guppy/internal/interpreter"
)

// TODO: All of this.

type methodLog struct {
	interpreter.Object
}

func (m methodLog) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodLog) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewLog(self), nil
	}
}

type Log struct {
	interpreter.Object

	source Stream
}

func NewLog(source Stream) Stream {
	return &Log{
		Object: newStreamObject(),
		source: unpublish(source),
	}
}

func (s *Log) RenderStream() string {
	return s.source.RenderStream() + ".log()"
}
