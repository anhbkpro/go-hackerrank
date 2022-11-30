package longest_increasing_subsequence

//Runtime: 8 ms, faster than 93.37% of Go online submissions for Longest Increasing Subsequence.
//Memory Usage: 3.6 MB, less than 95.15% of Go online submissions for Longest Increasing Subsequence.
func lengthOfLIS(nums []int) int {
	sub := make([]int, 0)
	sub = append(sub, nums[0])

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > sub[len(sub)-1] {
			sub = append(sub, num)
		} else {
			j := binarySearch(sub, num)
			sub[j] = num
		}
	}

	return len(sub)
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := (l + r) >> 1
		if arr[m] == target {
			return m
		}
		if arr[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
