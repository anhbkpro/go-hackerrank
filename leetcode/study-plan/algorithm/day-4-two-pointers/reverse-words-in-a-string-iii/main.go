package reverse_words_in_a_string_iii

import "strings"

func ReverseWords(s string) string {
	splitted := strings.Split(s, " ")
	var ans []string
	for _, split := range splitted {
		runes := []rune(split)
		reverse := make([]rune, len(runes))
		for i := 0; i < len(runes); i++ {
			reverse[i] = runes[len(runes)-1-i]
		}
		ans = append(ans, string(reverse))
	}
	return strings.Join(ans, " ")
}
