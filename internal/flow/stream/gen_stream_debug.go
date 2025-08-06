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

func (dw DebugWriter) VisitStreamAbove(sa StreamAbove) (any, error) {
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

func (dw DebugWriter) VisitStreamAbs(sa StreamAbs) (any, error) {
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

func (dw DebugWriter) VisitStreamAggregate(sa StreamAggregate) (any, error) {
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
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamAlerts(sa StreamAlerts) (any, error) {
	_s := "StreamAlerts(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sa.Object, sa.Object)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamBelow(sb StreamBelow) (any, error) {
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

func (dw DebugWriter) VisitStreamConstDouble(scd StreamConstDouble) (any, error) {
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

func (dw DebugWriter) VisitStreamConstInt(sci StreamConstInt) (any, error) {
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

func (dw DebugWriter) VisitStreamData(sd StreamData) (any, error) {
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
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamEvents(se StreamEvents) (any, error) {
	_s := "StreamEvents(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", se.Object, se.Object)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamFill(sf StreamFill) (any, error) {
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

func (dw DebugWriter) VisitStreamGeneric(sg StreamGeneric) (any, error) {
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

func (dw DebugWriter) VisitStreamMax(sm StreamMax) (any, error) {
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

func (dw DebugWriter) VisitStreamMathOpDouble(smod StreamMathOpDouble) (any, error) {
	_s := "StreamMathOpDouble(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smod.Object, smod.Object)
	if smod.Stream != nil {
		_s += dw.p() + "Stream: " + s(smod.Stream.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Stream: nil\n"
	}
	_s += dw.p() + "Op: string(" + smod.Op + ")\n"
	// TODO: 3 Other float64
	_s += dw.p() + fmt.Sprintf("Other: %T(%v)\n", smod.Other, smod.Other)
	_s += dw.p() + "Reverse: bool(" + fmt.Sprintf("%t", smod.Reverse) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMathOpInt(smoi StreamMathOpInt) (any, error) {
	_s := "StreamMathOpInt(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smoi.Object, smoi.Object)
	if smoi.Stream != nil {
		_s += dw.p() + "Stream: " + s(smoi.Stream.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Stream: nil\n"
	}
	_s += dw.p() + "Op: string(" + smoi.Op + ")\n"
	// TODO: 3 Other int
	_s += dw.p() + fmt.Sprintf("Other: %T(%v)\n", smoi.Other, smoi.Other)
	_s += dw.p() + "Reverse: bool(" + fmt.Sprintf("%t", smoi.Reverse) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMathOpStream(smos StreamMathOpStream) (any, error) {
	_s := "StreamMathOpStream(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smos.Object, smos.Object)
	if smos.Left != nil {
		_s += dw.p() + "Left: " + s(smos.Left.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Left: nil\n"
	}
	_s += dw.p() + "Op: string(" + smos.Op + ")\n"
	if smos.Right != nil {
		_s += dw.p() + "Right: " + s(smos.Right.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamMathOpUnaryMinus(smoum StreamMathOpUnaryMinus) (any, error) {
	_s := "StreamMathOpUnaryMinus(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", smoum.Object, smoum.Object)
	if smoum.Stream != nil {
		_s += dw.p() + "Stream: " + s(smoum.Stream.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Stream: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamPercentile(sp StreamPercentile) (any, error) {
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

func (dw DebugWriter) VisitStreamPublish(sp StreamPublish) (any, error) {
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

func (dw DebugWriter) VisitStreamScale(ss StreamScale) (any, error) {
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

func (dw DebugWriter) VisitStreamThreshold(st StreamThreshold) (any, error) {
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

func (dw DebugWriter) VisitStreamTimeShift(sts StreamTimeShift) (any, error) {
	_s := "StreamTimeShift(\n"
	dw.i()
	// TODO: 0 Object interpreter.Object
	_s += dw.p() + fmt.Sprintf("Object: %T(%v)\n", sts.Object, sts.Object)
	if sts.Source != nil {
		_s += dw.p() + "Source: " + s(sts.Source.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Source: nil\n"
	}
	_s += dw.p() + "Offset: string(" + sts.Offset + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamTop(st StreamTop) (any, error) {
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

func (dw DebugWriter) VisitStreamTransform(st StreamTransform) (any, error) {
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
	_s += dw.p() + "Over: string(" + st.Over + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStreamUnion(su StreamUnion) (any, error) {
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
