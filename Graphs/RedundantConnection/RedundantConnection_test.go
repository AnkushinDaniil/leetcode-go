package redundantConnection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		edges  [][]int
		output []int
	}{
		{
			edges:  [][]int{{1, 2}, {1, 3}, {2, 3}},
			output: []int{2, 3},
		},
		{
			edges:  [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			output: []int{1, 4},
		},
		{
			edges: [][]int{
				{21, 22},
				{4, 7},
				{12, 13},
				{13, 25},
				{12, 15},
				{17, 23},
				{15, 16},
				{8, 21},
				{23, 24},
				{3, 9},
				{19, 21},
				{13, 21},
				{4, 10},
				{5, 6},
				{1, 20},
				{10, 16},
				{1, 4},
				{10, 14},
				{5, 20},
				{1, 2},
				{3, 24},
				{2, 11},
				{11, 24},
				{24, 25},
				{17, 18},
			},
			output: []int{24, 25},
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, findRedundantConnection(testCase.edges))
		})
	}
}
