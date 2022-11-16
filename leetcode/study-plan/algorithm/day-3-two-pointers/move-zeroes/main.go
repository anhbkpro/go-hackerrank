package move_zeroes

func MoveZeroes(nums []int) {
	insertIndex := 0
	for _, num := range nums {
		if num != 0 {
			nums[insertIndex] = num
			insertIndex++
		}
	}
	for insertIndex < len(nums) {
		nums[insertIndex] = 0
		insertIndex++
	}
}
