package debug

import (
	"fmt"
	"strings"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type FFIPrint struct {
	itypes.Object
}

func (f FFIPrint) Repr() string {
	return "_print"
}

func (f FFIPrint) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		StarParam: "star",
		//KWParam: "kw",
	}, nil
}

func (f FFIPrint) Call(i itypes.Interpreter) (itypes.Object, error) {
	if args, err := itypes.ArgAs[*interpreter.ObjectTuple](i, "star"); err != nil {
		return nil, err
	} else {
		var sb strings.Builder
		for _, arg := range args.Items {
			sb.WriteString(itypes.Repr(arg))
		}
		fmt.Printf("%s\n", sb.String())
	}
	return primitive.NewObjectNone(), nil
}
