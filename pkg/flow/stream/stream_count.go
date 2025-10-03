package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFICount struct {
	interpreter.Object
}

func (f FFICount) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		StarParam: "streams",
	}, nil
}

func (f FFICount) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if streamsRaw, err := interpreter.ArgAs[*interpreter.ObjectTuple](i, "streams"); err != nil {
		return nil, err
	} else {
		var streams []Stream
		for idx, streamRaw := range streamsRaw.Items {
			if stream, ok := streamRaw.(Stream); ok {
				streams = append(streams, unpublish(stream))
			} else {
				return nil, fmt.Errorf("argument %d is %T not a Stream", idx, stream)
			}
		}
		return NewStreamCount(newStreamObject(), streams), nil
	}
}

var _ = interpreter.FlowCall(FFICount{})
