package combination_sum

func CombinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var cur []int
	var total int
	dfs(&ans, candidates, cur, 0, total, target)
	return ans
}

func dfs(ans *[][]int, candidates, cur []int, i, total, target int) {
	if total == target {
		dst := make([]int, len(cur))
		copy(dst, cur)
		*ans = append(*ans, dst)
		return
	}

	if i >= len(candidates) || total > target {
		return
	}

	cur = append(cur, candidates[i])
	dfs(ans, candidates, cur, i, total+candidates[i], target)
	cur = cur[:len(cur)-1]
	dfs(ans, candidates, cur, i+1, total, target)
}
