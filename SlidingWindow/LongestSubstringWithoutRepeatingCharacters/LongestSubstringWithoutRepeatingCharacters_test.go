package longestSubstringWithoutRepeatingCharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		s      string
		output int
	}{
		{
			name:   "Example 1",
			s:      "abcabcbb",
			output: 3,
		},
		{
			name:   "Example 2",
			s:      "bbbbb",
			output: 1,
		},
		{
			name:   "Example 3",
			s:      "pwwkew",
			output: 3,
		},
		{
			name:   "Example 4",
			s:      "",
			output: 0,
		},
		{
			name:   "Example 5",
			s:      " ",
			output: 1,
		},
		{
			name:   "Example 6",
			s:      "au",
			output: 2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, lengthOfLongestSubstring(testCase.s), testCase.output)
		})
	}
}
