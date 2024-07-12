package ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	alphabet := [26]uint16{}
	for i := range magazine {
		alphabet[magazine[i]-'a']++
	}
	for i := range ransomNote {
		idx := ransomNote[i] - 'a'
		if alphabet[idx] <= 0 {
			return false
		}
		alphabet[idx]--
	}
	return true
}
