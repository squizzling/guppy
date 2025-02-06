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

func (p *Parser) peekMatch(n int, tokenTypes ...tokenizer.TokenType) bool {
	tok := p.tokens.Peek(n)
	if !tok.Ok() {
		return false
	}
	return slices.Contains(tokenTypes, tok.Value().Type)
}

func (p *Parser) match(tokenType tokenizer.TokenType) bool {
	if tokenType == tokenizer.TokenTypeEOF && p.isAtEnd() {
		return true
	} else if nextToken := p.tokens.Peek(0); !nextToken.Ok() {
		return false
	} else if nextToken.Value().Type == tokenType {
		_ = p.tokens.Get()
		return true
	} else {
		return false
	}
}

func (p *Parser) capture(tts ...tokenizer.TokenType) (tokenizer.Token, bool) {
	if p.isAtEnd() {
		return tokenizer.Token{}, false
	}

	nextToken := p.tokens.Peek(0)
	if !nextToken.Ok() {
		return tokenizer.Token{}, false
	}
	for _, tt := range tts {
		if nextToken.Value().Type == tt {
			return p.tokens.Get().Value(), true
		}
	}
	return tokenizer.Token{}, false
}

func (p *Parser) isAtEnd() bool {
	if nextToken := p.tokens.Peek(0); !nextToken.Ok() {
		return true
	} else {
		return nextToken.Value().Type == tokenizer.TokenTypeEOF
	}
}

/*func (p *Parser) dumpTokens(n int) {
	n = min(len(p.tokens), p.next+n)
	for i := p.next; i < n; i++ {
		_, _ = fmt.Fprintf(os.Stderr, "[%d] - %v\n", i, p.tokens[i])
	}
	_, _ = fmt.Fprintf(os.Stderr, "------\n")
}*/
