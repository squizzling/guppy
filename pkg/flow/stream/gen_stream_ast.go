package stream

import (
	"time"

	"guppy/pkg/flow/filter"
	"guppy/pkg/interpreter"
)

type VisitorStream interface {
	VisitStreamFuncAbs(sfa *StreamFuncAbs) (any, error)
	VisitStreamFuncAlerts(sfa *StreamFuncAlerts) (any, error)
	VisitStreamFuncCombine(sfc *StreamFuncCombine) (any, error)
	VisitStreamFuncConstDouble(sfcd *StreamFuncConstDouble) (any, error)
	VisitStreamFuncConstInt(sfci *StreamFuncConstInt) (any, error)
	VisitStreamFuncCount(sfc *StreamFuncCount) (any, error)
	VisitStreamFuncData(sfd *StreamFuncData) (any, error)
	VisitStreamFuncDetect(sfd *StreamFuncDetect) (any, error)
	VisitStreamFuncEvents(sfe *StreamFuncEvents) (any, error)
	VisitStreamFuncMax(sfm *StreamFuncMax) (any, error)
	VisitStreamFuncMean(sfm *StreamFuncMean) (any, error)
	VisitStreamFuncMedian(sfm *StreamFuncMedian) (any, error)
	VisitStreamFuncMin(sfm *StreamFuncMin) (any, error)
	VisitStreamFuncSum(sfs *StreamFuncSum) (any, error)
	VisitStreamFuncThreshold(sft *StreamFuncThreshold) (any, error)
	VisitStreamFuncUnion(sfu *StreamFuncUnion) (any, error)
	VisitStreamFuncWhen(sfw *StreamFuncWhen) (any, error)
	VisitStreamMethodAbove(sma *StreamMethodAbove) (any, error)
	VisitStreamMethodAbs(sma *StreamMethodAbs) (any, error)
	VisitStreamMethodAggregate(sma *StreamMethodAggregate) (any, error)
	VisitStreamMethodBelow(smb *StreamMethodBelow) (any, error)
	VisitStreamMethodFill(smf *StreamMethodFill) (any, error)
	VisitStreamMethodGeneric(smg *StreamMethodGeneric) (any, error)
	VisitStreamMethodPercentile(smp *StreamMethodPercentile) (any, error)
	VisitStreamMethodPublish(smp *StreamMethodPublish) (any, error)
	VisitStreamMethodScale(sms *StreamMethodScale) (any, error)
	VisitStreamMethodTimeShift(smts *StreamMethodTimeShift) (any, error)
	VisitStreamMethodTop(smt *StreamMethodTop) (any, error)
	VisitStreamMethodTransform(smt *StreamMethodTransform) (any, error)
	VisitStreamMethodTransformCycle(smtc *StreamMethodTransformCycle) (any, error)
	VisitStreamBinaryOpDouble(sbod *StreamBinaryOpDouble) (any, error)
	VisitStreamBinaryOpInt(sboi *StreamBinaryOpInt) (any, error)
	VisitStreamBinaryOpStream(sbos *StreamBinaryOpStream) (any, error)
	VisitStreamIsNone(sin *StreamIsNone) (any, error)
	VisitStreamTernary(st *StreamTernary) (any, error)
	VisitStreamUnaryOpMinus(suom *StreamUnaryOpMinus) (any, error)
}

type Stream interface {
	interpreter.Object
	Accept(vs VisitorStream) (any, error)
	CloneTimeShift(amount time.Duration) Stream
}

type StreamFuncAbs struct {
	interpreter.Object
	Sources []Stream
}

func NewStreamFuncAbs(
	Object interpreter.Object,
	Sources []Stream,
) *StreamFuncAbs {
	return &StreamFuncAbs{
		Object:  Object,
		Sources: Sources,
	}
}

func (sfa *StreamFuncAbs) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncAbs(sfa)
}

func (sfa *StreamFuncAbs) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncAbs{
		Object:  sfa.Object,
		Sources: sfa.Sources,
	}
}

type StreamFuncAlerts struct {
	interpreter.Object
}

func NewStreamFuncAlerts(
	Object interpreter.Object,
) *StreamFuncAlerts {
	return &StreamFuncAlerts{
		Object: Object,
	}
}

