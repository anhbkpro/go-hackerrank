package sliding_window_maximum

import (
	"github.com/gammazero/deque"
)

func MaxSlidingWindow(nums []int, k int) []int {
	var ans []int
	var q deque.Deque[int] // index
	l, r := 0, 0
	for r < len(nums) {
		// pop smaller values from the queue
		for q.Len() > 0 && nums[q.Back()] < nums[r] {
			q.PopBack()
		}
		q.PushBack(r)
		//fmt.Println("deque: ", q)

		// remove the left most from the queue
		if l > q.Front() {
			q.PopFront()
		}

		if r+1 >= k {
			ans = append(ans, nums[q.Front()])
			l++
		}
		r++
	}

	return ans
}
