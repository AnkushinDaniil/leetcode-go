package thirdMaximumNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	testTable := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "Example 3",
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
		{
			name: "Example 4",
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			name: "Example 5",
			nums: []int{1, 2, 2, 5, 3, 5},
			want: 2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, thirdMax(testCase.nums), testCase.want)
		})
	}
}
