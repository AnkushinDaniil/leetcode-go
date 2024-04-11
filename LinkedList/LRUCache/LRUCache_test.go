package lruCache

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
				"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get",
			},
			input:  [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			output: [][]int{{}, {}, {}, {1}, {}, {-1}, {}, {-1}, {3}, {4}},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var lruCache LRUCache
			res := make([][]int, 0, len(testCase.functions))
			for i := 0; i < len(testCase.functions); i++ {
				switch testCase.functions[i] {
				case "LRUCache":
					lruCache = Constructor(testCase.input[i][0])
					res = append(res, []int{})
				case "get":
					res = append(res, []int{lruCache.Get(testCase.input[i][0])})
				case "put":
					lruCache.Put(testCase.input[i][0], testCase.input[i][1])
					res = append(res, []int{})
				}
			}
			assert.Equal(t, res, testCase.output)
		})
	}
}
