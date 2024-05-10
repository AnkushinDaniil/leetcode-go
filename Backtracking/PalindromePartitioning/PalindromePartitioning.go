package palindromePartitioning

// func getFirstComposition(n, k int, composition []int) bool {
// 	if n < k {
// 		return false
// 	}
// 	for i := 0; i < k-1; i++ {
// 		composition[i] = 1
// 	}
// 	composition[k-1] = n - k + 1
// 	return true
// }

// func getNextComposition(n, k int, composition []int) bool {
// 	if composition[0] == n-k+1 {
// 		return false
// 	}
// 	last := k - 1
// 	for composition[last] == 1 {
// 		last--
// 	}

// 	z := composition[last]
// 	composition[last-1]++
// 	composition[last] = 1
// 	composition[k-1] = z - 1
// 	return true
// }

// func allCompositions(n int) [][]int {
// 	compositions := make([][]int, 0, 1<<n-1)
// 	for k := 1; k <= n; k++ {
// 		composition := make([]int, k)
// 		for exists := getFirstComposition(n, k, composition); exists; exists = getNextComposition(n, k, composition) {
// 			compositions = append(compositions, make([]int, len(composition)))
// 			copy(compositions[len(compositions)-1], composition)
// 		}
// 	}
// 	return compositions
// }

// func split(s string, composition []int) []string {
// 	res := make([]string, len(composition))
// 	index := 0
// 	for i := 0; i < len(composition); i++ {
// 		res[i] = s[index : index+composition[i]]
// 		index += composition[i]
// 	}
// 	return res
// }

func partitionIter(s string) [][]string {
	res := make([][]string, 0, 1<<len(s)-1)
	n := len(s)
	compositions := make([][]int, 0, 1<<n-1)

	for k := 1; k <= n; k++ {
		compositions = append(compositions, make([]int, k))
		for i := 0; i < k-1; i++ {
			compositions[k-1][i] = 1
		}
		compositions[k-1][k-1] = n - k + 1
		for compositions[k-1][0] != n-k+1 {
			last := k - 1
			for compositions[k-1][last] == 1 {
				last--
			}

			z := compositions[k-1][last]
			compositions[k-1][last-1]++
			compositions[k-1][last] = 1
			compositions[k-1][k-1] = z - 1
		}
	}
	for i := 0; i < len(compositions); i++ {
		isAllPalindrome := true
		index := 0
		for j := 0; j < len(compositions[i]); j++ {
			isPalindrome := true
			for a, b := index, index+compositions[i][j]-1; a < b; a, b = a+1, b-1 {
				if s[a] != s[b] {
					isPalindrome = false
					break
				}
			}
			if !isPalindrome {
				isAllPalindrome = false
				break
			}
			index += compositions[i][j]
		}
		if isAllPalindrome {
			// res = append(res, split(s, compositions[i]))
			res = append(res, make([]string, len(compositions[i])))
			index := 0
			for a := 0; a < len(compositions[i]); a++ {
				res[len(res)-1][a] = s[index : index+compositions[i][a]]
				index += compositions[i][a]
			}
		}
	}
	return res
}

func partition(s string) [][]string {
	res := make([][]string, 0, 1<<len(s)-1)
	stack := make([]string, 0, len(s))

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			res = append(res, make([]string, len(stack)))
			copy(res[len(res)-1], stack)
			return
		}

		for i := idx; i < len(s); i++ {
			if isPalindrome(s[idx : i+1]) {
				stack = append(stack, s[idx:i+1])
				backtrack(i + 1)
				stack = stack[:len(stack)-1]
			}
		}
	}

	backtrack(0)

	return res
}

func isPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
