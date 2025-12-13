package stream

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

// TODO: All of this.

type methodAbs struct {
	itypes.Object
}

func (m methodAbs) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
		},
	}, nil
}

func (m methodAbs) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else {
		return NewStreamMethodAbs(prototypeStreamDouble, unpublish(self)), nil
	}
}

func (ma methodAbs) Repr() string {
	// TODO: Better
	return "methodAbs()"
}

func (sma *StreamMethodAbs) Repr() string {
	// TODO: Better
	return ".abs()"
}
