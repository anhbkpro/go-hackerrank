package sender_with_largest_word_count

import (
	"fmt"
	"strings"
)

func LargestWordCount(messages []string, senders []string) string {
	freqWords := make(map[string]int)
	totalWords := 0
	for i, sender := range senders {
		wordCount := len(strings.Split(messages[i], " "))
		freqWords[sender] += wordCount
		totalWords += wordCount
	}
	fmt.Println(freqWords)

	bucket := make([][]string, totalWords+1)
	for u, f := range freqWords {
		userSameFrequency := bucket[f]
		userSameFrequency = append(userSameFrequency, u)
		bucket[f] = userSameFrequency
	}

	fmt.Println(bucket)

	var user string
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			user = bucket[i][0]
			for _, res := range bucket[i] {
				if res > user {
					user = res
				}
			}
			return user
		}
	}

	return ""
}

func LargestWordCountDiscuss(messages []string, senders []string) string {
	freqWords := make(map[string]int)
	maxCount := 0
	var res string
	for i := 0; i < len(messages); i++ {
		count := len(strings.Split(messages[i], " "))
		freqWords[senders[i]] += count
		total := freqWords[senders[i]]
		if total > maxCount || (total == maxCount && senders[i] > res) {
			maxCount = total
			res = senders[i]
		}
	}
	return res
}
