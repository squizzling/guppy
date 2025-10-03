package stream

import (
	"fmt"
	"strings"
)

type DebugWriter struct {
	depth int
}

func (dw DebugWriter) p() string {
	return strings.Repeat(" ", 2*dw.depth)
}

func (dw *DebugWriter) i() {
	dw.depth++
}

func (dw *DebugWriter) o() {
	dw.depth--
}

func s(a any, err error) string {
	return a.(string)
}

func (dw DebugWriter) VisitStreamFuncAbs(sfa *StreamFuncAbs) (any, error) {
	_s := "StreamFuncAbs(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfa.Object, sfa.Object)
	if sfa.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sfa.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sfa.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncAlerts(sfa *StreamFuncAlerts) (any, error) {
	_s := "StreamFuncAlerts(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfa.Object, sfa.Object)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncCombine(sfc *StreamFuncCombine) (any, error) {
	_s := "StreamFuncCombine(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfc.Object, sfc.Object)
	if sfc.Source != nil {
		_s += dw.p() + "Source: " + s(sfc.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Mode: string(" + sfc.Mode + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncConstDouble(sfcd *StreamFuncConstDouble) (any, error) {
	_s := "StreamFuncConstDouble(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfcd.Object, sfcd.Object)
	// TODO: 1 Value float64
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sfcd.Value, sfcd.Value)
	// TODO: 2 Key map[string]string
	_s += dw.p() + fmt.Sprintf("Key: %T(%v)\n", sfcd.Key, sfcd.Key)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncConstInt(sfci *StreamFuncConstInt) (any, error) {
	_s := "StreamFuncConstInt(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfci.Object, sfci.Object)
	// TODO: 1 Value int
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sfci.Value, sfci.Value)
	// TODO: 2 Key map[string]string
	_s += dw.p() + fmt.Sprintf("Key: %T(%v)\n", sfci.Key, sfci.Key)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncCount(sfc *StreamFuncCount) (any, error) {
	_s := "StreamFuncCount(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfc.Object, sfc.Object)
	if sfc.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sfc.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sfc.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncData(sfd *StreamFuncData) (any, error) {
	_s := "StreamFuncData(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfd.Object, sfd.Object)
	_s += dw.p() + "MetricName: string(" + sfd.MetricName + ")\n"
	// TODO: 2 Filter filter.Filter
	_s += dw.p() + fmt.Sprintf("Filter: %T(%v)\n", sfd.Filter, sfd.Filter)
	_s += dw.p() + "Rollup: string(" + sfd.Rollup + ")\n"
	_s += dw.p() + "Extrapolation: string(" + sfd.Extrapolation + ")\n"
	// TODO: 5 MaxExtrapolations int
	_s += dw.p() + fmt.Sprintf("MaxExtrapolations: %T(%v)\n", sfd.MaxExtrapolations, sfd.MaxExtrapolations)
	// TODO: 6 TimeShift time.Duration
	_s += dw.p() + fmt.Sprintf("TimeShift: %T(%v)\n", sfd.TimeShift, sfd.TimeShift)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncDetect(sfd *StreamFuncDetect) (any, error) {
	_s := "StreamFuncDetect(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfd.Object, sfd.Object)
	if sfd.On != nil {
		_s += dw.p() + "On: " + s(sfd.On.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "On: nil\n"
	}
	if sfd.Off != nil {
		_s += dw.p() + "Off: " + s(sfd.Off.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Off: nil\n"
	}
	_s += dw.p() + "Mode: string(" + sfd.Mode + ")\n"
	// TODO: 4 Annotations interpreter.Object
	_s += dw.p() + fmt.Sprintf("Annotations: %T(%v)\n", sfd.Annotations, sfd.Annotations)
	// TODO: 5 EventAnnotations interpreter.Object
	_s += dw.p() + fmt.Sprintf("EventAnnotations: %T(%v)\n", sfd.EventAnnotations, sfd.EventAnnotations)
	// TODO: 6 AutoResolveAfter *time.Duration
	_s += dw.p() + fmt.Sprintf("AutoResolveAfter: %T(%v)\n", sfd.AutoResolveAfter, sfd.AutoResolveAfter)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncEvents(sfe *StreamFuncEvents) (any, error) {
	_s := "StreamFuncEvents(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfe.Object, sfe.Object)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncMax(sfm *StreamFuncMax) (any, error) {
	_s := "StreamFuncMax(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfm.Object, sfm.Object)
	if sfm.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sfm.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sfm.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Value interpreter.Object
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sfm.Value, sfm.Value)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncMean(sfm *StreamFuncMean) (any, error) {
	_s := "StreamFuncMean(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfm.Object, sfm.Object)
	if sfm.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sfm.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sfm.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Constants []interpreter.Object
	_s += dw.p() + fmt.Sprintf("Constants: %T(%v)\n", sfm.Constants, sfm.Constants)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncMedian(sfm *StreamFuncMedian) (any, error) {
	_s := "StreamFuncMedian(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfm.Object, sfm.Object)
	if sfm.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sfm.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sfm.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Constants []interpreter.Object
	_s += dw.p() + fmt.Sprintf("Constants: %T(%v)\n", sfm.Constants, sfm.Constants)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncMin(sfm *StreamFuncMin) (any, error) {
	_s := "StreamFuncMin(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfm.Object, sfm.Object)
	if sfm.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sfm.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sfm.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Value interpreter.Object
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sfm.Value, sfm.Value)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncSum(sfs *StreamFuncSum) (any, error) {
	_s := "StreamFuncSum(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfs.Object, sfs.Object)
	if sfs.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sfs.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sfs.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Constant float64
	_s += dw.p() + fmt.Sprintf("Constant: %T(%v)\n", sfs.Constant, sfs.Constant)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncThreshold(sft *StreamFuncThreshold) (any, error) {
	_s := "StreamFuncThreshold(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sft.Object, sft.Object)
	// TODO: 1 Value float64
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sft.Value, sft.Value)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncUnion(sfu *StreamFuncUnion) (any, error) {
	_s := "StreamFuncUnion(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfu.Object, sfu.Object)
	if sfu.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sfu.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sfu.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFuncWhen(sfw *StreamFuncWhen) (any, error) {
	_s := "StreamFuncWhen(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sfw.Object, sfw.Object)
	if sfw.Predicate != nil {
		_s += dw.p() + "Predicate: " + s(sfw.Predicate.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Predicate: nil\n"
	}
	// TODO: 2 Lasting *time.Duration
	_s += dw.p() + fmt.Sprintf("Lasting: %T(%v)\n", sfw.Lasting, sfw.Lasting)
	// TODO: 3 AtLeast float64
	_s += dw.p() + fmt.Sprintf("AtLeast: %T(%v)\n", sfw.AtLeast, sfw.AtLeast)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodAbove(sma *StreamMethodAbove) (any, error) {
	_s := "StreamMethodAbove(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sma.Object, sma.Object)
	if sma.Source != nil {
		_s += dw.p() + "Source: " + s(sma.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodAbs(sma *StreamMethodAbs) (any, error) {
	_s := "StreamMethodAbs(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sma.Object, sma.Object)
	if sma.Source != nil {
		_s += dw.p() + "Source: " + s(sma.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodAggregate(sma *StreamMethodAggregate) (any, error) {
	_s := "StreamMethodAggregate(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sma.Object, sma.Object)
	if sma.Source != nil {
		_s += dw.p() + "Source: " + s(sma.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Fn: string(" + sma.Fn + ")\n"
	if sma.By == nil {
		_s += dw.p() + "By: nil\n"
	} else if len(sma.By) == 0 {
		_s += dw.p() + "By: []\n"
	} else {
		_s += dw.p() + "By: [\n"
		dw.i()
		for _, _r := range sma.By {
			_s += dw.p() + _r + "\n" // []string
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	_s += dw.p() + "AllowAllMissing: bool(" + fmt.Sprintf("%t", sma.AllowAllMissing) + ")\n"
	if sma.AllowMissing == nil {
		_s += dw.p() + "AllowMissing: nil\n"
	} else if len(sma.AllowMissing) == 0 {
		_s += dw.p() + "AllowMissing: []\n"
	} else {
		_s += dw.p() + "AllowMissing: [\n"
		dw.i()
		for _, _r := range sma.AllowMissing {
			_s += dw.p() + _r + "\n" // []string
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodBelow(smb *StreamMethodBelow) (any, error) {
	_s := "StreamMethodBelow(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smb.Object, smb.Object)
	if smb.Source != nil {
		_s += dw.p() + "Source: " + s(smb.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodFill(smf *StreamMethodFill) (any, error) {
	_s := "StreamMethodFill(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smf.Object, smf.Object)
	if smf.Source != nil {
		_s += dw.p() + "Source: " + s(smf.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	// TODO: 2 Value interpreter.Object
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", smf.Value, smf.Value)
	// TODO: 3 Duration int
	_s += dw.p() + fmt.Sprintf("Duration: %T(%v)\n", smf.Duration, smf.Duration)
	// TODO: 4 MaxCount int
	_s += dw.p() + fmt.Sprintf("MaxCount: %T(%v)\n", smf.MaxCount, smf.MaxCount)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodGeneric(smg *StreamMethodGeneric) (any, error) {
	_s := "StreamMethodGeneric(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smg.Object, smg.Object)
	if smg.Source != nil {
		_s += dw.p() + "Source: " + s(smg.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Call: string(" + smg.Call + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodPercentile(smp *StreamMethodPercentile) (any, error) {
	_s := "StreamMethodPercentile(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smp.Object, smp.Object)
	if smp.Source != nil {
		_s += dw.p() + "Source: " + s(smp.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodPublish(smp *StreamMethodPublish) (any, error) {
	_s := "StreamMethodPublish(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smp.Object, smp.Object)
	if smp.Source != nil {
		_s += dw.p() + "Source: " + s(smp.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Label: string(" + smp.Label + ")\n"
	_s += dw.p() + "Enable: bool(" + fmt.Sprintf("%t", smp.Enable) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodScale(sms *StreamMethodScale) (any, error) {
	_s := "StreamMethodScale(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sms.Object, sms.Object)
	if sms.Source != nil {
		_s += dw.p() + "Source: " + s(sms.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	// TODO: 2 Multiple float64
	_s += dw.p() + fmt.Sprintf("Multiple: %T(%v)\n", sms.Multiple, sms.Multiple)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodTimeShift(smts *StreamMethodTimeShift) (any, error) {
	_s := "StreamMethodTimeShift(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smts.Object, smts.Object)
	if smts.Source != nil {
		_s += dw.p() + "Source: " + s(smts.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	// TODO: 2 Offset time.Duration
	_s += dw.p() + fmt.Sprintf("Offset: %T(%v)\n", smts.Offset, smts.Offset)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodTop(smt *StreamMethodTop) (any, error) {
	_s := "StreamMethodTop(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smt.Object, smt.Object)
	if smt.Source != nil {
		_s += dw.p() + "Source: " + s(smt.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodTransform(smt *StreamMethodTransform) (any, error) {
	_s := "StreamMethodTransform(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smt.Object, smt.Object)
	if smt.Source != nil {
		_s += dw.p() + "Source: " + s(smt.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Fn: string(" + smt.Fn + ")\n"
	// TODO: 3 Over time.Duration
	_s += dw.p() + fmt.Sprintf("Over: %T(%v)\n", smt.Over, smt.Over)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMethodTransformCycle(smtc *StreamMethodTransformCycle) (any, error) {
	_s := "StreamMethodTransformCycle(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smtc.Object, smtc.Object)
	if smtc.Source != nil {
		_s += dw.p() + "Source: " + s(smtc.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Fn: string(" + smtc.Fn + ")\n"
	_s += dw.p() + "Cycle: string(" + smtc.Cycle + ")\n"
	// TODO: 4 CycleStart *string
	_s += dw.p() + fmt.Sprintf("CycleStart: %T(%v)\n", smtc.CycleStart, smtc.CycleStart)
	// TODO: 5 Timezone *string
	_s += dw.p() + fmt.Sprintf("Timezone: %T(%v)\n", smtc.Timezone, smtc.Timezone)
	_s += dw.p() + "PartialValues: bool(" + fmt.Sprintf("%t", smtc.PartialValues) + ")\n"
	// TODO: 7 ShiftCycles int
	_s += dw.p() + fmt.Sprintf("ShiftCycles: %T(%v)\n", smtc.ShiftCycles, smtc.ShiftCycles)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamBinaryOpDouble(sbod *StreamBinaryOpDouble) (any, error) {
	_s := "StreamBinaryOpDouble(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sbod.Object, sbod.Object)
	if sbod.Stream != nil {
		_s += dw.p() + "Stream: " + s(sbod.Stream.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Stream: nil\n"
	}
	_s += dw.p() + "Op: string(" + sbod.Op + ")\n"
	// TODO: 3 Other float64
	_s += dw.p() + fmt.Sprintf("Other: %T(%v)\n", sbod.Other, sbod.Other)
	_s += dw.p() + "Reverse: bool(" + fmt.Sprintf("%t", sbod.Reverse) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamBinaryOpInt(sboi *StreamBinaryOpInt) (any, error) {
	_s := "StreamBinaryOpInt(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sboi.Object, sboi.Object)
	if sboi.Stream != nil {
		_s += dw.p() + "Stream: " + s(sboi.Stream.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Stream: nil\n"
	}
	_s += dw.p() + "Op: string(" + sboi.Op + ")\n"
	// TODO: 3 Other int
	_s += dw.p() + fmt.Sprintf("Other: %T(%v)\n", sboi.Other, sboi.Other)
	_s += dw.p() + "Reverse: bool(" + fmt.Sprintf("%t", sboi.Reverse) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamBinaryOpStream(sbos *StreamBinaryOpStream) (any, error) {
	_s := "StreamBinaryOpStream(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sbos.Object, sbos.Object)
	if sbos.Left != nil {
		_s += dw.p() + "Left: " + s(sbos.Left.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Left: nil\n"
	}
	_s += dw.p() + "Op: string(" + sbos.Op + ")\n"
	if sbos.Right != nil {
		_s += dw.p() + "Right: " + s(sbos.Right.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamIsNone(sin *StreamIsNone) (any, error) {
	_s := "StreamIsNone(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sin.Object, sin.Object)
	if sin.Source != nil {
		_s += dw.p() + "Source: " + s(sin.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Invert: bool(" + fmt.Sprintf("%t", sin.Invert) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamTernary(st *StreamTernary) (any, error) {
	_s := "StreamTernary(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", st.Object, st.Object)
	if st.Condition != nil {
		_s += dw.p() + "Condition: " + s(st.Condition.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Condition: nil\n"
	}
	if st.Left != nil {
		_s += dw.p() + "Left: " + s(st.Left.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Left: nil\n"
	}
	if st.Right != nil {
		_s += dw.p() + "Right: " + s(st.Right.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamUnaryOpMinus(suom *StreamUnaryOpMinus) (any, error) {
	_s := "StreamUnaryOpMinus(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", suom.Object, suom.Object)
	if suom.Stream != nil {
		_s += dw.p() + "Stream: " + s(suom.Stream.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Stream: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}
