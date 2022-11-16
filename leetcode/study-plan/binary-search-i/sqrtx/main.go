package sqrtx

// MySqrt My Solution
// 	Runtime: 0 ms, faster than 100.00% of Go online submissions for Sqrt(x).
//	Memory Usage: 2.1 MB, less than 23.26% of Go online submissions for Sqrt(x).
func MySqrt(x int) int {
	lo, hi := 0, x
	for lo <= hi {
		mid := (lo + hi) >> 1
		if mid*mid > x {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo - 1
}
