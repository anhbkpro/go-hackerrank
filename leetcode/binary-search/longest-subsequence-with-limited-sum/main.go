package longest_subsequence_with_limited_sum

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for _, query := range queries {
		l, r := 0, len(nums)-1
		for l <= r {
			m := (l + r) >> 1
			if nums[m] > query {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		ans = append(ans, l)
	}
	return ans
}
