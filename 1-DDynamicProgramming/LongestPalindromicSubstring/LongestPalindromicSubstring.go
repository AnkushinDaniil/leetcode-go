package longestPalindromicSubstring

func longestPalindrome(s string) string {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		for l, r := i, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if r-l > right-left {
				right = r
				left = l
			}
		}
		for l, r := i, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if r-l > right-left {
				right = r
				left = l
			}
		}
	}
	return s[left : right+1]
}

// func longestPalindrome2(s string) string {
// 	var l, r, i int
// 	n := len(s)

// 	checkAndUpdate := func() func(int) {
// 		maximum := -1
// 		return func(one int) {
// 			for j := 0; i-j-1 >= 0 && i+j+one < n; j++ {
// 				if s[i-j-1] == s[i+j+one] {
// 					if j*2+one > maximum {
// 						l = i - j - 1
// 						r = i + j + one
// 						maximum = j*2 + one
// 					}
// 				} else {
// 					break
// 				}
// 			}
// 		}
// 	}()

// 	for i = 1; i <= n; i++ {
// 		checkAndUpdate(0)
// 		checkAndUpdate(1)
// 	}

// 	return s[l : r+1]
// }

func longestPalindrome2(s string) string {
	left, right, length, n := 0, 0, 0, len(s)
	for i := 0; i < n; i++ {
		for l, r := i, i; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
			if r-l > length {
				right = r
				left = l
				length = r - l
			}
		}
		for l, r := i, i+1; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
			if r-l > length {
				right = r
				left = l
				length = r - l
			}
		}
	}
	return s[left : right+1]
}
