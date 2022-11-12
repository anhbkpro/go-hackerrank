package binary_search

func Search(nums []int, target int) int {
	mid := len(nums) / 2
	res := 0
	switch {
	case len(nums) == 0:
		res = -1 // not found
	case nums[mid] > target:
		res = Search(nums[:mid], target)
	case nums[mid] < target:
		res = Search(nums[mid+1:], target)
		if res >= 0 {
			res += mid + 1
		}
	default: // nums[mid] == search
		res = mid // found
	}
	return res
}
