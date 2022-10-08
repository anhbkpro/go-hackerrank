package letter_combinations_of_a_phone_number

import (
	"strconv"
)

func LetterCombinations(digits string) []string {
	ans := make([]string, 0)
	digitToChar := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	if digits != "" {
		backtrack(&ans, 0, digits, "", digitToChar)
	}
	return ans
}

func backtrack(ans *[]string, i int, digits, curStr string, digitToChar map[int]string) {
	if len(curStr) == len(digits) {
		*ans = append(*ans, curStr)
		return
	}
	runes := []rune(digits)
	digitStr := string(runes[i])
	digit, _ := strconv.Atoi(digitStr)
	for _, r := range digitToChar[digit] {
		backtrack(ans, i+1, digits, curStr+string(r), digitToChar)
	}
}
