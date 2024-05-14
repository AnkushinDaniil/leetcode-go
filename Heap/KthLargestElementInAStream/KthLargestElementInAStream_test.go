package kthLargestElementInAStream

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		funcs            []string
		constructorParam []int
		params           []int
		output           []int
	}{
		{
			funcs:            []string{"KthLargest", "add", "add", "add", "add", "add"},
			constructorParam: []int{3, 4, 5, 8, 2},
			params:           []int{0, 3, 5, 10, 9, 4},
			output:           []int{0, 4, 5, 5, 8, 8},
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			var heap KthLargest
			res := make([]int, len(testCase.funcs))
			for i := 0; i < len(testCase.funcs); i++ {
				switch testCase.funcs[i] {
				case "KthLargest":
					heap = Constructor(testCase.constructorParam[0], testCase.constructorParam[1:])
					res[i] = 0
				case "add":
					res[i] = heap.Add(testCase.params[i])
				}
			}
			assert.Equal(t, testCase.output, res)
		})
	}
}

func BenchmarkHeap(b *testing.B) {
	params := []int{0, 3, 5, 10, 9, 4}
	heap := Constructor(3, []int{4, 5, 8, 2})
	for i := 0; i < b.N; i++ {
		heap.Add(params[i%len(params)])
	}
}
