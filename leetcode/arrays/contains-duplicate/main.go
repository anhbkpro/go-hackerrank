package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = struct{}{}
		}
	}
	return false
}
