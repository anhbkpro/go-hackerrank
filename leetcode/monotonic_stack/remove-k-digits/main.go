package remove_k_digits

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []int

func (s *stack) push(item int) {
	*s = append(*s, item)
}

func (s *stack) pop() int {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func RemoveKdigits(num string, k int) string {
	ans := ""
	st := stack{}
	runes := []rune(num)
	for i := 0; i < len(runes); i++ {
		number, _ := strconv.Atoi(string(runes[i]))
		for len(st) != 0 && k > 0 && number < st.top() {
			st.pop()
			k--
		}
		st.push(number)
	}

	for k > 0 {
		st.pop()
		k--
	}
	fmt.Println(st)
	for len(st) != 0 {
		ans = strconv.Itoa(st.pop()) + ans
	}
	for strings.HasPrefix(ans, "0") {
		ans = ans[1:]
	}
	if len(ans) == 0 {
		return "0"
	}

	return ans
}
