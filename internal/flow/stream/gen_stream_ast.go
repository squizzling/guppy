package stream

import (
	"guppy/internal/flow/filter"
	"guppy/internal/interpreter"
)

type VisitorStream interface {
	VisitStreamAbove(sa StreamAbove) (any, error)
	VisitStreamAbs(sa StreamAbs) (any, error)
	VisitStreamAggregate(sa StreamAggregate) (any, error)
	VisitStreamAlerts(sa StreamAlerts) (any, error)
	VisitStreamBelow(sb StreamBelow) (any, error)
	VisitStreamConstDouble(scd StreamConstDouble) (any, error)
	VisitStreamConstInt(sci StreamConstInt) (any, error)
	VisitStreamData(sd StreamData) (any, error)
	VisitStreamEvents(se StreamEvents) (any, error)
	VisitStreamFill(sf StreamFill) (any, error)
	VisitStreamGeneric(sg StreamGeneric) (any, error)
	VisitStreamMax(sm StreamMax) (any, error)
	VisitStreamMathOpDouble(smod StreamMathOpDouble) (any, error)
	VisitStreamMathOpInt(smoi StreamMathOpInt) (any, error)
	VisitStreamMathOpStream(smos StreamMathOpStream) (any, error)
	VisitStreamMathOpUnaryMinus(smoum StreamMathOpUnaryMinus) (any, error)
	VisitStreamPercentile(sp StreamPercentile) (any, error)
	VisitStreamPublish(sp StreamPublish) (any, error)
	VisitStreamScale(ss StreamScale) (any, error)
	VisitStreamThreshold(st StreamThreshold) (any, error)
	VisitStreamTimeShift(sts StreamTimeShift) (any, error)
	VisitStreamTop(st StreamTop) (any, error)
	VisitStreamTransform(st StreamTransform) (any, error)
	VisitStreamUnion(su StreamUnion) (any, error)
}

type Stream interface {
	interpreter.Object
	Accept(vs VisitorStream) (any, error)
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

func (sa StreamAbove) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamAbove(sa)
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

func (sa StreamAbs) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamAbs(sa)
}

type StreamAggregate struct {
	interpreter.Object
	Source Stream
	Fn     string
	By     []string
}

func NewStreamAggregate(
	Object interpreter.Object,
	Source Stream,
	Fn string,
	By []string,
) *StreamAggregate {
	return &StreamAggregate{
		Object: Object,
		Source: Source,
		Fn:     Fn,
		By:     By,
	}
}

func (sa StreamAggregate) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamAggregate(sa)
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

func (sa StreamAlerts) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamAlerts(sa)
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

func (sb StreamBelow) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamBelow(sb)
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

func (scd StreamConstDouble) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamConstDouble(scd)
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

func (sci StreamConstInt) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamConstInt(sci)
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

func (sd StreamData) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamData(sd)
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

func (se StreamEvents) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamEvents(se)
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

func (sf StreamFill) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFill(sf)
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

func (sg StreamGeneric) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamGeneric(sg)
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

func (sm StreamMax) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMax(sm)
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

func (smod StreamMathOpDouble) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMathOpDouble(smod)
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

func (smoi StreamMathOpInt) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMathOpInt(smoi)
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

func (smos StreamMathOpStream) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMathOpStream(smos)
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

func (smoum StreamMathOpUnaryMinus) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMathOpUnaryMinus(smoum)
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

func (sp StreamPercentile) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamPercentile(sp)
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

func (sp StreamPublish) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamPublish(sp)
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

func (ss StreamScale) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamScale(ss)
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

func (st StreamThreshold) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamThreshold(st)
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

func (sts StreamTimeShift) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTimeShift(sts)
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

func (st StreamTop) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTop(st)
}

type StreamTransform struct {
	interpreter.Object
	Source Stream
	Fn     string
	Over   string
}

func NewStreamTransform(
	Object interpreter.Object,
	Source Stream,
	Fn string,
	Over string,
) *StreamTransform {
	return &StreamTransform{
		Object: Object,
		Source: Source,
		Fn:     Fn,
		Over:   Over,
	}
}

func (st StreamTransform) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTransform(st)
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

func (su StreamUnion) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamUnion(su)
}
