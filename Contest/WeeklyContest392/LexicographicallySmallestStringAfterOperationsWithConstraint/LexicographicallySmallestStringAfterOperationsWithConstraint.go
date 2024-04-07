package binarySearch

func getSmallestString(s string, k int) string {
	t := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		target := s[i]
		for ch := byte('a'); ch <= s[i]; ch++ {
			distLeft, distRight := int(ch+26-s[i]), int(s[i]-ch)

			if distLeft <= k || distRight <= k {
				if distLeft < distRight {
					k -= distLeft
				} else {
					k -= distRight
				}
				target = ch
				break
			}
		}
		t = append(t, target)
	}
	return string(t)
}
