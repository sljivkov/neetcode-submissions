func maxProfit(prices []int) int {

	var (
		res = 0
		l = 0
		r = 1
	)

	for l <= len(prices) - 2 {

		for r <= len(prices) - 1 && prices[l] < prices[r] {
			res = max(res, prices[r] - prices[l])
			r += 1
		}
		l = r
		r += 1

	}
	return res
}
