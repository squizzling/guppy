package main

import (
	"fmt"
	"os"

	iflow "guppy/pkg/flow"
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

	errProgram := iflow.NewInterpreter(false).Execute(program)
	if errProgram != nil {
		fmt.Printf("%v\n", errProgram)
	}
}
