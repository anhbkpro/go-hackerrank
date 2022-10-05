package subsets

func Subsets(nums []int) [][]int {
	var res [][]int
	var subset []int

	dfs(&nums, &res, 0, subset)

	return res
}

func dfs(nums *[]int, res *[][]int, i int, subset []int) {
	if i >= len(*nums) {
		dst := make([]int, len(subset))
		copy(dst, subset)
		*res = append(*res, dst)
		return
	}

	// decision to include nums[i]
	subset = append(subset, (*nums)[i])
	dfs(nums, res, i+1, subset)

	// decision NOT to include nums[i]
	subset = (subset)[:len(subset)-1]
	dfs(nums, res, i+1, subset)
}
