package first_unique_character_in_a_string

import (
	"math"
)

func FirstUniqChar(s string) int {
	ans := math.MaxInt
	seen := make(map[rune]int)
	for _, r := range []rune(s) {
		seen[r]++
	}

	for i := 0; i < len(s); i++ {
		if seen[rune(s[i])] == 1 {
			return i
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
