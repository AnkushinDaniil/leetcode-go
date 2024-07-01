package numberOf1Bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	n      int
	output int
}

var tests = []test{
	{
		n:      11,
		output: 3,
	},
	{
		n:      128,
		output: 1,
	},
	{
		n:      2147483645,
		output: 30,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := hammingWeight(testCase.n)
	assert.Equal(t, testCase.output, res)
}
