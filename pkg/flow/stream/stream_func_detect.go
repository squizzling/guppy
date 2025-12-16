package stream

import (
	"fmt"
	"time"

	"github.com/squizzling/guppy/pkg/flow/annotate"
	"github.com/squizzling/guppy/pkg/flow/duration"
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/ftypes"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiDetect struct {
	/**
	detect(None)               : <'detect_with_label function'> unsupported type 'none' object for argument 'on'
	detect(1)                  : <'detect_with_label function'> unsupported type 'long' object for argument 'on'
	detect(1.1)                : <'detect_with_label function'> unsupported type 'double' object for argument 'on'
	detect('1')                : <'detect_with_label function'> unsupported type 'string' object for argument 'on'
	detect(const(1))           : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'on'
	detect(const(1) > const(2)):
	detect(data('x'))          : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'on'
	detect(duration('1m'))     : <'detect_with_label function'> unsupported type 'duration' object for argument 'on'
	detect(True)               : <'detect_with_label function'> unsupported type 'boolean' object for argument 'on'
	detect(False)              : <'detect_with_label function'> unsupported type 'boolean' object for argument 'on'
	detect([])                 : <'detect_with_label function'> unsupported type 'list' object for argument 'on'
	detect(['1'])              : <'detect_with_label function'> unsupported type 'list' object for argument 'on'
	detect([1])                : <'detect_with_label function'> unsupported type 'list' object for argument 'on'
	detect(())                 : <'detect_with_label function'> unsupported type 'tuple' object for argument 'on'
	detect(('1',))             : <'detect_with_label function'> unsupported type 'tuple' object for argument 'on'
	detect(('1',1))            : <'detect_with_label function'> unsupported type 'tuple' object for argument 'on'
	detect({})                 : <'detect_with_label function'> unsupported type 'dict' object for argument 'on'
	detect({'x': 'x'})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'on'
	detect({'x': ['x']})       : <'detect_with_label function'> unsupported type 'dict' object for argument 'on'
	detect({'x': 1})           : <'detect_with_label function'> unsupported type 'dict' object for argument 'on'
	detect({'x': [1]})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'on'
	detect(lambda x: x)        : <'detect_with_label function'> unsupported type <lambda function> for argument 'on'

	detect(const(1) > const(2), None)               :
	detect(const(1) > const(2), 1)                  : <'detect_with_label function'> unsupported type 'long' object for argument 'off'
	detect(const(1) > const(2), 1.1)                : <'detect_with_label function'> unsupported type 'double' object for argument 'off'
	detect(const(1) > const(2), '1')                : <'detect_with_label function'> unsupported type 'string' object for argument 'off'
	detect(const(1) > const(2), const(1))           : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'off'
	detect(const(1) > const(2), const(1) > const(2)):
	detect(const(1) > const(2), data('x'))          : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'off'
	detect(const(1) > const(2), duration('1m'))     : <'detect_with_label function'> unsupported type 'duration' object for argument 'off'
	detect(const(1) > const(2), True)               : <'detect_with_label function'> unsupported type 'boolean' object for argument 'off'
	detect(const(1) > const(2), False)              : <'detect_with_label function'> unsupported type 'boolean' object for argument 'off'
	detect(const(1) > const(2), [])                 : <'detect_with_label function'> unsupported type 'list' object for argument 'off'
	detect(const(1) > const(2), ['1'])              : <'detect_with_label function'> unsupported type 'list' object for argument 'off'
	detect(const(1) > const(2), [1])                : <'detect_with_label function'> unsupported type 'list' object for argument 'off'
	detect(const(1) > const(2), ())                 : <'detect_with_label function'> unsupported type 'tuple' object for argument 'off'
	detect(const(1) > const(2), ('1',))             : <'detect_with_label function'> unsupported type 'tuple' object for argument 'off'
	detect(const(1) > const(2), ('1',1))            : <'detect_with_label function'> unsupported type 'tuple' object for argument 'off'
	detect(const(1) > const(2), {})                 : <'detect_with_label function'> unsupported type 'dict' object for argument 'off'
	detect(const(1) > const(2), {'x': 'x'})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'off'
	detect(const(1) > const(2), {'x': ['x']})       : <'detect_with_label function'> unsupported type 'dict' object for argument 'off'
	detect(const(1) > const(2), {'x': 1})           : <'detect_with_label function'> unsupported type 'dict' object for argument 'off'
	detect(const(1) > const(2), {'x': [1]})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'off'
	detect(const(1) > const(2), lambda x: x)        : <'detect_with_label function'> unsupported type <lambda function> for argument 'off'

	detect(const(1) > const(2), None, None)               : <'detect_with_label function'> unsupported type 'none' object for argument 'mode'
	detect(const(1) > const(2), None, 1)                  : <'detect_with_label function'> unsupported type 'long' object for argument 'mode'
	detect(const(1) > const(2), None, 1.1)                : <'detect_with_label function'> unsupported type 'double' object for argument 'mode'
	detect(const(1) > const(2), None, '1')                : <'detect_with_label function'> argument 'mode' got invalid value '1'; expected [paired, split]
	detect(const(1) > const(2), None, const(1))           : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'mode'
	detect(const(1) > const(2), None, const(1) > const(2)): <'detect_with_label function'> unsupported type <stream of BOOLEAN> for argument 'mode'
	detect(const(1) > const(2), None, data('x'))          : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'mode'
	detect(const(1) > const(2), None, duration('1m'))     : <'detect_with_label function'> unsupported type 'duration' object for argument 'mode'
	detect(const(1) > const(2), None, True)               : <'detect_with_label function'> unsupported type 'boolean' object for argument 'mode'
	detect(const(1) > const(2), None, False)              : <'detect_with_label function'> unsupported type 'boolean' object for argument 'mode'
	detect(const(1) > const(2), None, [])                 : <'detect_with_label function'> unsupported type 'list' object for argument 'mode'
	detect(const(1) > const(2), None, ['1'])              : <'detect_with_label function'> unsupported type 'list' object for argument 'mode'
	detect(const(1) > const(2), None, [1])                : <'detect_with_label function'> unsupported type 'list' object for argument 'mode'
	detect(const(1) > const(2), None, ())                 : <'detect_with_label function'> unsupported type 'tuple' object for argument 'mode'
	detect(const(1) > const(2), None, ('1',))             : <'detect_with_label function'> unsupported type 'tuple' object for argument 'mode'
	detect(const(1) > const(2), None, ('1',1))            : <'detect_with_label function'> unsupported type 'tuple' object for argument 'mode'
	detect(const(1) > const(2), None, {})                 : <'detect_with_label function'> unsupported type 'dict' object for argument 'mode'
	detect(const(1) > const(2), None, {'x': 'x'})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'mode'
	detect(const(1) > const(2), None, {'x': ['x']})       : <'detect_with_label function'> unsupported type 'dict' object for argument 'mode'
	detect(const(1) > const(2), None, {'x': 1})           : <'detect_with_label function'> unsupported type 'dict' object for argument 'mode'
	detect(const(1) > const(2), None, {'x': [1]})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'mode'
	detect(const(1) > const(2), None, lambda x: x)        : <'detect_with_label function'> unsupported type <lambda function> for argument 'mode'

	detect(const(1) > const(2), None, 'paired', None)               :
	detect(const(1) > const(2), None, 'paired', 1)                  : <'detect_with_label function'> unsupported type 'long' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', 1.1)                : <'detect_with_label function'> unsupported type 'double' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', '1')                : <'detect_with_label function'> unsupported type 'string' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', const(1))           : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', const(1) > const(2)): <'detect_with_label function'> unsupported type <stream of BOOLEAN> for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', data('x'))          : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', duration('1m'))     : <'detect_with_label function'> unsupported type 'duration' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', True)               : <'detect_with_label function'> unsupported type 'boolean' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', False)              : <'detect_with_label function'> unsupported type 'boolean' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', [])                 :
	detect(const(1) > const(2), None, 'paired', ['1'])              : Error executing SignalFlow program (ref: G8TqJeXAwA4): [error parsing program detect(const(1) > const(2), None, 'paired', ['1']): parameter 'annotations' must be a list of annotated() values, got ''string' object']
	detect(const(1) > const(2), None, 'paired', [1])                : Error executing SignalFlow program (ref: G8TqJ4nA0AA): [error parsing program detect(const(1) > const(2), None, 'paired', [1]): parameter 'annotations' must be a list of annotated() values, got ''long' object']
	detect(const(1) > const(2), None, 'paired', ())                 :
	detect(const(1) > const(2), None, 'paired', ('1',))             : Error executing SignalFlow program (ref: G8TqKb3AwAA): [error parsing program detect(const(1) > const(2), None, 'paired', ('1',)): parameter 'annotations' must be a list of annotated() values, got ''string' object']
	detect(const(1) > const(2), None, 'paired', ('1',1))            : Error executing SignalFlow program (ref: G8TqKzqA0AA): [error parsing program detect(const(1) > const(2), None, 'paired', ('1',1)): parameter 'annotations' must be a list of annotated() values, got ''string' object']
	detect(const(1) > const(2), None, 'paired', {})                 : <'detect_with_label function'> unsupported type 'dict' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', {'x': 'x'})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', {'x': ['x']})       : <'detect_with_label function'> unsupported type 'dict' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', {'x': 1})           : <'detect_with_label function'> unsupported type 'dict' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', {'x': [1]})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'annotations'
	detect(const(1) > const(2), None, 'paired', lambda x: x)        : <'detect_with_label function'> unsupported type <lambda function> for argument 'annotations'

	detect(const(1) > const(2), None, 'paired', None, None)               :
	detect(const(1) > const(2), None, 'paired', None, 1)                  : <'detect_with_label function'> unsupported type 'long' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, 1.1)                : <'detect_with_label function'> unsupported type 'double' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, '1')                : <'detect_with_label function'> unsupported type 'string' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, const(1))           : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, const(1) > const(2)): <'detect_with_label function'> unsupported type <stream of BOOLEAN> for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, data('x'))          : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, duration('1m'))     : <'detect_with_label function'> unsupported type 'duration' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, True)               : <'detect_with_label function'> unsupported type 'boolean' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, False)              : <'detect_with_label function'> unsupported type 'boolean' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, [])                 : <'detect_with_label function'> unsupported type 'list' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, ['1'])              : <'detect_with_label function'> unsupported type 'list' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, [1])                : <'detect_with_label function'> unsupported type 'list' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, ())                 : <'detect_with_label function'> unsupported type 'tuple' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, ('1',))             : <'detect_with_label function'> unsupported type 'tuple' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, ('1',1))            : <'detect_with_label function'> unsupported type 'tuple' object for argument 'event_annotations'
	detect(const(1) > const(2), None, 'paired', None, {})                 :
	detect(const(1) > const(2), None, 'paired', None, {'x': 'x'})         :
	detect(const(1) > const(2), None, 'paired', None, {'x': ['x']})       : Error executing SignalFlow program (ref: G8TrHS0AwBU): [block 02-DETECT: argument event_annotations expects dictionary of string to object that can be converted to a string, got key or value that could not be converted to a string]
	detect(const(1) > const(2), None, 'paired', None, {'x': 1})           :
	detect(const(1) > const(2), None, 'paired', None, {'x': [1]})         : Error executing SignalFlow program (ref: G8TrHS0AwBc): [block 02-DETECT: argument event_annotations expects dictionary of string to object that can be converted to a string, got key or value that could not be converted to a string]
	detect(const(1) > const(2), None, 'paired', None, lambda x: x)        : <'detect_with_label function'> unsupported type <lambda function> for argument 'event_annotations'

	detect(const(1) > const(2), None, 'paired', None, None, None)               :
	detect(const(1) > const(2), None, 'paired', None, None, 1)                  : <'detect_with_label function'> argument 'auto_resolve_after' got invalid value '1'; expected must be between 1s and 52w2d
	detect(const(1) > const(2), None, 'paired', None, None, 1.1)                : <'detect_with_label function'> unsupported type 'double' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, '1')                : Error executing SignalFlow program (ref: G8TqxYkA4AA): [error parsing program detect(const(1) > const(2), None, 'paired', None, None, '1'): invalid duration string 1]
	detect(const(1) > const(2), None, 'paired', None, None, const(1))           : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, const(1) > const(2)): <'detect_with_label function'> unsupported type <stream of BOOLEAN> for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, data('x'))          : <'detect_with_label function'> unsupported type <stream of DOUBLE> for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, duration('1m'))     :
	detect(const(1) > const(2), None, 'paired', None, None, True)               : <'detect_with_label function'> unsupported type 'boolean' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, False)              : <'detect_with_label function'> unsupported type 'boolean' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, [])                 : <'detect_with_label function'> unsupported type 'list' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, ['1'])              : <'detect_with_label function'> unsupported type 'list' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, [1])                : <'detect_with_label function'> unsupported type 'list' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, ())                 : <'detect_with_label function'> unsupported type 'tuple' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, ('1',))             : <'detect_with_label function'> unsupported type 'tuple' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, ('1',1))            : <'detect_with_label function'> unsupported type 'tuple' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, {})                 : <'detect_with_label function'> unsupported type 'dict' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, {'x': 'x'})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, {'x': ['x']})       : <'detect_with_label function'> unsupported type 'dict' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, {'x': 1})           : <'detect_with_label function'> unsupported type 'dict' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, {'x': [1]})         : <'detect_with_label function'> unsupported type 'dict' object for argument 'auto_resolve_after'
	detect(const(1) > const(2), None, 'paired', None, None, lambda x: x)        : <'detect_with_label function'> unsupported type <lambda function> for argument 'auto_resolve_after'
	*/

	On          Stream                     `ffi:"on"`
	Off         ftypes.ThingOrNone[Stream] `ffi:"off"`
	Mode        *primitive.ObjectString    `ffi:"mode"`
	Annotations struct {
		None  *primitive.ObjectNone
		List  *primitive.ObjectList
		Tuple *primitive.ObjectTuple
	} `ffi:"annotations"`
	EventAnnotations ftypes.ThingOrNone[*primitive.ObjectDict] `ffi:"event_annotations"`
	AutoResolveAfter struct {
		None     *primitive.ObjectNone
		Duration *duration.Duration
		Int      *primitive.ObjectInt
		String   *primitive.ObjectString
	} `ffi:"auto_resolve_after"`
}

