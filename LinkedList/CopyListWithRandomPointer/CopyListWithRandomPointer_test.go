package copyListWithRandomPointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		head   *Node
		output *Node
	}{
		{
			name: "Example 1",
			head: CreateListFromArray([][]int{
				{7, -1},
				{13, 0},
				{11, 4},
				{10, 2},
				{1, 0},
			}),
			output: CreateListFromArray([][]int{
				{7, -1},
				{13, 0},
				{11, 4},
				{10, 2},
				{1, 0},
			}),
		},
		{
			name: "Example 2",
			head: CreateListFromArray([][]int{
				{1, 1},
				{2, 1},
			}),
			output: CreateListFromArray([][]int{
				{1, 1},
				{2, 1},
			}),
		},
		{
			name: "Example 3",
			head: CreateListFromArray([][]int{
				{3, -1},
				{3, 0},
				{3, -1},
			}),
			output: CreateListFromArray([][]int{
				{3, -1},
				{3, 0},
				{3, -1},
			}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, copyRandomList(testCase.head), testCase.output)
		})
	}
}

func CreateListFromArray(arr [][]int) *Node {
	if len(arr) == 0 {
		return nil
	}

	res := make([]Node, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		res[i] = Node{
			Val:  arr[i][0],
			Next: &res[i+1],
		}
	}
	res[len(arr)-1] = Node{
		Val:  arr[len(arr)-1][0],
		Next: nil,
	}
	for i := 0; i < len(arr); i++ {
		if arr[i][1] == -1 {
			res[i].Random = nil
		} else {
			res[i].Random = &res[arr[i][1]]
		}
	}

	return &res[0]
}
