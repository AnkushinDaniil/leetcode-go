package binarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		matrix [][]int
		target int
		output bool
	}{
		{
			name: "Example 1",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			output: true,
		},
		{
			name: "Example 2",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 13,
			output: false,
		},
		{
			name: "Example 3",
			matrix: [][]int{
				{1, 1},
			},
			target: 2,
			output: false,
		},
		{
			name: "Example 4",
			matrix: [][]int{
				{1},
			},
			target: 1,
			output: true,
		},
		{
			name: "Example 5",
			matrix: [][]int{
				{1},
				{3},
			},
			target: 3,
			output: true,
		},
		{
			name: "Example 6",
			matrix: [][]int{
				{1},
			},
			target: 0,
			output: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, searchMatrix(testCase.matrix, testCase.target), testCase.output)
		})
	}
}
