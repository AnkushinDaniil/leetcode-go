package serializeAndDeserializeBinaryTree

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]interface{}) *TreeNode = trees.CreateTreeFromArray

func Test(t *testing.T) {
	testTable := []struct {
		name         string
		root         *TreeNode
		outputTree   *TreeNode
		outputString string
	}{
		{
			name:         "Example 1",
			root:         CreateTreeFromArray([]interface{}{1, 2, 3, nil, nil, 4, 5}),
			outputTree:   CreateTreeFromArray([]interface{}{1, 2, 3, nil, nil, 4, 5}),
			outputString: "1,2,nil,nil,3,4,nil,nil,5,nil,nil,",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ser := Constructor()
			deser := Constructor()
			data := ser.serialize(testCase.root)
			ans := deser.deserialize(data)
			assert.Equal(
				t,
				testCase.outputTree,
				ans,
			)
			assert.Equal(
				t,
				testCase.outputString,
				data,
			)
		})
	}
}

// func Benchmark(b *testing.B) {
// 	root := CreateTreeFromArray([]interface{}{-10, 9, 20, nil, nil, 15, 7})
// 	for i := 0; i < b.N; i++ {
// 		maxPathSum(root)
// 	}
// }

// func BenchmarkMax(b *testing.B) {
// 	root := CreateTreeFromArray([]interface{}{-10, 9, 20, nil, nil, 15, 7})
// 	for i := 0; i < b.N; i++ {
// 		maxPathSum_(root)
// 	}
// }
