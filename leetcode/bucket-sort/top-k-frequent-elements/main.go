package top_k_frequent_elements

import "fmt"

// TopKFrequent My Solution (refer Drawing Explanation of https://www.youtube.com/watch?v=YPTqKIgVk-k):
//	Runtime: 11 ms, faster than 95.07% of Go online submissions for Top K Frequent Elements.
//	Memory Usage: 6.6 MB, less than 8.55% of Go online submissions for Top K Frequent Elements.
func TopKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums)+1)
	freq := make(map[int]int) // map[frequent][]{numbers}
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for k, v := range freq {
		sameFrequentElements := bucket[v]
		sameFrequentElements = append(sameFrequentElements, k)
		bucket[v] = sameFrequentElements
	}

	fmt.Println(bucket)

	ans := make([]int, 0)
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 && k > 0 {
			capacity := k
			if capacity > len(bucket[i]) {
				capacity = len(bucket[i])
			}
			ans = append(ans, bucket[i][:capacity]...)
			k -= capacity
		}
	}
	return ans
}
