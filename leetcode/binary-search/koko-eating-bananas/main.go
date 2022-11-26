package koko_eating_bananas

import "math"

//Runtime: 47 ms, faster than 90.83% of Go online submissions for Koko Eating Bananas.
//Memory Usage: 6.6 MB, less than 96.33% of Go online submissions for Koko Eating Bananas.
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, max(piles) // slowest speed is 1 banana per hour
	for l < r {
		m := (l + r) >> 1
		if feasible(piles, h, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func feasible(piles []int, H, m int) bool {
	totalHours := 0 // Hours take to eat all bananas at speed m.
	for _, p := range piles {
		totalHours += (p + m - 1) / m
	}
	return totalHours <= H
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
