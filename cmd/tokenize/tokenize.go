package main

import (
	"fmt"
	"os"
	"strings"

	"guppy/internal/parser/tokenizer"
)

func escape(s string) string {
	return strings.Replace(s, "\n", "\\n", -1)
}

func main() {
	d, _ := os.ReadFile(os.Args[1])
	t := tokenizer.NewTokenizer(string(d))
	for {
		token := t.Get()
		fmt.Printf("%s %s\n", token.Type, escape(token.Lexeme))
		if token.Type == tokenizer.TokenTypeEOF || token.Type == tokenizer.TokenTypeError {
			break
		}
	}
}
