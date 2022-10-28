package two_sum

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if _, ok := seen[target-num]; ok {
			return []int{seen[target-num], i}
		} else {
			seen[num] = i
		}
	}
	return []int{}
}
