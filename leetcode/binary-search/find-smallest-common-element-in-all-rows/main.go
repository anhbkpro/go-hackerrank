package find_smallest_common_element_in_all_rows

import (
	"fmt"
	"math"
)

//Runtime: 170 ms, faster than 27.27% of Go online submissions for Find Smallest Common Element in All Rows.
//Memory Usage: 10.8 MB, less than 58.59% of Go online submissions for Find Smallest Common Element in All Rows.
func smallestCommonElement(mat [][]int) int {
	m := make(map[int]int)
	for _, row := range mat {
		for _, i := range row {
			m[i]++
		}
	}
	ans := make([]int, 0)
	for k, v := range m {
		if v == len(mat) {
			ans = append(ans, k)
		}
	}
	fmt.Println(ans)
	if len(ans) == 0 {
		return -1
	}
	min := math.MaxInt
	for _, i := range ans {
		if i < min {
			min = i
		}
	}
	return min
}
