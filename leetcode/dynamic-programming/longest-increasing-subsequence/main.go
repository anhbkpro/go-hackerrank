package longest_increasing_subsequence

import "math"

func LengthOfLIS(nums []int) int {
	LIS := make([]int, len(nums))
	for i := range LIS {
		LIS[i] = 1
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				LIS[i] = int(math.Max(float64(LIS[i]), 1+float64(LIS[j])))
			}
		}
	}

	var len int
	for _, val := range LIS {
		if len < val {
			len = val
		}
	}

	return len
}
