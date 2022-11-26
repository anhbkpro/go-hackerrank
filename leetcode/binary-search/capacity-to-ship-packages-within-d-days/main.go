package capacity_to_ship_packages_within_d_days

//Runtime: 49 ms, faster than 66.67% of Go online submissions for Capacity To Ship Packages Within D Days.
//Memory Usage: 7 MB, less than 37.50% of Go online submissions for Capacity To Ship Packages Within D Days.
// This solution can be used to solve: Leetcode 410. Split Array Largest Sum
// How about Leetcode 875. Koko Eating Bananas?
// weights = piles
// days = hours
// find the least weight capacity of the ship = find the minimum integer k (her bananas-per-hour eating speed of k)
func shipWithinDays(weights []int, days int) int {
	l, r := maxAndSum(weights)
	for l < r {
		m := (l + r) >> 1
		if feasible(weights, days, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// condition function:
//if we can successfully ship all packages within D days with capacity m,
//then we can definitely ship them all with any capacity larger than m
func feasible(weights []int, D, capacity int) bool {
	days, total := 1, 0
	for _, weight := range weights {
		total += weight
		if total > capacity { // too heavy, wait for the next day
			total = weight
			days++
		}
	}
	return days <= D // can ship within D days
}

// we need to initialize our boundary correctly
func maxAndSum(arr []int) (int, int) {
	max, sum := 0, 0
	for _, item := range arr {
		sum += item
		if max < item {
			max = item
		}
	}
	return max, sum
}
