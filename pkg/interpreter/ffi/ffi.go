package ffi

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type funcFieldSetter func(i itypes.Interpreter, tgt reflect.Value) error

// This amazing syntax brought to you courtesy of the golang design team
var typeObject = reflect.TypeOf((*itypes.Object)(nil)).Elem()

type FFICall interface {
	Call(i itypes.Interpreter) (itypes.Object, error)
}

type ffi[T FFICall] struct {
	itypes.Object
	params    *itypes.Params
	setFields []funcFieldSetter
	defaults  T
}

func (f *ffi[T]) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return f.params, nil
}

func (f *ffi[T]) Call(i itypes.Interpreter) (itypes.Object, error) {
	data := f.defaults
	valueDestination := reflect.ValueOf(&data).Elem()

	for _, fn := range f.setFields {
		if err := fn(i, valueDestination); err != nil {
			// TODO: We could accumulate errors here if we wanted, not sure if we do
			return nil, err
		}
	}

	return data.Call(i)
}

func NewFFI[T FFICall](defaults T) itypes.FlowCall {
	ffiDefaults := reflect.ValueOf(defaults)
	if ffiDefaults.Kind() == reflect.Pointer {
		ffiDefaults = ffiDefaults.Elem()
	}
	if ffiDefaults.Kind() != reflect.Struct {
		panic(fmt.Sprintf(
			"NewFFI: type %s has unknown kind=%s, type=%s",
			ffiDefaults.Type().Name(),
			ffiDefaults.Type().Kind().String(),
			ffiDefaults.Type().String(),
		))
	}

	var setFields []funcFieldSetter

	params := &itypes.Params{}

	for idx := 0; idx < ffiDefaults.NumField(); idx++ {
		fieldType := ffiDefaults.Type().Field(idx)

		argNameType, ok := fieldType.Tag.Lookup("ffi")
		if !ok {
			continue
		}

		argParts := strings.Split(argNameType, ",")
		argName := argParts[0]

		structFieldName := ffiDefaults.Type().Name() + "." + fieldType.Name

		fieldValue := ffiDefaults.Field(idx)
		var defaultValue itypes.Object
		if fieldValue.Kind() == reflect.Struct {
			for idx := 0; idx < fieldValue.Type().NumField(); idx++ {
				fld := fieldValue.Field(idx)
				if !fld.CanConvert(typeObject) {
					panic(fmt.Sprintf(
						"NewFFI: %s.%s does not have an underlying itypes.Object type, kind=%s, type=%s",
						structFieldName,
						fieldType.Type.Field(idx).Name,
						fieldType.Type.Field(idx).Type.Kind().String(),
						fieldType.Type.Field(idx).Type.String(),
					))
				} else if !fld.IsNil() {
					defaultValue = fld.Interface().(itypes.Object)
					break
				}
			}
		} else if !fieldValue.Type().ConvertibleTo(typeObject) {
			panic(fmt.Sprintf(
				"NewFFI: %s does not have an underlying itypes.Object or struct type, kind=%s, type=%s",
				structFieldName,
				fieldType.Type.Kind().String(),
				fieldType.Type.String(),
			))
		} else if !fieldValue.IsNil() {
			defaultValue = fieldValue.Interface().(itypes.Object)
		}

		switch {
		case len(argParts) == 1:
			params.Params = append(params.Params, itypes.ParamDef{Name: argParts[0], Default: defaultValue})
		case argParts[1] == "star":
			// TODO: A StarParam will always come in as a *primitive.ObjectTuple, it would be nice if we could somehow
			//       turn this in to a slice of a specific type, or slice of Object, and do more type enforcement.
			//       It's complicated by the import cycle (among other things) though.  For now, this is sufficient.
			params.StarParam = argParts[0]
			if fieldType.Type.String() != "*primitive.ObjectTuple" {
				// Can't use the actual type due to import cycles
				panic(fmt.Sprintf("NewFFI: %s must be *primitive.ObjectTuple as it is a star-args", structFieldName))
			}
		case argParts[1] == "kw":
			params.KWParams = append(params.KWParams, itypes.ParamDef{Name: argParts[0], Default: defaultValue})
		case argParts[1] == "kwargs":
			// TODO: We need to do type checking on the destination
			panic("kwargs not yet supported")
			params.StarParam = argParts[0]
		default:
			panic("NewFFI: unrecognized param type: " + argParts[1])
		}

		// Compute the typeString when we create the FFI, this lets us move type checking in to the
		// startup path (ie, fail-fast), and we only need to do this once, assuming we re-use FFI
		// objects between interpreters.
		typeString := newExpectedTypeString(fieldType.Type)
		if fieldType.Type.Kind() == reflect.Struct {
			setFields = append(setFields, oneOfHandler(idx, argName, structFieldName, typeString))
		} else {
			setFields = append(setFields, singleHandler(idx, argName, structFieldName, typeString))
		}
	}

	return &ffi[T]{
		Object:    itypes.NewObject(nil),
		params:    params,
		defaults:  defaults,
		setFields: setFields,
	}
}

