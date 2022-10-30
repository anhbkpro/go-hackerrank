package substring_with_concatenation_of_all_words

// FindSubstring https://bohenan.gitbooks.io/leetcode/content/chapter1/substring-with-concatenation-of-all-words.html
func FindSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordsLen := len(words) * wordLen
	ans := make([]int, 0)
	freqWord := mapOf(words)
	for l := 0; l+wordsLen <= len(s); l++ {
		count := make(map[string]int)
		c := 0
		for r := l; r < l+wordsLen; r += wordLen {
			word := s[r : r+wordLen]
			count[word]++
			if freqWord[word] == 0 || count[word] > freqWord[word] {
				break
			}
			c += wordLen
		}
		if c == wordsLen {
			ans = append(ans, l)
		}
	}
	return ans
}

func mapOf(words []string) map[string]int {
	res := make(map[string]int)
	for _, word := range words {
		res[word]++
	}
	return res
}