func NewFFIDetect() itypes.FlowCall {
	return ffi.NewFFI(ffiDetect{
		Off:  ftypes.NewThingOrNoneNone[Stream](),
		Mode: primitive.NewObjectString("paired"),
		Annotations: struct {
			None  *primitive.ObjectNone
			List  *primitive.ObjectList
			Tuple *primitive.ObjectTuple
		}{
			None: primitive.NewObjectNone(),
		},
		EventAnnotations: ftypes.NewThingOrNoneNone[*primitive.ObjectDict](),
		AutoResolveAfter: struct {
			None     *primitive.ObjectNone
			Duration *duration.Duration
			Int      *primitive.ObjectInt
			String   *primitive.ObjectString
		}{
			None: primitive.NewObjectNone(),
		},
	})
}

func (f ffiDetect) Call(i itypes.Interpreter) (itypes.Object, error) {
	if mode, err := f.resolveMode(); err != nil {
		return nil, err
	} else if annotations, err := f.resolveAnnotations(); err != nil {
		return nil, err
	} else if eventAnnotations, err := f.resolveEventAnnotations(); err != nil {
		return nil, err
	} else if autoResolveAfter, err := f.resolveAutoResolveAfter(); err != nil {
		return nil, err
	} else {
		return NewStreamFuncDetect(
			prototypeStreamAlert,
			f.On,
			f.resolveOff(),
			mode,
			annotations,
			eventAnnotations,
			autoResolveAfter,
		), nil
	}
}

