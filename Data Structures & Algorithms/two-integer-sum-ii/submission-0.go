func twoSum(numbers []int, target int) []int {
	res := make([]int, 0)
	mapa := make(map[int]int)

	for idx, num := range numbers {

		if val, ok := mapa[num]; ok {
			return []int{val,idx+1}
		} 
		mapa[target - num] = idx + 1

	}

	return res
}
