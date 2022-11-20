package check_if_n_and_its_double_exist

func checkIfExist(arr []int) bool {
	seen := make(map[int]bool)
	for _, e := range arr {
		if seen[e*2] || (e&1 == 0 && seen[e/2]) { // same as e%2 == 0
			return true
		}
		seen[e] = true
	}
	return false
}
