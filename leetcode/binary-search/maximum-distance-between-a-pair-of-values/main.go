package maximum_distance_between_a_pair_of_values

import (
	"sort"
)

func maxDistance(nums1 []int, nums2 []int) int {
	ans := 0
	for i, num1 := range nums1 {
		j := sort.Search(len(nums2), func(idx int) bool {
			return nums2[idx] < num1
		})
		if j < len(nums2) {
			j--
		} else {
			j = len(nums2) - 1
		}
		if j > i {
			if j-i > ans {
				ans = j - i
			}
		}
	}

	return ans
}
