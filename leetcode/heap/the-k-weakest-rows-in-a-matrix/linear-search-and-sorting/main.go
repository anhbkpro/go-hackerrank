package linear_search_and_sorting

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	// 1. count soldiers in each row
	pairs := make([][2]int, len(mat))
	for i, row := range mat {
		soldiers := 0
		for _, v := range row {
			if v == 1 {
				soldiers++
			}
		}
		pairs[i] = [2]int{i, soldiers}
	}

	// 2. sort rows by soldiers count
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})

	// 3. return first k rows
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i][0]
	}
	return result
}
