package find_target_indices_after_sorting_array

import (
	"fmt"
	"sort"
)

func findLeftIndex(nums []int, target int) int {
	l, r := 0, len(nums)-1
	lb := len(nums)
	for l <= r {
		m := (l + r) >> 1
		if nums[m] == target {
			lb = m
			r = m - 1
		}
		if nums[m] >= target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return lb
}

func findRightIndex(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	rb := -1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			rb = mid
			lo = mid + 1
		}
		if nums[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return rb
}

func targetIndices(nums []int, target int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	leftIndex := findLeftIndex(nums, target)
	rightIndex := findRightIndex(nums, target)
	fmt.Println(leftIndex, rightIndex)
	ans := make([]int, 0)
	for leftIndex <= rightIndex {
		ans = append(ans, leftIndex)
		leftIndex++
	}
	return ans
}
