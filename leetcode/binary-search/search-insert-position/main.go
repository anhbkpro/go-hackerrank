package search_insert_position

//Runtime: 5 ms, faster than 78.97% of Go online submissions for Search Insert Position.
//Memory Usage: 2.9 MB, less than 77.59% of Go online submissions for Search Insert Position.
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) >> 1
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
