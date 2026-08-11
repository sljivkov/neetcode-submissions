func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {

		if row[len(row)-1] >= target {
			l, r := 0, len(row)-1

			for l <= r {
				curr := (l + r) / 2

				if row[curr] == target {
					return true

				} else if row[curr] < target {
					l = curr + 1

				} else {
					r = curr - 1

				}
			}
		}
	}

	return false
}