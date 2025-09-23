package stream

import (
	"time"

	"guppy/internal/flow/filter"
	"guppy/internal/interpreter"
)

type VisitorStream interface {
	VisitStreamAbove(sa *StreamAbove) (any, error)
	VisitStreamAbs(sa *StreamAbs) (any, error)
	VisitStreamAggregate(sa *StreamAggregate) (any, error)
	VisitStreamAlerts(sa *StreamAlerts) (any, error)
	VisitStreamBelow(sb *StreamBelow) (any, error)
	VisitStreamConstDouble(scd *StreamConstDouble) (any, error)
	VisitStreamConstInt(sci *StreamConstInt) (any, error)
	VisitStreamData(sd *StreamData) (any, error)
	VisitStreamDetect(sd *StreamDetect) (any, error)
	VisitStreamEvents(se *StreamEvents) (any, error)
	VisitStreamFill(sf *StreamFill) (any, error)
	VisitStreamGeneric(sg *StreamGeneric) (any, error)
	VisitStreamIsNone(sin *StreamIsNone) (any, error)
	VisitStreamMax(sm *StreamMax) (any, error)
	VisitStreamMean(sm *StreamMean) (any, error)
	VisitStreamMedian(sm *StreamMedian) (any, error)
	VisitStreamMin(sm *StreamMin) (any, error)
	VisitStreamMathOpDouble(smod *StreamMathOpDouble) (any, error)
	VisitStreamMathOpInt(smoi *StreamMathOpInt) (any, error)
	VisitStreamMathOpStream(smos *StreamMathOpStream) (any, error)
	VisitStreamMathOpUnaryMinus(smoum *StreamMathOpUnaryMinus) (any, error)
	VisitStreamPercentile(sp *StreamPercentile) (any, error)
	VisitStreamPublish(sp *StreamPublish) (any, error)
	VisitStreamScale(ss *StreamScale) (any, error)
	VisitStreamTernary(st *StreamTernary) (any, error)
	VisitStreamThreshold(st *StreamThreshold) (any, error)
	VisitStreamTimeShift(sts *StreamTimeShift) (any, error)
	VisitStreamTop(st *StreamTop) (any, error)
	VisitStreamTransform(st *StreamTransform) (any, error)
	VisitStreamTransformCycle(stc *StreamTransformCycle) (any, error)
	VisitStreamUnion(su *StreamUnion) (any, error)
	VisitStreamWhen(sw *StreamWhen) (any, error)
}

type Stream interface {
	interpreter.Object
	Accept(vs VisitorStream) (any, error)
	Clone() Stream
}

type StreamAbove struct {
	interpreter.Object
	Source Stream
}

func NewStreamAbove(
	Object interpreter.Object,
	Source Stream,
) *StreamAbove {
	return &StreamAbove{
		Object: Object,
		Source: Source,
	}
}

func (sa *StreamAbove) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamAbove(sa)
}

func (sa *StreamAbove) Clone() Stream {
	return &StreamAbove{
		Object: sa.Object,
		Source: sa.Source.Clone(),
	}
}

type StreamAbs struct {
	interpreter.Object
	Source Stream
}

func NewStreamAbs(
	Object interpreter.Object,
	Source Stream,
) *StreamAbs {
	return &StreamAbs{
		Object: Object,
		Source: Source,
	}
}

func (sa *StreamAbs) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamAbs(sa)
}

func (sa *StreamAbs) Clone() Stream {
	return &StreamAbs{
		Object: sa.Object,
		Source: sa.Source.Clone(),
	}
}

type StreamAggregate struct {
	interpreter.Object
	Source          Stream
	Fn              string
	By              []string
	AllowAllMissing bool
	AllowMissing    []string
}

func NewStreamAggregate(
	Object interpreter.Object,
	Source Stream,
	Fn string,
	By []string,
	AllowAllMissing bool,
	AllowMissing []string,
) *StreamAggregate {
	return &StreamAggregate{
		Object:          Object,
		Source:          Source,
		Fn:              Fn,
		By:              By,
		AllowAllMissing: AllowAllMissing,
		AllowMissing:    AllowMissing,
	}
}

