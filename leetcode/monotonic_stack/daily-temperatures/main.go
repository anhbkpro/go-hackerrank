package daily_temperatures

// DailyTemperatures https://1e9.medium.com/monotonic-queue-notes-980a019d5793
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	nearest := make([]int, n) // nearest biggest indexes from right to left

	for left := n - 1; left >= 0; left-- {
		right := left + 1
		for right < n && temperatures[right] <= temperatures[left] {
			right = nearest[right]
		}
		nearest[left] = right

		if nearest[left] == n {
			ans[left] = 0
		} else {
			ans[left] = nearest[left] - left
		}
	}
	return ans
}
