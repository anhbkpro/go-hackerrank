package minimum_consecutive_cards_to_pick_up

import "math"

func MinimumCardPickup(cards []int) int {
	l, ans := 0, math.MaxInt
	freqWindow := make(map[int]bool)
	for r := 0; r < len(cards); r++ {
		for freqWindow[cards[r]] {
			ans = int(math.Min(float64(ans), float64(r-l+1)))
			delete(freqWindow, cards[l])
			l++
		}
		freqWindow[cards[r]] = true
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
