func minEatingSpeed(piles []int, h int) int {
	l, r := 1, maxInSlice(piles)

	for l <= r {
		curr := (l + r) / 2
		temp := 0

		for _, p := range piles {
			temp += (curr + p - 1) / curr
		}

		if temp > h {
			l = curr + 1
		} else {
			r = curr - 1
		}
	}

	return l
}
func maxInSlice(nums []int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}