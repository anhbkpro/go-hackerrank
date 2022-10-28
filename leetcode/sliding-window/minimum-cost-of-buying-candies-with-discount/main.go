package minimum_cost_of_buying_candies_with_discount

import (
	"fmt"
	"sort"
)

func MinimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] >= cost[j]
	})
	fmt.Println("Sorted array:", cost)
	ans := 0
	for i := 0; i < len(cost); i++ {
		if i%3 == 0 || i%3 == 1 {
			ans += cost[i]
		}
	}
	return ans
}
