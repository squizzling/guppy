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
		if token.Ok() {
			fmt.Printf("%s %s\n", token.Value().Type, escape(token.Value().Lexeme))
			if token.Value().Type == tokenizer.TokenTypeEOF {
				break
			}
		} else {
			fmt.Printf("Err: %s\n", token.Err())
			break
		}
	}
}
