package missingNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	nums   []int
	output int
}

var tests = []test{
	{
		nums:   []int{3, 0, 1},
		output: 2,
	},
	{
		nums:   []int{0, 1},
		output: 2,
	},
	{
		nums:   []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
		output: 8,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := missingNumber(testCase.nums)
	assert.Equal(t, testCase.output, res)
}
