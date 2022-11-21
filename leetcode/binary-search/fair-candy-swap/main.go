package fair_candy_swap

//Runtime: 173 ms, faster than 42.10% of Go online submissions for Fair Candy Swap.
//Memory Usage: 9 MB, less than 5.26% of Go online submissions for Fair Candy Swap.
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aM, bM := make(map[int]bool), make(map[int]bool)

	totalA := 0
	totalB := 0
	for _, size := range aliceSizes {
		totalA += size
		aM[size] = true
	}
	for _, size := range bobSizes {
		totalB += size
		bM[size] = true
	}
	missing := (totalA - totalB) / 2
	for _, size := range bobSizes {
		if aM[size+missing] {
			return []int{size + missing, size}
		}
	}

	return nil
}
