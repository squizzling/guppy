package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFICount struct {
	itypes.Object
}

func (f FFICount) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "streams",
	}, nil
}

func (f FFICount) Call(i itypes.Interpreter) (itypes.Object, error) {
	if streamsRaw, err := itypes.ArgAs[*interpreter.ObjectTuple](i, "streams"); err != nil {
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
		return NewStreamFuncCount(newStreamObject(), streams), nil
	}
}

var _ = interpreter.FlowCall(FFICount{})
