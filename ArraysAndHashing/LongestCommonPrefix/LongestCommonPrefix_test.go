package longestCommonPrefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	strs   []string
	output string
}

var tests = []test{
	{
		strs:   []string{"flower", "flow", "flight"},
		output: "fl",
	},
	{
		strs:   []string{"dog", "racecar", "car"},
		output: "",
	},
	{
		strs:   []string{"ab", "a"},
		output: "a",
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := longestCommonPrefix(testCase.strs)
	assert.Equal(t, testCase.output, res)
}