func (sfa *StreamFuncAlerts) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncAlerts(sfa)
}

func (sfa *StreamFuncAlerts) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncAlerts{
		Object: sfa.Object,
	}
}

type StreamFuncCombine struct {
	interpreter.Object
	Source Stream
	Mode   string
}

func NewStreamFuncCombine(
	Object interpreter.Object,
	Source Stream,
	Mode string,
) *StreamFuncCombine {
	return &StreamFuncCombine{
		Object: Object,
		Source: Source,
		Mode:   Mode,
	}
}

func (sfc *StreamFuncCombine) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncCombine(sfc)
}

func (sfc *StreamFuncCombine) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncCombine{
		Object: sfc.Object,
		Source: cloneTimeshift(sfc.Source, amount),
		Mode:   sfc.Mode,
	}
}

type StreamFuncConstDouble struct {
	interpreter.Object
	Value float64
	Key   map[string]string
}

func NewStreamFuncConstDouble(
	Object interpreter.Object,
	Value float64,
	Key map[string]string,
) *StreamFuncConstDouble {
	return &StreamFuncConstDouble{
		Object: Object,
		Value:  Value,
		Key:    Key,
	}
}

func (sfcd *StreamFuncConstDouble) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncConstDouble(sfcd)
}

func (sfcd *StreamFuncConstDouble) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncConstDouble{
		Object: sfcd.Object,
		Value:  sfcd.Value,
		Key:    sfcd.Key,
	}
}

type StreamFuncConstInt struct {
	interpreter.Object
	Value int
	Key   map[string]string
}

func NewStreamFuncConstInt(
	Object interpreter.Object,
	Value int,
	Key map[string]string,
) *StreamFuncConstInt {
	return &StreamFuncConstInt{
		Object: Object,
		Value:  Value,
		Key:    Key,
	}
}

func (sfci *StreamFuncConstInt) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncConstInt(sfci)
}

func (sfci *StreamFuncConstInt) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncConstInt{
		Object: sfci.Object,
		Value:  sfci.Value,
		Key:    sfci.Key,
	}
}

type StreamFuncCount struct {
	interpreter.Object
	Sources []Stream
}

func NewStreamFuncCount(
	Object interpreter.Object,
	Sources []Stream,
) *StreamFuncCount {
	return &StreamFuncCount{
		Object:  Object,
		Sources: Sources,
	}
}

func (sfc *StreamFuncCount) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncCount(sfc)
}

func (sfc *StreamFuncCount) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncCount{
		Object:  sfc.Object,
		Sources: sfc.Sources,
	}
}

type StreamFuncData struct {
	interpreter.Object
	MetricName        string
	Filter            filter.Filter
	Rollup            string
	Extrapolation     string
	MaxExtrapolations int
	TimeShift         time.Duration
}

func NewStreamFuncData(
	Object interpreter.Object,
	MetricName string,
	Filter filter.Filter,
	Rollup string,
	Extrapolation string,
	MaxExtrapolations int,
	TimeShift time.Duration,
) *StreamFuncData {
	return &StreamFuncData{
		Object:            Object,
		MetricName:        MetricName,
		Filter:            Filter,
		Rollup:            Rollup,
		Extrapolation:     Extrapolation,
		MaxExtrapolations: MaxExtrapolations,
		TimeShift:         TimeShift,
	}
}

func (sfd *StreamFuncData) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncData(sfd)
}

func (sfd *StreamFuncData) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncData{
		Object:            sfd.Object,
		MetricName:        sfd.MetricName,
		Filter:            sfd.Filter,
		Rollup:            sfd.Rollup,
		Extrapolation:     sfd.Extrapolation,
		MaxExtrapolations: sfd.MaxExtrapolations,
		TimeShift:         sfd.TimeShift,
	}
}

type StreamFuncDetect struct {
	interpreter.Object
	On               Stream
	Off              Stream
	Mode             string
	Annotations      interpreter.Object
	EventAnnotations interpreter.Object
	AutoResolveAfter *time.Duration
}

