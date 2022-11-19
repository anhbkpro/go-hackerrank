package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] > nums[hi] { // smaller values must be in the right sorted portion
			lo = mid + 1
		} else { // min value must be in the left sorted portion
			hi = mid
		}
	}
	return nums[lo]
}
