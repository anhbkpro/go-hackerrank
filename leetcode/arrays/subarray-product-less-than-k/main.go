package subarray_product_less_than_k

func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	prod := 1
	var left int
	var ans int
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}

	return ans
}