func (sa *StreamAggregate) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamAggregate(sa)
}

func (sa *StreamAggregate) Clone() Stream {
	return &StreamAggregate{
		Object:          sa.Object,
		Source:          sa.Source.Clone(),
		Fn:              sa.Fn,
		By:              sa.By,
		AllowAllMissing: sa.AllowAllMissing,
		AllowMissing:    sa.AllowMissing,
	}
}

type StreamAlerts struct {
	interpreter.Object
}

func NewStreamAlerts(
	Object interpreter.Object,
) *StreamAlerts {
	return &StreamAlerts{
		Object: Object,
	}
}

func (sa *StreamAlerts) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamAlerts(sa)
}

func (sa *StreamAlerts) Clone() Stream {
	return &StreamAlerts{
		Object: sa.Object,
	}
}

type StreamBelow struct {
	interpreter.Object
	Source Stream
}

func NewStreamBelow(
	Object interpreter.Object,
	Source Stream,
) *StreamBelow {
	return &StreamBelow{
		Object: Object,
		Source: Source,
	}
}

func (sb *StreamBelow) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamBelow(sb)
}

func (sb *StreamBelow) Clone() Stream {
	return &StreamBelow{
		Object: sb.Object,
		Source: sb.Source.Clone(),
	}
}

type StreamConstDouble struct {
	interpreter.Object
	Value float64
	Key   map[string]string
}

func NewStreamConstDouble(
	Object interpreter.Object,
	Value float64,
	Key map[string]string,
) *StreamConstDouble {
	return &StreamConstDouble{
		Object: Object,
		Value:  Value,
		Key:    Key,
	}
}

func (scd *StreamConstDouble) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamConstDouble(scd)
}

func (scd *StreamConstDouble) Clone() Stream {
	return &StreamConstDouble{
		Object: scd.Object,
		Value:  scd.Value,
		Key:    scd.Key,
	}
}

type StreamConstInt struct {
	interpreter.Object
	Value int
	Key   map[string]string
}

func NewStreamConstInt(
	Object interpreter.Object,
	Value int,
	Key map[string]string,
) *StreamConstInt {
	return &StreamConstInt{
		Object: Object,
		Value:  Value,
		Key:    Key,
	}
}

func (sci *StreamConstInt) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamConstInt(sci)
}

func (sci *StreamConstInt) Clone() Stream {
	return &StreamConstInt{
		Object: sci.Object,
		Value:  sci.Value,
		Key:    sci.Key,
	}
}

type StreamData struct {
	interpreter.Object
	MetricName        string
	Filter            filter.Filter
	Rollup            string
	Extrapolation     string
	MaxExtrapolations int
}

func NewStreamData(
	Object interpreter.Object,
	MetricName string,
	Filter filter.Filter,
	Rollup string,
	Extrapolation string,
	MaxExtrapolations int,
) *StreamData {
	return &StreamData{
		Object:            Object,
		MetricName:        MetricName,
		Filter:            Filter,
		Rollup:            Rollup,
		Extrapolation:     Extrapolation,
		MaxExtrapolations: MaxExtrapolations,
	}
}

func (sd *StreamData) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamData(sd)
}

func (sd *StreamData) Clone() Stream {
	return &StreamData{
		Object:            sd.Object,
		MetricName:        sd.MetricName,
		Filter:            sd.Filter,
		Rollup:            sd.Rollup,
		Extrapolation:     sd.Extrapolation,
		MaxExtrapolations: sd.MaxExtrapolations,
	}
}

type StreamDetect struct {
	interpreter.Object
	On               Stream
	Off              Stream
	Mode             string
	Annotations      interpreter.Object
	EventAnnotations interpreter.Object
	AutoResolveAfter *time.Duration
}

