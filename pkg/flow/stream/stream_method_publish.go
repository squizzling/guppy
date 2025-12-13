package stream

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type methodPublish struct {
	itypes.Object
}

func (mp methodPublish) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "label", Default: primitive.NewObjectString("")}, // TODO: Validate "" vs None
			{Name: "enable", Default: primitive.NewObjectBool(true)},
		},
		//KWParam: "additional_dimensions", // Maybe, I don't fully know this one.
	}, nil
}

func (mp methodPublish) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if label, err := interpreter.ArgAsString(i, "label"); err != nil {
		return nil, err
	} else if enable, err := itypes.ArgAs[*primitive.ObjectBool](i, "enable"); err != nil {
		return nil, err
	} else {
		// TODO: This whole thing is a hack to expose published data
		pub := NewStreamMethodPublish(newStreamObject(), unpublish(self), label, enable.Value)
		if rawPublished, err := i.GetGlobal("_published"); err != nil {
			return nil, err
		} else if published, ok := rawPublished.(interface{ Append(s *StreamMethodPublish) }); !ok {
			return nil, fmt.Errorf("invalid type")
		} else {
			published.Append(pub)
		}
		return pub, nil
	}
}

func (smp *StreamMethodPublish) Repr() string {
	// TODO: Better
	return ".publish()"
}
