package koko_eating_bananas

import "math"

//Runtime: 52 ms, faster than 85.84% of Go online submissions for Koko Eating Bananas.
//Memory Usage: 6.6 MB, less than 95.58% of Go online submissions for Koko Eating Bananas.
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
		//totalHours += (p + m - 1) / m // totalHours += (p+m-1)/m equal to totalHours += Math.ceil(p/m)
		totalHours += p / m
		if p%m != 0 {
			totalHours++
		}
		//if totalHours > H {
		//	return false
		//}
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
