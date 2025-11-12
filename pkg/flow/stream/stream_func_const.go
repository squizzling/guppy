package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIConst struct {
	itypes.Object
}

func (f FFIConst) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "value"},
			{Name: "key", Default: interpreter.NewObjectMissing()},
		},
	}, nil
}

func (f FFIConst) resolveValue(i itypes.Interpreter) (itypes.Object, error) {
	if value, err := i.GetArg("value"); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (f FFIConst) resolveKey(i itypes.Interpreter) (map[string]string, error) {
	if key, err := i.GetArg("key"); err != nil {
		return nil, err
	} else {
		switch key := key.(type) {
		case *interpreter.ObjectMissing:
			return nil, nil
		case *interpreter.ObjectDict:
			if m, err := key.AsMapStringString(); err != nil {
				return nil, err
			} else if len(m) == 0 {
				return nil, fmt.Errorf("key is empty map")
			} else {
				return m, nil
			}
		default:
			return nil, fmt.Errorf("key is %T not *interpreter.ObjectMissing, or *interpreter.ObjectDict", key)
		}
	}
}

func (f FFIConst) Call(i itypes.Interpreter) (itypes.Object, error) {
	if value, err := f.resolveValue(i); err != nil {
		return nil, err
	} else if key, err := f.resolveKey(i); err != nil {
		return nil, err
	} else {
		switch value := value.(type) {
		case *interpreter.ObjectInt:
			return NewStreamFuncConstInt(newStreamObject(), value.Value, key), nil
		case *interpreter.ObjectDouble:
			return NewStreamFuncConstDouble(newStreamObject(), value.Value, key), nil
		default:
			return nil, fmt.Errorf("value is %T not *interpreter.ObjectInt, or *interpreter.ObjectDouble", value)
		}
	}
}

var _ = itypes.FlowCall(FFIConst{})
