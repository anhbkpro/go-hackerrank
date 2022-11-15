package valid_perfect_square

// IsPerfectSquare
//	My Solution
//	Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Perfect Square.
//	Memory Usage: 1.9 MB, less than 86.52% of Go online submissions for Valid Perfect Square.
func IsPerfectSquare(num int) bool {
	i, j := 0, num
	for i <= j {
		h := (i + j) >> 1 // avoid overflow when computing h. Refer naming src/sort/search.go
		if h*h == num {
			return true
		}
		if h*h < num {
			i = h + 1
		} else {
			j = h - 1
		}
	}
	return false
}
