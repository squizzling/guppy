package renderer

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"guppy/pkg/flow/stream"
	"guppy/pkg/interpreter/primitive"
)

type GraphWriter struct {
	NextID     int
	DataBlocks int
	Writer     io.Writer

	StreamNodes map[string]string
}

func escape(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\l")
	return s
}

func (g *GraphWriter) GetNode(stream stream.Stream) (string, bool) {
	key := fmt.Sprintf("%p", stream)
	nodeId, ok := g.StreamNodes[key]
	return nodeId, ok
}

type opt func(n *node)

type node struct {
	shape     string
	color     string
	fillColor string
	fontColor string
}

func Shape(s string) opt {
	return func(n *node) {
		n.shape = s
	}
}

func Color(s string) opt {
	return func(n *node) {
		n.color = s
	}
}

func FillColor(s string) opt {
	return func(n *node) {
		n.fillColor = s
	}
}

func FontColor(s string) opt {
	return func(n *node) {
		n.fontColor = s
	}
}

func (g *GraphWriter) DefineNode(stream stream.Stream, label string, nodeOpt ...opt) string {

	key := fmt.Sprintf("%p", stream)
	if stream != nil {
		if nodeId, ok := g.StreamNodes[key]; ok {
			return nodeId
		}
	}
	nodeId := "_" + strconv.Itoa(g.NextID)
	g.NextID++
	n := &node{
		shape: "box",
	}

	for _, o := range nodeOpt {
		o(n)
	}

	var sb strings.Builder
	sb.WriteString("  " + nodeId + "[label=\"" + escape(label) + "\"")
	if n.shape != "" {
		sb.WriteString(", shape=\"" + n.shape + "\"")
	}
	if n.color != "" {
		sb.WriteString(", color=\"" + n.color + "\"")
	}
	if n.fillColor != "" {
		sb.WriteString(", fillcolor=\"" + n.fillColor + "\"")
	}
	if n.fontColor != "" {
		sb.WriteString(", fontcolor=\"" + n.fontColor + "\"")
	}
	sb.WriteString("]\n")

	//_, _ = g.Writer.Write([]byte(fmt.Sprintf("  %s [label=\"%s\", shape=\"box\"]\n", nodeId, escape(label))))
	_, _ = g.Writer.Write([]byte(sb.String()))
	if stream != nil {
		g.StreamNodes[key] = nodeId
	}
	return nodeId
}

func (g *GraphWriter) DefineEdge(to string, from string, label string) {
	_, _ = g.Writer.Write([]byte(fmt.Sprintf("  %s -> %s [label=\"%s\"]\n", from, to, escape(label))))
}

