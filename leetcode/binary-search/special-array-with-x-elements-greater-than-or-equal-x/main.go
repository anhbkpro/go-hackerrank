package special_array_with_x_elements_greater_than_or_equal_x

func specialArray(nums []int) int {
	for i := 0; i <= len(nums); i++ {
		count := 0
		for _, num := range nums {
			if num >= i {
				count++
			}
		}
		if count == i {
			return i
		}
	}
	return -1
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Special Array With X Elements Greater Than or Equal X.
//Memory Usage: 2.2 MB, less than 26.47% of Go online submissions for Special Array With X Elements Greater Than or Equal X.
func specialArrayBs(nums []int) int {
	l, r := 0, len(nums)
	for l <= r {
		mid := (l + r) >> 1
		curr := count(nums, mid)
		if curr == mid {
			return mid
			//If the count of >=mid is less than the mid
		} else if curr < mid {
			//decrease it;
			r = mid - 1
		} else {
			//Increase it;as the value count>mid and we need to reduce the count of values ">=" the mid
			l = mid + 1
		}
	}
	return -1
}

func count(nums []int, t int) int {
	c := 0
	for _, num := range nums {
		if num >= t {
			c++
		}
	}
	return c
}
