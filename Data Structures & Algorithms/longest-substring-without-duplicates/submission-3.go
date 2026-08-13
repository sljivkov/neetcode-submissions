func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 { return 0}
	res := 0
	l, r := 0, 0
	mapa := make(map[byte]bool)

	for r < len(s) {
		for mapa[s[r]] {
			mapa[s[l]] = false
			l++ 
		}
		mapa[s[r]] = true
		res = max(res, r-l+1)
		r++
	}

	return res
}