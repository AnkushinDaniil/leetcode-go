package heightChecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParity(t *testing.T) {
	testTable := []struct {
		name string
		nums []int
		want int
	}{
		// {
		// 	name: "Example 1",
		// 	nums: []int{1, 1, 4, 2, 1, 3},
		// 	want: 3,
		// },
		// {
		// 	name: "Example 2",
		// 	nums: []int{5, 1, 2, 3, 4},
		// 	want: 5,
		// },
		// {
		// 	name: "Example 3",
		// 	nums: []int{1, 2, 3, 4, 5},
		// 	want: 0,
		// },
		{
			name: "Example 4",
			nums: []int{7, 4, 5, 6, 4, 2, 1, 4, 6, 5, 4, 8, 3, 1, 8, 2, 7, 6, 3, 2},
			want: 17,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, heightChecker(testCase.nums), testCase.want)
		})
	}
}

func TestQuickSortAlgorithm(t *testing.T) {
	testTable := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 4, 2, 1, 3},
			want: []int{1, 1, 1, 2, 3, 4},
		},
		{
			name: "Example 2",
			nums: []int{5, 1, 2, 3, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Example 3",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			quickSortAlgorithm(&testCase.nums)
			assert.Equal(t, testCase.nums, testCase.want)
		})
	}
}
