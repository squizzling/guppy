package debug

import (
	"fmt"
	"strings"

	"guppy/internal/interpreter"
)

type FFIPrint struct {
	interpreter.Object
}

func (f FFIPrint) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		StarParam: "star",
		//KWParam: "kw",
	}, nil
}

func (f FFIPrint) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if args, err := interpreter.ArgAs[*interpreter.ObjectTuple](i, "star"); err != nil {
		return nil, err
	} else {
		var sb strings.Builder
		for idx, arg := range args.Items {
			switch arg := arg.(type) {
			case *interpreter.ObjectString:
				sb.WriteString(arg.Value)
			case *interpreter.ObjectInt:
				sb.WriteString(fmt.Sprintf("%d", arg.Value))
			default:
				sb.WriteString(fmt.Sprintf("unknown[%d/%T](%v)", idx, arg, arg))
			}
		}
		fmt.Printf("%s\n", sb.String())
	}
	return interpreter.NewObjectNone(), nil
}
