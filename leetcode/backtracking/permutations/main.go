package permutations

func Permute(nums []int) [][]int {
	var ans [][]int
	if len(nums) == 1 {
		dst := make([]int, len(nums))
		copy(dst, nums)
		return [][]int{dst}
	}

	for i := 0; i < len(nums); i++ {
		n := nums[0]
		nums = nums[1:]
		perms := Permute(nums)
		for _, perm := range perms {
			perm = append(perm, n)
			ans = append(ans, perm)
		}
		nums = append(nums, n)
	}

	return ans
}
