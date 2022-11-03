package final_prices_with_a_special_discount_in_a_shop

type stack []int

func (s *stack) pop() int {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) push(item int) {
	*s = append(*s, item)
}

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func FinalPrices(prices []int) []int {
	st := stack{}
	ans := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		for len(st) != 0 && prices[i] <= prices[st.top()] {
			ans[st.top()] = prices[st.pop()] - prices[i]
		}
		st.push(i)
	}

	for len(st) != 0 && ans[st.top()] == 0 {
		ans[st.top()] = prices[st.pop()]
	}

	return ans
}
