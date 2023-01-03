package sequential_digits

import (
	"fmt"
	"strconv"
)

func sequentialDigits(low int, high int) []int {
	sample := "123456789"
	n := 10
	nums := make([]int, 0)

	lowLen := len(fmt.Sprint(low))
	highLen := len(fmt.Sprint(high))
	for length := lowLen; length < highLen+1; length++ {
		for start := 0; start < n-length; start++ {
			num, _ := strconv.Atoi(sample[start : start+length])
			if num >= low && num <= high {
				nums = append(nums, num)
			}
		}
	}

	return nums
}
