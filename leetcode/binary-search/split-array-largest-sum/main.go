package split_array_largest_sum

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Split Array Largest Sum.
//Memory Usage: 2.1 MB, less than 56.86% of Go online submissions for Split Array Largest Sum.
func splitArray(nums []int, k int) int {
	l, r := maxSumTuple(nums)
	for l < r {
		m := (l + r) >> 1
		if canSplit(nums, k, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func canSplit(nums []int, k, threshold int) bool {
	count, total := 1, 0
	for _, num := range nums {
		total += num
		if total > threshold {
			count++
			total = num
		}
	}
	return count <= k
}

func maxSumTuple(nums []int) (int, int) {
	max := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if max < num {
			max = num
		}
	}
	return max, sum
}
