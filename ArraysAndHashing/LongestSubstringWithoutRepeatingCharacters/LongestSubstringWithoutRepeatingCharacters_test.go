package longestSubstringWithoutRepeatingCharacters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	s      string
	output int
}

var tests = []test{
	{
		s:      "pwwkew",
		output: 3,
	},
	{
		s:      "aabaab!bb",
		output: 3,
	},
	{
		s:      "dvdf",
		output: 3,
	},
	{
		s:      "",
		output: 0,
	},
	{
		s:      "abba",
		output: 2,
	},
	{
		s:      "abcabcbb",
		output: 3,
	},
	{
		s:      "bbbbb",
		output: 1,
	},
	{
		s:      " ",
		output: 1,
	},
	{
		s:      "aa",
		output: 1,
	},
	{
		s:      "ab",
		output: 2,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	t.Run(fmt.Sprintf("%#v", testCase), func(t *testing.T) {
		res := lengthOfLongestSubstring(testCase.s)
		assert.Equal(t, testCase.output, res)
	})
}
