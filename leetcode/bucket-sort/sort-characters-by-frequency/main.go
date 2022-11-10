package sort_characters_by_frequency

import "strings"

// FrequencySort My Solution
// 	Runtime: 11 ms, faster than 69.35% of Go online submissions for Sort Characters By Frequency.
//	Memory Usage: 5.3 MB, less than 58.06% of Go online submissions for Sort Characters By Frequency.
func FrequencySort(s string) string {
	freq := make(map[rune]int)
	runes := []rune(s)
	for _, r := range runes {
		freq[r]++
	}
	bucket := make([][]rune, len(runes)+1)
	for k, v := range freq {
		sameFreqChars := bucket[v]
		sameFreqChars = append(sameFreqChars, k)
		bucket[v] = sameFreqChars
	}

	var ans strings.Builder
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			for _, char := range bucket[i] {
				ans.WriteString(strings.Repeat(string(char), i))
			}
		}
	}

	return ans.String()
}
