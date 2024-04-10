package linkedListCycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		head   *ListNode
		output bool
	}{
		{
			name:   "Example 1",
			head:   CreateListWithCycleFromArray([]int{3, 2, 0, -4}, 1),
			output: true,
		},
		{
			name:   "Example 2",
			head:   CreateListWithCycleFromArray([]int{1, 2}, 0),
			output: true,
		},
		{
			name:   "Example 3",
			head:   CreateListWithCycleFromArray([]int{1}, -1),
			output: false,
		},
		{
			name:   "Example 4",
			head:   CreateListWithCycleFromArray([]int{1}, 0),
			output: true,
		},
		{
			name:   "Example 5",
			head:   CreateListWithCycleFromArray([]int{}, -1),
			output: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, hasCycle(testCase.head), testCase.output)
		})
	}
}

func CreateListWithCycleFromArray(arr []int, pos int) *ListNode {
	res := make([]ListNode, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		res[i] = ListNode{
			Val:  arr[i],
			Next: &res[i+1],
		}
	}
	if len(arr) > 0 {
		if pos >= 0 {
			res[len(arr)-1] = ListNode{
				Val:  arr[len(arr)-1],
				Next: &res[pos],
			}
		} else {
			res[len(arr)-1] = ListNode{
				Val:  arr[len(arr)-1],
				Next: nil,
			}
		}
		return &res[0]
	}

	return nil
}
