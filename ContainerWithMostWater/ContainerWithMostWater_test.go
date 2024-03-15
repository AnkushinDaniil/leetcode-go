package containerWithMostWater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want: 49,
		},
		{
			name: "Example 2",
			nums: []int{1, 1},
			want: 1,
		},
		{
			name: "Example 3",
			nums: []int{1, 2, 4, 3},
			want: 4,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, maxArea(testCase.nums), testCase.want)
		})
	}
}