func NewStreamDetect(
	Object interpreter.Object,
	On Stream,
	Off Stream,
	Mode string,
	Annotations interpreter.Object,
	EventAnnotations interpreter.Object,
	AutoResolveAfter *time.Duration,
) *StreamDetect {
	return &StreamDetect{
		Object:           Object,
		On:               On,
		Off:              Off,
		Mode:             Mode,
		Annotations:      Annotations,
		EventAnnotations: EventAnnotations,
		AutoResolveAfter: AutoResolveAfter,
	}
}

func (sd *StreamDetect) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamDetect(sd)
}

func (sd *StreamDetect) Clone() Stream {
	return &StreamDetect{
		Object:           sd.Object,
		On:               sd.On.Clone(),
		Off:              sd.Off.Clone(),
		Mode:             sd.Mode,
		Annotations:      sd.Annotations,
		EventAnnotations: sd.EventAnnotations,
		AutoResolveAfter: sd.AutoResolveAfter,
	}
}

type StreamEvents struct {
	interpreter.Object
}

func NewStreamEvents(
	Object interpreter.Object,
) *StreamEvents {
	return &StreamEvents{
		Object: Object,
	}
}

func (se *StreamEvents) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamEvents(se)
}

func (se *StreamEvents) Clone() Stream {
	return &StreamEvents{
		Object: se.Object,
	}
}

type StreamFill struct {
	interpreter.Object
	Source   Stream
	Value    interpreter.Object
	Duration int
	MaxCount int
}

func NewStreamFill(
	Object interpreter.Object,
	Source Stream,
	Value interpreter.Object,
	Duration int,
	MaxCount int,
) *StreamFill {
	return &StreamFill{
		Object:   Object,
		Source:   Source,
		Value:    Value,
		Duration: Duration,
		MaxCount: MaxCount,
	}
}

func (sf *StreamFill) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFill(sf)
}

func (sf *StreamFill) Clone() Stream {
	return &StreamFill{
		Object:   sf.Object,
		Source:   sf.Source.Clone(),
		Value:    sf.Value,
		Duration: sf.Duration,
		MaxCount: sf.MaxCount,
	}
}

type StreamGeneric struct {
	interpreter.Object
	Source Stream
	Call   string
}

func NewStreamGeneric(
	Object interpreter.Object,
	Source Stream,
	Call string,
) *StreamGeneric {
	return &StreamGeneric{
		Object: Object,
		Source: Source,
		Call:   Call,
	}
}

func (sg *StreamGeneric) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamGeneric(sg)
}

func (sg *StreamGeneric) Clone() Stream {
	return &StreamGeneric{
		Object: sg.Object,
		Source: sg.Source.Clone(),
		Call:   sg.Call,
	}
}

type StreamIsNone struct {
	interpreter.Object
	Source Stream
	Invert bool
}

func NewStreamIsNone(
	Object interpreter.Object,
	Source Stream,
	Invert bool,
) *StreamIsNone {
	return &StreamIsNone{
		Object: Object,
		Source: Source,
		Invert: Invert,
	}
}

func (sin *StreamIsNone) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamIsNone(sin)
}

func (sin *StreamIsNone) Clone() Stream {
	return &StreamIsNone{
		Object: sin.Object,
		Source: sin.Source.Clone(),
		Invert: sin.Invert,
	}
}

type StreamMax struct {
	interpreter.Object
	Sources []Stream
	Value   interpreter.Object
}

func NewStreamMax(
	Object interpreter.Object,
	Sources []Stream,
	Value interpreter.Object,
) *StreamMax {
	return &StreamMax{
		Object:  Object,
		Sources: Sources,
		Value:   Value,
	}
}

func (sm *StreamMax) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMax(sm)
}

func (sm *StreamMax) Clone() Stream {
	return &StreamMax{
		Object:  sm.Object,
		Sources: sm.Sources,
		Value:   sm.Value,
	}
}

type StreamMean struct {
	interpreter.Object
	Sources   []Stream
	Constants []interpreter.Object
}

