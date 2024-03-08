package thirdMaximumNumber

import "math"

func thirdMax(nums []int) int {
	n := 3
	k := len(nums)
	c := 0

	m := make([]int, n)

	for i := 0; i < n; i++ {
		m[i] = math.MinInt
	}

	for i := 0; i < k; i++ {
	out:
		for j := 0; j < n; j++ {
			if nums[i] == m[j] {
				break out
			}
			if nums[i] > m[j] {
				for k := n - 1; k > j; k-- {
					m[k] = m[k-1]
				}
				m[j] = nums[i]
				c++
				break
			}
		}
	}

	if c >= 3 {
		return m[n-1]
	} else {
		return m[0]
	}
}
