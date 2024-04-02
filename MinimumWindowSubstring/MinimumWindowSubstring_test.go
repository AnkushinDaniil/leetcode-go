package minimumWindowSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		s      string
		t      string
		output string
	}{
		{
			name:   "Example 1",
			s:      "ADOBECODEBANC",
			t:      "ABC",
			output: "BANC",
		},
		{
			name:   "Example 2",
			s:      "a",
			t:      "a",
			output: "a",
		},
		{
			name:   "Example 3",
			s:      "a",
			t:      "aa",
			output: "",
		},
		{
			name:   "Example 4",
			s:      "cabwefgewcwaefgcf",
			t:      "cae",
			output: "cwae",
		},
		{
			name:   "Example 5",
			s:      "a",
			t:      "b",
			output: "",
		},
		{
			name:   "Example 6",
			s:      "aa",
			t:      "aa",
			output: "aa",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, minWindow(testCase.s, testCase.t), testCase.output)
		})
	}
}