func singleHandler(idx int, argName string, structFieldName string, typeString string) funcFieldSetter {
	return func(i itypes.Interpreter, valueFFIStruct reflect.Value) error {
		valueDestination := valueFFIStruct.Field(idx)
		if argValue, err := i.GetArg(argName); err != nil {
			return handleArgMissing(argName, structFieldName, typeString, valueDestination)
		} else {
			return singlePresent(argName, structFieldName, typeString, valueDestination, argValue)
		}
	}
}

func oneOfHandler(idx int, argName string, structFieldName string, typeString string) funcFieldSetter {
	return func(i itypes.Interpreter, valueFFIStruct reflect.Value) error {
		valueOneOf := valueFFIStruct.Field(idx)
		if argValue, err := i.GetArg(argName); err != nil {
			return handleArgMissing(argName, structFieldName, typeString, valueOneOf)
		} else {
			return oneOfPresent(argName, structFieldName, typeString, valueOneOf, argValue)
		}
	}
}

func handleArgMissing(argName string, structFieldName string, typeString string, valueDestination reflect.Value) error {
	// In theory, this should always be handled by VisitExpressionCall, however we keep it as a fail-safe.
	// We also keep it because we may migrate the behavior from VisitExpressionCall to this entirely later.
	if hasDefault(valueDestination) {
		return nil
	} else {
		return fmt.Errorf("param `%s` for %s is missing, expecting %s", argName, structFieldName, typeString)
	}
}

func singlePresent(argName string, structFieldName string, typeString string, valueDestination reflect.Value, argValue itypes.Object) error {
	if va := reflect.ValueOf(argValue); !va.CanConvert(valueDestination.Type()) {
		return fmt.Errorf("param `%s` for %s is %T not %s", argName, structFieldName, argValue, typeString)
	} else {
		valueDestination.Set(va)
		return nil
	}
}

func oneOfPresent(argName string, structFieldName string, typeString string, valueOneOf reflect.Value, argValue itypes.Object) error {
	va := reflect.ValueOf(argValue)
	isSet := false

	for idx := 0; idx < valueOneOf.NumField(); idx++ {
		oneOfFldV := valueOneOf.Field(idx)
		switch oneOfFldV.Kind() {
		case reflect.Pointer:
			if isSet || va.Type() != oneOfFldV.Type() {
				oneOfFldV.Set(reflect.Zero(oneOfFldV.Type()))
			} else {
				oneOfFldV.Set(va)
				isSet = true
			}
		case reflect.Interface:
			if isSet || !va.CanConvert(oneOfFldV.Type()) {
				oneOfFldV.Set(reflect.Zero(oneOfFldV.Type()))
			} else {
				oneOfFldV.Set(va)
				isSet = true
			}
		default:
			continue
		}
	}
	if !isSet {
		return fmt.Errorf("param `%s` for %s is %T not %s", argName, structFieldName, argValue, typeString)
	}
	return nil
}

func hasDefault(valueDestination reflect.Value) bool {
	switch valueDestination.Kind() {
	case reflect.Pointer:
		return !valueDestination.IsNil()
	case reflect.Struct:
		for idx := 0; idx < valueDestination.NumField(); idx++ {
			oneOfField := valueDestination.Field(idx)
			if oneOfField.Kind() != reflect.Pointer {
				// TODO: We shouldn't panic in non-startup code
				panic(fmt.Sprintf(
					"hasDefault: field %s has unknown kind=%s, type=%s",
					valueDestination.Type().Field(idx).Name,
					oneOfField.Kind().String(),
					valueDestination.Type().Field(idx).Type.Name(),
				))
			}
			if !oneOfField.IsNil() {
				return true
			}
		}
		return false
	default:
		// TODO: We shouldn't panic in non-startup code
		panic(fmt.Sprintf(
			"hasDefault: type %s has unknown kind=%s",
			valueDestination.Type().Name(),
			valueDestination.Kind(),
		))
	}
}

func newExpectedTypeString(fldT reflect.Type) string {
	switch fldT.Kind() {
	case reflect.Pointer:
		return fldT.String()
	case reflect.Interface:
		return fldT.String()
	case reflect.Struct:
		var sb strings.Builder
		for oneOfIdx := 0; oneOfIdx < fldT.NumField(); oneOfIdx++ {
			oneOfFldT := fldT.Field(oneOfIdx)
			if oneOfFldT.Type.Kind() != reflect.Pointer && oneOfFldT.Type.Kind() != reflect.Interface {
				panic(fmt.Sprintf(
					"newExpectedTypeString: field %s has unknown kind=%s, type=%s",
					oneOfFldT.Name,
					oneOfFldT.Type.Kind().String(),
					oneOfFldT.Type.String(),
				))
			}
			sb.WriteString(oneOfFldT.Type.String())
			if oneOfIdx != fldT.NumField()-1 {
				sb.WriteString(", ")
				if oneOfIdx == fldT.NumField()-2 {
					sb.WriteString("or ")
				}
			}
		}
		return sb.String()
	default:
		panic(fmt.Sprintf(
			"newExpectedTypeString: type %s has unknown kind=%s",
			fldT.String(),
			fldT.Kind(),
		))
	}
}
