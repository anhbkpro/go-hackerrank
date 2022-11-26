package minimized_maximum_of_products_distributed_to_any_store

import "math"

//Runtime: 411 ms, faster than 40.00% of Go online submissions for Minimized Maximum of Products Distributed to Any Store.
//Memory Usage: 9.2 MB, less than 20.00% of Go online submissions for Minimized Maximum of Products Distributed to Any Store.
func minimizedMaximum(n int, quantities []int) int {
	l, r := 1, max(quantities)
	for l < r {
		m := (l + r) >> 1
		if feasible(n, quantities, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func feasible(stores int, quantities []int, m int) bool {
	total := 0
	for _, quantity := range quantities {
		total += quantity / m
		if quantity%m != 0 {
			total++
		}
	}

	return stores >= total
}

func max(arr []int) int {
	res := math.MinInt
	for _, item := range arr {
		if item > res {
			res = item
		}
	}
	return res
}
