package sum_of_subarray_minimums

const MOD = 1000000007

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

func SumSubarrayMins(arr []int) int {
	sumOfMinimums := 0
	n := len(arr)
	st := stack{}
	for i := 0; i <= n; i++ {
		for len(st) != 0 && (i == n || arr[st.top()] >= arr[i]) {
			mid := st.pop()
			leftBoundary := 0
			rightBoundary := i
			if len(st) == 0 {
				leftBoundary = -1
			} else {
				leftBoundary = st.top()
			}
			count := (rightBoundary - mid) * (mid - leftBoundary) % MOD
			sumOfMinimums += (count * arr[mid]) % MOD
			sumOfMinimums %= MOD
		}
		st.push(i)
	}
	return sumOfMinimums
}
