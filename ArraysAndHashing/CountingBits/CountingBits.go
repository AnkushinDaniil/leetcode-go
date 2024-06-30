package countingBits

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = res[i>>1] + (i & 1)
	}
	return res
}

func countBitsOld(n int) []int {
	res := make([]int, n+1)
	prev := 1
	cur := 2
	for i := 1; i <= n; i++ {
		if i < cur {
			res[i] = res[i-prev] + 1
		} else {
			res[i] = 1
			prev, cur = cur, cur<<1
		}
	}
	return res
}
