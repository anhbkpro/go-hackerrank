package rotate_array

func Rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}
	temp := make([]int, len(nums))
	for i := 0; i < n; i++ {
		temp[i] = nums[i]
	}

	for i := 0; i < n; i++ {
		nums[(i+k)%n] = temp[i]
	}
}
