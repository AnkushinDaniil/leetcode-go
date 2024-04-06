package medianOfTwoSortedArrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		nums1  []int
		nums2  []int
		output float64
	}{
		{
			name:   "Example 1",
			nums1:  []int{1, 3},
			nums2:  []int{2},
			output: 2.0,
		},
		{
			name:   "Example 2",
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			output: 2.5,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, findMedianSortedArrays(testCase.nums1, testCase.nums2), testCase.output)
		})
	}
}
