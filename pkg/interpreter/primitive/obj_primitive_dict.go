package primitive

import (
	"errors"
	"fmt"
	"strings"

	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
)

type DictItem struct {
	Key   itypes.Object
	Value itypes.Object
}

type ObjectDict struct {
	itypes.Object

	Items []DictItem
}

var prototypeObjectDict = itypes.NewObject(map[string]itypes.Object{
	"get":           ffi.NewFFI(ffiObjectDictGet{Default: NewObjectNone()}),
	"__subscript__": ffi.NewFFI(ffiObjectDictSubscript{}),
})

func NewObjectDict(items []DictItem) *ObjectDict {
	return &ObjectDict{
		Object: prototypeObjectDict,

		Items: items,
	}
}

func NewObjectDictFromMap(items map[string]itypes.Object) *ObjectDict {
	var itemList []DictItem
	for key, value := range items {
		itemList = append(itemList, DictItem{
			Key:   NewObjectString(key),
			Value: value,
		})
	}
	return NewObjectDict(itemList)
}

func (od *ObjectDict) AsMapStringString() (map[string]string, error) {
	m := make(map[string]string)
	for idx, item := range od.Items {
		sKey := ""
		sValue := ""
		switch key := item.Key.(type) {
		case *ObjectString:
			sKey = key.Value
		default:
			return nil, fmt.Errorf("dict idx %d (%s) is %T not *interpreter.ObjectString", idx, itypes.Repr(key), key)
		}
		switch value := item.Value.(type) {
		case *ObjectString:
			sValue = value.Value
		default:
			return nil, fmt.Errorf("dict idx %d (%s) is %T not *interpreter.ObjectString", idx, sKey, value)
		}
		m[sKey] = sValue
	}
	return m, nil
}

func (od *ObjectDict) getKeyOrDefault(key string, def itypes.Object) (itypes.Object, error) {
	if val := od.getStringOrNil(key); val == nil {
		return def, nil
	} else {
		return val, nil
	}
}

func (od *ObjectDict) getStringOrNil(key string) itypes.Object {
	for _, item := range od.Items {
		if itemKey, ok := item.Key.(*ObjectString); ok {
			if itemKey.Value == key {
				return item.Value
			}
		}
	}
	return nil
}

func (od *ObjectDict) Repr() string {
	var sb strings.Builder
	sb.WriteString("{")
	for idx, item := range od.Items {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(itypes.Repr(item.Key))
		sb.WriteString(": ")
		sb.WriteString(itypes.Repr(item.Value))
	}
	sb.WriteString("}")
	return sb.String()
}

type ffiObjectDictGet struct {
	Self    *ObjectDict   `ffi:"self"`
	Key     *ObjectString `ffi:"key"`
	Default itypes.Object `ffi:"default"`
}

func (f ffiObjectDictGet) Call(i itypes.Interpreter) (itypes.Object, error) {
	return f.Self.getKeyOrDefault(f.Key.Value, f.Default)
}

type ffiObjectDictSubscript struct {
	Self  *ObjectDict   `ffi:"self"`
	Start *ObjectString `ffi:"start"`
}

func (f ffiObjectDictSubscript) Call(i itypes.Interpreter) (itypes.Object, error) {
	if val := f.Self.getStringOrNil(f.Start.Value); val != nil {
		return val, nil
	} else {
		return nil, errors.New("key not found in dict")
	}
}
