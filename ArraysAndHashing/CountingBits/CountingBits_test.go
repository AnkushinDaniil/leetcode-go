package countingBits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	n      int
	output []int
}

var tests = []test{
	{
		n:      2,
		output: []int{0, 1, 1},
	},
	{
		n:      5,
		output: []int{0, 1, 1, 2, 1, 2},
	},
	{
		n:      0,
		output: []int{0},
	},
	{
		n:      16,
		output: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1},
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := countBits(testCase.n)
	assert.Equal(t, testCase.output, res)
}
