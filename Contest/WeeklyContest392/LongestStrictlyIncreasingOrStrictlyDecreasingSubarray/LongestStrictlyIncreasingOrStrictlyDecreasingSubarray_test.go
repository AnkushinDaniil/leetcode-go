package binarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		nums   []int
		output int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 4, 3, 3, 2},
			output: 2,
		},
		{
			name:   "Example 2",
			nums:   []int{3, 3, 3, 3},
			output: 1,
		},
		{
			name:   "Example 3",
			nums:   []int{3, 2, 1},
			output: 3,
		},
		{
			name:   "Example 4",
			nums:   []int{3, 3, 2, 1},
			output: 3,
		},
		{
			name:   "Example 5",
			nums:   []int{3, 3, 4, 5},
			output: 3,
		},
		{
			name:   "Example 6",
			nums:   []int{4, 3, 3},
			output: 2,
		},
		{
			name:   "Example 7",
			nums:   []int{1, 9, 7, 1},
			output: 3,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, longestMonotonicSubarray(testCase.nums), testCase.output)
		})
	}
}
