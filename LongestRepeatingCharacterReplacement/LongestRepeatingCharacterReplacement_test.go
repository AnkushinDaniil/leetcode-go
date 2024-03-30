package longestRepeatingCharacterReplacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		s      string
		k      int
		output int
	}{
		{
			name:   "Example 1",
			s:      "ABAB",
			k:      2,
			output: 4,
		},
		{
			name:   "Example 2",
			s:      "AABABBA",
			k:      1,
			output: 4,
		},
		{
			name:   "Example 3",
			s:      "",
			k:      1,
			output: 0,
		},
		{
			name:   "Example 4",
			s:      "A",
			k:      1,
			output: 1,
		},
		{
			name:   "Example 5",
			s:      "ABBB",
			k:      2,
			output: 4,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, characterReplacement(testCase.s, testCase.k), testCase.output)
		})
	}
}
