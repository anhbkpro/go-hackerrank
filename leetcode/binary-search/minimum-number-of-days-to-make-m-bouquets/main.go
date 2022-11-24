package minimum_number_of_days_to_make_m_bouquets

import "math"

//Runtime: 189 ms, faster than 80.00% of Go online submissions for Minimum Number of Days to Make m Bouquets.
//Memory Usage: 10.3 MB, less than 20.00% of Go online submissions for Minimum Number of Days to Make m Bouquets.
func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}
	l, r := minMax(bloomDay)
	for l < r {
		mid := (l + r) >> 1
		if feasible(bloomDay, mid, m, k) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func feasible(bloomDay []int, threshold, M, K int) bool {
	bouquets, flowers := 0, 0
	for _, bloom := range bloomDay {
		if bloom <= threshold {
			flowers++
		} else {
			flowers = 0
		}
		if flowers >= K {
			bouquets++
			flowers = 0
		}
	}
	return bouquets >= M
}

func minMax(arr []int) (int, int) {
	min, max := math.MaxInt, math.MinInt
	for _, item := range arr {
		if min > item {
			min = item
		}
		if max < item {
			max = item
		}
	}
	return min, max
}
