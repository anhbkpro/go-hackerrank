package find_all_lonely_numbers_in_the_array

func FindLonely(nums []int) []int {
	ans := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		if v1, v2, v3 := m[nums[i]], m[nums[i]-1], m[nums[i]+1]; v1 == 1 && v2+v3 == 0 {
			ans = append(ans, nums[i])
		}
	}
	return ans
}
