package medianOfTwoSortedArrays

import "math"

func leftEval(nums *[]int, i int) int {
	if i > 0 {
		return (*nums)[i-1]
	} else {
		return math.MinInt
	}
}

func rightEval(nums *[]int, i, n int) int {
	if i < n {
		return (*nums)[i]
	} else {
		return math.MaxInt
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)

	if nums2Len < nums1Len {
		nums1, nums2 = nums2, nums1
		nums1Len, nums2Len = nums2Len, nums1Len
	}

	totalLen := nums1Len + nums2Len
	left := 0
	right := nums1Len
	realMedian := (totalLen + 1) / 2

	for {
		m1 := (left + right) / 2
		m2 := realMedian - m1

		nums1MedLeft := leftEval(&nums1, m1)
		nums2MedLeft := leftEval(&nums2, m2)

		nums1MedRight := rightEval(&nums1, m1, nums1Len)
		nums2MedRight := rightEval(&nums2, m2, nums2Len)

		if nums1MedLeft > nums2MedRight {
			right = m1 - 1
		} else if nums2MedLeft > nums1MedRight {
			left = m1 + 1
		} else {
			maxLeft := max(nums1MedLeft, nums2MedLeft)
			minRight := min(nums1MedRight, nums2MedRight)
			if totalLen&1 == 1 {
				return float64(maxLeft)
			} else {
				return float64(maxLeft+minRight) / 2.0
			}
		}
	}
}