func NewStreamMean(
	Object interpreter.Object,
	Sources []Stream,
	Constants []interpreter.Object,
) *StreamMean {
	return &StreamMean{
		Object:    Object,
		Sources:   Sources,
		Constants: Constants,
	}
}

func (sm *StreamMean) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMean(sm)
}

func (sm *StreamMean) Clone() Stream {
	return &StreamMean{
		Object:    sm.Object,
		Sources:   sm.Sources,
		Constants: sm.Constants,
	}
}

type StreamMedian struct {
	interpreter.Object
	Sources   []Stream
	Constants []interpreter.Object
}

func NewStreamMedian(
	Object interpreter.Object,
	Sources []Stream,
	Constants []interpreter.Object,
) *StreamMedian {
	return &StreamMedian{
		Object:    Object,
		Sources:   Sources,
		Constants: Constants,
	}
}

func (sm *StreamMedian) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMedian(sm)
}

func (sm *StreamMedian) Clone() Stream {
	return &StreamMedian{
		Object:    sm.Object,
		Sources:   sm.Sources,
		Constants: sm.Constants,
	}
}

type StreamMin struct {
	interpreter.Object
	Sources []Stream
	Value   interpreter.Object
}

func NewStreamMin(
	Object interpreter.Object,
	Sources []Stream,
	Value interpreter.Object,
) *StreamMin {
	return &StreamMin{
		Object:  Object,
		Sources: Sources,
		Value:   Value,
	}
}

func (sm *StreamMin) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMin(sm)
}

func (sm *StreamMin) Clone() Stream {
	return &StreamMin{
		Object:  sm.Object,
		Sources: sm.Sources,
		Value:   sm.Value,
	}
}

type StreamMathOpDouble struct {
	interpreter.Object
	Stream  Stream
	Op      string
	Other   float64
	Reverse bool
}

func NewStreamMathOpDouble(
	Object interpreter.Object,
	Stream Stream,
	Op string,
	Other float64,
	Reverse bool,
) *StreamMathOpDouble {
	return &StreamMathOpDouble{
		Object:  Object,
		Stream:  Stream,
		Op:      Op,
		Other:   Other,
		Reverse: Reverse,
	}
}

func (smod *StreamMathOpDouble) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMathOpDouble(smod)
}

func (smod *StreamMathOpDouble) Clone() Stream {
	return &StreamMathOpDouble{
		Object:  smod.Object,
		Stream:  smod.Stream.Clone(),
		Op:      smod.Op,
		Other:   smod.Other,
		Reverse: smod.Reverse,
	}
}

type StreamMathOpInt struct {
	interpreter.Object
	Stream  Stream
	Op      string
	Other   int
	Reverse bool
}

func NewStreamMathOpInt(
	Object interpreter.Object,
	Stream Stream,
	Op string,
	Other int,
	Reverse bool,
) *StreamMathOpInt {
	return &StreamMathOpInt{
		Object:  Object,
		Stream:  Stream,
		Op:      Op,
		Other:   Other,
		Reverse: Reverse,
	}
}

func (smoi *StreamMathOpInt) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMathOpInt(smoi)
}

func (smoi *StreamMathOpInt) Clone() Stream {
	return &StreamMathOpInt{
		Object:  smoi.Object,
		Stream:  smoi.Stream.Clone(),
		Op:      smoi.Op,
		Other:   smoi.Other,
		Reverse: smoi.Reverse,
	}
}

type StreamMathOpStream struct {
	interpreter.Object
	Left  Stream
	Op    string
	Right Stream
}

func NewStreamMathOpStream(
	Object interpreter.Object,
	Left Stream,
	Op string,
	Right Stream,
) *StreamMathOpStream {
	return &StreamMathOpStream{
		Object: Object,
		Left:   Left,
		Op:     Op,
		Right:  Right,
	}
}

func (smos *StreamMathOpStream) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMathOpStream(smos)
}

func (smos *StreamMathOpStream) Clone() Stream {
	return &StreamMathOpStream{
		Object: smos.Object,
		Left:   smos.Left.Clone(),
		Op:     smos.Op,
		Right:  smos.Right.Clone(),
	}
}

