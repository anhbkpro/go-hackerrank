package two_sum

func TwoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := prevMap[target-nums[i]]; !ok {
			prevMap[nums[i]] = i
		} else {
			return []int{prevMap[target-nums[i]], i}
		}
	}
	return nil
}
