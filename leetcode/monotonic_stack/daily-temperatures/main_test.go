package daily_temperatures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
