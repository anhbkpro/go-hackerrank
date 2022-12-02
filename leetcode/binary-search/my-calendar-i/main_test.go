package my_calendar_i

import "testing"

func TestMyCalendar_Book(t *testing.T) {
	c := Constructor()
	c.Book(10, 20)
	c.Book(15, 25)
	c.Book(20, 30)
}
