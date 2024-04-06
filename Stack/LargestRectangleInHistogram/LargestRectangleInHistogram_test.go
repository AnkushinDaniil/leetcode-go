package largestRectangleInHistogram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name    string
		heights []int
		output  int
	}{
		{
			name:    "Example 1",
			heights: []int{2, 1, 5, 6, 2, 3},
			output:  10,
		},
		{
			name:    "Example 2",
			heights: []int{2, 4},
			output:  4,
		},
		{
			name:    "Example 3",
			heights: []int{1, 1},
			output:  2,
		},
		{
			name:    "Example 4",
			heights: []int{2, 1, 2},
			output:  3,
		},
		{
			name:    "Example 5",
			heights: []int{5, 4, 1, 2},
			output:  8,
		},
		{
			name:    "Example 6",
			heights: []int{4, 2, 0, 3, 2, 5},
			output:  6,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, largestRectangleArea(testCase.heights), testCase.output)
		})
	}
}
