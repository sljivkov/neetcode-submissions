type temperature struct {
	idx   int
	value int
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]temperature, 0)

	for idx, temp := range temperatures {

		for len(stack) > 0 && temp > stack[len(stack)-1].value {
			poped := stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			res[poped.idx] = idx - poped.idx

		}
		stack = append(stack, temperature{idx, temp})

	}
	return res
}