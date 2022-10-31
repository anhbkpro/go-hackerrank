package daily_temperatures

type stack []int

func (s *stack) Push(element int) {
	*s = append(*s, element)
}

func (s *stack) Pop() int {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

// DailyTemperaturesStack https://leetcode.com/problems/daily-temperatures/discuss/1574806/C%2B%2B-EASY-standard-sol-oror-INTUITIVE-approach-with-DRY-RUN-oror-STACK-appraoch
func DailyTemperaturesStack(temperatures []int) []int {
	// we will simply pop the top of the stack out until we find an element greater than the current index
	// or until the stack gets empty (This, means no greater element on the right).
	// ==> stack contains non-increasing temperature indexes
	n := len(temperatures)
	stack := stack{} // indexes
	warmerTemperatureDays := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		// pop until we find next greater element to the right
		for len(stack) != 0 && temperatures[stack.Peek()] <= temperatures[i] {
			stack.Pop()
		}
		// if the stack is not empty, we found the next greater element
		// else can not find any element larger than the current temperature (default day is 0)
		if len(stack) != 0 {
			warmerTemperatureDays[i] = stack.Peek() - i
		}

		// push the index of current temperature in the stack,
		// same as pushing current temperature in stack
		stack.Push(i)
	}

	return warmerTemperatureDays
}
