package minimum_deletions_to_make_character_frequencies_unique

// MinDeletions
// My Solution:
// 	Runtime: 206 ms, faster than 8.00% of Go online submissions for Minimum Deletions to Make Character Frequencies Unique.
// 	Memory Usage: 7.2 MB, less than 28.00% of Go online submissions for Minimum Deletions to Make Character Frequencies Unique.
func MinDeletions(s string) int {
	freq := make(map[rune]int)
	runes := []rune(s)
	for _, r := range runes {
		freq[r]++
	}

	seenFreq := make(map[int]bool)
	seenChar := make(map[rune]bool)
	seenFreq[freq[runes[0]]] = true
	seenChar[runes[0]] = true
	toBeDeleted := 0
	for i := 0; i < len(runes); i++ {
		if seenChar[runes[i]] { // if we see this char before
			continue
		}
		// if not
		seenChar[runes[i]] = true // set seen
		for seenFreq[freq[runes[i]]] {
			freq[runes[i]]--
			toBeDeleted++
		}
		if freq[runes[i]] == 0 {
			delete(freq, runes[i])
		} else {
			seenFreq[freq[runes[i]]] = true
		}
	}
	return toBeDeleted
}

func MinDeletionsDiscuss(s string) int {
	cnt := make([]int, 26)
	res := 0
	used := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-97]++
	}
	for i := 0; i < 26; i++ {
		for cnt[i] > 0 && used[cnt[i]] {
			cnt[i]-- // delete characters until we find unused count
			res++
		}
		used[cnt[i]] = true
	}
	return res
}
