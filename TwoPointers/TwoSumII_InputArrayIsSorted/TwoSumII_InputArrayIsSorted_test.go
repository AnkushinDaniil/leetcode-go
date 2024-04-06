package twoSumII_InputArrayIsSorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "Example 1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "Example 2",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "Example 3",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, twoSum(testCase.numbers, testCase.target), testCase.want)
		})
	}
}
