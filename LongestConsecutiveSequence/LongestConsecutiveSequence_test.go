package longestConsecutiveSequence

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
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, longestConsecutive(testCase.nums), testCase.want)
		})
	}
}
