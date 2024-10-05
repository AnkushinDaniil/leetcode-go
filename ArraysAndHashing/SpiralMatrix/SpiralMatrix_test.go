package spiralmatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected []int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d, %d", len(test.matrix), len(test.matrix[0])), func(t *testing.T) {
			assert.Equal(t, test.expected, spiralOrder(test.matrix), "spiralOrder(%v) should return %v", test.matrix, test.expected)
		})
	}
}
