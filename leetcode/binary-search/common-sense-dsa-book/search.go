package common_sense_dsa_book

func Search(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result, searchCount = Search(a[:mid], search)
	case a[mid] < search:
		result, searchCount = Search(a[mid+1:], search)
		if result >= 0 {
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}
