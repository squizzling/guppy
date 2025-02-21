package main

import (
	"fmt"
	"os"

	"guppy/internal/flow/filter"
	"guppy/internal/flow/stream"
	"guppy/internal/interpreter"
	"guppy/internal/parser/flow"
	"guppy/internal/parser/parser"
	"guppy/internal/parser/tokenizer"
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

	i := interpreter.NewInterpreter()
	i.Scope.DeclareSet("filter", &filter.FFIFilter{Object: interpreter.NewObject(nil)})
	i.Scope.DeclareSet("data", &stream.FFIData{Object: interpreter.NewObject(nil)})

	errProgram := i.Execute(program)
	if errProgram != nil {
		fmt.Printf("%v\n", errProgram)
	}
}
