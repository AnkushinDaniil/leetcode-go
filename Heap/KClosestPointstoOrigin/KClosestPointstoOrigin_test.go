package kClosestPointstoOrigin

import (
	"fmt"
	"slices"
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
		{
			points: [][]int{{-5, 4}, {-3, 2}, {0, 1}, {-3, 7}, {-2, 0}, {-4, -6}, {0, -5}},
			k:      6,
			output: [][]int{{0, 1}, {-2, 0}, {-3, 2}, {0, -5}, {-5, 4}, {-4, -6}},
		},
		{
			points: [][]int{{89, 6}, {-39, -4}, {-13, 91}, {97, -61}, {1, 7}, {-66, 69}, {-51, 68}, {82, -6}, {-21, 44}, {-58, -83}, {-40, 73}, {-88, -24}},
			k:      8,
			output: [][]int{{1, 7}, {-39, -4}, {-21, 44}, {82, -6}, {-40, 73}, {-51, 68}, {89, 6}, {-88, -24}},
		},
		{
			points: [][]int{{-66, 42}, {-67, 94}, {46, 10}, {35, 27}, {-9, -6}, {-84, -97}, {38, -22}, {3, -39}, {71, -97}, {-40, -86}, {-45, 56}, {-91, 59}, {24, -11}, {91, 100}, {-68, 43}, {34, 27}},
			k:      6,
			output: [][]int{{-9, -6}, {24, -11}, {3, -39}, {34, 27}, {38, -22}, {35, 27}},
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			res := kClosest(testCase.points, testCase.k)[:]
			slices.SortFunc(res, compare)
			expected := testCase.output[:]
			slices.SortFunc(expected, compare)
			assert.Equal(t, expected, res)
		})
	}
}

func compare(a, b []int) int {
	return (a[0]*a[0] + a[1]*a[1]) - (b[0]*b[0] + b[1]*b[1])
}
