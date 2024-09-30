package verifyinganaliendictionary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlienSorted(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		words  []string
		order  string
		output bool
	}{
		{
			words:  []string{"hello", "leetcode"},
			order:  "hlabcdefgijkmnopqrstuvwxyz",
			output: true,
		},
		{
			words:  []string{"word", "world", "row"},
			order:  "worldabcefghijkmnpqstuvxyz",
			output: false,
		},
		{
			words:  []string{"apple", "app"},
			order:  "abcdefghijklmnopqrstuvwxyz",
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test %v", tc.words), func(t *testing.T) {
			assert.Equal(t, tc.output, isAlienSorted(tc.words, tc.order))
		})
	}
}
