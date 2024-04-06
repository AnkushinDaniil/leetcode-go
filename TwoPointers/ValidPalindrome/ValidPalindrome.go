package validPalindrome

func isPalindrome(s string) bool {
	b := make([]byte, 0, len(s))

	for i := range s {
		if s[i] == ' ' {
			continue
		} else if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			b = append(b, s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			b = append(b, s[i]-'A'+'a')
		}
	}

	n := len(b)

	for i := 0; i < n/2; i++ {
		if b[i] != b[n-1-i] {
			return false
		}
	}

	return true
}
