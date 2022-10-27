package fruit_into_baskets

import "math"

func TotalFruit(fruits []int) int {
	l, ans := 0, 0
	freqFruit := make(map[int]int)
	for r := 0; r < len(fruits); r++ {
		fruitType := fruits[r]
		if _, ok := freqFruit[fruitType]; ok {
			freqFruit[fruitType]++
		} else {
			freqFruit[fruitType] = 1
		}

		if len(freqFruit) > 2 {
			leftFruitType := fruits[l]
			freqFruit[leftFruitType]--
			if freqFruit[leftFruitType] == 0 {
				delete(freqFruit, leftFruitType)
			}
			l++
		}
		ans = int(math.Max(float64(ans), float64(r-l+1)))
	}

	return ans
}
