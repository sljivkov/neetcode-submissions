func findMin(nums []int) int {
	res := 99999
	l, r := 0, len(nums)-1

	for l <= r {
		curr := (l + r) / 2
		res = min(res, nums[curr])

		if nums[curr] > nums[r] {
			l = curr + 1 
		} else  {
			r = curr - 1
		}

	}

	return res
}
