package combinationSumII

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
			nums:   []int{10, 1, 2, 7, 6, 1, 5},
			target: 8,
			output: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			nums:   []int{2, 5, 2, 1, 2},
			target: 5,
			output: [][]int{{1, 2, 2}, {5}},
		},
	}

	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, combinationSum2(testCase.nums, testCase.target))
		})
	}
}
