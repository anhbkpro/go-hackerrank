package find_first_and_last_position_of_element_in_sorted_array

import "fmt"

//Runtime: 11 ms, faster than 71.83% of Go online submissions for Find First and Last Position of Element in Sorted Array.
//Memory Usage: 4 MB, less than 14.67% of Go online submissions for Find First and Last Position of Element in Sorted Array.
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	firstIndex := firstOccur(nums, target)
	fmt.Println(firstIndex)
	if firstIndex == -1 {
		return []int{-1, -1}
	}
	secondIndex := secondOccur(nums, target)

	return []int{firstIndex, secondIndex}
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
		if l >= len(nums) || (l == r && nums[l] != target) {
			return -1
		}
	}
	return l
}

func secondOccur(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) >> 1
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
