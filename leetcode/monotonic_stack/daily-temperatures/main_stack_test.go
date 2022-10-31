package daily_temperatures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperaturesStack(t *testing.T) {
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, DailyTemperaturesStack([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