func (f ffiDetect) resolveOff() Stream {
	if f.Off.None != nil {
		return nil
	} else {
		return f.Off.Thing
	}
}

func (f ffiDetect) resolveMode() (string, error) {
	switch f.Mode.Value {
	case "paired", "split":
		return f.Mode.Value, nil
	default:
		return "", fmt.Errorf("ffiDetect.resolveMode: param `mode` is %s, expected [paired, split]", f.Mode.Value)
	}
}

func (f ffiDetect) resolveAnnotations() ([]*annotate.Annotated, error) {
	var items []itypes.Object
	if f.Annotations.None != nil {
		return nil, nil
	} else if f.Annotations.List != nil {
		items = f.Annotations.List.Items
	} else {
		items = f.Annotations.Tuple.Items
	}

	var out []*annotate.Annotated
	for idx, item := range items {
		if annotation, ok := item.(*annotate.Annotated); !ok {
			return nil, fmt.Errorf("ffiDetect.resolveAnnotations: idx %d is %T not %T", idx, item, annotation)
		} else {
			out = append(out, annotation)
		}
	}
	return out, nil
}

func (f ffiDetect) resolveEventAnnotations() (itypes.Object, error) {
	if f.EventAnnotations.None != nil {
		return nil, nil
	} else {
		panic("handle event annotations")
	}
}

func (f ffiDetect) resolveAutoResolveAfter() (*time.Duration, error) {
	if f.AutoResolveAfter.None != nil {
		return nil, nil
	} else if f.AutoResolveAfter.Int != nil {
		d := time.Duration(f.AutoResolveAfter.Int.Value) * time.Millisecond
		return &d, nil
	} else if f.AutoResolveAfter.Duration != nil {
		return &f.AutoResolveAfter.Duration.Duration, nil
	} else if d, err := duration.ParseDuration(f.AutoResolveAfter.String.Value); err != nil {
		return nil, fmt.Errorf("ffiDetect.resolveAutoResolveAfter: param `auto_resolve_after` failed to parse: %w", err)
	} else {
		return &d, nil
	}
}