type StreamMathOpUnaryMinus struct {
	interpreter.Object
	Stream Stream
}

func NewStreamMathOpUnaryMinus(
	Object interpreter.Object,
	Stream Stream,
) *StreamMathOpUnaryMinus {
	return &StreamMathOpUnaryMinus{
		Object: Object,
		Stream: Stream,
	}
}

func (smoum *StreamMathOpUnaryMinus) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMathOpUnaryMinus(smoum)
}

func (smoum *StreamMathOpUnaryMinus) Clone() Stream {
	return &StreamMathOpUnaryMinus{
		Object: smoum.Object,
		Stream: smoum.Stream.Clone(),
	}
}

type StreamPercentile struct {
	interpreter.Object
	Source Stream
}

func NewStreamPercentile(
	Object interpreter.Object,
	Source Stream,
) *StreamPercentile {
	return &StreamPercentile{
		Object: Object,
		Source: Source,
	}
}

func (sp *StreamPercentile) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamPercentile(sp)
}

func (sp *StreamPercentile) Clone() Stream {
	return &StreamPercentile{
		Object: sp.Object,
		Source: sp.Source.Clone(),
	}
}

type StreamPublish struct {
	interpreter.Object
	Source Stream
	Label  string
	Enable bool
}

func NewStreamPublish(
	Object interpreter.Object,
	Source Stream,
	Label string,
	Enable bool,
) *StreamPublish {
	return &StreamPublish{
		Object: Object,
		Source: Source,
		Label:  Label,
		Enable: Enable,
	}
}

func (sp *StreamPublish) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamPublish(sp)
}

func (sp *StreamPublish) Clone() Stream {
	return &StreamPublish{
		Object: sp.Object,
		Source: sp.Source.Clone(),
		Label:  sp.Label,
		Enable: sp.Enable,
	}
}

type StreamScale struct {
	interpreter.Object
	Source   Stream
	Multiple float64
}

func NewStreamScale(
	Object interpreter.Object,
	Source Stream,
	Multiple float64,
) *StreamScale {
	return &StreamScale{
		Object:   Object,
		Source:   Source,
		Multiple: Multiple,
	}
}

func (ss *StreamScale) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamScale(ss)
}

func (ss *StreamScale) Clone() Stream {
	return &StreamScale{
		Object:   ss.Object,
		Source:   ss.Source.Clone(),
		Multiple: ss.Multiple,
	}
}

type StreamTernary struct {
	interpreter.Object
	Condition Stream
	Left      Stream
	Right     Stream
}

func NewStreamTernary(
	Object interpreter.Object,
	Condition Stream,
	Left Stream,
	Right Stream,
) *StreamTernary {
	return &StreamTernary{
		Object:    Object,
		Condition: Condition,
		Left:      Left,
		Right:     Right,
	}
}

func (st *StreamTernary) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTernary(st)
}

func (st *StreamTernary) Clone() Stream {
	return &StreamTernary{
		Object:    st.Object,
		Condition: st.Condition.Clone(),
		Left:      st.Left.Clone(),
		Right:     st.Right.Clone(),
	}
}

type StreamThreshold struct {
	interpreter.Object
	Value float64
}

func NewStreamThreshold(
	Object interpreter.Object,
	Value float64,
) *StreamThreshold {
	return &StreamThreshold{
		Object: Object,
		Value:  Value,
	}
}

func (st *StreamThreshold) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamThreshold(st)
}

func (st *StreamThreshold) Clone() Stream {
	return &StreamThreshold{
		Object: st.Object,
		Value:  st.Value,
	}
}

type StreamTimeShift struct {
	interpreter.Object
	Source Stream
	Offset string
}

func NewStreamTimeShift(
	Object interpreter.Object,
	Source Stream,
	Offset string,
) *StreamTimeShift {
	return &StreamTimeShift{
		Object: Object,
		Source: Source,
		Offset: Offset,
	}
}

