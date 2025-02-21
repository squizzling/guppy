package filter

import (
	"fmt"

	"guppy/internal/interpreter"
)

type FFIFilter struct {
	interpreter.Object
}

func (f FFIFilter) Params(i *interpreter.Interpreter) ([]interpreter.ParamData, error) {
	return []interpreter.ParamData{
		{Name: "key"},
		{Name: "value"},
	}, nil
}

func (f FFIFilter) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if key, err := interpreter.ArgAsString(i, "key"); err != nil {
		return nil, err
	} else if value, err := interpreter.ArgAsString(i, "value"); err != nil {
		// TODO: value should read an array as `*value`
		return nil, err
	} else {
		return NewKV(key, []string{value}), nil
	}
}

type kv struct {
	interpreter.Object

	key    string
	values []string
}

func NewKV(key string, values []string) Filter {
	return &kv{
		Object: newFilterObject(),
		key:    key,
		values: values,
	}
}

func (fkc *kv) RenderFilter() string {
	if len(fkc.values) != 1 {
		panic("stop")
	}
	return fmt.Sprintf("filter('%s', '%s')", fkc.key, fkc.values[0])
}
