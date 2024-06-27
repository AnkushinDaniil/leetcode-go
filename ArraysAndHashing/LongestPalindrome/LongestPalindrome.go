package longestPalindrome

const MASK = ^0 - 1

func longestPalindrome(s string) int {
	tables := [2][26]int{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' {
			tables[0][s[i]-'a']++
		} else {
			tables[1][s[i]-'A']++
		}
	}
	res := 0

	i := 0
	solo := false
	for ; i < 26 && !solo; i++ {
		if tables[0][i]&1 == 1 {
			solo = true
		}
		if tables[1][i]&1 == 1 {
			solo = true
		}
		res += tables[0][i] & MASK
		res += tables[1][i] & MASK
	}

	for ; i < 26; i++ {
		res += tables[0][i] & MASK
		res += tables[1][i] & MASK
	}

	if solo {
		res++
	}

	return res
}
