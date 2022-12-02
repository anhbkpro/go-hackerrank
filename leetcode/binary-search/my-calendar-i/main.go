package my_calendar_i

type event struct {
	start, end int
	next       *event
}

type MyCalendar struct {
	head *event
}

func Constructor() MyCalendar {
	var head *event
	return MyCalendar{head: head}
}

// Book Runtime: 153 ms, faster than 50.00% of Go online submissions for My Calendar I.
//Memory Usage: 7.6 MB, less than 65.25% of Go online submissions for My Calendar I.
//Solution: https://github.com/chai2010/LeetCode-in-Go/blob/master/Algorithms/0729.my-calendar-i/my-calendar-i.go
func (c *MyCalendar) Book(start int, end int) bool {
	e := &event{
		start: start,
		end:   end,
	}

	if c.head == nil ||
		e.end <= c.head.start {
		e.next, c.head = c.head, e // Insert new node as head
		return true
	}

	t := &event{next: c.head}
	for t.next != nil && t.next.end <= e.start {
		t = t.next // travel to the last node which is considering to insert
	}

	if t.next == nil ||
		e.end <= t.next.start {
		e.next, t.next = t.next, e
		return true
	}

	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
