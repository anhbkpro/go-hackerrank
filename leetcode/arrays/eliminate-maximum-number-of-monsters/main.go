package eliminate_maximum_number_of_monsters

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	for i := 0; i < len(dist); i++ {
		dist[i] = (dist[i] - 1) / speed[i]
	}

	sort.Slice(dist, func(i, j int) bool {
		return dist[i] < dist[j]
	})

	for i := 0; i < len(dist); i++ {
		if i > dist[i] {
			return i
		}
	}

	return len(dist)
}
