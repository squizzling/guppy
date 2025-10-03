package stream

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFIConst struct {
	interpreter.Object
}

func (f FFIConst) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "value"},
			{Name: "key", Default: interpreter.NewObjectMissing()},
		},
	}, nil
}

func (f FFIConst) resolveValue(i *interpreter.Interpreter) (interpreter.Object, error) {
	if value, err := i.Scope.Get("value"); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (f FFIConst) resolveKey(i *interpreter.Interpreter) (map[string]string, error) {
	if key, err := i.Scope.Get("key"); err != nil {
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

func (f FFIConst) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
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

var _ = interpreter.FlowCall(FFIConst{})
