package kClosestPointstoOrigin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		points [][]int
		k      int
		output [][]int
	}{
		{
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			output: [][]int{{-2, 2}},
		},
		{
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			output: [][]int{{-2, 4}, {3, 3}},
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, kClosest(testCase.points, testCase.k))
		})
	}
}
