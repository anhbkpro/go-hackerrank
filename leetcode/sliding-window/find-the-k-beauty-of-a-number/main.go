package find_the_k_beauty_of_a_number

import "strconv"

func DivisorSubstrings(num int, k int) int {
	numStr := strconv.Itoa(num)
	ans := 0
	n := len(numStr)
	for r := 0; r <= n-k; r++ {
		subString := numStr[r : r+k]
		subInt, _ := strconv.Atoi(subString)
		if subInt != 0 && num%subInt == 0 {
			ans++
		}
	}
	return ans
}
