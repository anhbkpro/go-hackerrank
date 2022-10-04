package subsets

func Subsets(nums []int) [][]int {
	var res [][]int
	var subset []int

	dfs(0, subset, &nums, &res)

	return res
}

func dfs(i int, subset []int, nums *[]int, res *[][]int) {
	if i >= len(*nums) {
		dst := make([]int, len(subset))
		copy(dst, subset)
		*res = append(*res, dst)
		return
	}

	// decision to include nums[i]
	subset = append(subset, (*nums)[i])
	dfs(i+1, subset, nums, res)

	// decision NOT to include nums[i]
	subset = (subset)[:len(subset)-1]
	dfs(i+1, subset, nums, res)
}
