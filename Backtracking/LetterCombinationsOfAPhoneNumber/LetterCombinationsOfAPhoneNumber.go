package letterCombinationsOfAPhoneNumber

var table = [8]string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	n := 1
	for i := 0; i < len(digits); i++ {
		n *= len(table[digits[i]-'2'])
	}
	res := make([]string, 0, n)

	substring := make([]byte, 0, len(digits))

	var backtrack func(int)
	backtrack = func(i int) {
		if i == len(digits) {
			res = append(res, string(substring))
			return
		}
		for j := 0; j < len(table[digits[i]-'2']); j++ {
			substring = append(substring, table[digits[i]-'2'][j])
			backtrack(i + 1)
			substring = substring[:len(substring)-1]
		}
	}

	backtrack(0)

	return res
}

func letterCombinationsEmpty(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	res := make([]string, 0)

	substring := make([]byte, 0, len(digits))

	var backtrack func(int)
	backtrack = func(i int) {
		if i == len(digits) {
			res = append(res, string(substring))
			return
		}
		for j := 0; j < len(table[digits[i]-'2']); j++ {
			substring = append(substring, table[digits[i]-'2'][j])
			backtrack(i + 1)
			substring = substring[:len(substring)-1]
		}
	}

	backtrack(0)

	return res
}

func letterCombinationsRange(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	n := 1
	for _, digit := range digits {
		n *= len(table[digit-'2'])
	}
	res := make([]string, 0, n)

	substring := make([]rune, 0, len(digits))

	var backtrack func(int)
	backtrack = func(i int) {
		if i == len(digits) {
			res = append(res, string(substring))
			return
		}
		for _, letter := range table[digits[i]-'2'] {
			substring = append(substring, letter)
			backtrack(i + 1)
			substring = substring[:len(substring)-1]
		}
	}

	backtrack(0)

	return res
}

func letterCombinationsMap(digits string) []string {
	mapper := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ln := len(digits)
	result := make([]string, 0, ln)

	var walker func(cursor int, combo []rune)
	walker = func(cursor int, combo []rune) {
		if cursor >= ln && len(combo) > 0 {
			c := make([]rune, len(combo))
			copy(c, combo)
			result = append(result, string(c))
			return
		} else if cursor >= ln {
			return
		}

		for _, v := range mapper[digits[cursor]] {
			walker(cursor+1, append(combo, v))
		}
	}

	walker(0, make([]rune, 0, ln))

	return result
}

func letterCombinationsRecurse(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	alphabet := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}
	recurse(&res, "", digits, alphabet)

	return res
}

func recurse(res *[]string, word, digits string, alphabet []string) {
	if len(digits) > 0 {
		letters := alphabet[digits[0]-'2']
		for _, letter := range letters {
			recurse(res, word+string(letter), digits[1:], alphabet)
		}
	} else {
		*res = append(*res, word)
	}
}
