package search_insert_position

// SearchInsert My Solution:
// Runtime: 4 ms, faster than 82.89% of Go online submissions for Search Insert Position.
//	Memory Usage: 3 MB, less than 8.61% of Go online submissions for Search Insert Position.
func SearchInsert(nums []int, target int) int {
	mid := len(nums) / 2
	res := 0
	switch {
	case len(nums) == 0:
		res = 0 // not found
	case nums[mid] > target:
		res = SearchInsert(nums[:mid], target)
	case nums[mid] < target:
		res = SearchInsert(nums[mid+1:], target)
		res += mid + 1

	default:
		res = mid
	}
	return res
}
