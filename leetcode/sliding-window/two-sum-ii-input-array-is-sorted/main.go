package two_sum_ii_input_array_is_sorted

func TwoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	sum := numbers[l] + numbers[r]
	for sum != target {
		if sum > target {
			r--
		} else {
			l++
		}
		sum = numbers[l] + numbers[r]
	}

	return []int{l + 1, r + 1}
}
