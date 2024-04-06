package searchInRotatedSortedArray

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
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			output: 4,
		},
		{
			name:   "Example 2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			output: -1,
		},
		{
			name:   "Example 3",
			nums:   []int{1},
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
