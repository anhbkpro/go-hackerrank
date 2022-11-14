package peak_index_in_a_mountain_array

func PeakIndexInMountainArray(arr []int) int {
	lo, hi := 1, len(arr)-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] < arr[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
