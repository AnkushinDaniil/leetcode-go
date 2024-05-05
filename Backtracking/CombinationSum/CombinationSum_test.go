package combinationSum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		nums   []int
		target int
		output [][]int
	}{
		{
			nums:   []int{2, 3, 6, 7},
			target: 7,
			output: [][]int{{2, 2, 3}, {7}},
		},
		{
			nums:   []int{2, 3, 5},
			target: 8,
			output: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			nums:   []int{2},
			target: 1,
			output: [][]int{},
		},
	}

	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, combinationSum(testCase.nums, testCase.target))
		})
	}
}
