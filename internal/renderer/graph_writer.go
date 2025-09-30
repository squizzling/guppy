package renderer

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"guppy/internal/flow/stream"
	"guppy/internal/interpreter"
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

func (g *GraphWriter) VisitStreamAbove(sa *stream.StreamAbove) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamAbs(sa *stream.StreamAbs) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamAggregate(sa *stream.StreamAggregate) (any, error) {
	if nodeId, ok := g.GetNode(sa); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("aggregate block\n")

	sb.WriteString("Fn: " + sa.Fn + "\n")
	if len(sa.By) > 0 {
		sb.WriteString(fmt.Sprintf("By: [%s]\n", sa.By))
	}
	if sa.AllowAllMissing {
		sb.WriteString("AllowMissing: true")
	} else if len(sa.AllowMissing) > 0 {
		sb.WriteString(fmt.Sprintf("AllowMissing: [%s]\n", sa.AllowMissing))
	}

	nodeId := g.DefineNode(sa, sb.String())
	sourceNodeId, err := sa.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamAlerts(sa *stream.StreamAlerts) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamBelow(sb *stream.StreamBelow) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamConstDouble(scd *stream.StreamConstDouble) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamConstInt(sci *stream.StreamConstInt) (any, error) {
	if nodeId, ok := g.GetNode(sci); ok {
		return nodeId, nil
	}

	var sb strings.Builder

	sb.WriteString("const block\n")
	sb.WriteString("Value: " + strconv.Itoa(sci.Value))
	return g.DefineNode(sci, sb.String()), nil
}

func (g *GraphWriter) VisitStreamData(sd *stream.StreamData) (any, error) {
	if nodeId, ok := g.GetNode(sd); ok {
		return nodeId, nil
	}

	g.DataBlocks++

	var sb strings.Builder
	sb.WriteString("data block\n")
	sb.WriteString("Metric: " + sd.MetricName + "\n")
	if sd.Filter != nil {
		sb.WriteString("Filter: " + sd.Filter.RenderFilter() + "\n")
	}
	if sd.Rollup != "" {
		sb.WriteString("Rollup: " + sd.Rollup + "\n")
	}
	if sd.Extrapolation != "null" {
		sb.WriteString("Extrapolation: " + sd.Extrapolation + "\n")
		if sd.MaxExtrapolations != -1 {
			sb.WriteString("MaxExtrapolation: " + strconv.Itoa(sd.MaxExtrapolations) + "\n")
		}
	}

	return g.DefineNode(sd, sb.String(), Color("red")), nil
}

func (g *GraphWriter) VisitStreamDetect(sd *stream.StreamDetect) (any, error) {
	if nodeId, ok := g.GetNode(sd); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("detect block\n")
	if sd.Mode != "paired" {
		sb.WriteString("Mode: " + sd.Mode + "\n")
	}
	if sd.Annotations != nil {
		if _, isNone := sd.Annotations.(*interpreter.ObjectNone); !isNone {
			sb.WriteString(fmt.Sprintf("Annotations: %T\n", sd.Annotations))
		}
	}
	if sd.EventAnnotations != nil {
		if _, isNone := sd.EventAnnotations.(*interpreter.ObjectNone); !isNone {
			sb.WriteString(fmt.Sprintf("EventAnnotations: %T\n", sd.EventAnnotations))
		}
	}

	if sd.AutoResolveAfter != nil {
		// TODO: Render using SFX methods
		sb.WriteString("AutoResolveAfter: " + sd.AutoResolveAfter.String() + "\n")
	}

	nodeId := g.DefineNode(sd, sb.String())

	nodeOnId, err := sd.On.Accept(g)
	if err != nil {
		return "", err
	}
	g.DefineEdge(nodeId, nodeOnId.(string), "On")
	if sd.Off != nil {
		nodeOffId, err := sd.Off.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, nodeOffId.(string), "Off")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamEvents(se *stream.StreamEvents) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamFill(sf *stream.StreamFill) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamGeneric(sg *stream.StreamGeneric) (any, error) {
	if nodeId, ok := g.GetNode(sg); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("generic block\n")
	sb.WriteString("Call: " + sg.Call + "\n")

	nodeId := g.DefineNode(sg, sb.String())

	if sg.Source != nil {
		sourceNodeId, err := sg.Source.Accept(g)
		if err != nil {
			return nil, err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamIsNone(sin *stream.StreamIsNone) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMax(sm *stream.StreamMax) (any, error) {
	if nodeId, ok := g.GetNode(sm); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("max block\n")
	if sm.Value != nil {
		sb.WriteString("Value: ")
		switch sm.Value.(type) {
		default:
			panic(fmt.Sprintf("Unknown type: %T", sm.Value))
		}
		sb.WriteString("\n")
	}

	nodeId := g.DefineNode(sm, sb.String())

	for _, source := range sm.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMean(sm *stream.StreamMean) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMedian(sm *stream.StreamMedian) (any, error) {
	if nodeId, ok := g.GetNode(sm); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("median block\n")
	if len(sm.Constants) > 0 {
		sb.WriteString("Constants: [\n")
		for _, constant := range sm.Constants {
			switch constant.(type) {
			default:
				panic(fmt.Sprintf("Unknown type: %T", constant))
			}
		}
		sb.WriteString("]\n")
	}

	nodeId := g.DefineNode(sm, sb.String())

	for _, source := range sm.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMin(sm *stream.StreamMin) (any, error) {
	if nodeId, ok := g.GetNode(sm); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("min block\n")
	if sm.Value != nil {
		sb.WriteString("Value: ")
		switch sm.Value.(type) {
		default:
			panic(fmt.Sprintf("Unknown type: %T", sm.Value))
		}
		sb.WriteString("\n")
	}

	nodeId := g.DefineNode(sm, sb.String())

	for _, source := range sm.Sources {
		sourceNodeId, err := source.Accept(g)
		if err != nil {
			return "", err
		}
		g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMathOpDouble(smod *stream.StreamMathOpDouble) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamMathOpInt(smoi *stream.StreamMathOpInt) (any, error) {
	if nodeId, ok := g.GetNode(smoi); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("math op int block\n")
	sb.WriteString("Op: " + smoi.Op + "\n")
	sb.WriteString("Other: " + strconv.Itoa(smoi.Other) + "\n")
	sb.WriteString("Reverse: " + strconv.FormatBool(smoi.Reverse) + "\n")

	nodeId := g.DefineNode(smoi, sb.String())

	if smoi.Stream != nil {
		otherNodeId, err := smoi.Stream.Accept(g)
		if err != nil {
			return nil, err
		}
		g.DefineEdge(nodeId, otherNodeId.(string), "Other")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMathOpStream(smos *stream.StreamMathOpStream) (any, error) {
	if nodeId, ok := g.GetNode(smos); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("math op stream block\n")
	sb.WriteString("Op: " + smos.Op + "\n")

	nodeId := g.DefineNode(smos, sb.String())

	if smos.Left != nil {
		leftNodeId, err := smos.Left.Accept(g)
		if err != nil {
			return nil, err
		}
		g.DefineEdge(nodeId, leftNodeId.(string), "Left")
	}
	if smos.Right != nil {
		rightNodeId, err := smos.Right.Accept(g)
		if err != nil {
			return nil, err
		}
		g.DefineEdge(nodeId, rightNodeId.(string), "Right")
	}
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamMathOpUnaryMinus(smoum *stream.StreamMathOpUnaryMinus) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamPercentile(sp *stream.StreamPercentile) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamPublish(sp *stream.StreamPublish) (any, error) {
	if nodeId, ok := g.GetNode(sp); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("publish block\n")
	sb.WriteString(fmt.Sprintf("Enable: %t\n", sp.Enable))
	sb.WriteString("Label: " + sp.Label + "\n")

	nodeId := g.DefineNode(sp, sb.String(), Color("green"))

	sourceNodeId, err := sp.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamScale(ss *stream.StreamScale) (any, error) {
	//TODO implement me
	panic("implement me")
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

func (g *GraphWriter) VisitStreamThreshold(st *stream.StreamThreshold) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamTimeShift(sts *stream.StreamTimeShift) (any, error) {
	if nodeId, ok := g.GetNode(sts); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("timeshift block\n")
	sb.WriteString("Offset: " + sts.Offset.String() + "\n")
	nodeId := g.DefineNode(sts, sb.String())

	sourceNodeId, err := sts.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamTop(st *stream.StreamTop) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamTransform(st *stream.StreamTransform) (any, error) {
	if nodeId, ok := g.GetNode(st); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("transform block\n")

	sb.WriteString("Fn: " + st.Fn + "\n")
	sb.WriteString("Over: " + st.Over.String() + "\n")

	nodeId := g.DefineNode(st, sb.String())
	sourceNodeId, err := st.Source.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, sourceNodeId.(string), "Source")
	return nodeId, nil
}

func (g *GraphWriter) VisitStreamTransformCycle(stc *stream.StreamTransformCycle) (any, error) {
	//TODO implement me
	panic(stc.Fn)
	panic("implement me")
}

func (g *GraphWriter) VisitStreamUnion(su *stream.StreamUnion) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GraphWriter) VisitStreamWhen(sw *stream.StreamWhen) (any, error) {
	if nodeId, ok := g.GetNode(sw); ok {
		return nodeId, nil
	}

	var sb strings.Builder
	sb.WriteString("when block\n")

	if sw.Lasting != nil {
		sb.WriteString("Lasting: " + sw.Lasting.String() + "\n")
	}
	if sw.AtLeast != 1.0 {
		sb.WriteString("AtLeast: " + strconv.FormatFloat(sw.AtLeast, 'g', 6, 64) + "\n")
	}

	nodeId := g.DefineNode(sw, sb.String())
	predicateNodeId, err := sw.Predicate.Accept(g)
	if err != nil {
		return nil, err
	}
	g.DefineEdge(nodeId, predicateNodeId.(string), "Predicate")
	return nodeId, nil
}
