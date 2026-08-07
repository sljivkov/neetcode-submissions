func isPalindrome(s string) bool {
	l,r := 0, len(s)-1
	runes := []rune(s)

	

	for l < r {

		for l < r && !isAlnum(unicode.ToLower(runes[l])) {
			l++
		}
		for l < r && !isAlnum(unicode.ToLower(runes[r])) {
			r--
		}

		if unicode.ToLower(runes[l]) != unicode.ToLower(runes[r]){
			return false
		}

		l++
		r--

	}

	return true
}

func isAlnum(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}