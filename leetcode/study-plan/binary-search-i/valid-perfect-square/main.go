package valid_perfect_square

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
