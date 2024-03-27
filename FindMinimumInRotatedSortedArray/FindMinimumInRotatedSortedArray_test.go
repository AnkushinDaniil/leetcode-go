package findMinimumInRotatedSortedArray

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
			nums:   []int{3, 4, 5, 1, 2},
			output: 1,
		},
		{
			name:   "Example 2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			output: 0,
		},
		{
			name:   "Example 3",
			nums:   []int{11, 13, 15, 17},
			output: 11,
		},
		{
			name:   "Example 4",
			nums:   []int{3, 1, 2},
			output: 1,
		},
		{
			name:   "Example 5",
			nums:   []int{2, 1},
			output: 1,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, findMin(testCase.nums), testCase.output)
		})
	}
}
