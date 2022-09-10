package lisa_workbook

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkbook(t *testing.T) {
	input := []int32{4, 2, 6, 1, 10}
	got := Workbook(5, 3, input)
	assert.Equal(t, int32(4), got)

	input2 := []int32{3, 8, 15, 11, 14, 1, 9, 2, 24, 31}
	got2 := Workbook(10, 5, input2)
	assert.Equal(t, int32(8), got2)
}
