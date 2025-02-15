package parser

import (
	"slices"

	"guppy/internal/parser/tokenizer"
)

type Parser struct {
	tokens *tokenizer.Tokenizer
	Next   tokenizer.Token
}

func NewParser(tokens *tokenizer.Tokenizer) *Parser {
	return &Parser{
		tokens: tokens,
		Next:   tokens.Peek(0),
	}
}

func (p *Parser) RemainingTokens() int {
	return p.tokens.RemainingTokens()
}

func (p *Parser) PeekMatch(n int, tokenTypes ...tokenizer.TokenType) bool {
	return slices.Contains(tokenTypes, p.tokens.Peek(n).Type)
}

func (p *Parser) Match(tokenType tokenizer.TokenType) bool {
	if p.Next.Type == tokenType {
		p.tokens.Advance()
		p.Next = p.tokens.Peek(0)
		return true
	}
	return false
}

func (p *Parser) Capture(tts ...tokenizer.TokenType) (tokenizer.Token, bool) {
	for _, tt := range tts {
		if p.Next.Type == tt {
			p.tokens.Advance()
			nextToken := p.Next
			p.Next = p.tokens.Peek(0)
			return nextToken, true
		}
	}
	return tokenizer.Token{}, false
}
