package single_row_keyboard

func calculateTime(keyboard string, word string) int {
	idx := make(map[rune]int)
	for i, r := range []rune(keyboard) {
		idx[r] = i
	}

	time := 0
	prev := 0
	for _, r := range []rune(word) {
		time += abs(idx[r], prev)
		prev = idx[r]
	}
	return time
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
