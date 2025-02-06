package filter

import (
	"fmt"

	"github.com/squizzling/types/pkg/result"

	"guppy/internal/interpreter"
)

type FFIFilter struct {
	interpreter.Object
}

func (f FFIFilter) Args(i *interpreter.Interpreter) result.Result[[]interpreter.ArgData] {
	return result.Ok([]interpreter.ArgData{
		{Name: "key"},
		{Name: "value"},
	})
}

func (f FFIFilter) Call(i *interpreter.Interpreter) result.Result[interpreter.Object] {
	if resultKey := interpreter.ArgAsString(i, "key"); !resultKey.Ok() {
		return result.Err[interpreter.Object](resultKey.Err())
	} else if resultValue := interpreter.ArgAsString(i, "value"); !resultValue.Ok() {
		// TODO: value should read an array as `*value`
		return result.Err[interpreter.Object](resultValue.Err())
	} else {
		return result.Ok[interpreter.Object](NewKV(resultKey.Value(), []string{resultValue.Value()}))
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
