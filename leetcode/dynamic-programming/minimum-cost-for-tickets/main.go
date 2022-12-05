package minimum_cost_for_tickets

import (
	"fmt"
	"math"
	"sort"
)

//Runtime: 3 ms, faster than 70.00% of Go online submissions for Minimum Cost For Tickets.
//Memory Usage: 2.4 MB, less than 74.00% of Go online submissions for Minimum Cost For Tickets.
func mincostTickets(days []int, costs []int) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	isTravelDay := func(day int) bool {
		idx := sort.SearchInts(days, day)
		return days[idx] == day
	}

	sums := make([]int, days[len(days)-1]+1)
	var lastDay, tally int
	for i := 1; i <= days[len(days)-1]; i++ {
		if !isTravelDay(i) {
			sums[i] = tally
			fmt.Printf("=== Not a travel day %d tally = %d\n", i, tally)
			continue
		}
		fmt.Printf("=== A travel day %d tally = %d\n", i, tally)
		sums[i] = math.MaxInt32
		for j := 0; j < 3; j++ {
			switch j {
			case 0:
				lastDay = i - 1
			case 1:
				lastDay = i - 7
			case 2:
				lastDay = i - 30
			}
			if lastDay < 0 {
				lastDay = 0
			}
			sums[i] = min(sums[i], sums[lastDay]+costs[j])
			tally = sums[i]
		}
	}
	return sums[len(sums)-1]
}
