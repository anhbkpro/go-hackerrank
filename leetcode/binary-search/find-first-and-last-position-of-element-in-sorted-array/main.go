package find_first_and_last_position_of_element_in_sorted_array

//Runtime: 12 ms, faster than 68.52% of Go online submissions for Find First and Last Position of Element in Sorted Array.
//Memory Usage: 4 MB, less than 14.67% of Go online submissions for Find First and Last Position of Element in Sorted Array.
func searchRange(nums []int, target int) []int {
	firstOccurrence := findBound(nums, target, true)
	if firstOccurrence == -1 {
		return []int{-1, -1}
	}

	lastOccurrence := findBound(nums, target, false)

	return []int{firstOccurrence, lastOccurrence}
}

func findBound(nums []int, target int, isFirst bool) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] == target {
			if isFirst {
				// This means we found our lower bound.
				if m == l || nums[m-1] != target {
					return m
				}
				// Search on the left side for the bound.
				r = m - 1
			} else {
				// This means we found our upper bound.
				if m == r || nums[m+1] != target {
					return m
				}
				// Search on the right side for the bound.
				l = m + 1
			}
		} else if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		}
	}
	return -1
}
