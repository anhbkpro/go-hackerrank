package common_sense_dsa_book

func Search(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2 // left/lower mid
	switch {
	case len(a) == 0:
		result = -1 // not found (result < 0)

	// if the target is less than mid,
	//there's no way mid will be our answer, and we can exclude it very confidently using a[:mid] (exclude mid)
	// same as the target is larger than mid
	case a[mid] > search:
		result, searchCount = Search(a[:mid], search) // exclude mid
	case a[mid] < search:
		result, searchCount = Search(a[mid+1:], search) // exclude mid
		if result >= 0 {                                // expect the search always on the right of the mid if found (result >=0)
			// And this all comes from the simple fact that in a sorted list,
			// everything to the right of n will be greater or equal to it, and vice versa.
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found (result >=0)
	}
	searchCount++
	return
}
