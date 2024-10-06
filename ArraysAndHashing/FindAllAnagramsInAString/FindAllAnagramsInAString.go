package findallanagramsinastring

func index(ch byte) byte {
	return ch - 'a'
}

func findAnagrams(s string, p string) []int {
  if len(s) < len(p) {
    return []int{}
  }

	var alphabet [26]int
	nonZeros := 0

	for i := range p {
		idx := index(p[i])
		if alphabet[idx] == 0 {
			nonZeros++
		}
		alphabet[idx]++
	}

	for i := range p {
		idx := index(s[i])
		alphabet[idx]--
		if alphabet[idx] == 0 {
			nonZeros--
		}
	}

	res := make([]int, 0, len(s)-len(p)+1)

	if nonZeros == 0 {
		res = append(res, 0)
	}

	for i := len(p); i < len(s); i++ {
		leftIdx := index(s[i-len(p)])
		if alphabet[leftIdx] == 0 {
			nonZeros++
		}
		alphabet[leftIdx]++

		rightIxd := index(s[i])
		alphabet[rightIxd]--
		if alphabet[rightIxd] == 0 {
			nonZeros--
		}

    if nonZeros == 0 {
      res = append(res, i-len(p)+1)
    }
	}

  return res
}
