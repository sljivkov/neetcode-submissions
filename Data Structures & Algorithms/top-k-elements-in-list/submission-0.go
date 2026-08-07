func topKFrequent(nums []int, k int) []int {

    mapa := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		mapa[num] += 1
	}

	for key, value := range mapa {
		freq[value] = append(freq[value], key)
	}
	res := []int{}
	for i := len(freq)-1; i >= 0; i-- {
		for _, value := range freq[i]{
			res = append(res, value)
			if len(res) == k{
				return res
			}
		}
	}
	return nil
}
