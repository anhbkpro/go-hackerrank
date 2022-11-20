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
