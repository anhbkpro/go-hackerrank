package min_cost_climbing_stairs

import "math"

func MinCostClimbingStairs(cost []int32) int32 {
	cost = append(cost, 0)
	for i := len(cost) - 3; i > -1; i-- {
		cost[i] += int32(math.Min(float64(cost[i+1]), float64(cost[i+2])))
	}

	return int32(math.Min(float64(cost[0]), float64(cost[1])))
}
