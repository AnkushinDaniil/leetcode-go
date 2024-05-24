package rottingOranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		grid   [][]int
		output int
	}{
		{
			grid:   [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			output: 4,
		},
		{
			grid:   [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			output: -1,
		},
		{
			grid:   [][]int{{0, 2}},
			output: 0,
		},
		{
			grid:   [][]int{{0}},
			output: 0,
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, orangesRotting(testCase.grid))
		})
	}
}
