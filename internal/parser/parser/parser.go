package parser

import (
	"slices"

	"guppy/internal/parser/tokenizer"
)

type Parser struct {
	tokens *tokenizer.Tokenizer
}

func NewParser(tokens *tokenizer.Tokenizer) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) RemainingTokens() int {
	return p.tokens.RemainingTokens()
}

func (p *Parser) Peek(n int) tokenizer.Token {
	return p.tokens.Peek(n)
}

func (p *Parser) PeekMatch(n int, tokenTypes ...tokenizer.TokenType) bool {
	return slices.Contains(tokenTypes, p.tokens.Peek(n).Type)
}

func (p *Parser) Match(tokenType tokenizer.TokenType) bool {
	nextToken := p.tokens.Peek(0)
	if nextToken.Type == tokenType {
		p.tokens.Advance()
		return true
	}
	return false
}

func (p *Parser) Capture(tts ...tokenizer.TokenType) (tokenizer.Token, bool) {
	if p.isAtEnd() {
		return tokenizer.Token{}, false
	}

	nextToken := p.tokens.Peek(0)
	for _, tt := range tts {
		if nextToken.Type == tt {
			p.tokens.Advance()
			return nextToken, true
		}
	}
	return tokenizer.Token{}, false
}

func (p *Parser) isAtEnd() bool {
	nextToken := p.tokens.Peek(0)
	return nextToken.Type == tokenizer.TokenTypeEOF || nextToken.Type == tokenizer.TokenTypeError
}
