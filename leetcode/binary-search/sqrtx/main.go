package sqrtx

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Sqrt(x).
//Memory Usage: 2.1 MB, less than 23.43% of Go online submissions for Sqrt(x).
func mySqrt(x int) int {
	l, r := 0, x+1
	for l < r {
		m := (l + r) >> 1
		if m*m > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1 // `l` is the minimum k value, `k - 1` is the answer
}
