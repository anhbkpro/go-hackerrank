package fruit_into_baskets

import "math"

func TotalFruit(fruits []int) int {
	l, r, maxFruits := 0, 0, 0
	typeFruits := make(map[int]int)
	for r < len(fruits) {
		fruitType := fruits[r]
		if _, ok := typeFruits[fruitType]; ok {
			typeFruits[fruitType]++
		} else {
			typeFruits[fruitType] = 1
		}

		if len(typeFruits) <= 2 {
			maxFruits = int(math.Max(float64(maxFruits), float64(r-l+1)))
			r++
		} else {
			leftFruitType := fruits[l]
			typeFruits[leftFruitType]--
			if typeFruits[leftFruitType] == 0 {
				delete(typeFruits, leftFruitType)
			}
			r++
			l++
		}
	}
	return maxFruits
}
