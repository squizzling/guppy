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
	return slices.Contains(tokenTypes, p.tokens.Peek(n).Type)
}

func (p *Parser) match(tokenType tokenizer.TokenType) bool {
	nextToken := p.tokens.Peek(0)
	if nextToken.Type == tokenType {
		p.tokens.Advance()
		return true
	}
	return false
}

func (p *Parser) capture(tts ...tokenizer.TokenType) (tokenizer.Token, bool) {
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

/*func (p *Parser) dumpTokens(n int) {
	n = min(len(p.tokens), p.next+n)
	for i := p.next; i < n; i++ {
		_, _ = fmt.Fprintf(os.Stderr, "[%d] - %v\n", i, p.tokens[i])
	}
	_, _ = fmt.Fprintf(os.Stderr, "------\n")
}*/
