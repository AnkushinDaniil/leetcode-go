package kthLargestElementInAnArray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		array  []int
		k      int
		output int
	}{
		{
			array:  []int{3, 2, 1, 5, 6, 4},
			k:      2,
			output: 5,
		},
		{
			array:  []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:      4,
			output: 4,
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, findKthLargest(testCase.array, testCase.k))
		})
	}
}
