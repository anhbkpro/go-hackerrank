package unique_paths

func UniquePaths(m int, n int) int {
	row := make([]int, n)
	for i := 0; i < n; i++ {
		row[i] = 1
	}
	for i := 0; i < m-1; i++ {
		newRow := make([]int, n)
		for j := 0; j < n; j++ {
			newRow[j] = 1
		}
		for k := n - 2; k >= 0; k-- {
			newRow[k] = newRow[k+1] + row[k]
		}
		row = newRow
	}

	return row[0]
}
