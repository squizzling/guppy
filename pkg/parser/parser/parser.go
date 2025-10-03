package parser

import (
	"slices"

	"guppy/pkg/parser/tokenizer"
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

func (p *Parser) MatchErr(tokenType tokenizer.TokenType) *ParseError {
	if p.Next.Type == tokenType {
		p.tokens.Advance()
		p.Next = p.tokens.Peek(0)
		return nil
	}
	return FailMsgf("expecting %s in %s, found %s", tokenType, callerName(0), p.Next.Type)
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

func (p *Parser) CaptureErr(tts ...tokenizer.TokenType) (tokenizer.Token, *ParseError) {
	for _, tt := range tts {
		if p.Next.Type == tt {
			p.tokens.Advance()
			nextToken := p.Next
			p.Next = p.tokens.Peek(0)
			return nextToken, nil
		}
	}
	return p.Next, FailMsgf("expecting %s in %s, found %s", tts, callerName(0), p.Next.Type)
}
