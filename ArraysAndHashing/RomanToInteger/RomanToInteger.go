package romanToInteger

var hashTable = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	prev := hashTable[s[0]]
	res := 0 + prev
	for i := 1; i < len(s); i++ {
		val := hashTable[s[i]]
		if val > prev {
			res -= 2 * prev
		}
		res += val
		prev = val
	}
	return res
}
