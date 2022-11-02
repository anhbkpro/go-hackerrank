package best_time_to_buy_and_sell_stock_ii

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

func (s *stack) clear() {
	*s = make([]int, 0)
}

func MaxProfit(prices []int) int {
	st := stack{}
	profit := 0
	for i := 0; i < len(prices); i++ {
		for len(st) != 0 && prices[i] > st.top() {
			profit += prices[i] - st.pop()
		}
		st.clear()
		st.push(prices[i])
	}

	return profit
}
