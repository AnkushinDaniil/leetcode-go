package permutationInString

func checkInclusion(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 > l2 {
		return false
	}

	table1 := [26]uint16{}
	table2 := [26]uint16{}

	for i := range s1 {
		table1[s1[i]-'a']++
	}

	for i := range s2[:l1] {
		table2[s2[i]-'a']++
	}

	for i := range s2[:l2-l1] {
		if table1 == table2 {
			return true
		} else {
			table2[s2[i]-'a']--
			table2[s2[i+l1]-'a']++
		}
	}
	return table1 == table2
}
