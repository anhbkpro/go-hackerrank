package find_smallest_letter_greater_than_target

func NextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if target >= letters[n-1] {
		return letters[0]
	}

	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if letters[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return letters[hi]
}
