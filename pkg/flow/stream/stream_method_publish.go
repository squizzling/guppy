package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type methodPublish struct {
	interpreter.Object
}

func (mp methodPublish) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "self"},
			{Name: "label", Default: interpreter.NewObjectString("")}, // TODO: Validate "" vs None
			{Name: "enable", Default: interpreter.NewObjectBool(true)},
		},
		//KWParam: "additional_dimensions", // Maybe, I don't fully know this one.
	}, nil
}

func (mp methodPublish) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if self, err := interpreter.ArgAs[Stream](i, "self"); err != nil {
		return nil, err
	} else if label, err := interpreter.ArgAsString(i, "label"); err != nil {
		return nil, err
	} else if enable, err := interpreter.ArgAsBool(i, "enable"); err != nil {
		return nil, err
	} else {
		// TODO: This whole thing is a hack to expose published data
		pub := NewStreamMethodPublish(newStreamObject(), unpublish(self), label, enable)
		if rawPublished, err := i.Globals.Get("_published"); err != nil {
			return nil, err
		} else if published, ok := rawPublished.(interface{ Append(s *StreamMethodPublish) }); !ok {
			return nil, fmt.Errorf("invalid type")
		} else {
			published.Append(pub)
		}
		return pub, nil
	}
}
