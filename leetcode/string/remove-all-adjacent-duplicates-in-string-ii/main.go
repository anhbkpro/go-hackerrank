package remove_all_adjacent_duplicates_in_string_ii

import "strings"

type CharFrequency struct {
	c string
	f int
}

func RemoveDuplicates(s string, k int) string {
	var stack []CharFrequency
	for _, v := range s {
		curStr := string(v)
		if stack != nil && len(stack) > 0 && stack[len(stack)-1].c == curStr {
			stack[len(stack)-1].f++
		} else {
			stack = append(stack, CharFrequency{
				c: curStr,
				f: 1,
			})
		}

		if stack[len(stack)-1].f == k {
			stack = stack[:len(stack)-1] // pop item from stack
		}
	}
	var ans string
	for _, cf := range stack {
		ans += strings.Repeat(cf.c, cf.f)
	}
	return ans
}
