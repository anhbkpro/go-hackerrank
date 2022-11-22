package intersection_of_three_sorted_arrays

import "sort"

//Runtime: 8 ms, faster than 85.00% of Go online submissions for Intersection of Three Sorted Arrays.
//Memory Usage: 6.8 MB, less than 5.00% of Go online submissions for Intersection of Three Sorted Arrays.
func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	m := make(map[int]int)
	for _, item := range arr1 {
		m[item]++
	}
	for _, item := range arr2 {
		m[item]++
	}
	for _, item := range arr3 {
		m[item]++
	}

	ans := make([]int, 0)
	for k, v := range m {
		if v == 3 {
			ans = append(ans, k)
		}
	}

	// O(nlog(n))
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	return ans
}

