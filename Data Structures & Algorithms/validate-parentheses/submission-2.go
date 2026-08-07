func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	res := make([]rune, 0)
	mapa := map[rune]rune {
		'}': '{',
		']': '[',
		')': '(',
	}

	for _,a := range s {
		if val, ok := mapa[a]; !ok {
			res = append(res, a)
		} else if len(res) > 0 && val == res[len(res)-1] {
			res = res[:len(res)-1]
		} else {
			return false
		}

	}

	return len(res) == 0
}
