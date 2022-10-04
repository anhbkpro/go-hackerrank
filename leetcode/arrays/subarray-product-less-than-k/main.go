package subarray_product_less_than_k

func NumSubarrayProductLessThanK(nums []int, k int) int {
	prod := 1
	var windowStart int
	var res int
	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		prod *= nums[windowEnd]
		for prod >= k {
			prod /= nums[windowStart]
			windowStart++
		}
		res += windowEnd - windowStart + 1
	}

	return res
}
