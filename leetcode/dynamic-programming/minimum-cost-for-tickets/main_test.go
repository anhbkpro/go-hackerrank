package minimum_cost_for_tickets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMincostTickets(t *testing.T) {
	assert.Equal(t, 11, mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}
