package main

import (
	"fmt"
	"os"

	"guppy/internal/renderer"
	pflow "guppy/pkg/flow"
	"guppy/pkg/parser/flow"
	"guppy/pkg/parser/parser"
	"guppy/pkg/parser/tokenizer"
)

func main() {
	d, _ := os.ReadFile(os.Args[1])
	t := tokenizer.NewTokenizer(string(d))
	p := parser.NewParser(t)
	program, err := flow.ParseProgram(p)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		ss := err.Stack()
		for _, s := range ss {
			_, _ = fmt.Fprintf(os.Stderr, "%s %s\n", s.Location, s.Message)
		}
		os.Exit(1)
	}

	i := pflow.NewInterpreter(false)
	errProgram := i.Execute(program)
	if errProgram != nil {
		fmt.Printf("%v\n", errProgram)
	}

	rawPublished, err2 := i.Globals.Get("_published")
	if err2 != nil {
		fmt.Printf("Failed to get _published: %s", err)
	}

	published := rawPublished.(*pflow.Published)

	fmt.Printf("digraph G {\n")
	fmt.Printf("rankdir=\"BT\"")
	fmt.Printf("graph [bgcolor=\"black\"]\n")
	fmt.Printf("node [fontcolor=\"white\", color=\"white\", shape=\"box\", fontname=\"Helvetica\"]\n")
	fmt.Printf("edge [fontcolor=\"white\", fillcolor=\"white\", color=\"white\"]\n")
	gw := &renderer.GraphWriter{Writer: os.Stdout, StreamNodes: map[string]string{}}
	for _, stream := range published.Streams {
		_, err2 := stream.Accept(gw)
		if err2 != nil {
			panic(err2)
		}
	}
	fmt.Printf("}\n")
	_, _ = fmt.Fprintf(os.Stderr, "Data Nodes: %d\n", gw.DataBlocks)
	_, _ = fmt.Fprintf(os.Stderr, "Nodes: %d\n", gw.NextID)
}