func NewStreamFuncDetect(
	Object interpreter.Object,
	On Stream,
	Off Stream,
	Mode string,
	Annotations interpreter.Object,
	EventAnnotations interpreter.Object,
	AutoResolveAfter *time.Duration,
) *StreamFuncDetect {
	return &StreamFuncDetect{
		Object:           Object,
		On:               On,
		Off:              Off,
		Mode:             Mode,
		Annotations:      Annotations,
		EventAnnotations: EventAnnotations,
		AutoResolveAfter: AutoResolveAfter,
	}
}

func (sfd *StreamFuncDetect) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncDetect(sfd)
}

func (sfd *StreamFuncDetect) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncDetect{
		Object:           sfd.Object,
		On:               cloneTimeshift(sfd.On, amount),
		Off:              cloneTimeshift(sfd.Off, amount),
		Mode:             sfd.Mode,
		Annotations:      sfd.Annotations,
		EventAnnotations: sfd.EventAnnotations,
		AutoResolveAfter: sfd.AutoResolveAfter,
	}
}

type StreamFuncEvents struct {
	interpreter.Object
}

func NewStreamFuncEvents(
	Object interpreter.Object,
) *StreamFuncEvents {
	return &StreamFuncEvents{
		Object: Object,
	}
}

func (sfe *StreamFuncEvents) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncEvents(sfe)
}

func (sfe *StreamFuncEvents) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncEvents{
		Object: sfe.Object,
	}
}

type StreamFuncMax struct {
	interpreter.Object
	Sources []Stream
	Value   interpreter.Object
}

func NewStreamFuncMax(
	Object interpreter.Object,
	Sources []Stream,
	Value interpreter.Object,
) *StreamFuncMax {
	return &StreamFuncMax{
		Object:  Object,
		Sources: Sources,
		Value:   Value,
	}
}

func (sfm *StreamFuncMax) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncMax(sfm)
}

func (sfm *StreamFuncMax) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncMax{
		Object:  sfm.Object,
		Sources: sfm.Sources,
		Value:   sfm.Value,
	}
}

type StreamFuncMean struct {
	interpreter.Object
	Sources   []Stream
	Constants []interpreter.Object
}

func NewStreamFuncMean(
	Object interpreter.Object,
	Sources []Stream,
	Constants []interpreter.Object,
) *StreamFuncMean {
	return &StreamFuncMean{
		Object:    Object,
		Sources:   Sources,
		Constants: Constants,
	}
}

func (sfm *StreamFuncMean) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncMean(sfm)
}

func (sfm *StreamFuncMean) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncMean{
		Object:    sfm.Object,
		Sources:   sfm.Sources,
		Constants: sfm.Constants,
	}
}

type StreamFuncMedian struct {
	interpreter.Object
	Sources   []Stream
	Constants []interpreter.Object
}

func NewStreamFuncMedian(
	Object interpreter.Object,
	Sources []Stream,
	Constants []interpreter.Object,
) *StreamFuncMedian {
	return &StreamFuncMedian{
		Object:    Object,
		Sources:   Sources,
		Constants: Constants,
	}
}

func (sfm *StreamFuncMedian) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncMedian(sfm)
}

func (sfm *StreamFuncMedian) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncMedian{
		Object:    sfm.Object,
		Sources:   sfm.Sources,
		Constants: sfm.Constants,
	}
}

type StreamFuncMin struct {
	interpreter.Object
	Sources []Stream
	Value   interpreter.Object
}

func NewStreamFuncMin(
	Object interpreter.Object,
	Sources []Stream,
	Value interpreter.Object,
) *StreamFuncMin {
	return &StreamFuncMin{
		Object:  Object,
		Sources: Sources,
		Value:   Value,
	}
}

func (sfm *StreamFuncMin) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncMin(sfm)
}

func (sfm *StreamFuncMin) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncMin{
		Object:  sfm.Object,
		Sources: sfm.Sources,
		Value:   sfm.Value,
	}
}

type StreamFuncSum struct {
	interpreter.Object
	Sources  []Stream
	Constant float64
}

func NewStreamFuncSum(
	Object interpreter.Object,
	Sources []Stream,
	Constant float64,
) *StreamFuncSum {
	return &StreamFuncSum{
		Object:   Object,
		Sources:  Sources,
		Constant: Constant,
	}
}

func (sfs *StreamFuncSum) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncSum(sfs)
}

