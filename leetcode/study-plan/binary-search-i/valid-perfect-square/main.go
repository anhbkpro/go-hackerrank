package valid_perfect_square

// IsPerfectSquare
//	My Solution
//	Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Perfect Square.
//	Memory Usage: 1.9 MB, less than 86.52% of Go online submissions for Valid Perfect Square.
func IsPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	lo, hi := 0, num
	for lo < hi {
		mid := (lo + hi) >> 1
		if mid*mid < num {
			lo = mid + 1
		} else if mid*mid > num {
			hi = mid
		} else {
			return true
		}
	}
	return false
}
