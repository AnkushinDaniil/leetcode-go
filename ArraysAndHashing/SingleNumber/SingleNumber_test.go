package singleNumber

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
		nums:   []int{2, 2, 1},
		output: 1,
	},
	{
		nums:   []int{4, 1, 2, 1, 2},
		output: 4,
	},
	{
		nums:   []int{1},
		output: 1,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := singleNumber(testCase.nums)
	assert.Equal(t, testCase.output, res)
}
