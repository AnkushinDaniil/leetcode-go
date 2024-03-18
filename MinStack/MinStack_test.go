package minStack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name      string
		functions []string
		input     [][]int
		output    [][]int
	}{
		{
			name: "Example 1",
			functions: []string{
				"MinStack",
				"push",
				"push",
				"push",
				"getMin",
				"pop",
				"top",
				"getMin",
			},
			input:  [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			output: [][]int{{}, {}, {}, {}, {-3}, {}, {0}, {-2}},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var minStack MinStack
			res := make([][]int, 0, len(testCase.functions))
			for i := 0; i < len(testCase.functions); i++ {
				switch testCase.functions[i] {
				case "MinStack":
					minStack = Constructor()
					res = append(res, []int{})
				case "push":
					minStack.Push(testCase.input[i][0])
					res = append(res, []int{})
				case "pop":
					minStack.Pop()
					res = append(res, []int{})
				case "getMin":
					res = append(res, []int{minStack.GetMin()})
				case "top":
					res = append(res, []int{minStack.Top()})
				}
			}
			assert.Equal(t, res, testCase.output)
		})
	}
}
