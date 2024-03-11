package topKFrequentElements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	testTable := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "Example 2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, topKFrequent(testCase.nums, testCase.k), testCase.want)
		})
	}
}
