package find_peak_element

// FindPeakElement My Solution
//	Runtime: 7 ms, faster than 40.21% of Go online submissions for Find Peak Element.
//	Memory Usage: 2.8 MB, less than 70.98% of Go online submissions for Find Peak Element.
func FindPeakElement(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] < nums[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
