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

func (dw DebugWriter) VisitStreamAbove(sa *StreamAbove) (any, error) {
	_s := "StreamAbove(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sa.Object, sa.Object)
	if sa.Source != nil {
		_s += dw.p() + "Source: " + s(sa.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamAbs(sa *StreamAbs) (any, error) {
	_s := "StreamAbs(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sa.Object, sa.Object)
	if sa.Source != nil {
		_s += dw.p() + "Source: " + s(sa.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamAggregate(sa *StreamAggregate) (any, error) {
	_s := "StreamAggregate(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sa.Object, sa.Object)
	if sa.Source != nil {
		_s += dw.p() + "Source: " + s(sa.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Fn: string(" + sa.Fn + ")\n"
	if sa.By == nil {
		_s += dw.p() + "By: nil\n"
	} else if len(sa.By) == 0 {
		_s += dw.p() + "By: []\n"
	} else {
		_s += dw.p() + "By: [\n"
		dw.i()
		for _, _r := range sa.By {
			_s += dw.p() + _r + "\n" // []string
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	_s += dw.p() + "AllowAllMissing: bool(" + fmt.Sprintf("%t", sa.AllowAllMissing) + ")\n"
	if sa.AllowMissing == nil {
		_s += dw.p() + "AllowMissing: nil\n"
	} else if len(sa.AllowMissing) == 0 {
		_s += dw.p() + "AllowMissing: []\n"
	} else {
		_s += dw.p() + "AllowMissing: [\n"
		dw.i()
		for _, _r := range sa.AllowMissing {
			_s += dw.p() + _r + "\n" // []string
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamAlerts(sa *StreamAlerts) (any, error) {
	_s := "StreamAlerts(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sa.Object, sa.Object)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamBelow(sb *StreamBelow) (any, error) {
	_s := "StreamBelow(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sb.Object, sb.Object)
	if sb.Source != nil {
		_s += dw.p() + "Source: " + s(sb.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamCombine(sc *StreamCombine) (any, error) {
	_s := "StreamCombine(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sc.Object, sc.Object)
	if sc.Source != nil {
		_s += dw.p() + "Source: " + s(sc.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Mode: string(" + sc.Mode + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamConstDouble(scd *StreamConstDouble) (any, error) {
	_s := "StreamConstDouble(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", scd.Object, scd.Object)
	// TODO: 1 Value float64
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", scd.Value, scd.Value)
	// TODO: 2 Key map[string]string
	_s += dw.p() + fmt.Sprintf("Key: %T(%v)\n", scd.Key, scd.Key)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamConstInt(sci *StreamConstInt) (any, error) {
	_s := "StreamConstInt(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sci.Object, sci.Object)
	// TODO: 1 Value int
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sci.Value, sci.Value)
	// TODO: 2 Key map[string]string
	_s += dw.p() + fmt.Sprintf("Key: %T(%v)\n", sci.Key, sci.Key)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamData(sd *StreamData) (any, error) {
	_s := "StreamData(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sd.Object, sd.Object)
	_s += dw.p() + "MetricName: string(" + sd.MetricName + ")\n"
	// TODO: 2 Filter filter.Filter
	_s += dw.p() + fmt.Sprintf("Filter: %T(%v)\n", sd.Filter, sd.Filter)
	_s += dw.p() + "Rollup: string(" + sd.Rollup + ")\n"
	_s += dw.p() + "Extrapolation: string(" + sd.Extrapolation + ")\n"
	// TODO: 5 MaxExtrapolations int
	_s += dw.p() + fmt.Sprintf("MaxExtrapolations: %T(%v)\n", sd.MaxExtrapolations, sd.MaxExtrapolations)
	// TODO: 6 TimeShift time.Duration
	_s += dw.p() + fmt.Sprintf("TimeShift: %T(%v)\n", sd.TimeShift, sd.TimeShift)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamDetect(sd *StreamDetect) (any, error) {
	_s := "StreamDetect(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sd.Object, sd.Object)
	if sd.On != nil {
		_s += dw.p() + "On: " + s(sd.On.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "On: nil\n"
	}
	if sd.Off != nil {
		_s += dw.p() + "Off: " + s(sd.Off.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Off: nil\n"
	}
	_s += dw.p() + "Mode: string(" + sd.Mode + ")\n"
	// TODO: 4 Annotations interpreter.Object
	_s += dw.p() + fmt.Sprintf("Annotations: %T(%v)\n", sd.Annotations, sd.Annotations)
	// TODO: 5 EventAnnotations interpreter.Object
	_s += dw.p() + fmt.Sprintf("EventAnnotations: %T(%v)\n", sd.EventAnnotations, sd.EventAnnotations)
	// TODO: 6 AutoResolveAfter *time.Duration
	_s += dw.p() + fmt.Sprintf("AutoResolveAfter: %T(%v)\n", sd.AutoResolveAfter, sd.AutoResolveAfter)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamEvents(se *StreamEvents) (any, error) {
	_s := "StreamEvents(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", se.Object, se.Object)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFill(sf *StreamFill) (any, error) {
	_s := "StreamFill(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sf.Object, sf.Object)
	if sf.Source != nil {
		_s += dw.p() + "Source: " + s(sf.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	// TODO: 2 Value interpreter.Object
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sf.Value, sf.Value)
	// TODO: 3 Duration int
	_s += dw.p() + fmt.Sprintf("Duration: %T(%v)\n", sf.Duration, sf.Duration)
	// TODO: 4 MaxCount int
	_s += dw.p() + fmt.Sprintf("MaxCount: %T(%v)\n", sf.MaxCount, sf.MaxCount)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamGeneric(sg *StreamGeneric) (any, error) {
	_s := "StreamGeneric(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sg.Object, sg.Object)
	if sg.Source != nil {
		_s += dw.p() + "Source: " + s(sg.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Call: string(" + sg.Call + ")\n"
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

func (dw DebugWriter) VisitStreamMax(sm *StreamMax) (any, error) {
	_s := "StreamMax(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sm.Object, sm.Object)
	if sm.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sm.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sm.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Value interpreter.Object
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sm.Value, sm.Value)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMean(sm *StreamMean) (any, error) {
	_s := "StreamMean(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sm.Object, sm.Object)
	if sm.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sm.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sm.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Constants []interpreter.Object
	_s += dw.p() + fmt.Sprintf("Constants: %T(%v)\n", sm.Constants, sm.Constants)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMedian(sm *StreamMedian) (any, error) {
	_s := "StreamMedian(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sm.Object, sm.Object)
	if sm.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sm.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sm.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Constants []interpreter.Object
	_s += dw.p() + fmt.Sprintf("Constants: %T(%v)\n", sm.Constants, sm.Constants)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMin(sm *StreamMin) (any, error) {
	_s := "StreamMin(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sm.Object, sm.Object)
	if sm.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sm.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sm.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	// TODO: 2 Value interpreter.Object
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", sm.Value, sm.Value)
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

func (dw DebugWriter) VisitStreamCount(sc *StreamCount) (any, error) {
	_s := "StreamCount(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sc.Object, sc.Object)
	if sc.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(sc.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range sc.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
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

func (dw DebugWriter) VisitStreamPercentile(sp *StreamPercentile) (any, error) {
	_s := "StreamPercentile(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sp.Object, sp.Object)
	if sp.Source != nil {
		_s += dw.p() + "Source: " + s(sp.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamPublish(sp *StreamPublish) (any, error) {
	_s := "StreamPublish(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sp.Object, sp.Object)
	if sp.Source != nil {
		_s += dw.p() + "Source: " + s(sp.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Label: string(" + sp.Label + ")\n"
	_s += dw.p() + "Enable: bool(" + fmt.Sprintf("%t", sp.Enable) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamScale(ss *StreamScale) (any, error) {
	_s := "StreamScale(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", ss.Object, ss.Object)
	if ss.Source != nil {
		_s += dw.p() + "Source: " + s(ss.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	// TODO: 2 Multiple float64
	_s += dw.p() + fmt.Sprintf("Multiple: %T(%v)\n", ss.Multiple, ss.Multiple)
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

func (dw DebugWriter) VisitStreamThreshold(st *StreamThreshold) (any, error) {
	_s := "StreamThreshold(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", st.Object, st.Object)
	// TODO: 1 Value float64
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", st.Value, st.Value)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamTimeShift(sts *StreamTimeShift) (any, error) {
	_s := "StreamTimeShift(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sts.Object, sts.Object)
	if sts.Source != nil {
		_s += dw.p() + "Source: " + s(sts.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	// TODO: 2 Offset time.Duration
	_s += dw.p() + fmt.Sprintf("Offset: %T(%v)\n", sts.Offset, sts.Offset)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamTop(st *StreamTop) (any, error) {
	_s := "StreamTop(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", st.Object, st.Object)
	if st.Source != nil {
		_s += dw.p() + "Source: " + s(st.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamTransform(st *StreamTransform) (any, error) {
	_s := "StreamTransform(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", st.Object, st.Object)
	if st.Source != nil {
		_s += dw.p() + "Source: " + s(st.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Fn: string(" + st.Fn + ")\n"
	// TODO: 3 Over time.Duration
	_s += dw.p() + fmt.Sprintf("Over: %T(%v)\n", st.Over, st.Over)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamTransformCycle(stc *StreamTransformCycle) (any, error) {
	_s := "StreamTransformCycle(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", stc.Object, stc.Object)
	if stc.Source != nil {
		_s += dw.p() + "Source: " + s(stc.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Fn: string(" + stc.Fn + ")\n"
	_s += dw.p() + "Cycle: string(" + stc.Cycle + ")\n"
	// TODO: 4 CycleStart *string
	_s += dw.p() + fmt.Sprintf("CycleStart: %T(%v)\n", stc.CycleStart, stc.CycleStart)
	// TODO: 5 Timezone *string
	_s += dw.p() + fmt.Sprintf("Timezone: %T(%v)\n", stc.Timezone, stc.Timezone)
	_s += dw.p() + "PartialValues: bool(" + fmt.Sprintf("%t", stc.PartialValues) + ")\n"
	// TODO: 7 ShiftCycles int
	_s += dw.p() + fmt.Sprintf("ShiftCycles: %T(%v)\n", stc.ShiftCycles, stc.ShiftCycles)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamUnion(su *StreamUnion) (any, error) {
	_s := "StreamUnion(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", su.Object, su.Object)
	if su.Sources == nil {
		_s += dw.p() + "Sources: nil\n"
	} else if len(su.Sources) == 0 {
		_s += dw.p() + "Sources: []\n"
	} else {
		_s += dw.p() + "Sources: [\n"
		dw.i()
		for _, _r := range su.Sources {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamWhen(sw *StreamWhen) (any, error) {
	_s := "StreamWhen(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sw.Object, sw.Object)
	if sw.Predicate != nil {
		_s += dw.p() + "Predicate: " + s(sw.Predicate.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Predicate: nil\n"
	}
	// TODO: 2 Lasting *time.Duration
	_s += dw.p() + fmt.Sprintf("Lasting: %T(%v)\n", sw.Lasting, sw.Lasting)
	// TODO: 3 AtLeast float64
	_s += dw.p() + fmt.Sprintf("AtLeast: %T(%v)\n", sw.AtLeast, sw.AtLeast)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}
