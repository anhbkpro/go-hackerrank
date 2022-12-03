package container_with_most_water

func MaxArea(height []int) int {
	result := 0
	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		result = max(result, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
