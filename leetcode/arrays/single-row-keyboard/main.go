package single_row_keyboard

func calculateTime(keyboard string, word string) int {
	keyIndices := make(map[rune]int, 26)
	for i, c := range []rune(keyboard) {
		keyIndices[c-'a'] = i // actually `keyIndices[c] = i` is enough
	}

	prev := 0
	result := 0
	for _, c := range []rune(word) {
		result += abs(keyIndices[c-'a'], prev)
		prev = keyIndices[c-'a']
	}
	return result
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
