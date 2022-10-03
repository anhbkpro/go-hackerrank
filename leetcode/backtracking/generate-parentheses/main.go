package generate_parentheses

import "strings"

func GenerateParenthesis(n int) []string {
	var stack []string
	var res []string

	backtrack(0, 0, n, &stack, &res)

	return res
}

func backtrack(open, closed, n int, stack, res *[]string) {
	if open == closed && closed == n {
		*res = append(*res, strings.Join(*stack, ""))
		return
	}
	if open < n {
		*stack = append(*stack, "(")
		backtrack(open+1, closed, n, stack, res)
		*stack = (*stack)[0 : len(*stack)-1]
	}

	if closed < open {
		*stack = append(*stack, ")")
		backtrack(open, closed+1, n, stack, res)
		*stack = (*stack)[0 : len(*stack)-1]
	}
}
