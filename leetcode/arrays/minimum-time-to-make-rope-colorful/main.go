package minimum_time_to_make_rope_colorful

import "math"

func MinCost(colors string, neededTime []int) int {
	var ans int
	var left int
	var windowSum int
	max := neededTime[0]
	prevColor := colors[0]
	for right := 0; right < len(colors); right++ {
		currentColor := colors[right]
		windowSum += neededTime[right]
		for currentColor != prevColor {
			if right-left > 1 {
				ans += windowSum - neededTime[right] - max
			}

			prevColor = currentColor
			left = right
			max = neededTime[left]
			windowSum = neededTime[left]
		}

		if right == len(colors)-1 && right > left {
			max = int(math.Max(float64(max), float64(neededTime[right])))
			ans += windowSum - max
			break
		}

		if max < neededTime[right] {
			max = neededTime[right]
		}
	}

	return ans
}
