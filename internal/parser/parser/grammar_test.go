package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"guppy/internal/parser/tokenizer"
)

func TestParseIdList(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		input           string
		expectedIds     []string
		expectedError   string
		remainingTokens int
	}{
		{input: "a", expectedIds: []string{"a"}, remainingTokens: 1},
		{input: "a,", expectedIds: []string{"a"}, remainingTokens: 1},
		{input: "a,,", expectedIds: []string{"a"}, remainingTokens: 2},
		{input: "a,b", expectedIds: []string{"a", "b"}, remainingTokens: 1},
		{input: "a,b,", expectedIds: []string{"a", "b"}, remainingTokens: 1},
		{input: "a,b,,", expectedIds: []string{"a", "b"}, remainingTokens: 2},
		{input: "a,b,c", expectedIds: []string{"a", "b", "c"}, remainingTokens: 1},
		{input: "a,b,c,", expectedIds: []string{"a", "b", "c"}, remainingTokens: 1},
		{input: "a,b,c,,", expectedIds: []string{"a", "b", "c"}, remainingTokens: 2},
		{input: ",a", expectedError: "expecting ID, found COMMA", remainingTokens: 3},
	} {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			p := NewParser(tokenizer.NewTokenizer(tc.input))
			ids, err := p.parseIdList()
			if tc.expectedError != "" {
				require.NotNil(t, err)
				assert.Equal(t, tc.expectedError, err.Message)
			} else {
				require.Nil(t, err)
			}
			assert.Equal(t, tc.remainingTokens, p.tokens.RemainingTokens())
			assert.Equal(t, tc.expectedIds, ids)
		})
	}
}