func (sfs *StreamFuncSum) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncSum{
		Object:   sfs.Object,
		Sources:  sfs.Sources,
		Constant: sfs.Constant,
	}
}

type StreamFuncThreshold struct {
	interpreter.Object
	Value float64
}

func NewStreamFuncThreshold(
	Object interpreter.Object,
	Value float64,
) *StreamFuncThreshold {
	return &StreamFuncThreshold{
		Object: Object,
		Value:  Value,
	}
}

func (sft *StreamFuncThreshold) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncThreshold(sft)
}

func (sft *StreamFuncThreshold) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncThreshold{
		Object: sft.Object,
		Value:  sft.Value,
	}
}

type StreamFuncUnion struct {
	interpreter.Object
	Sources []Stream
}

func NewStreamFuncUnion(
	Object interpreter.Object,
	Sources []Stream,
) *StreamFuncUnion {
	return &StreamFuncUnion{
		Object:  Object,
		Sources: Sources,
	}
}

func (sfu *StreamFuncUnion) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncUnion(sfu)
}

func (sfu *StreamFuncUnion) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncUnion{
		Object:  sfu.Object,
		Sources: sfu.Sources,
	}
}

type StreamFuncWhen struct {
	interpreter.Object
	Predicate Stream
	Lasting   *time.Duration
	AtLeast   float64
}

func NewStreamFuncWhen(
	Object interpreter.Object,
	Predicate Stream,
	Lasting *time.Duration,
	AtLeast float64,
) *StreamFuncWhen {
	return &StreamFuncWhen{
		Object:    Object,
		Predicate: Predicate,
		Lasting:   Lasting,
		AtLeast:   AtLeast,
	}
}

func (sfw *StreamFuncWhen) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamFuncWhen(sfw)
}

func (sfw *StreamFuncWhen) CloneTimeShift(amount time.Duration) Stream {
	return &StreamFuncWhen{
		Object:    sfw.Object,
		Predicate: cloneTimeshift(sfw.Predicate, amount),
		Lasting:   sfw.Lasting,
		AtLeast:   sfw.AtLeast,
	}
}

type StreamMethodAbove struct {
	interpreter.Object
	Source Stream
}

func NewStreamMethodAbove(
	Object interpreter.Object,
	Source Stream,
) *StreamMethodAbove {
	return &StreamMethodAbove{
		Object: Object,
		Source: Source,
	}
}

func (sma *StreamMethodAbove) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodAbove(sma)
}

func (sma *StreamMethodAbove) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodAbove{
		Object: sma.Object,
		Source: cloneTimeshift(sma.Source, amount),
	}
}

type StreamMethodAbs struct {
	interpreter.Object
	Source Stream
}

func NewStreamMethodAbs(
	Object interpreter.Object,
	Source Stream,
) *StreamMethodAbs {
	return &StreamMethodAbs{
		Object: Object,
		Source: Source,
	}
}

func (sma *StreamMethodAbs) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodAbs(sma)
}

func (sma *StreamMethodAbs) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodAbs{
		Object: sma.Object,
		Source: cloneTimeshift(sma.Source, amount),
	}
}

type StreamMethodAggregate struct {
	interpreter.Object
	Source          Stream
	Fn              string
	By              []string
	AllowAllMissing bool
	AllowMissing    []string
}

func NewStreamMethodAggregate(
	Object interpreter.Object,
	Source Stream,
	Fn string,
	By []string,
	AllowAllMissing bool,
	AllowMissing []string,
) *StreamMethodAggregate {
	return &StreamMethodAggregate{
		Object:          Object,
		Source:          Source,
		Fn:              Fn,
		By:              By,
		AllowAllMissing: AllowAllMissing,
		AllowMissing:    AllowMissing,
	}
}

func (sma *StreamMethodAggregate) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodAggregate(sma)
}

func (sma *StreamMethodAggregate) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodAggregate{
		Object:          sma.Object,
		Source:          cloneTimeshift(sma.Source, amount),
		Fn:              sma.Fn,
		By:              sma.By,
		AllowAllMissing: sma.AllowAllMissing,
		AllowMissing:    sma.AllowMissing,
	}
}

