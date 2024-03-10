package containsDuplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	testTable := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "Example 3",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, containsDuplicate(testCase.nums), testCase.want)
		})
	}
}
