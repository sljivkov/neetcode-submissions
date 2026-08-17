func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	arr1 := make([]int, 26)
	arr2 := make([]int, 26)

	matches := 0
	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}

	for i := range 26 {
		if arr1[i] == arr2[i] {
			matches++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}

		curr := s2[r] - 'a'
		arr2[curr]++
		if arr1[curr] == arr2[curr] {
			matches++
		} else if arr1[curr]+1 == arr2[curr] {
			matches--
		}

		curr = s2[l] - 'a'
		arr2[curr]--
		if arr1[curr] == arr2[curr] {
			matches++
		} else if arr1[curr]-1 == arr2[curr] {
			matches--
		}
		l++
	}

	return matches == 26
}

