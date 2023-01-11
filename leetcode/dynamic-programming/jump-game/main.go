package jump_game

import "math"

func CanJump(nums []int) bool {
	var maxJump int
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] >= len(nums)-1 {
			return true
		}
		min := int(math.Min(float64(i+nums[i]), float64(len(nums)-1)))
		if maxJump < min {
			maxJump = min
		}
		if nums[i] == 0 && i >= maxJump {
			return false
		}
	}
	return false
}

func CanJumpGreedy(nums []int) bool {
	goal := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	if goal > 0 {
		return false
	}

	return true
}