func (sts *StreamTimeShift) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTimeShift(sts)
}

func (sts *StreamTimeShift) Clone() Stream {
	return &StreamTimeShift{
		Object: sts.Object,
		Source: sts.Source.Clone(),
		Offset: sts.Offset,
	}
}

type StreamTop struct {
	interpreter.Object
	Source Stream
}

func NewStreamTop(
	Object interpreter.Object,
	Source Stream,
) *StreamTop {
	return &StreamTop{
		Object: Object,
		Source: Source,
	}
}

func (st *StreamTop) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTop(st)
}

func (st *StreamTop) Clone() Stream {
	return &StreamTop{
		Object: st.Object,
		Source: st.Source.Clone(),
	}
}

type StreamTransform struct {
	interpreter.Object
	Source Stream
	Fn     string
	Over   time.Duration
}

func NewStreamTransform(
	Object interpreter.Object,
	Source Stream,
	Fn string,
	Over time.Duration,
) *StreamTransform {
	return &StreamTransform{
		Object: Object,
		Source: Source,
		Fn:     Fn,
		Over:   Over,
	}
}

func (st *StreamTransform) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTransform(st)
}

func (st *StreamTransform) Clone() Stream {
	return &StreamTransform{
		Object: st.Object,
		Source: st.Source.Clone(),
		Fn:     st.Fn,
		Over:   st.Over,
	}
}

type StreamTransformCycle struct {
	interpreter.Object
	Source        Stream
	Fn            string
	Cycle         string
	CycleStart    *string
	Timezone      *string
	PartialValues bool
	ShiftCycles   int
}

func NewStreamTransformCycle(
	Object interpreter.Object,
	Source Stream,
	Fn string,
	Cycle string,
	CycleStart *string,
	Timezone *string,
	PartialValues bool,
	ShiftCycles int,
) *StreamTransformCycle {
	return &StreamTransformCycle{
		Object:        Object,
		Source:        Source,
		Fn:            Fn,
		Cycle:         Cycle,
		CycleStart:    CycleStart,
		Timezone:      Timezone,
		PartialValues: PartialValues,
		ShiftCycles:   ShiftCycles,
	}
}

func (stc *StreamTransformCycle) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTransformCycle(stc)
}

func (stc *StreamTransformCycle) Clone() Stream {
	return &StreamTransformCycle{
		Object:        stc.Object,
		Source:        stc.Source.Clone(),
		Fn:            stc.Fn,
		Cycle:         stc.Cycle,
		CycleStart:    stc.CycleStart,
		Timezone:      stc.Timezone,
		PartialValues: stc.PartialValues,
		ShiftCycles:   stc.ShiftCycles,
	}
}

type StreamUnion struct {
	interpreter.Object
	Sources []Stream
}

func NewStreamUnion(
	Object interpreter.Object,
	Sources []Stream,
) *StreamUnion {
	return &StreamUnion{
		Object:  Object,
		Sources: Sources,
	}
}

func (su *StreamUnion) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamUnion(su)
}

func (su *StreamUnion) Clone() Stream {
	return &StreamUnion{
		Object:  su.Object,
		Sources: su.Sources,
	}
}

type StreamWhen struct {
	interpreter.Object
	Predicate Stream
	Lasting   *time.Duration
	AtLeast   float64
}

func NewStreamWhen(
	Object interpreter.Object,
	Predicate Stream,
	Lasting *time.Duration,
	AtLeast float64,
) *StreamWhen {
	return &StreamWhen{
		Object:    Object,
		Predicate: Predicate,
		Lasting:   Lasting,
		AtLeast:   AtLeast,
	}
}

func (sw *StreamWhen) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamWhen(sw)
}

func (sw *StreamWhen) Clone() Stream {
	return &StreamWhen{
		Object:    sw.Object,
		Predicate: sw.Predicate.Clone(),
		Lasting:   sw.Lasting,
		AtLeast:   sw.AtLeast,
	}
}
