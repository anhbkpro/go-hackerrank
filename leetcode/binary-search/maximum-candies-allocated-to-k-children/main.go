package maximum_candies_allocated_to_k_children

import "math"

//Runtime: 487 ms, faster than 11.76% of Go online submissions for Maximum Candies Allocated to K Children.
//Memory Usage: 9.5 MB, less than 58.82% of Go online submissions for Maximum Candies Allocated to K Children.
func maximumCandies(candies []int, k int64) int {
	l := 1
	r, total := maxSumTuple(candies)
	r++
	if int64(total) < k {
		return 0
	}
	for l < r {
		m := (l + r) >> 1
		if !feasible(candies, k, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}

func feasible(candies []int, k int64, m int) bool {
	count := 0
	for _, candy := range candies {
		count += candy / m
	}
	return k <= int64(count)
}

func maxSumTuple(arr []int) (int, int) {
	res := math.MinInt
	sum := 0
	for _, i := range arr {
		sum += i
		if res < i {
			res = i
		}
	}
	return res, sum
}
