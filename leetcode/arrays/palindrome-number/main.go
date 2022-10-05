package palindrome_number

import "fmt"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	runes := []rune(fmt.Sprint(x))
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}

	return true
}
