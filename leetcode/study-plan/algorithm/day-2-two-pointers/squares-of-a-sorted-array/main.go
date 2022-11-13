package squares_of_a_sorted_array

import "math"

func SortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if math.Abs(float64(nums[r])) > math.Abs(float64(nums[l])) {
			res[i] = nums[r] * nums[r]
			r--
		} else {
			res[i] = nums[l] * nums[l]
			l++
		}
	}
	return res
}
