package stream

import (
	"fmt"
	"strings"

	"guppy/internal/interpreter"
)

type FFIUnion struct {
	interpreter.Object
}

func (f FFIUnion) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		StarParam: "streams",
	}, nil
}

func (f FFIUnion) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if streamsRaw, err := interpreter.ArgAs[*interpreter.ObjectTuple](i, "streams"); err != nil {
		return nil, err
	} else {
		var streams []Stream
		for idx, streamRaw := range streamsRaw.Items {
			if stream, ok := streamRaw.(Stream); ok {
				streams = append(streams, stream)
			} else {
				return nil, fmt.Errorf("argument %d is %T not a Stream", idx, stream)
			}
		}
		return NewUnion(streams), nil
	}
}

var _ = interpreter.FlowCall(FFIData{})

type union struct {
	interpreter.Object

	streams []Stream
}

func NewUnion(streams []Stream) Stream {
	for idx, stream := range streams {
		streams[idx] = unpublish(stream)
	}
	return &union{
		Object: newStreamObject(),

		streams: streams,
	}
}

func (u *union) RenderStream() string {
	var sb strings.Builder
	sb.WriteString("union(")
	for idx, s := range u.streams {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(s.RenderStream())
	}
	sb.WriteString(")")
	return sb.String()
}
