
func characterReplacement(s string, k int) int {
	res, l, maxF := 0, 0, 0

	mapa := make(map[byte]int)

	for r := 0; r < len(s); r++{
		mapa[s[r]]++

		maxF = max(maxF, mapa[s[r]])

		for (r - l + 1) - maxF > k {
			mapa[s[l]]--
			l++

		}
		
		res = max(res, r-l+1)

	}

	return res
}

