package shortest_subarray_to_be_removed_to_make_array_sorted

import (
	"fmt"
	"math"
)

func FindLengthOfShortestSubarray(arr []int) int {
	ans := 0
	n := len(arr)
	left, right := 0, n-1
	for left+1 < n && arr[left] <= arr[left+1] {
		left++
	}
	if left == n-1 {
		return 0
	}

	for right > left && arr[right-1] <= arr[right] {
		right--
	}
	fmt.Println("left=", left, ";right=", right)
	ans = int(math.Min(float64(n-left-1), float64(right)))
	fmt.Println("ans", ans)
	i, j := 0, right
	for i <= left && j < n {
		if arr[j] >= arr[i] {
			ans = int(math.Min(float64(ans), float64(j-i-1)))
			fmt.Println("ans, i, j:", ans, i, j)
			i++
		} else {
			j++
		}
	}

	return ans
}
