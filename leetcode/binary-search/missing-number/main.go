package missing_number

import "sort"

func MissingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] > mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
