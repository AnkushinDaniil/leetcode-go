package reshapethematrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		mat  [][]int
		r    int
		c    int
		want [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
		{[][]int{{1, 2, 3, 4}}, 2, 2, [][]int{{1, 2}, {3, 4}}},
	}
	for _, test := range tests {
		assert.Equal(t, matrixReshape(test.mat, test.r, test.c), test.want)
	}
}
