package top_k_frequent_words

import (
	"sort"
)

// TopKFrequent My Solution
//	Runtime: 4 ms, faster than 96.00% of Go online submissions for Top K Frequent Words.
//	Memory Usage: 4.9 MB, less than 9.71% of Go online submissions for Top K Frequent Words.
func TopKFrequent(words []string, k int) []string {
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	bucket := make([][]string, len(words)+1)
	for w, f := range freq {
		sameFrequencyWords := bucket[f]
		sameFrequencyWords = append(sameFrequencyWords, w)
		sort.Slice(sameFrequencyWords, func(i, j int) bool {
			return sameFrequencyWords[i] < sameFrequencyWords[j]
		})
		bucket[f] = sameFrequencyWords
	}
	ans := make([]string, 0)
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
