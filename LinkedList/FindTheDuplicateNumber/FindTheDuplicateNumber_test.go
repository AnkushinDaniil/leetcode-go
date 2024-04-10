package reorderList

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
			nums:   []int{1, 3, 4, 2, 2},
			output: 2,
		},
		{
			name:   "Example 2",
			nums:   []int{3, 1, 3, 4, 2},
			output: 3,
		},
		{
			name:   "Example 3",
			nums:   []int{3, 3, 3, 3, 3},
			output: 3,
		},
		{
			name:   "Example 4",
			nums:   []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			output: 9,
		},
		{
			name:   "Example 5",
			nums:   []int{14, 16, 12, 1, 16, 17, 6, 8, 5, 19, 16, 13, 16, 3, 11, 16, 4, 16, 9, 7},
			output: 16,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, findDuplicate(testCase.nums), testCase.output)
		})
	}
}
