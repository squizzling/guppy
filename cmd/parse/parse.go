package main

import (
	"fmt"
	"os"

	"guppy/internal/parser/ast"
	"guppy/internal/parser/parser"
	"guppy/internal/parser/tokenizer"
)

func main() {
	d, _ := os.ReadFile(os.Args[1])
	t := tokenizer.NewTokenizer(string(d))
	p := parser.NewParser(t)
	program, err := p.ParseProgram()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		ss := err.Stack()
		for _, s := range ss {
			_, _ = fmt.Fprintf(os.Stderr, "%s %s\n", s.Location, s.Message)
		}
		os.Exit(1)
	}
	fmt.Printf("%s", program.Accept(ast.DebugWriter{}).(string))
}
