package container_with_most_water

import "math"

func MaxArea(height []int) int {
	result := 0
	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
		result = int(math.Max(float64(result), float64(area)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return result
}
