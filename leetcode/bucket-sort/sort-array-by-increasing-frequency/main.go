package sort_array_by_increasing_frequency

import (
	"sort"
)

// FrequencySort My Solution
//	Runtime: 8 ms, faster than 78.72% of Go online submissions for Sort Array by Increasing Frequency.
//	Memory Usage: 3.9 MB, less than 8.51% of Go online submissions for Sort Array by Increasing Frequency.
func FrequencySort(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	ans := make([]int, 0)
	bucket := make([][]int, len(nums)+1)
	for k, v := range freq {
		frequencyElements := bucket[v]
		frequencyElements = append(frequencyElements, k)
		sort.Slice(frequencyElements, func(i, j int) bool {
			return frequencyElements[i] > frequencyElements[j]
		})
		bucket[v] = frequencyElements
	}

	for i := 0; i <= len(bucket)-1; i++ {
		if len(bucket[i]) > 0 {
			for j := 0; j < len(bucket[i]); j++ {
				temp := i
				for temp > 0 {
					ans = append(ans, bucket[i][j])
					temp--
				}
			}
		}
	}
	return ans
}
