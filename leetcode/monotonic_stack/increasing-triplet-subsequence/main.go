package increasing_triplet_subsequence

import (
	"fmt"
	"math"
)

func IncreasingTriplet(nums []int) bool {
	fmt.Println("------")
	a, b := math.MaxInt, math.MaxInt
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num <= a {
			a = num
			fmt.Println("a", a)
		} else if num <= b {
			b = num
			fmt.Println("b", b)
		} else {
			fmt.Println("Found greater:", num)
			return true
		}
	}
	return false
}
