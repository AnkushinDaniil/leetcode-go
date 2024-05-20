package findMedianFromDataStream

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		funcs  []string
		params []int
		output []float64
	}{
		{
			funcs: []string{
				"MedianFinder",
				"addNum",
				"addNum",
				"findMedian",
				"addNum",
				"findMedian",
			},
			params: []int{0, 1, 2, 0, 3, 0},
			output: []float64{0, 0, 0, 1.5, 0, 2.0},
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			var obj MedianFinder
			res := make([]float64, len(testCase.funcs))
			for i := 0; i < len(testCase.funcs); i++ {
				switch testCase.funcs[i] {
				case "MedianFinder":
					obj = Constructor()
					res[i] = 0
				case "addNum":
					obj.AddNum(testCase.params[i])
					res[i] = 0
				case "findMedian":
					res[i] = obj.FindMedian()
				}
			}
			assert.Equal(t, testCase.output, res)
		})
	}
}
