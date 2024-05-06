package combinationSum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		nums   []int
		output [][]int
	}{
		{
			nums:   []int{1, 2, 3},
			output: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums:   []int{0, 1},
			output: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums:   []int{1},
			output: [][]int{{1}},
		},
	}

	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, permute(testCase.nums))
		})
	}
}
