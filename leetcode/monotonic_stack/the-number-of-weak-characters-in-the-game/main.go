package the_number_of_weak_characters_in_the_game

import (
	"fmt"
	"math"
	"sort"
)

func NumberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[j][1] < properties[i][1] // if same group, we sort defense score decreasingly
		}
		return properties[i][0] < properties[j][0]
	})
	fmt.Println(properties)

	count := 0
	max := math.MinInt
	for i := len(properties) - 1; i >= 0; i-- {
		if properties[i][1] < max {
			count++
		}
		max = int(math.Max(float64(max), float64(properties[i][1])))
	}

	return count
}
