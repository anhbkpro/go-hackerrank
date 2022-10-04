package subarray_sum_equals_k

func SubarraySum(nums []int, k int) int {
	var ans int
	var sum int
	m := make(map[int]int)
	m[0]++
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		ans += m[sum-k]
		m[sum]++
	}
	return ans
}
