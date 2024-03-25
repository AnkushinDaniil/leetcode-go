package binarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		nums   []int
		target int
		output int
	}{
		{
			name:   "Example 1",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			output: 4,
		},
		{
			name:   "Example 2",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			output: -1,
		},
		{
			name:   "Example 3",
			nums:   []int{2, 5},
			target: 2,
			output: 0,
		},
		{
			name:   "Example 4",
			nums:   []int{-1, 0, 5},
			target: 5,
			output: 2,
		},
		{
			name:   "Example 5",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 13,
			output: -1,
		},
		{
			name:   "Example 6",
			nums:   []int{5},
			target: 5,
			output: 0,
		},
		{
			name:   "Example 7",
			nums:   []int{-1, 0, 5},
			target: -1,
			output: 0,
		},
		{
			name:   "Example 7",
			nums:   []int{2, 5},
			target: 0,
			output: -1,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, search(testCase.nums, testCase.target), testCase.output)
		})
	}
}
