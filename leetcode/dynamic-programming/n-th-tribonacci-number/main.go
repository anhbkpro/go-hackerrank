package n_th_tribonacci_number

func Tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	one, two, three := 0, 1, 1

	for i := 3; i <= n; i++ {
		temp := one + two + three // top-down
		one = two
		two = three
		three = temp
	}
	return three // top-down
}
