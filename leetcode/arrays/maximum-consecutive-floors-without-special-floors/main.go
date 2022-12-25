package maximum_consecutive_floors_without_special_floors

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Slice(special, func(i, j int) bool {
		return special[i] < special[j]
	})

	maxBottom := special[0] - bottom
	maxTop := top - special[len(special)-1]
	res := max(maxBottom, maxTop)
	for i := 0; i < len(special)-1; i++ {
		res = max(res, special[i+1]-special[i]-1)
	}
	return res
}
