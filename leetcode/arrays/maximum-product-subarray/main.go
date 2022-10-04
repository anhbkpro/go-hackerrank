package maximum_product_subarray

import (
	"math"
)

func MaxProduct(nums []int) int {
	r := nums[0]
	iMax := r
	iMin := r
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			temP := iMin
			iMin = iMax
			iMax = temP
		}

		iMax = int(math.Max(float64(nums[i]), float64(iMax*nums[i])))
		iMin = int(math.Min(float64(nums[i]), float64(iMin*nums[i])))

		r = int(math.Max(float64(r), float64(iMax)))
	}

	return r
}
