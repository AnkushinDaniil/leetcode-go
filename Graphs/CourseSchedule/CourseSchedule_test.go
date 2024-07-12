package courseSchedule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		numCourses    int
		prerequisites [][]int
		output        bool
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			output:        true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			output:        false,
		},
		{
			numCourses:    5,
			prerequisites: [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}},
			output:        true,
		},
		{
			numCourses:    7,
			prerequisites: [][]int{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}},
			output:        true,
		},
		{
			numCourses:    20,
			prerequisites: [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}},
			output:        false,
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, canFinish(testCase.numCourses, testCase.prerequisites))
		})
	}
}