type StreamMethodBelow struct {
	interpreter.Object
	Source Stream
}

func NewStreamMethodBelow(
	Object interpreter.Object,
	Source Stream,
) *StreamMethodBelow {
	return &StreamMethodBelow{
		Object: Object,
		Source: Source,
	}
}

func (smb *StreamMethodBelow) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodBelow(smb)
}

func (smb *StreamMethodBelow) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodBelow{
		Object: smb.Object,
		Source: cloneTimeshift(smb.Source, amount),
	}
}

type StreamMethodFill struct {
	interpreter.Object
	Source   Stream
	Value    interpreter.Object
	Duration int
	MaxCount int
}

func NewStreamMethodFill(
	Object interpreter.Object,
	Source Stream,
	Value interpreter.Object,
	Duration int,
	MaxCount int,
) *StreamMethodFill {
	return &StreamMethodFill{
		Object:   Object,
		Source:   Source,
		Value:    Value,
		Duration: Duration,
		MaxCount: MaxCount,
	}
}

func (smf *StreamMethodFill) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodFill(smf)
}

func (smf *StreamMethodFill) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodFill{
		Object:   smf.Object,
		Source:   cloneTimeshift(smf.Source, amount),
		Value:    smf.Value,
		Duration: smf.Duration,
		MaxCount: smf.MaxCount,
	}
}

type StreamMethodGeneric struct {
	interpreter.Object
	Source Stream
	Call   string
}

func NewStreamMethodGeneric(
	Object interpreter.Object,
	Source Stream,
	Call string,
) *StreamMethodGeneric {
	return &StreamMethodGeneric{
		Object: Object,
		Source: Source,
		Call:   Call,
	}
}

func (smg *StreamMethodGeneric) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodGeneric(smg)
}

func (smg *StreamMethodGeneric) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodGeneric{
		Object: smg.Object,
		Source: cloneTimeshift(smg.Source, amount),
		Call:   smg.Call,
	}
}

type StreamMethodPercentile struct {
	interpreter.Object
	Source Stream
}

func NewStreamMethodPercentile(
	Object interpreter.Object,
	Source Stream,
) *StreamMethodPercentile {
	return &StreamMethodPercentile{
		Object: Object,
		Source: Source,
	}
}

func (smp *StreamMethodPercentile) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodPercentile(smp)
}

func (smp *StreamMethodPercentile) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodPercentile{
		Object: smp.Object,
		Source: cloneTimeshift(smp.Source, amount),
	}
}

type StreamMethodPublish struct {
	interpreter.Object
	Source Stream
	Label  string
	Enable bool
}

func NewStreamMethodPublish(
	Object interpreter.Object,
	Source Stream,
	Label string,
	Enable bool,
) *StreamMethodPublish {
	return &StreamMethodPublish{
		Object: Object,
		Source: Source,
		Label:  Label,
		Enable: Enable,
	}
}

func (smp *StreamMethodPublish) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodPublish(smp)
}

func (smp *StreamMethodPublish) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodPublish{
		Object: smp.Object,
		Source: cloneTimeshift(smp.Source, amount),
		Label:  smp.Label,
		Enable: smp.Enable,
	}
}

type StreamMethodScale struct {
	interpreter.Object
	Source   Stream
	Multiple float64
}

func NewStreamMethodScale(
	Object interpreter.Object,
	Source Stream,
	Multiple float64,
) *StreamMethodScale {
	return &StreamMethodScale{
		Object:   Object,
		Source:   Source,
		Multiple: Multiple,
	}
}

func (sms *StreamMethodScale) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodScale(sms)
}

func (sms *StreamMethodScale) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodScale{
		Object:   sms.Object,
		Source:   cloneTimeshift(sms.Source, amount),
		Multiple: sms.Multiple,
	}
}

type StreamMethodTimeShift struct {
	interpreter.Object
	Source Stream
	Offset time.Duration
}

func NewStreamMethodTimeShift(
	Object interpreter.Object,
	Source Stream,
	Offset time.Duration,
) *StreamMethodTimeShift {
	return &StreamMethodTimeShift{
		Object: Object,
		Source: Source,
		Offset: Offset,
	}
}

func (smts *StreamMethodTimeShift) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodTimeShift(smts)
}

