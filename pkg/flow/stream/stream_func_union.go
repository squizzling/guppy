package stream

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type FFIUnion struct {
	itypes.Object
}

func (f FFIUnion) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "streams",
	}, nil
}

func (f FFIUnion) Call(i itypes.Interpreter) (itypes.Object, error) {
	if streamsRaw, err := itypes.ArgAs[*primitive.ObjectTuple](i, "streams"); err != nil {
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
		return NewStreamFuncUnion(newStreamObject(), streams), nil
	}
}

var _ = itypes.FlowCall(FFIUnion{})
