package house_robber

import "math"

func Rob(nums []int) int {
	rob1, rob2 := 0, 0
	for _, num := range nums {
		temp := int(math.Max(float64(num+rob1), float64(rob2)))
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}
