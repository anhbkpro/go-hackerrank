package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // left sorted
			if target < nums[lo] || target > nums[mid] { // search to the right --> we need find lo
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // right sorted
			if target < nums[mid] || target > nums[hi] { // search to the left -> we need find hi
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}
