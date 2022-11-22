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

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Intersection of Three Sorted Arrays.
//Memory Usage: 4.1 MB, less than 85.00% of Go online submissions for Intersection of Three Sorted Arrays.
func arraysIntersection3Pointers(arr1 []int, arr2 []int, arr3 []int) []int {
	p1, p2, p3 := 0, 0, 0
	ans := make([]int, 0)
	for p1 < len(arr1) && p2 < len(arr2) && p3 < len(arr3) {
		if arr1[p1] == arr2[p2] && arr2[p2] == arr3[p3] {
			ans = append(ans, arr1[p1])
			p1++
			p2++
			p3++
		} else {
			if arr1[p1] < arr2[p2] {
				p1++
			} else if arr2[p2] < arr3[p3] {
				p2++
			} else {
				p3++
			}
		}
	}
	return ans
}