func (g *GraphWriter) VisitStreamFuncAbs(sfa *stream.StreamFuncAbs) (any, error) {
	if nodeId, ok := g.GetNode(sfa); ok {
		return nodeId, nil
	}

	var sb strings.Builder

	sb.WriteString("abs block\n")
	nodeId := g.DefineNode(sfa, sb.String())
	for _, source := range sfa.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamFuncAlerts(sfa *stream.StreamFuncAlerts) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamFuncCombine(sfc *stream.StreamFuncCombine) (any, error) {
	if nodeId, ok := g.GetNode(sfc); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("combine block\n")
	if sfc.Mode != "" {
		sb.WriteString("Mode: " + sfc.Mode)
	}

	nodeId := g.DefineNode(sfc, sb.String())

	sourceNodeId, err := sfc.Source.Accept(g)
	if err != nil {
		return "", err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamFuncConstDouble(sfcd *stream.StreamFuncConstDouble) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamFuncConstInt(sfci *stream.StreamFuncConstInt) (any, error) {
	if nodeId, ok := g.GetNode(sfci); ok {
		return nodeId, nil
	}

	var sb strings.Builder

	sb.WriteString("const block\n")
	sb.WriteString("Value: " + strconv.Itoa(sfci.Value))
	return g.DefineNode(sfci, sb.String()), nil
}

func (g *GraphWriter) VisitStreamFuncCount(sfc *stream.StreamFuncCount) (any, error) {
	if nodeId, ok := g.GetNode(sfc); ok {
		return nodeId, nil
	}

	var sb strings.Builder

	sb.WriteString("count block\n")
	nodeId := g.DefineNode(sfc, sb.String())
	for _, source := range sfc.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamFuncData(sfd *stream.StreamFuncData) (any, error) {
	if nodeId, ok := g.GetNode(sfd); ok {
		return nodeId, nil
	}

	g.DataBlocks++

	var sb strings.Builder
	sb.WriteString("data block\n")
	sb.WriteString("Metric: " + sfd.MetricName + "\n")
	if sfd.Filter != nil {
		sb.WriteString("Filter: " + sfd.Filter.RenderFilter() + "\n")
	}
	if sfd.Rollup != "" {
		sb.WriteString("Rollup: " + sfd.Rollup + "\n")
	}
	if sfd.Extrapolation != "null" {
		sb.WriteString("Extrapolation: " + sfd.Extrapolation + "\n")
		if sfd.MaxExtrapolations != -1 {
			sb.WriteString("MaxExtrapolation: " + strconv.Itoa(sfd.MaxExtrapolations) + "\n")
		}
	}

	return g.DefineNode(sfd, sb.String(), Color("red")), nil
}

func (g *GraphWriter) VisitStreamFuncDetect(sfd *stream.StreamFuncDetect) (any, error) {
	if nodeId, ok := g.GetNode(sfd); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("detect block\n")
	if sfd.Mode != "paired" {
		sb.WriteString("Mode: " + sfd.Mode + "\n")
	}
	if sfd.Annotations != nil {
		if _, isNone := sfd.Annotations.(*primitive.ObjectNone); !isNone {
			sb.WriteString(fmt.Sprintf("Annotations: %T\n", sfd.Annotations))
		}
	}
	if sfd.EventAnnotations != nil {
		if _, isNone := sfd.EventAnnotations.(*primitive.ObjectNone); !isNone {
			sb.WriteString(fmt.Sprintf("EventAnnotations: %T\n", sfd.EventAnnotations))
		}
	}

	if sfd.AutoResolveAfter != nil {
		// TODO: Render using SFX methods
		sb.WriteString("AutoResolveAfter: " + sfd.AutoResolveAfter.String() + "\n")
	}

	nodeId := g.DefineNode(sfd, sb.String())

	nodeOnId, err := sfd.On.Accept(g)
	if err != nil {
		return "", err
	}
	g.DefineEdge(nodeId, nodeOnId.(string), "On")
	if sfd.Off != nil {
		nodeOffId, err := sfd.Off.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, nodeOffId.(string), "Off")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamFuncEvents(sfe *stream.StreamFuncEvents) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamFuncMax(sfm *stream.StreamFuncMax) (any, error) {
	if nodeId, ok := g.GetNode(sfm); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("max block\n")
	if sfm.Value != nil {
		sb.WriteString("Value: ")
		switch sfm.Value.(type) {
		default:
			panic(fmt.Sprintf("Unknown type: %T", sfm.Value))
		}
		sb.WriteString("\n")
	}

	nodeId := g.DefineNode(sfm, sb.String())

	for _, source := range sfm.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamFuncMean(sfm *stream.StreamFuncMean) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamFuncMedian(sfm *stream.StreamFuncMedian) (any, error) {
	if nodeId, ok := g.GetNode(sfm); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("median block\n")
	if len(sfm.Constants) > 0 {
		sb.WriteString("Constants: [\n")
		for _, constant := range sfm.Constants {
			switch constant.(type) {
			default:
				panic(fmt.Sprintf("Unknown type: %T", constant))
			}
		}
		sb.WriteString("]\n")
	}

	nodeId := g.DefineNode(sfm, sb.String())

	for _, source := range sfm.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamFuncMin(sfm *stream.StreamFuncMin) (any, error) {
	if nodeId, ok := g.GetNode(sfm); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("min block\n")
	if sfm.Value != nil {
		sb.WriteString("Value: ")
		switch sfm.Value.(type) {
		default:
			panic(fmt.Sprintf("Unknown type: %T", sfm.Value))
		}
		sb.WriteString("\n")
	}

	nodeId := g.DefineNode(sfm, sb.String())

	for _, source := range sfm.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamFuncSum(sfs *stream.StreamFuncSum) (any, error) {
	if nodeId, ok := g.GetNode(sfs); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("sum block\n")
	if sfs.Constant != 0 {
		sb.WriteString("Constant: " + strconv.FormatFloat(sfs.Constant, 'f', 6, 64) + "\n ")
	}

	nodeId := g.DefineNode(sfs, sb.String())
	for _, source := range sfs.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamFuncThreshold(sft *stream.StreamFuncThreshold) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamFuncUnion(sfu *stream.StreamFuncUnion) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamFuncWhen(sfw *stream.StreamFuncWhen) (any, error) {
	if nodeId, ok := g.GetNode(sfw); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("when block\n")

	if sfw.Lasting != nil {
		sb.WriteString("Lasting: " + sfw.Lasting.String() + "\n")
	}
	if sfw.AtLeast != 1.0 {
		sb.WriteString("AtLeast: " + strconv.FormatFloat(sfw.AtLeast, 'g', 6, 64) + "\n")
	}

	nodeId := g.DefineNode(sfw, sb.String())
	predicateNodeId, err := sfw.Predicate.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, predicateNodeId.(string), "Predicate")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMethodAbove(sma *stream.StreamMethodAbove) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMethodAbs(sma *stream.StreamMethodAbs) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMethodAggregate(sma *stream.StreamMethodAggregate) (any, error) {
	if nodeId, ok := g.GetNode(sma); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("aggregate block\n")

	sb.WriteString("Fn: " + sma.Fn + "\n")
	if len(sma.By) > 0 {
		sb.WriteString(fmt.Sprintf("By: [%s]\n", sma.By))
	}
	if sma.AllowAllMissing {
		sb.WriteString("AllowMissing: true")
	} else if len(sma.AllowMissing) > 0 {
		sb.WriteString(fmt.Sprintf("AllowMissing: [%s]\n", sma.AllowMissing))
	}

	nodeId := g.DefineNode(sma, sb.String())
	sourceNodeId, err := sma.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMethodBelow(smb *stream.StreamMethodBelow) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMethodFill(smf *stream.StreamMethodFill) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMethodGeneric(smg *stream.StreamMethodGeneric) (any, error) {
	if nodeId, ok := g.GetNode(smg); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("generic block\n")
	sb.WriteString("Call: " + smg.Call + "\n")

	nodeId := g.DefineNode(smg, sb.String())

	if smg.Source != nil {
		sourceNodeId, err := smg.Source.Accept(g)
		if err != nil {
			return nil, err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMethodPercentile(smp *stream.StreamMethodPercentile) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMethodPublish(smp *stream.StreamMethodPublish) (any, error) {
	if nodeId, ok := g.GetNode(smp); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("publish block\n")
	sb.WriteString(fmt.Sprintf("Enable: %t\n", smp.Enable))
	sb.WriteString("Label: " + smp.Label + "\n")

	nodeId := g.DefineNode(smp, sb.String(), Color("green"))

	sourceNodeId, err := smp.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMethodScale(sms *stream.StreamMethodScale) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMethodTimeShift(smts *stream.StreamMethodTimeShift) (any, error) {
	if nodeId, ok := g.GetNode(smts); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("timeshift block\n")
	sb.WriteString("Offset: " + smts.Offset.String() + "\n")
	nodeId := g.DefineNode(smts, sb.String())

	sourceNodeId, err := smts.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMethodTop(smt *stream.StreamMethodTop) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMethodTransform(smt *stream.StreamMethodTransform) (any, error) {
	if nodeId, ok := g.GetNode(smt); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("transform block\n")

	sb.WriteString("Fn: " + smt.Fn + "\n")
	sb.WriteString("Over: " + smt.Over.String() + "\n")

	nodeId := g.DefineNode(smt, sb.String())
	sourceNodeId, err := smt.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMethodTransformCycle(smtc *stream.StreamMethodTransformCycle) (any, error) {
	//TODO implement me
	panic(smtc.Fn)
	panic("implement me")
}

func (g *GraphWriter) VisitStreamBinaryOpDouble(sbod *stream.StreamBinaryOpDouble) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamBinaryOpInt(sboi *stream.StreamBinaryOpInt) (any, error) {
	if nodeId, ok := g.GetNode(sboi); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("math op int block\n")
	sb.WriteString("Op: " + sboi.Op + "\n")
	sb.WriteString("Other: " + strconv.Itoa(sboi.Other) + "\n")
	sb.WriteString("Reverse: " + strconv.FormatBool(sboi.Reverse) + "\n")

	nodeId := g.DefineNode(sboi, sb.String())

	if sboi.Stream != nil {
		otherNodeId, err := sboi.Stream.Accept(g)
		if err != nil {
			return nil, err
		}
		g.DefineEdge(nodeId, otherNodeId.(string), "Other")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamBinaryOpStream(sbos *stream.StreamBinaryOpStream) (any, error) {
	if nodeId, ok := g.GetNode(sbos); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("binary op stream block\n")
	sb.WriteString("Op: " + sbos.Op + "\n")

	nodeId := g.DefineNode(sbos, sb.String())

	if sbos.Left != nil {
		leftNodeId, err := sbos.Left.Accept(g)
		if err != nil {
			return nil, err
		}
		g.DefineEdge(nodeId, leftNodeId.(string), "Left")
	}
	if sbos.Right != nil {
		rightNodeId, err := sbos.Right.Accept(g)
		if err != nil {
			return nil, err
		}
		g.DefineEdge(nodeId, rightNodeId.(string), "Right")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamIsNone(sin *stream.StreamIsNone) (any, error) {
	if nodeId, ok := g.GetNode(sin); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("stream is none block\n")
	if sin.Invert {
		sb.WriteString(fmt.Sprintf("Invert: %t\n", sin.Invert))
	}
	nodeId := g.DefineNode(sin, sb.String())

	sourceNodeId, err := sin.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamTernary(st *stream.StreamTernary) (any, error) {
	if nodeId, ok := g.GetNode(st); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("ternary block\n")
	nodeId := g.DefineNode(st, sb.String())

	leftNodeId, err := st.Left.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, leftNodeId.(string), "Left")

	conditionNodeId, err := st.Condition.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, conditionNodeId.(string), "Condition")

	rightNodeId, err := st.Right.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, rightNodeId.(string), "Right")

	return nodeId, nil
}

func (g *GraphWriter) VisitStreamUnaryOpMinus(suom *stream.StreamUnaryOpMinus) (any, error) {
	//TODO implement me
	panic("implement me")
}
