package subsetsII

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
			nums:   []int{1, 2, 2},
			output: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			nums:   []int{0},
			output: [][]int{{}, {0}},
		},
	}

	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, subsetsWithDup(testCase.nums))
		})
	}
}
