package frequency_of_the_most_frequent_element

import "sort"

func MaxFrequency(nums []int, k int) int {
	// sort the input array => max = last_item
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res, l, r, total := 1, 0, 0, 0
	for r = 0; r < len(nums); r++ {
		total += nums[r]

		// the key is to find out the valid condition: total + k < max*size = last_item * (r-l+1)
		for total+k < nums[r]*(r-l+1) {
			total -= nums[l]
			l++
		}
		
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
