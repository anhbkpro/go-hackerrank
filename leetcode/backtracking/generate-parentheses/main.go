package generate_parentheses

import "strings"

func GenerateParenthesis(n int) []string {
	var stack []string
	var res []string

	backtrack(&stack, &res, 0, 0, n)

	return res
}

func backtrack(stack, res *[]string, open, closed, n int) {
	if open == closed && closed == n {
		*res = append(*res, strings.Join(*stack, ""))
		return
	}
	if open < n {
		*stack = append(*stack, "(")
		backtrack(stack, res, open+1, closed, n)
		*stack = (*stack)[0 : len(*stack)-1]
	}

	if closed < open {
		*stack = append(*stack, ")")
		backtrack(stack, res, open, closed+1, n)
		*stack = (*stack)[0 : len(*stack)-1]
	}
}
