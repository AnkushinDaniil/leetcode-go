package trappingRainWater

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
			nums: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{4, 2, 0, 3, 2, 5},
			want: 9,
		},
		{
			name: "Example 3",
			nums: []int{2, 0, 2},
			want: 2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, trap(testCase.nums), testCase.want)
		})
	}
}