func (smts *StreamMethodTimeShift) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodTimeShift{
		Object: smts.Object,
		Source: cloneTimeshift(smts.Source, amount),
		Offset: smts.Offset,
	}
}

type StreamMethodTop struct {
	interpreter.Object
	Source Stream
}

func NewStreamMethodTop(
	Object interpreter.Object,
	Source Stream,
) *StreamMethodTop {
	return &StreamMethodTop{
		Object: Object,
		Source: Source,
	}
}

func (smt *StreamMethodTop) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodTop(smt)
}

func (smt *StreamMethodTop) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodTop{
		Object: smt.Object,
		Source: cloneTimeshift(smt.Source, amount),
	}
}

type StreamMethodTransform struct {
	interpreter.Object
	Source Stream
	Fn     string
	Over   time.Duration
}

func NewStreamMethodTransform(
	Object interpreter.Object,
	Source Stream,
	Fn string,
	Over time.Duration,
) *StreamMethodTransform {
	return &StreamMethodTransform{
		Object: Object,
		Source: Source,
		Fn:     Fn,
		Over:   Over,
	}
}

func (smt *StreamMethodTransform) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodTransform(smt)
}

func (smt *StreamMethodTransform) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodTransform{
		Object: smt.Object,
		Source: cloneTimeshift(smt.Source, amount),
		Fn:     smt.Fn,
		Over:   smt.Over,
	}
}

type StreamMethodTransformCycle struct {
	interpreter.Object
	Source        Stream
	Fn            string
	Cycle         string
	CycleStart    *string
	Timezone      *string
	PartialValues bool
	ShiftCycles   int
}

func NewStreamMethodTransformCycle(
	Object interpreter.Object,
	Source Stream,
	Fn string,
	Cycle string,
	CycleStart *string,
	Timezone *string,
	PartialValues bool,
	ShiftCycles int,
) *StreamMethodTransformCycle {
	return &StreamMethodTransformCycle{
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

func (smtc *StreamMethodTransformCycle) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamMethodTransformCycle(smtc)
}

func (smtc *StreamMethodTransformCycle) CloneTimeShift(amount time.Duration) Stream {
	return &StreamMethodTransformCycle{
		Object:        smtc.Object,
		Source:        cloneTimeshift(smtc.Source, amount),
		Fn:            smtc.Fn,
		Cycle:         smtc.Cycle,
		CycleStart:    smtc.CycleStart,
		Timezone:      smtc.Timezone,
		PartialValues: smtc.PartialValues,
		ShiftCycles:   smtc.ShiftCycles,
	}
}

type StreamBinaryOpDouble struct {
	interpreter.Object
	*ObjectStreamTernary
	Stream  Stream
	Op      string
	Other   float64
	Reverse bool
}

func NewStreamBinaryOpDouble(
	Object interpreter.Object,
	ObjectStreamTernary *ObjectStreamTernary,
	Stream Stream,
	Op string,
	Other float64,
	Reverse bool,
) *StreamBinaryOpDouble {
	return &StreamBinaryOpDouble{
		Object:              Object,
		ObjectStreamTernary: ObjectStreamTernary,
		Stream:              Stream,
		Op:                  Op,
		Other:               Other,
		Reverse:             Reverse,
	}
}

func (sbod *StreamBinaryOpDouble) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamBinaryOpDouble(sbod)
}

func (sbod *StreamBinaryOpDouble) CloneTimeShift(amount time.Duration) Stream {
	return &StreamBinaryOpDouble{
		Object:              sbod.Object,
		ObjectStreamTernary: sbod.ObjectStreamTernary,
		Stream:              cloneTimeshift(sbod.Stream, amount),
		Op:                  sbod.Op,
		Other:               sbod.Other,
		Reverse:             sbod.Reverse,
	}
}

type StreamBinaryOpInt struct {
	interpreter.Object
	*ObjectStreamTernary
	Stream  Stream
	Op      string
	Other   int
	Reverse bool
}

func NewStreamBinaryOpInt(
	Object interpreter.Object,
	ObjectStreamTernary *ObjectStreamTernary,
	Stream Stream,
	Op string,
	Other int,
	Reverse bool,
) *StreamBinaryOpInt {
	return &StreamBinaryOpInt{
		Object:              Object,
		ObjectStreamTernary: ObjectStreamTernary,
		Stream:              Stream,
		Op:                  Op,
		Other:               Other,
		Reverse:             Reverse,
	}
}

