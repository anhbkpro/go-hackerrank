package frequency_of_the_most_frequent_element

import "sort"

func MaxFrequency(nums []int, k int) int {
	// sort the input array => max = last_item
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res, i, j, sum := 1, 0, 0, 0
	for j = 0; j < len(nums); j++ {
		sum += nums[j]
		// the key is to find out the valid condition: sum + k < max*size = last_item * (j-i+1)
		for sum+k < nums[j]*(j-i+1) {
			sum -= nums[i]
			i++
		}
		if j-i+1 > res {
			res = j - i + 1
		}
	}
	return res
}
