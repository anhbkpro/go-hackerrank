package check_if_a_number_is_majority_element_in_a_sorted_array

//Runtime: 3 ms, faster than 76.92% of Go online submissions for Check If a Number Is Majority Element in a Sorted Array.
//Memory Usage: 2.5 MB, less than 42.31% of Go online submissions for Check If a Number Is Majority Element in a Sorted Array.
func isMajorityElement(nums []int, target int) bool {
	first := firstOccur(nums, target)
	second := first + len(nums)/2
	if second < len(nums) && nums[second] == target {
		return true
	}
	return false
}

func firstOccur(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) >> 1
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
