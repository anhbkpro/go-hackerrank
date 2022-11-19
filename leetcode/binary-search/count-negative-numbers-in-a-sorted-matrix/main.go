package count_negative_numbers_in_a_sorted_matrix

func countNegatives(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	lastNegative := cols - 1
	ans := 0
	for row := 0; row < rows; row++ {
		if grid[row][0] < 0 {
			ans += cols
			continue
		}
		if grid[row][cols-1] >= 0 {
			continue
		}

		lo, hi := 0, lastNegative
		for lo <= hi {
			mid := (lo + hi) >> 1
			if grid[row][mid] < 0 {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		ans += cols - lo
		lastNegative = lo
	}
	return ans
}
