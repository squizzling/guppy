package main

import (
	"fmt"
	"os"

	iflow "github.com/squizzling/guppy/pkg/flow"
	"github.com/squizzling/guppy/pkg/parser/flow"
	"github.com/squizzling/guppy/pkg/parser/parser"
	"github.com/squizzling/guppy/pkg/parser/tokenizer"
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

	trace := len(os.Args) == 3 && os.Args[2] == "-v"

	errProgram := iflow.NewInterpreter(trace).Execute(program)
	if errProgram != nil {
		fmt.Printf("%v\n", errProgram)
	}
}
