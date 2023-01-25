package number_of_different_positives

func minimumOperations(nums []int) int {
	count := make(map[int]bool)
	for _, v := range nums {
		if count[v] || v <= 0 {
			continue
		}
		count[v] = true
	}

	return len(count)
}
