package kth_missing_positive_number

import "fmt"

func findKthPositive(arr []int, k int) int {
	m := make(map[int]bool)
	for _, item := range arr {
		m[item] = true
	}
	fmt.Println(m)
	missingArr := make([]int, 0)
	for i := 1; i <= arr[len(arr)-1]+k+1; i++ {
		if !m[i] {
			missingArr = append(missingArr, i)
		}
	}
	fmt.Println(missingArr)
	lo, hi := 0, len(missingArr)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if mid == k {
			return missingArr[mid-1]
		} else if mid > k {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return 0
}

func findKthPositiveDiscuss(arr []int, k int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		m := (lo + hi) >> 1
		if arr[m]-1-m < k {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return lo + k
}
