package courseScheduleII

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		numCourses    int
		prerequisites [][]int
		output        []int
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			output:        []int{0, 1},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			output:        []int{0, 1, 2, 3},
		},
		{
			numCourses:    1,
			prerequisites: [][]int{},
			output:        []int{0},
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, findOrder(testCase.numCourses, testCase.prerequisites))
		})
	}
}
