func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))

	for i, _ := range nums{
		prefix[i] = 1
		postfix[i] = 1

	}

	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		postfix[i] = nums[i+1] * postfix[i+1]
	}

	for i := range nums {
		res[i] = prefix[i] * postfix[i]
	}

	return res

}