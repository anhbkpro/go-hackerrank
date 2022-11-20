package arranging_coins

// Runtime: 7 ms, faster than 69.03% of Go online submissions for Arranging Coins.
// Memory Usage: 2.2 MB, less than 70.80% of Go online submissions for Arranging Coins.
func arrangeCoins(n int) int {
	lo, hi := 0, n
	for lo <= hi {
		mid := (lo + hi) >> 1
		if mid*(mid+1)/2 <= n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Arranging Coins.
// Memory Usage: 2.2 MB, less than 70.80% of Go online submissions for Arranging Coins.
func arrangeCoinsDiscuss(n int) int {
	if n <= 1 {
		return n
	}
	if n <= 3 {
		if n == 3 {
			return 2
		}
		return 1
	}

	lo, hi := 0, n
	for lo <= hi {
		mid := (lo + hi) >> 1
		if mid*(mid+1)/2 <= n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}
