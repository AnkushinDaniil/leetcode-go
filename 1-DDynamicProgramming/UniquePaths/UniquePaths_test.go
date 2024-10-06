package uniquepaths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m        int
		n        int
		expected int
	}{
		{3, 2, 3},
		{7, 3, 28},
		{3, 3, 6},
		{3, 7, 28},
		{3, 1, 1},
		{1, 3, 1},
		{1, 1, 1},
		{1, 2, 1},
		{2, 1, 1},
		{2, 2, 2},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("(%d %d)", test.m, test.n) , func(t *testing.T) {
			assert.Equal(t, test.expected, uniquePaths(test.m, test.n))
		})
	}
}
