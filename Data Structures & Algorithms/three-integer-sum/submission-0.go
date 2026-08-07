func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := a + nums[l] + nums[r]

			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}