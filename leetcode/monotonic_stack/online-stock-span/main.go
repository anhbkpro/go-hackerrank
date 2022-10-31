package online_stock_span

type Item struct {
	price int
	ans   int // the answer for old calls
}

type stack []Item

func (s *stack) Push(element Item) {
	*s = append(*s, element)
}

func (s *stack) Pop() Item {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) Peek() Item {
	return (*s)[len(*s)-1]
}

type StockSpanner struct {
	s stack // a monotonic decreasing manner
}

func Constructor() StockSpanner {
	return StockSpanner{
		s: stack{},
	}
}

func (s *StockSpanner) Next(price int) int {
	ans := 1
	// While the top of the stack has a value (stock price) less than or equal to price,
	// it should be included in our answer, so pop it off the stack.
	for len(s.s) != 0 && s.s.Peek().price <= price {
		ans += s.s.Peek().ans
		s.s.Pop()
	}

	s.s.Push(Item{price: price, ans: ans}) // "remember" the answer for old calls
	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
