package maximumSubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	name   string
	nums   []int
	output int
}

var tests = []test{
	{
		name:   "[-1,1,2,1]",
		nums:   []int{-1, 1, 2, 1},
		output: 4,
	},
	{
		name:   "[-2,1,-3,4,-1,2,1,-5,4]",
		nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		output: 6,
	},
	{
		name:   "[1]",
		nums:   []int{1},
		output: 1,
	},
	{
		name:   "[5,4,-1,7,8]",
		nums:   []int{5, 4, -1, 7, 8},
		output: 23,
	},
	{
		name:   "[-1]",
		nums:   []int{-1},
		output: -1,
	},
	{
		name:   "[-2,1]",
		nums:   []int{-2, 1},
		output: 1,
	},
	{
		name:   "[-2,-1]",
		nums:   []int{-2, -1},
		output: -1,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.output, maxSubArray(testCase.nums))
		})
	}
}