func (sboi *StreamBinaryOpInt) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamBinaryOpInt(sboi)
}

func (sboi *StreamBinaryOpInt) CloneTimeShift(amount time.Duration) Stream {
	return &StreamBinaryOpInt{
		Object:              sboi.Object,
		ObjectStreamTernary: sboi.ObjectStreamTernary,
		Stream:              cloneTimeshift(sboi.Stream, amount),
		Op:                  sboi.Op,
		Other:               sboi.Other,
		Reverse:             sboi.Reverse,
	}
}

type StreamBinaryOpStream struct {
	interpreter.Object
	*ObjectStreamTernary
	Left  Stream
	Op    string
	Right Stream
}

func NewStreamBinaryOpStream(
	Object interpreter.Object,
	ObjectStreamTernary *ObjectStreamTernary,
	Left Stream,
	Op string,
	Right Stream,
) *StreamBinaryOpStream {
	return &StreamBinaryOpStream{
		Object:              Object,
		ObjectStreamTernary: ObjectStreamTernary,
		Left:                Left,
		Op:                  Op,
		Right:               Right,
	}
}

func (sbos *StreamBinaryOpStream) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamBinaryOpStream(sbos)
}

func (sbos *StreamBinaryOpStream) CloneTimeShift(amount time.Duration) Stream {
	return &StreamBinaryOpStream{
		Object:              sbos.Object,
		ObjectStreamTernary: sbos.ObjectStreamTernary,
		Left:                cloneTimeshift(sbos.Left, amount),
		Op:                  sbos.Op,
		Right:               cloneTimeshift(sbos.Right, amount),
	}
}

type StreamIsNone struct {
	interpreter.Object
	*ObjectStreamTernary
	Source Stream
	Invert bool
}

func NewStreamIsNone(
	Object interpreter.Object,
	ObjectStreamTernary *ObjectStreamTernary,
	Source Stream,
	Invert bool,
) *StreamIsNone {
	return &StreamIsNone{
		Object:              Object,
		ObjectStreamTernary: ObjectStreamTernary,
		Source:              Source,
		Invert:              Invert,
	}
}

func (sin *StreamIsNone) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamIsNone(sin)
}

func (sin *StreamIsNone) CloneTimeShift(amount time.Duration) Stream {
	return &StreamIsNone{
		Object:              sin.Object,
		ObjectStreamTernary: sin.ObjectStreamTernary,
		Source:              cloneTimeshift(sin.Source, amount),
		Invert:              sin.Invert,
	}
}

type StreamTernary struct {
	interpreter.Object
	Left      Stream
	Condition Stream
	Right     Stream
}

func NewStreamTernary(
	Object interpreter.Object,
	Left Stream,
	Condition Stream,
	Right Stream,
) *StreamTernary {
	return &StreamTernary{
		Object:    Object,
		Left:      Left,
		Condition: Condition,
		Right:     Right,
	}
}

func (st *StreamTernary) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamTernary(st)
}

func (st *StreamTernary) CloneTimeShift(amount time.Duration) Stream {
	return &StreamTernary{
		Object:    st.Object,
		Left:      cloneTimeshift(st.Left, amount),
		Condition: cloneTimeshift(st.Condition, amount),
		Right:     cloneTimeshift(st.Right, amount),
	}
}

type StreamUnaryOpMinus struct {
	interpreter.Object
	Stream Stream
}

func NewStreamUnaryOpMinus(
	Object interpreter.Object,
	Stream Stream,
) *StreamUnaryOpMinus {
	return &StreamUnaryOpMinus{
		Object: Object,
		Stream: Stream,
	}
}

func (suom *StreamUnaryOpMinus) Accept(vs VisitorStream) (any, error) {
	return vs.VisitStreamUnaryOpMinus(suom)
}

func (suom *StreamUnaryOpMinus) CloneTimeShift(amount time.Duration) Stream {
	return &StreamUnaryOpMinus{
		Object: suom.Object,
		Stream: cloneTimeshift(suom.Stream, amount),
	}
}
