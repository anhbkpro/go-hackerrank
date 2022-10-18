package climbing_stairs

func ClimbStairs(n int) int {
	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		temp := one     // 8, 5, 3, 2, 1 (two), 1 (one)
		one = one + two // bottom up
		two = temp
	}
	return one // bottom up
}
