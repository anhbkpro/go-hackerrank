package minimum_time_to_complete_trips

import "math"

//Runtime: 661 ms, faster than 58.33% of Go online submissions for Minimum Time to Complete Trips.
//Memory Usage: 9 MB, less than 66.67% of Go online submissions for Minimum Time to Complete Trips
func minimumTime(time []int, totalTrips int) int64 {
	l, r := int64(1), int64(max(time)*totalTrips)
	for l < r {
		m := (l + r) >> 1
		if feasible(time, totalTrips, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func feasible(time []int, totalTrips int, m int64) bool {
	total := int64(0)
	for _, t := range time {
		if m >= int64(t) {
			total += m / int64(t)
		}
	}
	return total >= int64(totalTrips)
}

func max(arr []int) int {
	res := math.MinInt
	for _, i := range arr {
		if i > res {
			res = i
		}
	}
	return res
}
