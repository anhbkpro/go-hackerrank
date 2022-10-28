package number_of_sub_arrays_of_size_k_and_average_greater_than_or_equal_to_threshold

func NumOfSubarrays(arr []int, k int, threshold int) int {
	l, windowSum, ans := 0, 0, 0
	for r := 0; r < len(arr); r++ {
		windowSum += arr[r]
		if r-l+1 == k {
			if windowSum/k >= threshold {
				ans++
			}
			windowSum -= arr[l]
			l++
		}
	}
	return ans
}
