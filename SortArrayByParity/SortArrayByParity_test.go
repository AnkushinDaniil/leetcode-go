package sortArrayByParity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParity(t *testing.T) {
	testTable := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "OK",
			nums: []int{3, 1, 2, 4},
			want: []int{4, 2, 1, 3},
		},
		{
			name: "0",
			nums: []int{0},
			want: []int{0},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, sortArrayByParity(testCase.nums), testCase.want)
		})
	}
}
