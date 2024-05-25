package pacificAtlanticWaterFlow

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		heights [][]int
		output  [][]int
	}{
		{
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			output: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			heights: [][]int{{1}},
			output:  [][]int{{0, 0}},
		},
		{
			heights: [][]int{{2, 1}, {1, 2}},
			output:  [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}
	for testCaseNum, testCase := range testTable {
		sortFunc := func(a, b []int) int {
			for i := 0; i < len(a); i++ {
				if a[i] == b[i] {
					continue
				}
				return a[i] - b[i]
			}
			return 0
		}
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			slices.SortFunc(testCase.output, sortFunc)
			res := pacificAtlantic(testCase.heights)
			slices.SortFunc(res, sortFunc)
			assert.Equal(t, testCase.output, res)
		})
	}
}
